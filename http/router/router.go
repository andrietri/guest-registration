package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApiV1(app *gin.RouterGroup, db *gorm.DB) {
	// Dependency injection
	diGuestRegistration := InitGuestRegistration(db)

	apiv1 := app.Group("/v1")
	{
		registrationGroup := apiv1.Group("registration")
		{
			registrationGroup.GET("/list", diGuestRegistration.FindAll)
			registrationGroup.GET("/:id_card_number", diGuestRegistration.FindByIdCardNumber)
			registrationGroup.POST("/", diGuestRegistration.Register)
		}
	}
}
