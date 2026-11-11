package persistence

import (
	"redis-lite/internal/store"
)

type Snapshot interface{
	Save()
	Load(filename string)
}

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


type Persist struct{
	Store *store.Store
}

func NewPersist(Store *store.Store) *Persist{
	return &Persist{
		Store: Store,
	}
}


func (p *Persist) CopyLiveStoreDataToBuffer() *StoreSnapshot{
	p.Store.Readlock.RLock()
	defer p.Store.Readlock.RUnlock()

	storeSnap := StoreSnapshot{
		StoreID: p.Store.StoreID,
		StoreName: p.Store.StoreName,
		Store: make(map[string]string),
	}

	for k, v := range p.Store.Store{
		storeSnap.Store[k] = v
	}

	return &storeSnap
}




