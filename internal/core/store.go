package core

import (
	"sync"
)

type Store struct {
	Data sync.Map
}

func NewStore() *Store {
	return &Store{}
}
