package registration

import "github.com/andrietri/guest-registration/database/model"

type RegistrationService struct {
	repository RegistrationRepository
}

func ProviderRegistrationService(repo RegistrationRepository) RegistrationService {
	return RegistrationService{
		repository: repo,
	}
}

func (service *RegistrationService) FindAll() (response ResponseList) {
	return service.repository.FindAll()
}

func (service *RegistrationService) FindByIdCardNumber(idCardNumber string) (response ResponseDetail) {
	return service.repository.FindByIdCardNumber(idCardNumber)
}

func (service *RegistrationService) Register(request model.Guest) (message string, err error) {
	message = "registration success"

	// check data by id card number if exist
	data := service.repository.FindByIdCardNumber(request.IdCardNumber)
	if data.Item.ID != 0 {
		message = "guest already registered"
		return
	}

	// register data
	err = service.repository.Register(request)
	if err != nil {
		message = "registration failed"
		return
	}

	return
}
