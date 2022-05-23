package submit

import (
	"log"
	"time"

	"github.com/rlarkin212/bjj-cs/configs"
	"github.com/rlarkin212/bjj-cs/internal/entities/instructionals"
	"github.com/rlarkin212/bjj-cs/internal/repository"
	"github.com/rlarkin212/bjj-cs/internal/repository/mongo"
	"github.com/rlarkin212/bjj-cs/util/searchable"
	"github.com/rlarkin212/bjj-cs/util/spliturl"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubmitService struct {
	repo repository.Repo
}

func NewSubmitService(config *configs.Config) *SubmitService {
	repo, err := mongo.NewMongoRepo(config)
	if err != nil {
		log.Fatal(err)
	}

	return &SubmitService{
		repo: repo,
	}
}

func (s *SubmitService) Submit(input *instructionals.NewInstructional) (*instructionals.Instructional, error) {
	newInstructional := &instructionals.Instructional{
		ObjectId:  primitive.NewObjectID(),
		Title:     input.Title,
		Presenter: input.Presenter,
		Part:      input.Part,
		CreatedAt: int64(time.Now().UTC().Unix()),
	}

	newInstructional.Id = newInstructional.ObjectId.Hex()

	download, watch, err := spliturl.SplitUrl(input.URL)
	if err != nil {
		return nil, err
	}

	newInstructional.DownloadUrl = download
	newInstructional.WatchUrl = watch
	newInstructional.SearchTitle = searchable.GenerateSearchable(newInstructional.Title)

	instructional, err := s.repo.Submit(newInstructional)

	return instructional, err
}
