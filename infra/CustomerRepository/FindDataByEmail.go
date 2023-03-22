package CustomerRepository

import (
	"context"
	"log"

	"github.com/CarlosEduardoAD/Go-ing_on/domain"
	"github.com/CarlosEduardoAD/Go-ing_on/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindDataByEmail(c *domain.Customer) (*mongo.Cursor, error) {
	filter := bson.D{{"email", c.Email}}

	client, err := utils.ConnectToMongo() // Mock up, mas vamos pegar a conexão daqui
	if err != nil {
		log.Fatal("Something bad happened ! -> ", err)
	}

	defer client.Disconnect(context.TODO())

	conn := client.Database("Go_ing_app").Collection("Clientes")
	result, err := conn.Find(context.TODO(), filter)
	if err != nil {
		log.Panic(err)
	}

	return result, nil
}
