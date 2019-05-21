package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Group struct {
	gorm.Model
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Location    string  `gorm:"default:null" json:"location"`
	PhoneNumber string  `gorm:"default:null" json:"phone_number"`
	IsActive    bool    `gorm:"default:1" json:"is_active"`
	Users       []User  `gorm:"ForeignKey:GroupID" json:"users"`
	Groups      []Group `gorm:"ForeignKey:GroupID" json:"locations"`
}

type Location struct {
	gorm.Model
	Name          string         `json:"name"`
	Address       string         `gorm:"default:null" json:"address"`
	Lat           float64        `gorm:"default:null" json:"lat"`
	Long          float64        `gorm:"default:null" json:"long"`
	GroupID       uint           `json:"group_id"`
	LocationUsers []LocationUser `gorm:"ForeignKey:LocationID" json:"location_users"`
}

type User struct {
	gorm.Model
	Name          string         `json:"name"`
	Phone         string         `json:"phone"`
	Email         string         `json:"email"`
	UserName      string         `json:"user_name"`
	Password      string         `json:"password"`
	GroupID       uint           `json:"group_id"`
	Roles         []Role         `gorm:"ForeignKey:UserID" json:"roles"`
	LocationUsers []LocationUser `gorm:"ForeignKey:UserID" json:"location_users"`
}

type Role struct {
	gorm.Model
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	UserID uint   `json:"user_id"`
}

type LocationUser struct {
	gorm.Model
	UserID     uint `json:"user_id"`
	LocationID uint `json:"location_id"`
}

type Booking struct {
	gorm.Model
	LocationID  uint    `json:"location_id"`
	CustomerID  string  `json:"customer_id"`
	Destination string  `json:"destination"`
	Lat         float64 `gorm:"default:null" json:"lat"`
	Long        float64 `gorm:"default:null" json:"long"`
	Status      string  `json:"status"`
	RouteID     uint    `json:"route_id"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(
		&Group{},
		&Location{},
		&User{},
		&Role{},
		&LocationUser{},
		&Booking{})
	db.Model(&User{}).AddForeignKey("group_id", "groups(id)", "CASCADE", "CASCADE")
	db.Model(&Role{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Location{}).AddForeignKey("group_id", "groups(id)", "CASCADE", "CASCADE")
	db.Model(&LocationUser{}).AddForeignKey("location_id", "locations(id)", "CASCADE", "CASCADE")
	db.Model(&LocationUser{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	return db
}
