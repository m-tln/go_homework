package storage

type IStorage interface {
	SearchByID(string) (any, bool)
	Add(any) string
}
