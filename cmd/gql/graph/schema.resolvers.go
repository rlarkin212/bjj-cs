package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rlarkin212/bjj-cs/cmd/gql/graph/generated"
	"github.com/rlarkin212/bjj-cs/internal/entities/instructionals"
)

func (r *mutationResolver) CreateInstructional(ctx context.Context, input instructionals.NewInstructional) (*instructionals.Instructional, error) {
	instructional, err := r.SubmitService.Submit(&input)

	return instructional, err
}

func (r *queryResolver) Instructionals(ctx context.Context) ([]*instructionals.Instructional, error) {
	instructionals, err := r.SearchService.Instructionals()

	return instructionals, err
}

func (r *queryResolver) Instructional(ctx context.Context, id string) (*instructionals.Instructional, error) {
	instructional, err := r.SearchService.Instructional(id)

	return instructional, err
}

func (r *queryResolver) Count(ctx context.Context) (int, error) {
	count, err := r.SearchService.Count()

	return count, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
