package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

var (
	issuedCertificateTableName = "issued_certificates"
)

// IssuedCertificate defines a model for user certificates for a course
type IssuedCertificate struct {
	utils.Base
	CertificateID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	UserID        string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	SerialNumber  string
	Certificate   *Certificate `gorm:"foreignkey:CertificateID"`
}

// TableName gorm standard table name
func (i *IssuedCertificate) TableName() string {
	return issuedCertificateTableName
}

// IssuedCertificateList defines array of certificate objects
type IssuedCertificateList []*IssuedCertificate

// TableName gorm standard table name
func (i *IssuedCertificateList) TableName() string {
	return issuedCertificateTableName
}

/**
CRUD functions
*/

// Create creates a new certificate record
func (i *IssuedCertificate) Create() error {
	possible := handler.NewRecord(i)
	if possible {
		if err := handler.Create(i).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches IssuedCertificate by id
func (i *IssuedCertificate) FetchByID() error {
	err := handler.First(i).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all IssuedCertificates
func (i *IssuedCertificate) FetchAll(il *IssuedCertificateList) error {
	err := handler.Find(il).Error
	return err
}

// UpdateOne updates a given certificate
func (i *IssuedCertificate) UpdateOne() error {
	err := handler.Save(i).Error
	return err
}

// Delete deletes certificate by id
func (i *IssuedCertificate) Delete() error {
	err := handler.Delete(i).Error
	return err
}
