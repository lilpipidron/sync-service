package models

import "time"

type Client struct {
	ID          int64     `gorm:"primary_key;column:id"`
	ClientName  string    `gorm:"column:client_name"`
	Version     int       `gorm:"column:version"`
	Image       string    `gorm:"column:image"`
	CPU         string    `gorm:"column:cpu"`
	Memory      string    `gorm:"column:memory"`
	NeedRestart bool      `gorm:"column:need_restart"`
	SpawnedAt   time.Time `gorm:"column:spawned_at"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
