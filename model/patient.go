package model

import (
	// "time"
	"covid19-api-gofiber/config"
	"errors"
	"fmt"
)

// Patient mapped from table <patients>
type Patient struct {
	Name      string    `gorm:"column:name;not null" json:"name" validate:"ne=null"`
	Phone     string    `gorm:"column:phone;not null" json:"phone" validate:"ne=null"`
	Address   string    `gorm:"column:address" json:"address" validate:"ne=null"`
	Status    string    `gorm:"column:status;not null" json:"status" validate:"ne=null"`
	InDateAt  string	`gorm:"column:in_date_at" json:"in_date_at" validate:"ne=null"`
	OutDateAt string 	`gorm:"column:out_date_at" json:"out_date_at" validate:"ne=null"`
	Timestamp string 	`gorm:"column:timestamp" json:"timestamp" validate:"ne=null"`
}


func (p *Patient) All() ([]Patient, error) {
	var patients []Patient;

	db := config.CreateConnection();

	// Get all records
	result := db.Find(&patients);
	
	if result.Error != nil {
		return nil, errors.New(`something went wrong, please check the details in the database`)
	}

	config.CloseConnection(db);

	return patients, nil
}

func (p *Patient) Store(patient Patient) error {
	db := config.CreateConnection();

    // Insert data into model
    result := db.Create(&patient);

	if result.Error != nil {
		return errors.New(`something went wrong, please check the details in the database`)
	}

	config.CloseConnection(db);

	return nil;
}