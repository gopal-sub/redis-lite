package persistence

import (
	"redis-lite/internal/store"
	"redis-lite/internal/snapshot"
)

type Snapshot interface{
	Save()
	Load(filename string)
}

type Persist struct{
	Store *store.Store
}

func NewPersist(Store *store.Store) *Persist{
	return &Persist{
		Store: Store,
	}
}


func (p *Persist) CopyLiveStoreDataToBuffer() *snapshot.StoreSnapshot{
	p.Store.Readlock.RLock()
	defer p.Store.Readlock.RUnlock()

	storeSnap := snapshot.StoreSnapshot{
		StoreID: p.Store.StoreID,
		StoreName: p.Store.StoreName,
		Store: make(map[string]string),
	}

	for k, v := range p.Store.Store{
		storeSnap.Store[k] = v
	}

	return &storeSnap
}




