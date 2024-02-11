package registration

import "github.com/andrietri/guest-registration/database/model"

type ResponseList struct {
	Items      []model.Guest `json:"items"`
	TotalCount int           `json:"total_count"`
}

type ResponseDetail struct {
	Item       model.Guest `json:"item"`
	TotalCount int         `json:"total_count"`
}

type RequestRegister struct {
	Name         string `json:"name"`
	IdCardNumber string `json:"id_card_number"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
}
