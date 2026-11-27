package server

import(
	"redis-lite/internal/set"
)


func (s *Server) AddSetToGlobalSetList(set *set.Set){
	s.Sets = append(s.Sets, set)
}

func (s *Server) FindSetInSetList(Name string)*set.Set {
	for _, set:=  range s.Sets{
		if set.Name == Name{
			return set
		}
	}
	return nil
}