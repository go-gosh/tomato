package handler

import (
	"github.com/go-gosh/tomato/app/context"
	"github.com/go-gosh/tomato/app/ent"

	"github.com/gin-gonic/gin"
)

func ginAdapter(fn func(*context.Context) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := &context.Context{
			Context: ctx,
		}
		c.SetLoginUser(&ent.User{
			ID:       0,
			Username: "onlyOne",
			Enabled:  true,
		})
		err := fn(c)
		if err != nil {
			ctx.AbortWithStatusJSON(200, gin.H{
				"code":    500,
				"message": err.Error(),
				"data":    nil,
			})
		}
	}
}
