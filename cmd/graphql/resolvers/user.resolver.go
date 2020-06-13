package resolvers

import (
	"fmt"
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/pkg/auth"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"github.com/pborman/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GetUser returns user object for a graphql
func GetUser(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		usr := &models.User{}
		usr.SetID(idQuery)
		usr.FetchByID()
		return usr, nil
	}
	return nil, errors.New("User ID not Provided")
}

// GetTimeFromStamp changes timestamp string to  *time.Time
func GetTimeFromStamp(ts string) *time.Time {
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return nil
	}
	tm := time.Unix(i, 0)
	return &tm
}

// SignUp creates a new user
func SignUp(p graphql.ResolveParams) (interface{}, error) {
	passwordSalt := uuid.NewRandom().String()
	saltedPassword := p.Args["password"].(string) + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "Error generating password hash")
	}

	timezoneArg := p.Args["timezone"]
	var timeZone *time.Time
	switch timezoneArg.(type) {
	case string:
		timeZone = GetTimeFromStamp(timezoneArg.(string))
	case time.Time:
		timeZone = timezoneArg.(*time.Time)
	}

	usr := &models.User{
		Name:         p.Args["name"].(string),
		Email:        p.Args["email"].(string),
		PasswordSalt: passwordSalt,
		PasswordHash: passwordHash,
		TimeZone:     timeZone,
	}
	usr.Create()

	claims := auth.Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   usr.ID,
			Audience:  "",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * 20000).Unix(),
		},
	}

	token, _ := authenticator.GenerateToken(claims)
	fmt.Println("****", token, "******")
	return token, nil
}

// SignIn returns token for a given user
func SignIn(p graphql.ResolveParams) (interface{}, error) {
	foundUser := models.User{
		Email: p.Args["email"].(string),
	}
	foundUser.FetchByEmail()
	if foundUser.GetID() == "" {
		return nil, nil
	}

	saltedPassword := p.Args["password"].(string) + foundUser.PasswordSalt
	if err := bcrypt.CompareHashAndPassword(foundUser.PasswordHash, []byte(saltedPassword)); err != nil {
		err = errors.WithStack(errors.New("ErrAuthenticationFailure"))
		return "", err
	}

	claims := auth.Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   foundUser.ID,
			Audience:  "",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * 20000).Unix(),
		},
	}
	token, _ := authenticator.GenerateToken(claims)

	c, _ := authenticator.ParseClaims(token)
	fmt.Println(c)

	return token, nil
}
