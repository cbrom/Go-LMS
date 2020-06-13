package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

var (
	certificateTableName = "certificates"
)

// Certificate defines a model for student certificates in a model
type Certificate struct {
	utils.Base
	CourseID      string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	IssuerID      string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	QRCorner      string
	QRScale       int
	Margin        int
	NameOffsetTop int
	FontSize      int
	Message       string
	Active        bool
}

// TableName gorm standard table name
func (u *Certificate) TableName() string {
	return certificateTableName
}

// CertificateList defines array of certificate objects
type CertificateList []*Certificate

// TableName gorm standard table name
func (u *CertificateList) TableName() string {
	return certificateTableName
}

/**
CRUD functions
*/

// Create creates a new certificate record
func (u *Certificate) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Certificate by id
func (u *Certificate) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Certificates
func (u *Certificate) FetchAll(ul *CertificateList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given certificate
func (u *Certificate) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes certificate by id
func (u *Certificate) Delete() error {
	err := handler.Delete(u).Error
	return err
}
