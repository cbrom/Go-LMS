package schemas

import (
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// CertificateSchema graphql schema of certificate model
var CertificateSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Certificate",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					c := p.Source.(*models.Certificate)
					return c.GetID(), nil
				},
			},
			"qr_corner": &graphql.Field{
				Type: graphql.Int,
			},
			"qr_scale": &graphql.Field{
				Type: graphql.Int,
			},
			"margin": &graphql.Field{
				Type: graphql.Int,
			},
			"name_offset_top": &graphql.Field{
				Type: graphql.Int,
			},
			"font_size": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"active": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	})

// CreateCertificateSchema contains fields to create a new certificate
var CreateCertificateSchema = graphql.FieldConfigArgument{
	"qr_corner": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"qr_scale": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"margin": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"name_offset_top": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"font_size": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"message": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"active": &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	},
	"course_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"course_author_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// CertificateFromSchema is an adapter for certificate schema
func CertificateFromSchema(p graphql.ResolveParams) models.Certificate {
	certificate := models.Certificate{
		QRCorner:       p.Args["qr_corner"].(string),
		QRScale:        p.Args["qr_scale"].(int),
		Margin:         p.Args["margin"].(int),
		NameOffsetTop:  p.Args["name_offset_top"].(int),
		FontSize:       p.Args["font_size"].(int),
		Message:        p.Args["message"].(string),
		Active:         p.Args["active"].(bool),
		CourseID:       p.Args["course_id"].(string),
		CourseAuthorID: p.Args["course_author_id"].(string),
	}
	return certificate
}
