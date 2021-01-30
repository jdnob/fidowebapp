package entity

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User exported
type User struct {
	ID              primitive.ObjectID `bson: "_id, omitempty"`
	UserName        string             `bson: "userName, omitempty"`
	UserDisplayName string             `bson: "userDisplayName, omitempty"`
	UserUUID        string             `bson: "userUUID, omitempty"`
	UserHandle      string             `bson: "userHandle, omitempty"`
}

const (
	UserCollection = "user"
)

// Find All entries of this collection
func FindAllUsers(ctx context.Context, database *mongo.Database) {
	log.Debug("Log debug!")
	// m := make(map[string]User)
	cursor, err := database.Collection(UserCollection).Find(ctx, bson.M{})
	if err != nil {
		// error
		log.Fatal(err)
	}

	var users []User

	if err := cursor.All(ctx, &users); err != nil {
		// error
		log.Error(err)
	}
	defer cursor.Close(ctx)
	log.Info(users)

}

func FindUser(ctx context.Context, database *mongo.Database, userUUID string) User {
	filter := bson.D{{"userUUID", userUUID}}
	collection := database.Collection("user")

	var result User
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.WithField("user", result).Info()
	return result
}
