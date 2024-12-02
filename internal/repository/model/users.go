package model

import "time"

type Users struct {
	Id                uint64 `db:"id" json:"id"`
	Name              string `db:"name" json:"name"`
	Email             string `db:"email" json:"email"`
	Password          string `db:"password" json:"password"`
	Age               uint   `db:"age" json:"age"`
	Gender            uint   `db:"gender" json:"gender"`
	Bio               string `db:"bio" json:"bio"`
	ProfilePictureUrl string `db:"profile_picture_url" json:"profile_picture_url"`
	IsVerified        bool   `db:"is_verified" json:"is_verified"`
	IsPremium         bool   `db:"is_premium" json:"is_premium"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
