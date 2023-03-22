package interfaces

import (
	"github.com/CarlosEduardoAD/Go-ing_on/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type CustomerInterface interface {
	FindData() (interface{}, error)
	FindDataByEmail(customer *domain.Customer) (interface{}, error)
	DeleteOneRecordByEmail(customer *domain.Customer) (string, error)
	InsertOneRecord(customer *domain.Customer) (string, error)
	UpdateOneRecordByEmail(customer *domain.Customer, filter bson.D) (interface{}, error)
}
