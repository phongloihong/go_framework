package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/phongloihong/go_framework/db"
	"github.com/phongloihong/go_framework/db/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func init() {
	collection = db.Database.Collection("student")
}

func Fetch() ([]*types.Student, error) {
	var students []*types.Student
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var student types.Student
		err := cur.Decode(&student)
		if err != nil {
			return nil, err
		}

		students = append(students, &student)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(context.TODO())

	return students, nil
}

func Insert(studentRequest types.StudentAddReq) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return collection.InsertOne(ctx, bson.M{
		"first_name": studentRequest.FirstName,
		"last_name":  studentRequest.LastName,
		"class_name": studentRequest.ClassName,
	})
}

func GetById(studentId primitive.ObjectID) (types.Student, error) {
	var student types.Student
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"_id": studentId}).Decode(&student)

	return student, err
}

func UpdateById(studentId primitive.ObjectID, studentReq types.StudentUpdateReq) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return collection.UpdateOne(
		ctx,
		bson.M{"_id": studentId},
		bson.M{
			"$set": bson.M{
				"first_name": studentReq.FirstName,
				"last_name":  studentReq.LastName,
				"class_name": studentReq.ClassName,
			},
		})
}

func DeleteById(studentId primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return collection.DeleteOne(ctx, bson.M{"_id": studentId})
}

func Find(studentRequest types.StudentSearchRequest) ([]*types.Student, error) {
	filters, err := getFilters(studentRequest)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, filters)
	if err != nil {
		log.Fatal(err)
	}

	var students []*types.Student
	for cur.Next(context.TODO()) {
		var student types.Student
		err := cur.Decode(&student)
		if err != nil {
			log.Fatal(err)
		}

		students = append(students, &student)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return students, nil
}

func getFilters(studentRequest types.StudentSearchRequest) (bson.D, error) {
	var filters bson.D

	filters = append(filters, bson.E{"first_name", studentRequest.FirstName})

	if studentRequest.Id != "" {
		objectID, err := primitive.ObjectIDFromHex(studentRequest.Id)
		if err != nil {
			return nil, err
		}

		filters = append(filters, bson.E{"_id", objectID})
	}

	if studentRequest.LastName != "" {
		filters = append(filters, bson.E{"last_name", studentRequest.LastName})
	}

	if studentRequest.ClassName != "" {
		filters = append(filters, bson.E{"class_name", studentRequest.ClassName})
	}

	return filters, nil
}
