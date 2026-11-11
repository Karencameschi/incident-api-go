package main

import (
	"sync"
	"time"
)

type Incident struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Store struct {
	mu        sync.Mutex
	incidents []Incident
	nextID    int
}

func (s *Store) List() []Incident {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]Incident, len(s.incidents))
	copy(result, s.incidents)
	return result
}

func (s *Store) Create(title, description string) Incident {
	s.mu.Lock()
	defer s.mu.Unlock()

	incident := Incident{
		ID:          s.nextID,
		Title:       title,
		Description: description,
		Status:      "open",
		CreatedAt:   time.Now(),
	}

	s.incidents = append(s.incidents, incident)
	s.nextID++
	return incident
}
