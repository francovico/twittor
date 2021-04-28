package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN es el objeto de conexion a la BD.
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:FiV434819@cluster0.w9an1.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

// ConectarBD es la funcion que me permite conectar la base de datos.
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la BD")
	return client
}

// Chequeo conexion es el Ping a la BD.
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
