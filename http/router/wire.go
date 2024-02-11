//go:build wireinject
// +build wireinject

package router

import (
	gr "github.com/andrietri/guest-registration/api/v1/registration"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitGuestRegistration(db *gorm.DB) gr.GuestRegistrationController {
	wire.Build(
		gr.ProviderMeetingRoomRepository,
		gr.ProviderMeetingRoomService,
		gr.ProviderMeetingRoomController,
	)

	return gr.RegistrationController{}
}
