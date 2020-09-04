package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// CreateCertificate creates a new certificate
func CreateCertificate(p graphql.ResolveParams) (interface{}, error) {
	certificate := schemas.CertificateFromSchema(p)
	if err := certificate.Create(); err == nil {
		return certificate.GetID(), nil
	}

	return nil, errors.New("Unable to create certificate")
}

// UpdateCertificate updates an existing certificate
func UpdateCertificate(p graphql.ResolveParams) (interface{}, error) {
	certificate := schemas. CertificateFromUpdateSchema(p)
	if err := certificate.UpdateOne(); err == nil {
		return certificate.GetID(), nil
	}

	return nil, errors.New("Unable to update certificate ")
}

// DeleteCertificate deletes an existing certificate
func DeleteCertificate(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		certificate := &models.Certificate{}
		certificate.SetID(idQuery)
		err := certificate.SoftDelete()
		return nil, err
	}

	return nil, errors.New("Certificate id not provided")
}

// CreateIssuedCertificate issues a new certificate
func CreateIssuedCertificate(p graphql.ResolveParams) (interface{}, error) {
	issuedCertificate := schemas.IssuedCertificateFromSchema(p)
	if err := issuedCertificate.Create(); err == nil {
		return issuedCertificate.GetID(), nil
	}

	return nil, errors.New("Unable to create issued certificate")
}

// UnissueCertificate deletes an issued certificate
func UnissueCertificate(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		issuedCertificate := &models.IssuedCertificate{}
		issuedCertificate.SetID(idQuery)
		err := issuedCertificate.SoftDelete()
		return nil, err
	}

	return nil, errors.New("Issued Certificate id not provided")
}
