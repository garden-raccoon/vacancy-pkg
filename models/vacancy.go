package models

import (
	proto "github.com/garden-raccoon/vacancy-pkg/protocols/vacancy"

	"github.com/gofrs/uuid"
)

type Vacancy struct {
	Name        string       `json:"name"`
	VacancyUUID uuid.UUID    `json:"vacancy_uuid"`
	UserUUID    uuid.UUID    `json:"user_uuid"`
	VacancyDesc *VacancyDesc `json:"vacancy_desc"`
}
type VacancyDesc struct {
	Responsibilities []string `json:"responsibilities"`
	Conditions       []string `json:"conditions"`
	Requirements     []string `json:"requirements"`
	Skills           []string `json:"skills"`
	Description      string   `json:"description"`
	Address          string   `json:"address"`
	Phone            int      `json:"phone"`
	PriceFrom        int      `json:"price_from"`
	PriceTo          int      `json:"price_to"`
}

// Proto is
func (v Vacancy) Proto() *proto.Vacancy {

	pb := &proto.Vacancy{
		VacancyUuid: v.VacancyUUID.Bytes(),
		UserUuid:    v.UserUUID.Bytes(),
		Name:        v.Name,
	}
	vacancyDesc := &proto.VacancyDesc{
		Conditions:       v.VacancyDesc.Conditions,
		Responsibilities: v.VacancyDesc.Responsibilities,
		Requirements:     v.VacancyDesc.Requirements,
		Skills:           v.VacancyDesc.Skills,
		Address:          v.VacancyDesc.Address,
		Phone:            int64(v.VacancyDesc.Phone),
		Description:      v.VacancyDesc.Description,
		PriceFrom:        int64(v.VacancyDesc.PriceFrom),
		PriceTo:          int64(v.VacancyDesc.PriceTo),
	}
	pb.VacancyDesc = vacancyDesc
	return pb
}
