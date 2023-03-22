package services

import (
	"log"
	"net/http"

	"github.com/CarlosEduardoAD/Go-ing_on/domain"
	"github.com/CarlosEduardoAD/Go-ing_on/infra/CustomerRepository"
)

func InsertUserService(c *domain.Customer) (int64, error) {
	res, err := CustomerRepository.InsertOneRecord(c)

	if err != nil {
		log.Panic("err[service]: invalid domain, repo failed => ", err)
		return http.StatusBadRequest, nil
	}

	return res, nil
}
