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

// UpdateCertificateSchema contains fields to update a certificate
var UpdateCertificateSchema = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
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

// CertificateFromUpdateSchema is an adapter for certificate update schema
func CertificateFromUpdateSchema(p graphql.ResolveParams) models.Certificate {
	certificate := models.Certificate{}

	if qrCorner, ok := p.Args["qr_corner"]; ok {
		certificate.QRCorner = qrCorner.(string)
	}
	if qrScale, ok := p.Args["qr_scale"]; ok {
		certificate.QRScale = qrScale.(int)
	}
	if margin, ok := p.Args["margin"]; ok {
		certificate.Margin = margin.(int)
	}
	if nameOffsetTop, ok := p.Args["name_offset_top"]; ok {
		certificate.NameOffsetTop = nameOffsetTop.(int)
	}
	if fontSize, ok := p.Args["font_size"]; ok {
		certificate.FontSize = fontSize.(int)
	}
	if message, ok := p.Args["message"]; ok {
		certificate.Message = message.(string)
	}
	if active, ok := p.Args["active"]; ok {
		certificate.Active = active.(bool)
	}
	if courseID, ok := p.Args["course_id"]; ok {
		certificate.CourseID = courseID.(string)
	}
	if courseAuthorID, ok := p.Args["course_author_id"]; ok {
		certificate.CourseAuthorID = courseAuthorID.(string)
	}
	certificate.SetID(p.Args["id"].(string))
	return certificate
}
