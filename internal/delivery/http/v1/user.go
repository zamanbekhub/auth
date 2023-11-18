package v1

import (
	"github.com/gin-gonic/gin"
	"service/internal/common/middleware"
	"service/internal/schema"
)

func (h *Handler) initUser(v1 *gin.RouterGroup) {
	group := v1.Group("/user")
	group.GET("/all", middleware.GinErrorHandle(h.GetAllUser))
	group.POST("/", middleware.GinErrorHandle(h.CreateUser))
}

// GetAllUser
// WhoAmi godoc
// @Summary Получить всех пользователей
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.User]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags user
// @Router /api/v1/user/all [get]
func (h *Handler) GetAllUser(c *gin.Context) error {
	users, err := h.services.User.GetAll(c.Request.Context())
	if err != nil {
		return err
	}
	return schema.Respond(users, c)
}

// CreateUser
// WhoAmi godoc
// @Summary Создание пользователыя
// @Accept json
// @Produce json
// @Param data body schema.UserCreate true "Создание пользователыя"
// @Success 200 {object} schema.Response[model.User]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags user
// @Router /api/v1/user [post]
func (h *Handler) CreateUser(c *gin.Context) error {
	var data schema.UserCreate
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	user, err := h.services.User.Create(c.Request.Context(), data)
	if err != nil {
		return err
	}
	return schema.Respond(user, c)
}
