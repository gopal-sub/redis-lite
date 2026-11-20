package snapshotcodec


const (
	Magicnumber = "RDLITE01"
	Formatversion uint16 = 1

	StoreCountSize uint32 = 4 //number of stores
	EntryCountSize uint32 = 4  // number of entries in each store max val for a store 2^32 -1
	LengthSize uint32 = 4 // lenght of each string or field
)

