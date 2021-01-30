package database

import (
	"context"
	"fidowebapp/config"
	"fmt"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	ContextKeyDatabase string = "ContextKey.Database"
)

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
