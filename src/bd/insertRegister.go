package bd

import (
	"context"
	"time"

	"github.com/fake_twitter/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertRegister crea un nuevo usuario en BD
func InsertRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter_db")

	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjId, _ := result.InsertedID.(primitive.ObjectID)

	return ObjId.String(), true, nil
}
