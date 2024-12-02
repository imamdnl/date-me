package entity

import (
	"date-me/domain/dto"
	"date-me/domain/value_object"
	"date-me/pkg/config"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id         uint64
	name       string
	email      string
	password   string
	age        uint
	gender     *value_object.UserGender
	bio        string
	photo      string
	isVerified bool
	isPremium  bool
}

func NewUser(d *dto.User) (*User, error) {
	gender, err := value_object.NewUserGender(d.Gender)
	if err != nil {
		return nil, err
	}
	if d.Gender == 0 {
		d.Gender = 1
	}

	return &User{
		id:         d.Id,
		name:       d.Name,
		email:      d.Email,
		password:   d.Password,
		age:        d.Age,
		gender:     gender,
		bio:        d.Bio,
		photo:      d.Photo,
		isVerified: d.IsVerified,
		isPremium:  d.IsPremium,
	}, nil
}

func BuildUserEntityFromLoginRequest(d *dto.User) (*User, error) {
	if d.Email == "" || d.Password == "" {
		return nil, errors.New("email or password cannot be empty")
	}
	return &User{
		email:    d.Email,
		password: d.Password,
	}, nil
}

func (u *User) GetId() uint64 {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetAge() uint {
	return u.age
}

func (u *User) GetGender() *value_object.UserGender {
	return u.gender
}

func (u *User) GetBio() string {
	return u.bio
}

func (u *User) GetPhoto() string {
	return u.photo
}

func (u *User) GetIsVerified() bool {
	return u.isVerified
}

func (u *User) GetIsPremium() bool {
	return u.isPremium
}

func (u *User) CheckPassword(payload string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.GetPassword()), []byte(payload))
	if err != nil {
		config.Logger().Error("error check password", zap.Error(err))
		return errors.New("invalid password")
	}
	return nil
}

func (u *User) SetPasswordToHash() {
	p, _ := bcrypt.GenerateFromPassword([]byte(u.GetPassword()), bcrypt.DefaultCost)
	u.password = string(p)
}
