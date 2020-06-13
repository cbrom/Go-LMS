package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

// IssuedCertificate defines a model for user certificates for a course
type IssuedCertificate struct {
	utils.Base
	CertificateID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	UserID        string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	SerialNumber  string
}
