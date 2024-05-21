package models

import (
	"context"
	"errors"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetCollection(uri, database, collection string) (*mongo.Collection, error) {
	viper.SetConfigType("toml")
	viper.SetConfigFile("configs/config.conf")
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New("服務器錯誤")
	}

	clientOptions := options.Client().ApplyURI(viper.GetString(uri))

	db, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, errors.New("服務器資料庫錯誤")
	}

	err = db.Ping(context.Background(), nil)
	if err != nil {
		return nil, errors.New("服務器資料庫錯誤")
	}

	return db.Database(viper.GetString(database)).Collection(viper.GetString(collection)), nil
}
