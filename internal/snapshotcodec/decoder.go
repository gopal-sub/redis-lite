package snapshotcodec

import (
	"encoding/binary"
	"os"
	"redis-lite/internal/snapshot"

)

type Decoder struct{
	File *os.File
	Dbsnapshot snapshot.DatabaseSnapshot
}

func NewDecoder(file *os.File) *Decoder{
	return &Decoder{
		File: file,
		Dbsnapshot: snapshot.DatabaseSnapshot{},
	}
}

func (d *Decoder) Decode() error{
	// verify the file formar -> magic number
	buff := make([]byte, 8)
	d.File.Read(buff)
	if string(buff) != Magicnumber{
		return INVALIDFILEFORMAT
	}
	clear(buff)

	// check the version
	buff = make([]byte, 2)
	d.File.Read(buff)
	version := binary.BigEndian.Uint16(buff)
	if version != 1 {
		return INVALIDFILEVERSION
	}
	clear(buff)

	//number of stores
	buff = make([]byte, StoreCountSize)
	d.File.Read(buff)
	numberOfStores := binary.BigEndian.Uint32(buff)
	clear(buff)

	//initialize DatabaseSnapshot

	DBSnapshot := snapshot.DatabaseSnapshot{
		Stores: make([]snapshot.StoreSnapshot, 0),
	}

	for range numberOfStores{
		//initialize StoreSnapshot this will get appended to DatabaseSnapshot
		storeSN := snapshot.StoreSnapshot{
			Store: make(map[string]string),
		}
		// read store id
		//read how many bytes to get store id

		buff = make([]byte, LengthSize)
		d.File.Read(buff)
		storeIdLen_readBytes := binary.BigEndian.Uint32(buff)
		clear(buff)

		buff = make([]byte, storeIdLen_readBytes)
		d.File.Read(buff)
		storeId := string(buff)
		storeSN.StoreID = storeId

		clear(buff)

		//read store name
		//how many bytes to read to get store name
		buff = make([]byte, LengthSize)
		d.File.Read(buff)
		storeNameLen_readBytes := binary.BigEndian.Uint32(buff)
		clear(buff)

		buff = make([]byte, storeNameLen_readBytes)
		d.File.Read(buff)
		storeName := string(buff)

		storeSN.StoreName = storeName
		clear(buff)

		//find number of entries and loop over all
		buff = make([]byte, EntryCountSize)
		d.File.Read(buff)

		numberEntriesInStore := binary.BigEndian.Uint32(buff)
		clear(buff)
		for range numberEntriesInStore{
			//how many byte to get the key
			//get key length
			buff = make([]byte, LengthSize)
			d.File.Read(buff)
			keyLen_readBytes := binary.BigEndian.Uint32(buff)
			clear(buff)

			//read the key
			buff = make([]byte, keyLen_readBytes)
			d.File.Read(buff)
			key := string(buff)


			// value
			//how many byte to get the val
			//get val length
			buff = make([]byte, LengthSize)
			d.File.Read(buff)
			valLen_readBytes := binary.BigEndian.Uint32(buff)
			clear(buff)

			//read the key
			buff = make([]byte, valLen_readBytes)
			d.File.Read(buff)
			val := string(buff)

			storeSN.Store[key] = val
			clear(buff)
		}

		DBSnapshot.Stores = append(DBSnapshot.Stores, storeSN)

	}

	return nil

}