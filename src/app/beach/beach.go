package beach

import (
	"time"
)

type Beach struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	
	Ranking int    `json:"ranking"`
	Name    string `json:"name"`
	State   string `json:"state"`
}

// List all beaches
func List() ([]Beach, error) {
	return persist.list()
}

// Create register a new beach
func (beach *Beach) Create() error {
	return persist.create(beach)
}
