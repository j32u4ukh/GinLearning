package middleware

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const userKey = "session_id"

// Use cookies to store session_id
func SetSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(userKey))
	return sessions.Sessions("mysession", store)
}

// User Auth Session Middleware
func AuthSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionId := session.Get(userKey)
		if sessionId == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		c.Next()
	}
}

// Save Session for user
func SaveSession(c *gin.Context, userId int) {
	session := sessions.Default(c)
	session.Set(userKey, userId)
	log.Printf("Saved session: %v -> %v", userKey, userId)
	log.Printf("session: %v", session)
	session.Save()
}

// Clear Session for user
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// Get Session for user
func GetSessionId(c *gin.Context) int {
	session := sessions.Default(c)
	sessionId := session.Get(userKey)
	if sessionId == nil {
		return 0
	}
	return sessionId.(int)
}

// Check Session for user
func CheckSession(c *gin.Context) bool {
	session := sessions.Default(c)
	sessionId := session.Get(userKey)
	return sessionId != nil
}
