package context

import (
	"net/http"

	"github.com/go-gosh/tomato/app/ent"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	loginUser *ent.User
}

func (c *Context) SetLoginUser(loginUser *ent.User) {
	c.loginUser = loginUser
}

func (c Context) LoginUser() (*ent.User, error) {
	return c.loginUser, nil
}

func (c Context) Response(code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
	})
}
