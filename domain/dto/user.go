package dto

type User struct {
	Id         uint64
	Name       string
	Email      string
	Password   string
	Age        uint
	Gender     uint
	Bio        string
	Photo      string
	IsVerified bool
	IsPremium  bool
}
