package CustomerRepository

import (
	"context"
	"log"

	"github.com/CarlosEduardoAD/Go-ing_on/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindData() (*mongo.Cursor, error) {
	client, err := utils.ConnectToMongo() // Mock up, mas vamos pegar a conexão daqui
	if err != nil {
		log.Fatal("Something bad happened ! -> ", err)
	}

	conn := client.Database("Go_ing_app").Collection("Clientes")
	result, err := conn.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Panic(err)
	}

	return result, nil
}
