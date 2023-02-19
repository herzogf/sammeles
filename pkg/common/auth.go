package common

import "github.com/gin-gonic/gin"

// magic user id used for admin access
const USER_ID_ADMIN = "_admin_"

// magic user id used for internal access, i.e. one type/service accessing another service
const USER_ID_INTERNAL = "_internal_"

// function to get the HTTP header "X-User-Id" from the gin context
func GetUserIdFromGinContext(c *gin.Context) string {
	return c.GetHeader("X-User-Id")
}