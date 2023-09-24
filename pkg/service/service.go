package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *serObj) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
