package search

import (
	"log"

	"github.com/rlarkin212/bjj-cs/configs"
	"github.com/rlarkin212/bjj-cs/internal/entities/instructionals"
	"github.com/rlarkin212/bjj-cs/internal/repository"
	"github.com/rlarkin212/bjj-cs/internal/repository/mongo"
)

type SearchService struct {
	repo repository.Repo
}

func NewSearchService(config *configs.Config) *SearchService {
	repo, err := mongo.NewMongoRepo(config)
	if err != nil {
		log.Fatal(err)
	}

	return &SearchService{
		repo: repo,
	}
}

func (s *SearchService) Instructionals() ([]*instructionals.Instructional, error) {
	instructionsls, err := s.repo.Instructionals()

	return instructionsls, err
}

func (s *SearchService) Instructional(id string) (*instructionals.Instructional, error) {
	instructional, err := s.repo.Instructional(id)

	return instructional, err
}

func (s *SearchService) Count() (int, error) {
	count, err := s.repo.Count()

	return count, err
}
