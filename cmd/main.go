package main

import (
	// "redis-lite/internal/server"
	"fmt"
	"os"
	"redis-lite/internal/snapshot"
	"redis-lite/internal/snapshotcodec"
)
	

func main(){

	// server := server.NewServer(":3000")
	// server.Start()
	dbsnap := snapshot.DatabaseSnapshot{}
	store1 := snapshot.StoreSnapshot{
		StoreID: "123A",
		StoreName: "localcache",
		Store: make(map[string]string, 0),
	}
	store1.Store["user1"] = "alex43@gmail.com"
	store1.Store["user2"] = "johnnyharris@newstudio.com"

	store2 := snapshot.StoreSnapshot{
		StoreID: "123B",
		StoreName: "cookies",
		Store: make(map[string]string, 0),
	}

	store2.Store["alex43@gmail.com"] = "312fedwfvf"
	dbsnap.Stores = append(dbsnap.Stores, store1)
	dbsnap.Stores = append(dbsnap.Stores, store2)
	f, err :=os.Create("./dump/1794371759.dump")
	if err != nil {
		fmt.Println(err)
	}

	encoder := snapshotcodec.NewEncoder(f)
	encoder.Encode(&dbsnap)
	fmt.Printf("% X\n", encoder.FileBytes)

}