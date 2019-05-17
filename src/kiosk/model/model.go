package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Group struct {
	gorm.Model
	Name     	string 		`json:"name"`
	Email  		string 		`json:"email"`
	Location  	string 		`gorm:"default:null" json:"location"`
	PhoneNumber string   	`gorm:"default:null" json:"phone_number"`
	IsActive	bool		`gorm:"default:1" json:"is_active"`
}

type Location struct {
	gorm.Model
	Name     	string     	`json:"name"`
	Address  	string     	`gorm:"default:null" json:"address"`
	Lat  		float64 	`gorm:"default:null" json:"lat"`
	Long      	float64    	`gorm:"default:null" json:"long"`
	Agents    	[]Agent 	`gorm:"ForeignKey:LocationID" json:"agents"`
}

type Agent struct {
	gorm.Model
	Name    	string     	`json:"name"`
	Phone  		string     	`json:"phone"`
	LocationID 	uint		`json:"location_id"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Group{}, &Location{}, &Agent{})
	db.Model(&Agent{}).AddForeignKey("location_id", "locations(id)", "CASCADE", "CASCADE")
	return db
}
