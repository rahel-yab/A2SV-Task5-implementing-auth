package data

import (
	"context"
	"task-management-rest-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddUser(ctx context.Context, client *mongo.Client, user models.User) error {
	collection := client.Database("task_manager").Collection("users")
	_, err := collection.InsertOne(ctx, user)
	return err
}

func GetUserByEmail(ctx context.Context, client *mongo.Client, email string) (*models.User, error) {
	collection := client.Database("task_manager").Collection("users")
	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(ctx context.Context, client *mongo.Client, username string) (*models.User, error) {
	collection := client.Database("task_manager").Collection("users")
	var user models.User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func IsUsersCollectionEmpty(ctx context.Context, client *mongo.Client) (bool, error) {
	collection := client.Database("task_manager").Collection("users")
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func UserExistsByEmail(ctx context.Context, client *mongo.Client, email string) (bool, error) {
	collection := client.Database("task_manager").Collection("users")
	count, err := collection.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func UserExistsByUsername(ctx context.Context, client *mongo.Client, username string) (bool, error) {
	collection := client.Database("task_manager").Collection("users")
	count, err := collection.CountDocuments(ctx, bson.M{"username": username})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func PromoteUserToAdmin(ctx context.Context, client *mongo.Client, identifier string) error {
	collection := client.Database("task_manager").Collection("users")
	filter := bson.M{"$or": []bson.M{
		{"username": identifier},
		{"email": identifier},
	}}
	update := bson.M{"$set": bson.M{"role": "admin"}}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}