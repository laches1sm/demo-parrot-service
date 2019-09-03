package infrastructure

import(
    "context"
    "fmt"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/satori/go.uuid"

)

type ParrotInfra interface{
    CreateParrot(*domain.Parrot) (*domain.Parrot, error)
    GetParrotByID(string) (*domain.Parrot, error)
}

type mongoRepo struct{
    Client options.Client()
}

// NewMongoRepo returhs an initialised instance of a mongoRepo with an established 
// connection to a mongo db
func NewMongoRepo() *mongoRepo{
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

    if err != nil{
        fmt.Printf(err)
	return nil
    }
    err := client.Ping(ctx, readpref.Primary())

    if err == nil{
	repo := &mongoRepo{
	    Client: client,
	}
	return repo
    } else{
	fmt.Printf(err)
	return nil
    }
}

func (repo *mongoRepo) CreateParrot(*domain.Parrot) (*domain.Parrot, error){}

func (repo *mongoRepo) GetParrotByID(parrotID string) (*domain.Parrot, error){} 
