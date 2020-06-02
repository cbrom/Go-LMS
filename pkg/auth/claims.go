package auth

import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// These are the expected values for Claims.Roles.
const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// Key is used to store/retrieve a Claims value from a context.Context
const Key ctxKey = 1

// Claims represents the authorization claims transmitted via a JWT
type Claims struct {
	jwt.StandardClaims
}

// ClaimPreferences defines preferences for the user
type ClaimPreferences struct {
	Timezone       string `json:"timezone"`
	DatetimeFormat string `json:"pref_datetime_format"`
	DateFormat     string `json:"pref_date_format"`
	TimeFormat     string `json:"pref_time_format"`
	tz             *time.Location
}

// NewClaims constructs a Claim value for the identified user.
/* The claims expire within a specifie duration of the provided time.
Additional fields of the Claims can be set after calling NewClaims is desired
*/
func NewClaims(userID string, now time.Time, expires time.Duration) Claims {
	c := Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userID,
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(expires).Unix(),
		},
	}

	return c
}

// NewClaimPreferences constructs ClaimPreferences for the user/account
func NewClaimPreferences(timezone *time.Location, datetimeFormat, dateFormat, timeFormat string) ClaimPreferences {
	p := ClaimPreferences{
		DatetimeFormat: datetimeFormat,
		DateFormat:     dateFormat,
		TimeFormat:     timeFormat,
	}

	if timezone != nil {
		p.Timezone = timezone.String()
	}

	return p
}

// Valid is called during the parsing of a token.
// func (c Claims) Valid() error {
// 	for _, r := range c.Roles {
// 		switch r {
// 		case RoleAdmin, RoleUser: // Role is valid.
// 		default:
// 			return fmt.Errorf("invalid role %q", r)
// 		}
// 	}
// 	if err := c.StandardClaims.Valid(); err != nil {
// 		return errors.Wrap(err, "validating standard claims")
// 	}
// 	return nil
// }

// HasAuth returns true if the user is authenticated
func (c Claims) HasAuth() bool {
	if c.Subject != "" {
		return true
	}
	return false
}

// HasRole returns true if the claims has atleast one of the provided roles.
// func (c Claims) HasRole(roles ...string) bool {
// 	for _, has := range c.Roles {
// 		for _, want := range roles {
// 			if has == want {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// TimeLocation returns the timezone used to format datetimes for the user.
func (c ClaimPreferences) TimeLocation() *time.Location {
	if c.tz == nil && c.Timezone != "" {
		c.tz, _ = time.LoadLocation(c.Timezone)
	}
	return c.tz
}

// TimeLocation returns the timezone used to format datetimes for user
// func (c Claims) TimeLocation() *time.Location {
// 	return c.Preferences.TimeLocation()
// }

// ClaimsFromContext loads the claims from context
func ClaimsFromContext(ctx context.Context) (Claims, error) {
	claims, ok := ctx.Value(Key).(Claims)
	if !ok {
		return Claims{}, errors.New("claims missing from context: HasRole called without/before Authenticate")
	}

	return claims, nil
}
