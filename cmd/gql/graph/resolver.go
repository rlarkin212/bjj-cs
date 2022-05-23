package graph

import (
	"github.com/rlarkin212/bjj-cs/internal/service/search"
	"github.com/rlarkin212/bjj-cs/internal/service/submit"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	SearchService *search.SearchService
	SubmitService *submit.SubmitService
}
