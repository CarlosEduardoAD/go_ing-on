package CustomerRepository

import (
	"context"
	"log"
	"net/http"

	"github.com/CarlosEduardoAD/Go-ing_on/domain"
	"github.com/CarlosEduardoAD/Go-ing_on/utils"
)

func InsertOneRecord(c *domain.Customer) (int64, error) {
	err := c.Validate()

	if err != nil {
		log.Panic("err[repo]: invalid customer => ", err)
	}

	client, err := utils.ConnectToMongo() // Mock up, mas vamos pegar a conexão daqui
	if err != nil {
		log.Fatal("Something bad happened ! -> ", err)
	}

	defer client.Disconnect(context.TODO())

	conn := client.Database("Go_ing_app").Collection("Clientes")

	_, err = conn.InsertOne(context.TODO(), c)
	if err != nil {
		log.Panic("err[repo]: failed inserting data => ", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil
}
