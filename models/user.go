package models

import "time"

type User struct {
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	IsActive        bool      `json:"isActive"`
	CreatedByGoogle bool      `json:"createdByGoogle"`
	CreatedAt       time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time `bson:"updatedAt" json:"updatedAt,omitempty"`
}

type Users []*User
