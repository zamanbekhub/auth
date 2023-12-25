package v1

import (
	"auth/internal/common/middleware"
	"auth/internal/schema"
	"auth/utils/token"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUser(v1 *gin.RouterGroup) {
	group := v1.Group("/user")
	group.GET("/me", middleware.GinErrorHandle(h.UserMe))
	group.POST("/login", middleware.GinErrorHandle(h.LoginUser))
	group.POST("/register", middleware.GinErrorHandle(h.RegisterUser))
}

// UserMe
// WhoAmi godoc
// @Summary Получить всех пользователей
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.User]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags user
// @Router /api/v1/user/me [get]
func (h *Handler) UserMe(c *gin.Context) error {
	userID, err := token.ExtractTokenID(c)
	if err != nil {
		return err
	}

	user, err := h.services.User.GetByID(c.Request.Context(), userID)
	if err != nil {
		return err
	}

	return schema.Respond(user, c)
}

// LoginUser
// WhoAmi godoc
// @Summary Получить всех пользователей
// @Accept json
// @Produce json
// @Param data body schema.LoginInput true "Логин пользователя"
// @Success 200 {object} schema.Response[[]model.User]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags user
// @Router /api/v1/user/login [post]
func (h *Handler) LoginUser(c *gin.Context) error {
	var data schema.LoginInput
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	users, err := h.services.User.Login(c.Request.Context(), data)
	if err != nil {
		return err
	}
	return schema.Respond(users, c)
}

// RegisterUser
// WhoAmi godoc
// @Summary Создание пользователыя
// @Accept json
// @Produce json
// @Param data body schema.UserCreate true "Создание пользователыя"
// @Success 200 {object} schema.Response[model.User]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags user
// @Router /api/v1/user/register [post]
func (h *Handler) RegisterUser(c *gin.Context) error {
	var data schema.UserCreate
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	user, err := h.services.User.Register(c.Request.Context(), data)
	if err != nil {
		return err
	}
	return schema.Respond(user, c)
}
