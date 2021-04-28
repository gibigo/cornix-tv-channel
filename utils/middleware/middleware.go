package middleware

import (
	"errors"
	"log"

	"github.com/gibigo/cornix-tv-channel/app/dal"
	"github.com/gibigo/cornix-tv-channel/utils/password"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"gorm.io/gorm"
)

var BasicAuth = basicauth.New(basicauth.Config{
	Users:      map[string]string{},
	Realm:      "Restricted",
	Authorizer: isAuthorized,
})

func isAuthorized(username, pwd string) bool {
	var user dal.User
	result := dal.FindUserByName(&user, username)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) && result.Error != nil {
		log.Printf("[middleware][authorize] error while reading user from database: %s", result.Error)
		return false
	}
	if err := password.Verify(user.Password, pwd); err != nil {
		return false
	} else {
		return true
	}
}
