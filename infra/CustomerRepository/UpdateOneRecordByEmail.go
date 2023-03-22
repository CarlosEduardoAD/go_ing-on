package CustomerRepository

import (
	"context"
	"log"
	"net/http"

	"github.com/CarlosEduardoAD/Go-ing_on/domain"
	"github.com/CarlosEduardoAD/Go-ing_on/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateOneRecordByEmail(c *domain.Customer, filter bson.D) (int64, error) {
	updateCustomer := bson.D{{"$set", bson.D{{"name", c.Name}, {"email", c.Email}, {"department", c.Departament}, {"role", c.Role}}}}

	client, err := utils.ConnectToMongo() // Mock up, mas vamos pegar a conexão daqui
	if err != nil {
		log.Fatal("Something bad happened ! -> ", err)
	}

	defer client.Disconnect(context.TODO())

	coll := client.Database("Go_ing_app").Collection("Clientes")
	_, err = coll.UpdateOne(context.TODO(), filter, updateCustomer)
	if err != nil {
		log.Panic(err)
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil

}
