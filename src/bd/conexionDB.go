package bd

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConnectBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://caritoApiGo:Wa4qzSBvwEOAi4eX@caroservices.yfii8.mongodb.net/twitter_db?retryWrites=true&w=majority")

//ConnectBD Conecta base de datos
func ConnectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("soy error")
		log.Fatal(err.Error())
		return client
	}
	fmt.Println("Conexión exitosa bbcita")
	return client
}

//CheckConnection chequea conexión
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
