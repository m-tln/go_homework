package storage

type StorageSlice struct {
	units []any
	id    []string
	NewID func() string
}

func (s *StorageSlice) SearchByID(id string) (any, bool) {
	for i, curId := range s.id {
		if curId == id {
			return s.units[i], true
		}
	}
	return nil, false
}

func (s *StorageSlice) Add(elem any) string {
	s.units = append(s.units, elem)
	key := s.NewID()
	s.id = append(s.id, key)
	return key
}

func (s *Storage) SetIDfunc(NewID func() string) {
	s.NewID = NewID
}
