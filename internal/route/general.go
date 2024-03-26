package route

import (
	"github.com/ali-shokoohi/notes/internal/handler"
	"github.com/gin-gonic/gin"
)

func LoadGeneralRoutes(group *gin.RouterGroup, gh handler.GeneralHandler) {
	group.GET("", gh.Home)
}
