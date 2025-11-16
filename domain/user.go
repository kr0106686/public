package domain

import "time"

type User struct {
	Provider   string    `json:"provider"`
	ProviderID string    `json:"provider_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Picture    string    `json:"picture"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
