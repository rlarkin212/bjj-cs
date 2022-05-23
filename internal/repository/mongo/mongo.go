package mongo

import (
	"context"
	"time"

	"github.com/rlarkin212/bjj-cs/configs"
	"github.com/rlarkin212/bjj-cs/internal/entities/instructionals"
	"github.com/rlarkin212/bjj-cs/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepo struct {
	config *configs.Config
	client *mongo.Client
}

func NewMongoRepo(config *configs.Config) (repository.Repo, error) {
	client, err := newClient(config)
	if err != nil {
		return nil, err
	}

	mongoRepo := &mongoRepo{
		config: config,
		client: client,
	}

	return mongoRepo, nil
}

func (r *mongoRepo) Instructionals() ([]*instructionals.Instructional, error) {
	ctx, cancel := r.context()
	defer cancel()

	res := []*instructionals.Instructional{}

	coll := r.collection()
	filter := bson.M{}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		instructional := &instructionals.Instructional{}

		err := cursor.Decode(instructional)
		if err != nil {
			return nil, err
		}

		res = append(res, instructional)
	}

	return res, nil
}

func (r *mongoRepo) Instructional(id string) (*instructionals.Instructional, error) {
	ctx, cancel := r.context()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	coll := r.collection()
	filter := bson.M{
		"_id": objectId,
	}

	instructional := &instructionals.Instructional{}

	res := coll.FindOne(ctx, filter)
	err = res.Decode(instructional)
	if err != nil {
		return nil, err
	}

	return instructional, err
}

func (r *mongoRepo) Count() (int, error) {
	ctx, cancel := r.context()
	defer cancel()

	coll := r.collection()
	filter := bson.M{}

	count, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return -1, err
	}

	return int(count), nil
}

func (r *mongoRepo) Submit(input *instructionals.Instructional) (*instructionals.Instructional, error) {
	ctx, cancel := r.context()
	defer cancel()

	coll := r.collection()

	_, err := coll.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func newClient(config *configs.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Mongo.Timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo.URL))
	if err != nil {
		return nil, err
	}

	//check if connected
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (r *mongoRepo) context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(r.config.Mongo.Timeout)*time.Second)
}

func (r *mongoRepo) collection() *mongo.Collection {
	return r.client.Database(r.config.Mongo.DB).Collection(r.config.Mongo.Collection)
}
