package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"

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
