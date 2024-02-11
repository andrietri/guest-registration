package registration

import (
	"github.com/andrietri/guest-registration/database/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RegistrationRepository struct {
	db *gorm.DB
}

func ProviderRegistrationRepository(db *gorm.DB) RegistrationRepository {
	return RegistrationRepository{db: db}
}

func (repo *RegistrationRepository) FindAll() (response ResponseList) {
	var (
		guest      []model.Guest
		totalCount int64
	)

	qx := repo.db.Model(&model.Guest{}).
		Count(&totalCount).Find(&guest)

	response.Items = guest
	response.TotalCount = int(totalCount)

	if qx.Error != nil {
		logrus.Warn("FindAll: ", qx.Error)
	}
	return
}

func (repo *RegistrationRepository) FindByIdCardNumber(idCardNumber string) (response ResponseDetail) {
	var (
		guest      model.Guest
		totalCount int64
	)

	qx := repo.db.Model(&model.Guest{}).
		Where("id_card_number = ?", idCardNumber).
		Count(&totalCount).
		First(&guest)

	response.Item = guest
	response.TotalCount = int(totalCount)

	if qx.Error != nil {
		logrus.Warn("FindByIdCardNumber: ", qx.Error)
	}

	return
}

func (repo *RegistrationRepository) Register(request model.Guest) (err error) {
	if err := repo.db.Create(&request).Error; err != nil {
		return err
	}

	return

}
