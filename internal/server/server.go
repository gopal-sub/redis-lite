package server

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"redis-lite/internal/client"
	"strings"
	"time"
	"redis-lite/internal/framer"
	"redis-lite/internal/parser"
	p "redis-lite/internal/persistence"
	"redis-lite/internal/protocol"
	"redis-lite/internal/store"
	"redis-lite/internal/snapshotcodec"
)



type Server struct{
	Address string
	Listner net.Listener
	Clients map[*client.Client]struct{}
	GlobalStores []*store.Store
	SnapshotWorker *p.SnapshotWorker
}

func NewServer(address string) *Server{
	listner, _ := net.Listen("tcp", address)
	return &Server{
		Address: address,
		Listner: listner,
		Clients: make(map[*client.Client]struct{}),
	}
}

func (s *Server) Start(){
	//create a default store called default
	defaultStore := store.NewStore(store.DEFAULTSTORE)
	persist := p.NewPersist(defaultStore)

	snapshotWorker := p.NewSnapshotWorker(3*time.Minute)
	s.SnapshotWorker = snapshotWorker
	snapshotWorker.AppendStorePersistenceToSnapShotWorker(persist)
	go snapshotWorker.SnapshotWorker()

	s.AddStoreToGlobalStoreList(defaultStore)

	fmt.Println("Server start")

	for {
		conn, _ := s.Listner.Accept()
		newClient := client.NewClient(conn, defaultStore)
		s.AddClientToGlobalList(newClient)
		go s.HandleClientConnection(newClient)

	}
}





func (s *Server) Recover(filename *os.File){
	decoder := snapshotcodec.NewDecoder(filename)
	err := decoder.Decode()
	if err != nil{
		fmt.Println(err)
		return
	}
	snapshotWorker := p.NewSnapshotWorker(3*time.Minute)
	s.SnapshotWorker = snapshotWorker
	for _, store_ := range decoder.Dbsnapshot.Stores{
		newStore := store.NewStore(store_.StoreName)
		newStore.StoreID = store_.StoreID
		newStore.Store = store_.Store

		persist := p.NewPersist(newStore)
		snapshotWorker.AppendStorePersistenceToSnapShotWorker(persist)
		s.AddStoreToGlobalStoreList(newStore)
	}
	go snapshotWorker.SnapshotWorker()

	ok, defaultStore := s.FindStore(store.DEFAULTSTORE)
	if !ok{
		defaultStore = store.NewStore(store.DEFAULTSTORE)
		persist := p.NewPersist(defaultStore)
		s.SnapshotWorker.AppendStorePersistenceToSnapShotWorker(persist)

	}

	fmt.Println("Recovery complete, Server started")
	for {
		conn, _ := s.Listner.Accept()
		newClient := client.NewClient(conn, defaultStore)
		s.AddClientToGlobalList(newClient)
		go s.HandleClientConnection(newClient)

	}

}	


func (s *Server) HandleClientConnection(client *client.Client) {
	buffer := make([]byte, 1024)
	clientframer := framer.Newframer()
	for {
		n , err := client.Conn.Read(buffer)
		if errors.Is(err, io.EOF){
			fmt.Println("client disconected")
			return
		}
		if err != nil {
			fmt.Println(err)

		}
		validmessages := clientframer.Feed(buffer[:n])
		//clear the buffer for safety
		clear(buffer)
		commands := []*protocol.COMMAND{}
		for _ , message :=  range validmessages{
			command, err := parser.Parser(message)
			if err != nil {
				break
				//if [valid, valid, invalid, valid] -> last valid wont be processed
			}
			commands = append(commands, command)
		}

		for _ , cmd := range commands{
			switch cmd.Type {
			case protocol.SET:
				keyIndex := strings.Index(cmd.Arg, " ")
				err := client.CurrentStore.Set(cmd.Arg[:keyIndex], cmd.Arg[keyIndex+1:])
				if err != nil {
					fmt.Println(err)
				}
			case protocol.GET:
				value, err := client.CurrentStore.Get(cmd.Arg)
				if err != nil {
					fmt.Println(err)
					client.SendDataToClient(err.Error())
					break
				}
				client.SendDataToClient(value)
				// push value to client
			case protocol.DELETE:
				err := client.CurrentStore.Delete(cmd.Arg)
				if err != nil {
					fmt.Println(err)
					client.SendDataToClient(err.Error())
				}
			case protocol.SELECT:
				storeName := cmd.Arg
				ok, store := s.FindStore(storeName)
				if !ok {
					client.SendDataToClient("Store does not exist")
					break
				}
				client.CurrentStore = store
				data := fmt.Sprintf("Current Store %v", client.CurrentStore.StoreName)
				client.SendDataToClient(data)

			case protocol.CREATESTORE:
				storeName := cmd.Arg
				ok, _ := s.FindStore(storeName)
				if ok {
					fmt.Println("store exist")
					//store exist cant create
					client.SendDataToClient("Store already exists")
					break
				}
				newStore := store.NewStore(storeName)
				s.AddStoreToGlobalStoreList(newStore)
				persist := p.NewPersist(newStore)
				s.SnapshotWorker.AppendStorePersistenceToSnapShotWorker(persist)
				data := fmt.Sprintf("Created Store %v", newStore.StoreName)
				client.SendDataToClient(data)
			}
		}

	}
	


}