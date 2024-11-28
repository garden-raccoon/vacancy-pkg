package models

import (
	proto "github.com/garden-raccoon/vacancy-pkg/protocols/vacancy"

	"github.com/gofrs/uuid"
)

type Vacancy struct {
	Name         string       `json:"name"`
	VacancyUUID  uuid.UUID    `json:"vacancy_uuid"`
	EmployerUUID uuid.UUID    `json:"employer_uuid"`
	VacancyDesc  *VacancyDesc `json:"vacancy_desc"`
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
		VacancyUuid:  v.VacancyUUID.Bytes(),
		EmployerUuid: v.EmployerUUID.Bytes(),
		Name:         v.Name,
	}
	vacancyDesc := &proto.VacancyDesc{
		Conditions:       pb.VacancyDesc.Conditions,
		Responsibilities: pb.VacancyDesc.Responsibilities,
		Requirements:     pb.VacancyDesc.Requirements,
		Skills:           pb.VacancyDesc.Skills,
		Address:          pb.VacancyDesc.Address,
		Phone:            pb.VacancyDesc.Phone,
		Description:      pb.VacancyDesc.Description,
		PriceFrom:        pb.VacancyDesc.PriceFrom,
		PriceTo:          pb.VacancyDesc.PriceTo,
	}
	pb.VacancyDesc = vacancyDesc
	return pb
}
