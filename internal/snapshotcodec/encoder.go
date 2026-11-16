package snapshotcodec

import (
	"encoding/binary"
	"fmt"
	"os"
	"redis-lite/internal/snapshot"
)

type Encoder struct {
	File *os.File
	FileBytes []byte
}

func NewEncoder(file *os.File) *Encoder{
	return &Encoder{
		File: file,
		FileBytes: make([]byte, 0),

	}
}


func (e *Encoder) Encode(dbsnapshot *snapshot.DatabaseSnapshot) {


	//first add Magicnumber this is 8 bytes
	magicBytes := []byte(Magicnumber)
	//magicByte is in asccii encoding
	fmt.Println("magicBytes")
	fmt.Println(magicBytes)
	e.FileBytes = append(e.FileBytes, magicBytes...)


	//add the version of the format, format is 2 bytes
	buff := make([]byte, 2)
	binary.BigEndian.PutUint16(buff, Formatversion)
	e.FileBytes = append(e.FileBytes, buff...)

	//number of stores 4 bytes
	numberStores := len(dbsnapshot.Stores)
	buff = make([]byte, StoreCountSize)
	binary.BigEndian.PutUint32(buff, uint32(numberStores))
	e.FileBytes = append(e.FileBytes, buff...)

	//encode stores
	for _, store := range dbsnapshot.Stores{
		storeIdBytes := []byte(store.StoreID)


		//represent storeID lenght in bytes and in LengthSize for v1 its 4 bytes
		storeIdLength := len(storeIdBytes)
		buff = make([]byte, LengthSize)
		binary.BigEndian.PutUint32(buff, uint32(storeIdLength))
		// 3 -> ABC
		// 6 -> ABCr43

		e.FileBytes = append(e.FileBytes, buff...)

		e.FileBytes = append(e.FileBytes, storeIdBytes...)

		// store name format
		// 5 -> Cache
		// 4 -> test
		// store 4 in 
		storeNameLen := len(store.StoreName)
		buff = make([]byte, LengthSize)
		binary.BigEndian.PutUint32(buff, uint32(storeNameLen))

		e.FileBytes = append(e.FileBytes, buff...)

		storeNameByte := []byte(store.StoreName)
		e.FileBytes = append(e.FileBytes, storeNameByte...)

		buff = make([]byte, EntryCountSize)
		numberEntriesStore := len(store.Store)

		binary.BigEndian.PutUint32(buff, uint32(numberEntriesStore))
		e.FileBytes = append(e.FileBytes, buff...)
		for key, val := range store.Store{
			// Key encoding
			// 10 -> alexcookie
			buff = make([]byte, LengthSize)
			keyLength := len(key)
			binary.BigEndian.PutUint32(buff, uint32(keyLength))
			e.FileBytes = append(e.FileBytes, buff...)
			keyByte := []byte(key)
			e.FileBytes = append(e.FileBytes, keyByte...)
			clear(buff)

			// value encoding
			// 4 -> 3cr4
			valLen := len(val)
			binary.BigEndian.PutUint32(buff, uint32(valLen))
			e.FileBytes = append(e.FileBytes, buff...)

			valByte := []byte(val)
			e.FileBytes = append(e.FileBytes, valByte...)
		}
	}
}