package model

import (
	"time"
	"covid19-api-gofiber/config"
)

// Patient mapped from table <patients>
type Patient struct {
	Name      string    `gorm:"column:name;not null" json:"name"`
	Phone     string    `gorm:"column:phone;not null" json:"phone"`
	Address   string    `gorm:"column:address" json:"address"`
	Status    string    `gorm:"column:status;not null" json:"status"`
	InDateAt  time.Time `gorm:"column:in_date_at" json:"in_date_at"`
	OutDateAt time.Time `gorm:"column:out_date_at" json:"out_date_at"`
	Timestamp time.Time `gorm:"column:timestamp" json:"timestamp"`
}

func (p *Patient) All() []Patient {
	var patients []Patient

	db := config.CreateConnection();

	// Get all records
	result := db.Find(&patients)

	if result.Error != nil {
		panic(`Something went wrong, please check the details in the database`)
	}

	return patients
}