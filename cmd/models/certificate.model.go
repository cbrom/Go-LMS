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

	Course             *Course               `gorm:"foreignkey:CourseID"`
	Issuer             *CourseAuthor         `gorm:"foreignkey:IssuerID"`
	IssuedCertificates IssuedCertificateList `gorm:"foreignkey:CertificateID"`
}

// TableName gorm standard table name
func (c *Certificate) TableName() string {
	return certificateTableName
}

// CertificateList defines array of certificate objects
type CertificateList []*Certificate

// TableName gorm standard table name
func (c *CertificateList) TableName() string {
	return certificateTableName
}

/**
CRUD functions
*/

// Create creates a new certificate record
func (c *Certificate) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Certificate by id
func (c *Certificate) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Certificates
func (c *Certificate) FetchAll(cl *CertificateList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given certificate
func (c *Certificate) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes certificate by id
func (c *Certificate) Delete() error {
	err := handler.Delete(c).Error
	return err
}
