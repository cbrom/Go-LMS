package schemas

import (
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// IssuedCertificateSchema graphql schema of issued certificate model
var IssuedCertificateSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IssuedCertificate",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					ic := p.Source.(*models.IssuedCertificate)
					return ic.GetID(), nil
				},
			},
			"serial_number": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

// CreateIssuedCertificateSchema contains fields to create a new issued certificate
var CreateIssuedCertificateSchema = graphql.FieldConfigArgument{
	"user_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"certificate_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// IssuedCertificateFromSchema is an adapter for issued certificate
func IssuedCertificateFromSchema(p graphql.ResolveParams) models.IssuedCertificate {
	issuedCertificate := models.IssuedCertificate{
		UserID:        p.Args["user_id"].(string),
		CertificateID: p.Args["certificate_id"].(string),
	}

	return issuedCertificate
}

// UpdateIssuedCertificateSchema contains fields to update an issued certificate
var UpdateIssuedCertificateSchema = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"user_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"certificate_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// IssuedCertificateFromUpdateSchema is an adapter for issued certificate
func IssuedCertificateFromUpdateSchema(p graphql.ResolveParams) models.IssuedCertificate {
	issuedCertificate := models.IssuedCertificate{}

	if userID, ok := p.Args["user_id"]; ok {
		issuedCertificate.UserID = userID.(string)
	}

	if certificateID, ok := p.Args["certificate_id"]; ok {
		issuedCertificate.CertificateID = certificateID.(string)
	}
	issuedCertificate.SetID(p.Args["id"].(string))

	return issuedCertificate
}
