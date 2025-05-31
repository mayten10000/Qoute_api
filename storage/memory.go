package storage

import (
	"math/rand"
	"sync"

	"quote-service/models"
)

type MemoryStorage struct {
	sync.Mutex
	quotes []models.Quote
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		quotes: []models.Quote{},
		nextID: 1,
	}
}

func (s *MemoryStorage) AddQuote(q models.Quote) models.Quote {
	s.Lock()
	defer s.Unlock()
	q.ID = s.nextID
	s.nextID++
	s.quotes = append(s.quotes, q)
	return q
}

func (s *MemoryStorage) GetAll() []models.Quote {
	s.Lock()
	defer s.Unlock()
	return append([]models.Quote(nil), s.quotes...)
}

func (s *MemoryStorage) GetRandom() (models.Quote, bool) {
	s.Lock()
	defer s.Unlock()
	if len(s.quotes) == 0 {
		return models.Quote{}, false
	}
	return s.quotes[rand.Intn(len(s.quotes))], true
}

func (s *MemoryStorage) GetByAuthor(author string) []models.Quote {
	s.Lock()
	defer s.Unlock()
	var result []models.Quote
	for _, q := range s.quotes {
		if q.Author == author {
			result = append(result, q)
		}
	}
	return result
}

func (s *MemoryStorage) DeleteByID(id int) bool {
	s.Lock()
	defer s.Unlock()
	for i, q := range s.quotes {
		if q.ID == id {
			s.quotes = append(s.quotes[:i], s.quotes[i+1:]...)
			return true
		}
	}
	return false
}
