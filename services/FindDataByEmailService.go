package services

import (
	"context"
	"log"

	"github.com/CarlosEduardoAD/Go-ing_on/domain"
	"github.com/CarlosEduardoAD/Go-ing_on/infra/CustomerRepository"
	"go.mongodb.org/mongo-driver/bson"
)

func FindDataByEmail(c *domain.Customer) (interface{}, error) {
	result, err := CustomerRepository.FindDataByEmail(c)

	if err != nil {
		log.Panic(err)
	}

	var resultToSend []interface{}

	for result.Next(context.TODO()) {
		var mongoMap bson.M
		if err := result.Decode(&mongoMap); err != nil {
			panic(err)
		}

		log.Println(mongoMap)

		bsonBytes, err := bson.Marshal(mongoMap)

		if err != nil {
			log.Panic("Erro no bson", err)
		}

		var jsonDoc interface{}
		err = bson.Unmarshal(bsonBytes, &jsonDoc)

		if err != nil {
			log.Panic("Erro no bson", err)
		}

		res, err := bson.MarshalExtJSON(jsonDoc, true, false)
		if err != nil {
			log.Panic("Erro no json", err)
		}

		resultToSend = append(resultToSend, string(res))

	}

	return resultToSend, nil
}
