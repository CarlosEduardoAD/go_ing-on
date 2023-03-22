package services

import (
	"net/http"

	"github.com/CarlosEduardoAD/Go-ing_on/domain"
	"github.com/CarlosEduardoAD/Go-ing_on/infra/CustomerRepository"
)

func DeleteOneRecordByEmail(c *domain.Customer) (int64, error) {
	res, err := CustomerRepository.DeleteOneRecordByEmail(c)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return res, nil
}
