package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	var em = EasyMongo{}
	em.Connect("mongodb://localhost:27017")

	//em.InsertOne(bson.D{{"name", "ppp"}, {"age", 18}})

	//em.InsertMany([]interface{}{
	//	bson.D{{"name", "y"}, {"age", 30}},
	//	bson.D{{"name", "k"}, {"age", 25}},
	//})

	//em.FindOne(bson.M{"name": "ppp"})

	//em.FindMany(bson.M{"age": 18})

	//em.UpdateOne(bson.M{"name": "pesar"}, bson.D{{"name", "pesarrr"}})
}

type PersonS struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
	Age  int                `json:"age,omitempty" bson:"age,omitempty"`
}

type EasyMongo struct {
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
}

func (em *EasyMongo) Connect(url string) {
	em.ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	em.client, _ = mongo.Connect(em.ctx, options.Client().ApplyURI(url))
	em.collection = em.client.Database("go-db").Collection("test")
	fmt.Println("connected")
}

func (em *EasyMongo) InsertOne(document bson.D) mongo.InsertOneResult {

	ObjectID, _ := em.collection.InsertOne(em.ctx, document)

	return *ObjectID
}

func (em *EasyMongo) InsertMany(documents []interface{}) {

	_, err := em.collection.InsertMany(em.ctx, documents)
	fmt.Println(err)

}

func (em *EasyMongo) FindOne(filter bson.M) PersonS {
	var res PersonS
	em.collection.FindOne(em.ctx, filter).Decode(&res)

	return res
}

func (em *EasyMongo) FindMany(filter bson.M) []PersonS {
	var res []PersonS
	cursor, _ := em.collection.Find(em.ctx, filter)
	err := cursor.All(context.TODO(), &res)
	if err != nil {
		panic(err)
	}

	return res
}

func (em *EasyMongo) DeleteOne(filter bson.M) *mongo.DeleteResult {
	res, _ := em.collection.DeleteOne(em.ctx, filter)

	return res
}

func (em *EasyMongo) DeleteMany(filter bson.M) *mongo.DeleteResult {
	res, _ := em.collection.DeleteMany(em.ctx, filter)

	return res
}

func (em *EasyMongo) UpdateOne(filter bson.M, document bson.D) *mongo.UpdateResult {
	res, _ := em.collection.UpdateOne(
		em.ctx,
		filter,
		bson.D{
			{"$set", document},
		},
	)

	return res
}

func (em *EasyMongo) UpdateMany(filter bson.M, document bson.D) *mongo.UpdateResult {
	res, _ := em.collection.UpdateMany(
		em.ctx,
		filter,
		bson.D{
			{"$set", document},
		},
	)

	return res
}

func (em *EasyMongo) Drop() bool {
	err := em.collection.Drop(em.ctx)

	if err == nil {
		return true
	} else {
		return false
	}
}
