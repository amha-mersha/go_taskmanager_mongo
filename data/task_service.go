package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/amha-mersha/go_taskmanager_mongo/models"
	"github.com/amha-mersha/go_taskmanager_mongo/route"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func int64ToObjID(ID int64) (primitive.ObjectID, error) {
	hexaString := fmt.Sprintf("%x", ID)
	if len(hexaString) < 24 {
		hexaString = fmt.Sprintf("%024s", hexaString)
	}
	return primitive.ObjectIDFromHex(hexaString)
}

type TaskError struct {
	message string
}

func (err TaskError) Error() string {
	return err.message
}

const (
	IDNotFound        = "no item found with the specified id"
	TaskAlreadyExists = "the task already exists in the database"
	MalformedJSON     = "sent a malfomed json"
	MismatchedFormat  = "the task have a mismatched stucture"
	MissingRequireds  = "there are some missing required feilds"
	MalformedID       = "the id sent is malformed"
)

func GetTasks() ([]models.Task, error) {
	var result []models.Task
	curr, err := route.Collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return []models.Task{}, err
	}
	err = curr.All(context.TODO(), &result)
	if err != nil {
		return []models.Task{}, err
	}
	return result, nil
}

func GetTaskByID(taskID int64) (models.Task, error) {
	var result models.Task
	ID, err := int64ToObjID(taskID)
	if err != nil {
		return models.Task{}, fmt.Errorf(MalformedID)
	}
	err = route.Collection.FindOne(context.TODO(), bson.D{{"_id", ID}}).Decode(&result)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return models.Task{}, fmt.Errorf(IDNotFound)
	} else if err != nil {
		return models.Task{}, err
	}
	return result, nil
}

func UpdateTask(taskID int64, updatedTask models.Task) (models.Task, error) {
	var result models.Task
	ID, err := int64ToObjID(taskID)
	if err != nil {
		return models.Task{}, fmt.Errorf(MalformedID)
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.D{{"_id", ID}}

	err = route.Collection.FindOneAndUpdate(context.TODO(), filter, updatedTask, opts).Decode(&result)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return models.Task{}, fmt.Errorf(IDNotFound)
	} else if err != nil {
		return models.Task{}, err
	}
	return result, nil
}

func DeleteTask(taskID int64) (models.Task, error) {
	var result models.Task
	ID, err := int64ToObjID(taskID)
	if err != nil {
		return models.Task{}, fmt.Errorf(MalformedID)
	}
	filter := bson.D{{"_id", ID}}
	err = route.Collection.FindOneAndDelete(context.TODO(), filter).Decode(&result)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return models.Task{}, fmt.Errorf(IDNotFound)
	} else if err != nil {
		return models.Task{}, err
	}
	return result, nil
}

func PostTask(newTask models.Task) (*mongo.InsertOneResult, error) {
	result, err := route.Collection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return &mongo.InsertOneResult{}, fmt.Errorf(MalformedID)
	}
	return result, err
}
