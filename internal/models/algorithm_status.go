package models

type AlgorithmStatus struct {
	ID       int64 `gorm:"primaryKey;auto_increment;column:id"`
	ClientID int64 `gorm:"column:client_id"`
	VWAP     bool  `gorm:"column:VWAP"`
	TWAP     bool  `gorm:"column:TWAP"`
	HFT      bool  `gorm:"column:HFT"`
}
