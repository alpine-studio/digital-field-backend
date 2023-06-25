package store

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"hackaton/internal/models"

	"strconv"
	"time"
)

type SessionRepository struct {
	store    *Store
	Sessions map[int]*models.Session `json:"sessions,omitempty"`
}

// CreateToken вспомогательная функция создания токена для сессии.
func (r *SessionRepository) createToken() string {
	b := md5.Sum([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	token := hex.EncodeToString(b[:])

	return token
}

func (r *SessionRepository) Create(role string) (*models.Session, error) {
	_, ok := models.PossibleRoles[role]
	if !ok {
		return nil, errors.New("роли не существует")
	}

	s := &models.Session{
		Token: r.createToken(),
	}
	s.ID = len(r.Sessions)
	r.Sessions[s.ID] = s

	return s, nil
}

func (r *SessionRepository) Find(token string) (*models.Session, error) {
	for _, s := range r.Sessions {
		if s.Token == token {
			return s, nil
		}
	}

	return nil, errors.New("токен не найден")
}
