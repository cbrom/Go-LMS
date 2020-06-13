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
}

// TableName gorm standard table name
func (u *IssuedCertificate) TableName() string {
	return issuedCertificateTableName
}

// IssuedCertificateList defines array of certificate objects
type IssuedCertificateList []*IssuedCertificate

// TableName gorm standard table name
func (u *IssuedCertificateList) TableName() string {
	return issuedCertificateTableName
}

/**
CRUD functions
*/

// Create creates a new certificate record
func (u *IssuedCertificate) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches IssuedCertificate by id
func (u *IssuedCertificate) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all IssuedCertificates
func (u *IssuedCertificate) FetchAll(ul *IssuedCertificateList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given certificate
func (u *IssuedCertificate) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes certificate by id
func (u *IssuedCertificate) Delete() error {
	err := handler.Delete(u).Error
	return err
}
