package CustomerRepository

import (
	"context"
	"log"
	"net/http"

	"github.com/CarlosEduardoAD/Go-ing_on/domain"
	"github.com/CarlosEduardoAD/Go-ing_on/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteOneRecordByEmail(c *domain.Customer) (int64, error) {
	filter := bson.D{{"email", c.Email}}

	client, err := utils.ConnectToMongo() // Mock up, mas vamos pegar a conexão daqui
	if err != nil {
		log.Fatal("Something bad happened ! -> ", err)
	}

	defer client.Disconnect(context.TODO())

	coll := client.Database("Go_ing_app").Collection("Clientes")
	_, err = coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Panic(err)
	}

	return http.StatusAccepted, nil
}
