package services

import (
	"net/http"

	"github.com/CarlosEduardoAD/Go-ing_on/domain"
	"github.com/CarlosEduardoAD/Go-ing_on/infra/CustomerRepository"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateOneRecordByEmail(c *domain.Customer, filter bson.D) (int64, error) {
	res, err := CustomerRepository.UpdateOneRecordByEmail(c, filter)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return res, nil
}
