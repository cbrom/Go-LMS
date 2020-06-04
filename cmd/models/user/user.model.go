package user

import (
	"go-lms-of-pupilfirst/pkg/utils"
	"time"
)

const (
	// Database table for User
	userTableName = "users"
)

type User struct {
	utils.Base
	Email                  string     `gorm:"type:varchar(100);unique_index" json:"email" validate:"required,unique,email"`
	PasswordSalt           string     `json:"-"`
	PasswordHash           []byte     `json:"-"`
	SiginInCount           int        `json:"sign_in_count" validate:"omitempty"`
	CurrentSignInAt        *time.Time `json:"current_sign_in_at" validate:"omitempty"`
	LastSignInAt           *time.Time `json:"last_sign_in_at" validate:"omitempty"`
	CurrentSignInIP        string     `json:"-" validate:"omitempty,ip"`
	LastSignInIP           string     `json:"-" validate:"omitempty,ip"`
	RememberToken          string     `json:"remember_token" validate:"omitempty"`
	ConfirmedAt            *time.Time `json:"confirmed_at" validate:"omitempty"`
	ConfirmationMailSentAt *time.Time `json:"confirmation_mail_sent_at" validate:"omitempty"`
	Name                   string     `json:"name" validate:"min=3,max=10,omitempty"`
	Phone                  string     `json:"phone" validate:"omitempty"`
	Title                  string     `json:"title" validate:"omitempty"`
	KeySkills              string     `json:"key_skills" validate:"omitempty"`
	About                  string     `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}

type UserCreateRequest struct {
	Name            string  `json:"name" validate:"required" example:"Groot"`
	Email           string  `json:"email" validate:"required,email,unique" example:"groot@golms.com"`
	Password        string  `json:"password" validate:"required" example:"GrootSecret"`
	PasswordConfirm string  `json:"password_confirm" validate:"required,eqfield=password" example:"GrootSecret"`
	Timezone        *string `json:"timezone" validate:"required" example:"America/Anchorage"`
}

type UserInfoUpdateRequest struct {
	ID        string `json:"id" validate:"required,uuid" example:"c01bdef7-173f-4d29-3edc60baf6a2"`
	Name      string `json:"name" validate:"min=3,max=10,omitempty"`
	Phone     string `json:"phone" validate:"omitempty"`
	Title     string `json:"title" validate:"omitempty"`
	KeySkills string `json:"key_skills", validate:"omitempty"`
	About     string `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`
}

// TableName gorm standard table name
func (u *User) TableName() string {
	return userTableName
}

// GetID returns Id of the user
func (u *User) GetID() string {
	return u.ID
}

// SetID sets Id of the user
func (u *User) SetID(id string) {
	u.ID = id
}

// SetCreatedAt sets field createdAt, should only be used in mongodb
func (u *User) SetCreatedAt(t time.Time) {
	u.CreatedAt = t
}

// SetUpdatedAt sets field UpdatedAt
func (u *User) SetUpdatedAt(t time.Time) {
	u.UpdatedAt = t
}

// SetArchivedAt sets field DeletedAt
func (u *User) SetArchivedAt(t *time.Time) {
	u.ArchivedAt = t
}

// List defines array of user objects
type List []*User

// TableName gorm standard table name
func (u *List) TableName() string {
	return userTableName
}
