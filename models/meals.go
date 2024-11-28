package models

import (
	"fmt"
	"github.com/gofrs/uuid"

	proto "github.com/garden-raccoon/orders-admin-package/protocols/orders-admin-package"
)

type MealsDb struct {
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Uuid        uuid.UUID `json:"uuid"`
	Description string    `json:"description"`
	Dummy       *Dummy    `json:"dummy"`
}
type Dummy struct {
	Dummies  []string
	Trumpies []string
	Puppies  []string
	Address  string
	Phone    int
}

// Proto is
func (mdb MealsDb) Proto() *proto.MealDb {

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

func MealDbFromProto(pb *proto.MealDb) *MealsDb {

	order := &MealsDb{
		Name:        pb.Name,
		Price:       float64(pb.Price),
		Description: pb.Description,
		Uuid:        uuid.FromBytesOrNil(pb.Uuid),
	}
	fmt.Println(" pb.Dummy are", pb.Dummy)
	dummy := &Dummy{Dummies: pb.Dummy.Dummies}
	order.Dummy = dummy
	return order
}

func MealsToProto(meals []*MealsDb) *proto.MealsDb {
	pb := &proto.MealsDb{}

	for _, b := range meals {
		pb.MealsDb = append(pb.MealsDb, b.Proto())
	}
	return pb
}

func MealsFromProto(pb *proto.MealsDb) (meals []*MealsDb) {
	for _, b := range pb.MealsDb {
		meals = append(meals, MealDbFromProto(b))
	}
	return
}
