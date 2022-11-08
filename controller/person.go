package controller

import (
	"cd_platform/ext"
	"cd_platform/pkg/person"
)

type PersonController struct {
	PersonService person.PersonService
}

func NewPersonController() *PersonController {
	psvc := person.NewService(ext.MiddleWare)
	return &PersonController{
		PersonService: psvc,
	}
}
