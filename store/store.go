package store

import (
	"hackaton/internal/models"
	"sync"
)

type Store struct {
	SessionRepo  *SessionRepository  `json:"sessionRepo"`
	VineyardRepo *VineyardRepository `json:"vineyardRepo"`
	sync.RWMutex
}

func New() *Store {
	return &Store{
		SessionRepo: &SessionRepository{
			Sessions: make(map[int]*models.Session, 0),
		},
		VineyardRepo: &VineyardRepository{
			Vineyards: make(map[int]*models.Vineyard, 0),
		},
	}
}

func (s *Store) Session() *SessionRepository {
	return s.SessionRepo
}

func (s *Store) Vineyard() *VineyardRepository {
	return s.VineyardRepo
}
