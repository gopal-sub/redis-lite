package snapshotcodec

import (
	"encoding/binary"
	"redis-lite/internal/snapshot"
	"os"
)

type Encoder struct {
	file *os.File
	fileBytes []byte
}

func NewEncoder(file *os.File) *Encoder{
	return &Encoder{
		file: file,
		fileBytes: make([]byte, 0),

	}
}


func (e *Encoder) Encode(dbsnapshot *snapshot.DatabaseSnapshot) {


	//first add Magicnumber this is 8 bytes
	magicBytes := []byte(Magicnumber)
	e.fileBytes = append(e.fileBytes, magicBytes...)


	//add the version of the format format is 2 bytes
	buff := make([]byte, 2)
	binary.BigEndian.PutUint16(buff, Formatversion)
	e.fileBytes = append(e.fileBytes, buff...)

	//number of stores 4 bytes
	numberStores := len(dbsnapshot.Stores)
	buff = make([]byte, StoreCountSize)
	binary.BigEndian.PutUint32(buff, uint32(numberStores))
	e.fileBytes = append(e.fileBytes, buff...)

	//encode stores
	for _, store := range dbsnapshot.Stores{
		storeIdBytes := []byte(store.StoreID)

		//represent storeID lenght in bytes and in LengthSize for v1 its 4 bytes
		storeIdLength := len(storeIdBytes)
		buff = make([]byte, LengthSize)
		binary.BigEndian.PutUint32(buff, uint32(storeIdLength))

		e.fileBytes = append(e.fileBytes, buff...)

		e.fileBytes = append(e.fileBytes, storeIdBytes...)

		

	}




}