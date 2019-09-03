package infrastructure

import (
	"context"
	"github.com/laches1sm/demo-parrot-service/src/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type ParrotInfra interface {
	CreateParrot(*domain.Parrot) (*domain.Parrot, error)
	GetParrotByID(string) (*domain.Parrot, error)
}

type mongoRepo struct {
	Client     mongo.Client
	Collection mongo.Collection
}

// NewMongoRepo returhs an initialised instance of a mongoRepo with an established
// connection to a mongo db
func NewMongoRepo() *mongoRepo {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return nil
	}
	errPing := client.Ping(ctx, readpref.Primary())

	if errPing == nil {
		collection := client.Database("parrot").Collection("parrots")
		repo := &mongoRepo{
			Client:     *client,
			Collection: *collection,
		}
		return repo
	} else {
		return nil
	}
}

func (repo *mongoRepo) CreateParrot(parrot *domain.Parrot) (*domain.Parrot, error) {
	parrotResult, err := repo.Collection.InsertOne(context.Background(), parrot)
	if err != nil {
		return nil, err
	}
	return repo.GetParrot(domain.Parrot(parrotResult))

}

func (repo *mongoRepo) GetParrot(parrot *domain.Parrot) (*domain.Parrot, error) {
	result := &domain.Parrot{}
	filter := bson.D{{"id", parrot.ID}}
	err := repo.Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
