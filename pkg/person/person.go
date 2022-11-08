package person

import (
	"cd_platform/mid"
)

type PersonService interface {
}

type Service struct {
	Mid *mid.Middle
}

func NewService(mid *mid.Middle) *Service {
	return &Service{
		Mid: mid,
	}
}
