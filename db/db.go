package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	TIMEOUT = 10
)

type DatabaseService struct {
	*mongo.Client
	DbName string
}

func (service *DatabaseService) GetCollection(collection string) *mongo.Collection {
	return service.Client.Database(service.DbName).Collection(collection)
}

func (service *DatabaseService) FilterOne(collection string, filter bson.M) *mongo.SingleResult {
	log.Println("Trying match by filter")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	findOpts := options.FindOneOptions{}
	findOpts.SetMaxTime(time.Second * (TIMEOUT / 5))
	result := service.GetCollection(collection).FindOne(ctx, filter, &findOpts)
	return result
}

func (service *DatabaseService) FilterMany(collection string, filter bson.M) *mongo.Cursor {
	log.Println("Trying match by filter")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	findOpts := options.FindOptions{}
	findOpts.SetMaxTime(time.Second * (TIMEOUT / 5))
	result, err := service.GetCollection(collection).Find(ctx, filter, &findOpts)
	checkError(err)
	return result
}

func (service *DatabaseService) FieldMatchesString(collection string, field string, matches string) *mongo.SingleResult {
	log.Println("Trying to filter by field " + field)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	filter := bson.M{}
	if len(field) > 0 && len(matches) > 0 {
		filter = bson.M{ field: matches }
	}
	findOpts := options.FindOneOptions{}
	findOpts.SetMaxTime(time.Second * (TIMEOUT / 5))
	result := service.GetCollection(collection).FindOne(ctx, filter, &findOpts)
	return result
}

func (service *DatabaseService) FindByFieldMatchesString(collection string, field string, matches string) *mongo.Cursor {
	log.Println("Trying to find all by field " + field)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	filter := bson.M{}
	if len(field) > 0 && len(matches) > 0 {
		filter = bson.M{ field: matches }
	}
	findOpts := options.FindOptions{}
	findOpts.SetMaxTime(time.Second * (TIMEOUT / 5))
	result, err := service.GetCollection(collection).Find(ctx, filter, &findOpts)
	checkError(err)
	return result
}

func (service *DatabaseService) InsertOne(collection string, payload interface{}) *mongo.InsertOneResult {
	log.Println("Inserting payload into collection " + collection)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	result, err := service.GetCollection(collection).InsertOne(ctx, payload)
	checkError(err)
	return result
}

func (service *DatabaseService) FindAll(collection string) *mongo.Cursor {
	log.Println("Finding all from collection: " + collection)
	return service.FindByFieldMatchesString(collection, "", "")
}

func (service *DatabaseService) GetById(collection string, id string) *mongo.SingleResult {
	log.Println("Getting by id from " + collection + " with _id " + id)
	return service.FieldMatchesString(collection, "_id", id)
}

func (service *DatabaseService) DeleteOneByFieldMatches(collection string, field string, matches string) *mongo.DeleteResult {
	log.Println("Deleting one by " + field + " from " + collection + " with " + field + "=" + matches)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	filter := bson.M{}
	if len(field) > 0 && len(matches) > 0 {
		filter = bson.M{ field: matches }
	}
	result, err := service.GetCollection(collection).DeleteOne(ctx, filter)
	checkError(err)
	return result
}

func (service *DatabaseService) DeleteManyByFieldMatches(collection string, field string, matches string) *mongo.DeleteResult {
	log.Println("Deleting many by " + field + " from " + collection + " with " + field + "=" + matches)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	filter := bson.M{}
	if len(field) > 0 && len(matches) > 0 {
		filter = bson.M{ field: matches }
	}
	result, err := service.GetCollection(collection).DeleteMany(ctx, filter)
	checkError(err)
	return result
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
