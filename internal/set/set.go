package set

type Set struct{
	Name string
	Set map[string]struct{}
}


func NewSet(name string) *Set{
	return &Set{
		Name: name,
		Set: map[string]struct{}{},
	}
}

func (s *Set) Add(key string) error{
	if s.KeyExists(key){
		return KEYEXISTS
	}
	s.Set[key] = struct{}{}
	return nil
}

func (s *Set) Delete(key string) error{
	if !s.KeyExists(key){
		return KEYDOESNOTEXISTS
	}

	delete(s.Set, key)
	return nil
}

func (s *Set) KeyExists(key string) bool{
	_, ok := s.Set[key]
	return ok
}


