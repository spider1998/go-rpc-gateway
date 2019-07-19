package service

import (
	v "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"sdkeji/person/pkg/app"
	"sdkeji/person/pkg/entity"
)

var Person PersonService

type PersonService struct{}

func (m *PersonService) CreatePerson(req entity.CreatePersonRequest) (person entity.Person, err error) {
	err = v.ValidateStruct(&req,
		v.Field(&req.Name, v.Required),
		v.Field(&req.Gender, v.Required),
		v.Field(&req.Age, v.Required),
	)
	if err != nil {
		return
	}
	exist, err := app.DB.Where("name = ?", req.Name).Get(&person)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if exist {
		app.Logger.Info("person already exist !", "person:", req.Name)
		return
	}
	person.ID = xid.New().String()
	person.Name = req.Name
	person.Age = req.Age
	person.Gender = req.Gender

	_, err = app.DB.Insert(&person)
	if err != nil {
		return
	}
	app.Logger.Info("create person success.", "person:", person.ID)
	return
}
