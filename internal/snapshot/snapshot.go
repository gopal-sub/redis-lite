package snapshot


type StoreSnapshot struct {
	StoreID string
	StoreName string
	Store map[string]string
}

type DatabaseSnapshot struct {
	Stores []StoreSnapshot
}

func NewDatabaseSnapshot() *DatabaseSnapshot{
	return &DatabaseSnapshot{
		Stores: make([]StoreSnapshot, 0),
	}
}

func (d *DatabaseSnapshot) AddStoreSnapToDatabaseSnapshot(storeSnap *StoreSnapshot){
	d.Stores = append(d.Stores, *storeSnap)
}