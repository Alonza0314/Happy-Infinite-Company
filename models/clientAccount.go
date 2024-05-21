package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	Username      string
	Email         string
	Password      string
}

func NewClient(u, e, p string) Client {
	return Client{Username: u, Email: e, Password: p}
}

func ProcessSignup(client Client) error {
	collection, err := SetCollection("mongodb.uri", "mongodb.database", "mongodb.clientsAccountCollection")
	if err != nil {
		return err
	}

	// 檢查username是否存在
	var existingUsernameClient Client
	err = collection.FindOne(context.Background(), bson.M{"username": client.Username}).Decode(&existingUsernameClient)
	if existingUsernameClient.Username != "" {
		return errors.New("使用者名稱已存在")
	} else if err != mongo.ErrNoDocuments {
		return err
	}

	// 檢查email是否存在
	var existingEmailClient Client
	err = collection.FindOne(context.Background(), bson.M{"email": client.Email}).Decode(&existingEmailClient)
	if existingEmailClient.Email != "" {
		return errors.New("電子郵件已存在")
	} else if err != mongo.ErrNoDocuments {
		return err
	}

	_, err = collection.InsertOne(context.Background(), client)
	if err != nil {
		return err
	}

	return nil
}

func ProcessLogin(client Client) error {
	collection, err := SetCollection("mongodb.uri", "mongodb.database", "mongodb.clientsAccountCollection")
	if err != nil {
		return err
	}

	// 檢查username是否存在資料庫裡面
	var existingUsernameClient Client
	err = collection.FindOne(context.Background(), bson.M{"username": client.Username}).Decode(&existingUsernameClient)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("使用者名稱不存在")
		}
		return err
	}

	// 檢查密碼
	if client.Password != existingUsernameClient.Password {
		return errors.New("密碼錯誤")
	}

	return nil
}

func ProcessFindpw(client Client) error {
	collection, err := SetCollection("mongodb.uri", "mongodb.database", "mongodb.clientsAccountCollection")
	if err != nil {
		return err
	}


	// 檢查username是否存在資料庫裡面
	var existingUsernameClient Client
	err = collection.FindOne(context.Background(), bson.M{"username": client.Username}).Decode(&existingUsernameClient)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("使用者名稱不存在")
		}
		return err
	}

	// 檢查email是否正確
	if client.Email != existingUsernameClient.Email {
		return errors.New("電子郵件錯誤")
	}

	return nil
}

func ProcessResetpw(client Client) error {
	collection, err := SetCollection("mongodb.uri", "mongodb.database", "mongodb.clientsAccountCollection")
	if err != nil {
		return err
	}

	// 設定filter
	filter := map[string]interface{} {
		"username": client.Username,
	}

	// 設定update資料
	update := map[string]interface{} {
		"$set": map[string]interface{} {
			"password": client.Password,
		},
	}

	// 限制時間
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// 更新密碼
	if _, err = collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.New("服務器密碼更新失敗")
	}

	return nil
}
