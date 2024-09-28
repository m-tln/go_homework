package storage

type StorageMap struct {
	units map[string]any
	NewID func() string
}

func (s *StorageMap) SearchByID(id string) (any, bool) {
	res, ok := s.units[id]
	return res, ok
}

func (s *StorageMap) Add(elem any) string {
	if len(s.units) == 0 {
		s.units = make(map[string]any)
	}
	key := s.NewID()
	s.units[key] = elem
	return key
}
