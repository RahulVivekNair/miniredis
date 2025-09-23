package core

import (
	"sync"
	"time"
)

type Store struct {
	Data   sync.Map
	Expiry sync.Map
}

func NewStore() *Store {
	store := &Store{}
	store.StartJanitor()
	return store
}

func (s *Store) StartJanitor() {
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			s.Expiry.Range(func(key, value any) bool {
				expTime := value.(time.Time)
				if time.Now().After(expTime) {
					s.Data.Delete(key)
					s.Expiry.Delete(key)
				}
				return true
			})
		}
	}()
}
