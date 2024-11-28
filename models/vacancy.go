package models

import (
	proto "github.com/garden-raccoon/vacancy/protocols/vacancy"

	"github.com/gofrs/uuid"
)

type Vacancy struct {
	Name        string    `json:"name"`
	VacancyUUID          uuid.UUID `json:"vacancy_uuid"`
	EmployerUUID        uuid.UUID `json:"employer_uuid"`
	VacancyDesc       *VacancyDesc   `json:"vacancy_desc"`
}
type VacancyDesc struct {
	Responsibilities []string `json:"responsibilities"`
	Conditions []string `json:"conditions"`
	Requirements []string `json:"requirements"`
	Skills []string `json:"skills"`
	Description string `json:"description"`
	Address  string `json:"address"`
	Phone    int `json:"phone"`
	PriceFrom int `json:"price_from"`
	PriceTo int `json:"price_to"`
}
// Proto is
func (v Vacancy) Proto() *proto. {

	order := &proto.MealDb{
		Uuid:        mdb.Uuid.Bytes(),
		Name:        mdb.Name,
		Description: mdb.Description,
		Price:       float32(mdb.Price),
	}
	dummy := &proto.Dummy{Dummies: mdb.Dummy.Dummies, Trumpies: mdb.Dummy.Trumpies, Puppies: mdb.Dummy.Puppies, Address: mdb.Dummy.Address, Phone: int64(mdb.Dummy.Phone)}
	order.Dummy = dummy
	return order
}