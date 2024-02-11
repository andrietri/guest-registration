package registration

import (
	"fmt"

	"github.com/andrietri/guest-registration/database/model"
	"github.com/andrietri/guest-registration/http/httpresponse"
	"github.com/andrietri/guest-registration/utils/constants"
	"github.com/erajayatech/go-helper"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type RegistrationController struct {
	service RegistrationService
}

func ProviderRegistrationController(rs RegistrationService) RegistrationController {
	return RegistrationController{
		service: rs,
	}
}

func (ctrl *RegistrationController) FindAll(c *gin.Context) {
	response := ctrl.service.FindAll()
	if len(response.Items) == 0 {
		httpresponse.NotFound(c)
		return
	}

	httpresponse.WithData(c, response)
}

func (ctrl *RegistrationController) FindByIdCardNumber(c *gin.Context) {
	idCardNumber := helper.ExpectedString(c.Param("id_card_number"))
	response := ctrl.service.FindByIdCardNumber(idCardNumber)
	if response.Item.ID == 0 {
		httpresponse.NotFound(c)
		return
	}

	httpresponse.WithData(c, response)
}

func (ctrl *RegistrationController) Register(c *gin.Context) {
	var (
		request model.Guest
	)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Warn("cannot should bind json : ", err.Error())
		httpresponse.BadRequest(c, "payload attribute not suitable")
		return
	}

	message, err := ctrl.service.Register(request)
	if err != nil {
		httpresponse.WithDataCustomResponseStatus(c, constants.HTTP_RESPONSE_BAD_REQUEST, interface{}(nil), fmt.Sprintf("%v", err))
		return
	}

	httpresponse.WithDataCustomResponseStatus(c, constants.HTTP_RESPONSE_OK, message, interface{}(nil))
}
