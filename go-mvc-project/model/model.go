package model

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client     *mongo.Client
	colPersons *mongo.Collection
}

type Person struct {
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
	Pnum string `json:"pnum" bson:"pnum" validate:"unique"`
}

// Failure Response Model
type Failure struct {
	Message string `json:"message" example:"Bad Request"`
}

// Failure Response Model
type Success struct {
	Message string `json:"message" example:"Success"`
}

func NewModel() (*Model, error) {
	r := &Model{}

	var err error
	mgUrl := "mongodb://127.0.0.1:27017"
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("go-ready")
		r.colPersons = db.Collection("tPerson")
		fmt.Println("Mongo DB Successful Connected")
	}

	return r, nil
}

/* 모든 유저 조회 */
func (p *Model) GetAllPersons() ([]Person, error) {
	var pers []Person
	filter := bson.D{}
	cursor, err := p.colPersons.Find(context.TODO(), filter)
	if err != nil {
		defer cursor.Close(context.TODO())
		fmt.Println("Finding all documents Error :", err)
		return pers, err
	} else {
		for cursor.Next(context.TODO()) {
			var result Person
			if err := cursor.Decode(&result); err != nil {
				fmt.Println("cursor.Next() error : ", err)
				os.Exit(1)
			} else {
				pers = append(pers, result)
			}

		}
	}
	return pers, err
}

func (p *Model) GetAllPersonsByName(name string) ([]Person, error) {
	var pers []Person
	filter := bson.M{"name": name} //bson.M 과 bson.D의 차이는 순서보장
	cursor, err := p.colPersons.Find(context.TODO(), filter)

	if err != nil {
		defer cursor.Close(context.TODO())
		fmt.Println("Finding all documents Error :", err)
		return pers, err
	} else {
		for cursor.Next(context.TODO()) {
			var result Person
			if err := cursor.Decode(&result); err != nil {
				fmt.Println("cursor.Next() error : ", err)
				os.Exit(1)
			} else {
				pers = append(pers, result)
			}

		}
	}
	return pers, nil
}

func (p *Model) GetPersonByPnum(pnum string) Person {
	var person Person
	filter := bson.M{"pnum": pnum}
	result := p.colPersons.FindOne(context.TODO(), filter)
	result.Decode(&person)

	return person
}

func (p *Model) CreatePerson(person Person) (*mongo.InsertOneResult, error) {
	/* InsertOne 반환값 ( ObjectId , error )*/
	ObjectId, err := p.colPersons.InsertOne(context.TODO(), person)
	return ObjectId, err
}

func (p *Model) DeletePerson(pnum string) (*mongo.DeleteResult, error) {
	filter := bson.M{"pnum": pnum}
	/* InsertOne 반환값 ( deleteCount , error )*/
	deleteCount, err := p.colPersons.DeleteOne(context.TODO(), filter)
	return deleteCount, err
}

func (p *Model) UpdateAgeByPnum(pnum string, age int) (*mongo.UpdateResult, error) {
	/* 조건 */
	filter := bson.M{"pnum": pnum}
	/* 수정 필드 */
	target := bson.D{
		{Key: "$set", Value: bson.D{{Key: "age", Value: age}}},
	}
	result, err := p.colPersons.UpdateOne(context.TODO(), filter, target)
	return result, err
}
