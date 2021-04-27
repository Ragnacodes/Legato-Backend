package middleware

import (
	"legato_server/authenticate"
	"legato_server/domain"

	"github.com/gin-gonic/gin"
)

const Authorization = "Authorization"
const UserKey = "UserKey"

func AuthMiddleware(u *domain.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(Authorization)

		// Allow unauthenticated users in
		if token == "" {
			c.Set(UserKey, nil)
			c.Next()
			return
		}

		// Check validation jwt token
		claim, err := authenticate.CheckToken(token)
		if err != nil {
			c.Set(UserKey, nil)
			c.Next()
			return
		}

		// Get user and check if the user exists in db
		user, err := (*u).GetUserByUsername(claim.Username)
		if err != nil {
			c.Set(UserKey, nil)
			c.Next()
			return
		}

		// put user in context
		c.Set(UserKey, &user)
		c.Next()
	}
}
