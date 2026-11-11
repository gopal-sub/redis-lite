package server


import (
	s "redis-lite/internal/store"
	// p "redis-lite/internal/persistence"
)

func (s *Server)AddStoreToGlobalStoreList(store *s.Store){
	s.GlobalStores = append(s.GlobalStores, store)
}

func (s *Server) FindStore(storeName string) (bool,*s.Store){
	for _, val := range s.GlobalStores{
		if val.StoreName == storeName{
			return true, val
		}
	}
	return false, nil
}