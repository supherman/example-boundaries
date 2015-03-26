package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/nicholasf/example/boundaries"
)

func Resource(boundary boundaries.Houses) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Params.ByName("address")

		collections, _ := boundary.ByAddress(name)
		ctx.JSON(200, collections[0])
	}
}
