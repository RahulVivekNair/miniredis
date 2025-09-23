package core

import (
	"sync"
)

type Store struct {
	data sync.Map
}

func NewStore() *Store {
	return &Store{}
}
