package data

import (
	"context"
	"errors"
	"task-management-rest-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getTaskCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("taskdb").Collection("tasks")
}


func GetAllTasks(ctx context.Context, client *mongo.Client) ([]models.Task, error) {
	collection := getTaskCollection(client)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// GetTaskByID returns a task by its ID from MongoDB
func GetTaskByID(ctx context.Context, client *mongo.Client, id string) (models.Task, error) {
	collection := getTaskCollection(client)
	var task models.Task
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return models.Task{}, errors.New("task not found")
	}
	return task, err
}

// AddTask inserts a new task into MongoDB
func AddTask(ctx context.Context, client *mongo.Client, task models.Task) error {
	collection := getTaskCollection(client)
	_, err := collection.InsertOne(ctx, task)
	return err
}

// UpdateTask updates a task by ID in MongoDB
func UpdateTask(ctx context.Context, client *mongo.Client, id string, updatedTask models.Task) error {
	collection := getTaskCollection(client)
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"id": id},
		bson.M{"$set": updatedTask},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

// RemoveTask deletes a task by ID from MongoDB
func RemoveTask(ctx context.Context, client *mongo.Client, id string) error {
	collection := getTaskCollection(client)
	result, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}


