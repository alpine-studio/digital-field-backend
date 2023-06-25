package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const SESSION_NAME = "session"

func (s *server) createSession(c *gin.Context) {
	role := c.PostForm("role")

	session, err := s.store.Session().Create(role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	c.SetCookie(SESSION_NAME, session.Token, 9999, "", "", false, false)
	c.JSON(http.StatusCreated, session)
}

func (s *server) logout(c *gin.Context) {
	c.SetCookie(SESSION_NAME, "", -1, "", "", false, false)
	c.Status(http.StatusOK)
}

func (s *server) whoAmI(c *gin.Context) {
	t := c.GetHeader(SESSION_NAME)

	session, err := s.store.Session().Find(t)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)

		return
	}

	c.JSON(http.StatusOK, session.Role)
}
