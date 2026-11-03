package store


type Store struct {
	Store map[string]string
	StoreID string
	StoreName string
}

func NewStore(storename string) *Store{
	return &Store{
		Store: map[string]string{},
		StoreID: "abdlandlandlnadls",// make this dynamic and uuid
		StoreName: storename,

	}
}


func (s *Store) Set(key string, value string) error{
	if key == "" || value == ""{
		return INVALIKEYORVAL
	}
	s.Store[key] = value
	return nil
}

func (s *Store) Get(key string) (string, error){
	if key == "" {
		return "", INVALIKEYORVAL
	}
	val, ok := s.Store[key]
	if !ok {
		return "", KEYDOESNOTEXIST

	}
	return val, nil
}



func (s *Store) Delete(key string) error{
	if key == "" {
		return INVALIKEYORVAL
	}
	_, ok := s.Store[key]
	if !ok {
		return KEYDOESNOTEXIST

	}
	delete(s.Store, key)
	return nil

}