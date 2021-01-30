package database

import (
	"context"
	"fidowebapp/config"
	"fmt"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	ContextKeyDatabase string = "ContextKey.Database"
)

// User exported
type User struct {
	ID              primitive.ObjectID `bson: "_id, omitempty"`
	UserName        string             `bson: "userName, omitempty"`
	UserDisplayName string             `bson: "userDisplayName, omitempty"`
	UserUUID        string             `bson: "userUUID, omitempty"`
	UserHandle      string             `bson: "userHandle, omitempty"`
}

func ContextWithDatabase(ctx context.Context, dbConfig config.DatabaseConfiguration) context.Context {
	return context.WithValue(ctx, ContextKeyDatabase, createDatabase(ctx, dbConfig))
}

func DatabaseFromContext(ctx context.Context) *mongo.Database {
	db, exists := ctx.Value(ContextKeyDatabase).(*mongo.Database)
	if db == nil || !exists {
		log.Panic("Database not found in context")
	}
	return db
}

func createDatabase(ctx context.Context, dbConfig config.DatabaseConfiguration) *mongo.Database {
	connectURL := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBURL, dbConfig.DBPort, dbConfig.DBName)
	clientOptions := options.Client().ApplyURI(connectURL)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(dbConfig.DBName)
}

// Find All entries of this collection
func FindAll(ctx context.Context, database *mongo.Database, collectionName string) {
	log.Debug("Log debug!")
	// m := make(map[string]User)
	cursor, err := database.Collection(collectionName).Find(ctx, bson.M{})
	if err != nil {
		// error
		log.Fatal(err)
	}

	var users []User
	// for cursor.Next(ctx) {
	// 	log.Info(cursor.Current)
	// 	u := User{}
	// 	err := cursor.Decode(&u)
	// 	if err != nil {
	// 		// Error
	// 	}
	// 	log.Info(u)

	// 	m[u.userUUID] = u
	// }

	if err := cursor.All(ctx, &users); err != nil {
		// error
	}
	defer cursor.Close(ctx)
	log.Info(users)

}

func FindUser(ctx context.Context, database *mongo.Database) User {
	filter := bson.D{{"userUUID", "4bacf836-3d6d-401e-99dc-54879cab1975"}}
	collection := database.Collection("user")

	var result User
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	// user := User{}

	log.WithField("user", result).Info()
	return result
}

// func main() {
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

// 	// Set client options
// 	clientOptions := options.Client().ApplyURI("mongodb://fidoUser:a-dX_j4j9Vo2RJ-VTKzk@localhost:27017/fido")

// 	// Connect to MongoDB
// 	client, err := mongo.Connect(ctx, clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Check the connection
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")
// 	database := client.Database("fido")
// 	// findAll(ctx, database, "user")
// 	// fmt.Println(database)
// 	// "4bacf836-3d6d-401e-99dc-54879cab1975"
// 	var xuser User = findUser(ctx, database)
// 	fmt.Println(xuser)
// 	defer func() {
// 		if err = client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()
// }
