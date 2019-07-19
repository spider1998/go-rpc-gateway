package server

import (
	"context"
	"sdkeji/person/pkg/entity"
	pb "sdkeji/person/pkg/proto"
	"sdkeji/person/pkg/service"
)

type PersonsServer struct{}

func (p *PersonsServer) CreatePerson(ctx context.Context, req *pb.CreatePersonRequest) (*pb.Person, error) {
	cond := entity.CreatePersonRequest{
		Age:    int(req.GetAge()),
		Name:   req.GetName(),
		Gender: entity.GenderState(req.GetGender()),
	}
	person, err := service.Person.CreatePerson(cond)
	if err != nil {
		return nil, nil
	}
	res := &pb.Person{
		Id:      person.ID,
		Name:    person.Name,
		Gender:  int32(person.Gender),
		Age:     int32(person.Age),
		Success: true,
	}
	return res, nil
}
