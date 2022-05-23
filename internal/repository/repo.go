package repository

import (
	"github.com/rlarkin212/bjj-cs/internal/entities/instructionals"
)

type Repo interface {
	Instructionals() ([]*instructionals.Instructional, error)
	Instructional(id string) (*instructionals.Instructional, error)
	Count() (int, error)
	Submit(input *instructionals.Instructional) (*instructionals.Instructional, error)
}
