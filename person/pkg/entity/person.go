package entity

import (
	"sdkeji/person/pkg/util"
)

type GenderState int8

const (
	GenderStateFemale GenderState = 1
	GenderStateMale   GenderState = 2
)

type Person struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Gender     GenderState `json:"gender"`
	Age        int         `json:"age"`
	CreateTime util.Time   `json:"create_time" xorm:"created"`
	UpdateTime util.Time   `json:"update_time" xorm:"updated"`
}

type CreatePersonRequest struct {
	Name   string      `json:"name"`
	Gender GenderState `json:"gender"`
	Age    int         `json:"age"`
}
