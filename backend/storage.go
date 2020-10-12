package backend

import js "github.com/vugu/vugu/js"


func NewStorage() Storage {
	return Storage{store: js.Global().Get("localStorage")}
}

type Storage struct {
	store js.Value
}

func (s *Storage) Test() string {

	//s.Set()

	return s.Get()
}

func (s *Storage) Get() string {
	return s.store.Get("test").String()
}

func (s *Storage) Set() {
	s.store.Set("test", "123")
}

type StorageRef struct {
	*Storage
}

type StorageSetter interface {
	SetStorage(s *Storage)
}

func (sr *StorageRef) SetStorage(s *Storage) {
	sr.Storage = s
}
