package db

import "gorm.io/gorm"

type Recovery struct {
	gorm.Model
	ChainId     uint64 `gorm:"type:bigint default 4690"`
	TxHash      string `gorm:"type:varchar(66)"`
	ComfirmHash string `gorm:"type:varchar(66)"`
	Account     string `gorm:"type:varchar(42)"`
	Status      uint
}
