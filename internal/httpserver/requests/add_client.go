package requests

import (
	"encoding/json"
	"time"
)

type AddClientRequest struct {
	ClientName  string    `json:"client_name"`
	Version     int       `json:"version"`
	Image       string    `json:"image"`
	CPU         string    `json:"cpu"`
	Memory      string    `json:"memory"`
	NeedRestart bool      `json:"need_restart"`
	SpawnedAt   time.Time `json:"spawned_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (addClient *AddClientRequest) UnmarshalJSON(data []byte) error {
	type tmp AddClientRequest
	aux := &struct {
		SpawnedAt string `json:"spawned_at"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		*tmp
	}{
		tmp: (*tmp)(addClient),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	addClient.SpawnedAt, err = time.Parse("2006-01-02", aux.SpawnedAt)
	if err != nil {
		return err
	}

	addClient.CreatedAt, err = time.Parse("2006-01-02", aux.CreatedAt)
	if err != nil {
		return err
	}

	addClient.UpdatedAt, err = time.Parse("2006-01-02", aux.UpdatedAt)
	return err
}
