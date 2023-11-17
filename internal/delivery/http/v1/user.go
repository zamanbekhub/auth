package v1

import (
	"github.com/gin-gonic/gin"
	"service/internal/common/middleware"
	"service/internal/schema"
)

func (h *Handler) initUser(v1 *gin.RouterGroup) {
	group := v1.Group("/user")
	group.GET("/all", middleware.GinErrorHandle(h.GetAllUser))
}

// GetAllUser
// WhoAmi godoc
// @Summary Получить всех пользователей
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.User]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags user
// @Router /api/v1/user/all [get]
func (h *Handler) GetAllUser(c *gin.Context) error {
	result, err := h.services.User.GetAll(c.Request.Context())
	if err != nil {
		return err
	}
	return schema.Respond(result, c)
}
