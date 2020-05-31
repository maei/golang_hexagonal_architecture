package mongodb

import (
	"context"
	"fmt"
	"github.com/maei/golang_hexagonal_architecture/src/service"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	collection = "sum"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoSumRepository(mongoURL, mongoDB string, mongoTimeout int) (service.SumRepositoryInterface, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}
	repo.client = client
	return repo, nil
}

func (m *mongoRepository) Store(req *service.SumResult) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	collection := m.client.Database(m.database).Collection(collection)
	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		return errors.New("error while writing to database")
	}
	return nil
}

func (m *mongoRepository) Find(code string) (*service.SumResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	sumResult := &service.SumResult{}
	collection := m.client.Database(m.database).Collection(collection)
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&sumResult)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(fmt.Sprintf("error finding collection for %v", code))
		}
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}
	return sumResult, nil
}

func (m *mongoRepository) Disconnect(ctx context.Context) {
	m.client.Disconnect(ctx)
}
