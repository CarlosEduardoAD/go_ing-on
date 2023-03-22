package domain

import (
	"errors"
	"regexp"
)

type Customer struct {
	Name        string `bson:"name" required:"true"`
	Email       string `bson:"email" required:"true"`
	Departament string `bson:"departament" required:"true"`
	Role        string `bson:"role" required:"true"`
}

type Roles string

const (
	ADMIN    Roles = "admin"
	EMPLOYEE Roles = "employee"
)

func CreateNewCustomer(name string, email string, department string, role Roles) *Customer {
	return &Customer{}
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return errors.New("err[domain]: empty name")
	}
	if c.Email != "" {
		rgx, _ := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		match := rgx.MatchString(c.Email)
		if !match {
			return errors.New("err[domain]: invalid email")
		}
	} else {
		return errors.New("err[domain]: empty e-mail")
	}

	if c.Departament == "" {
		return errors.New("err[domain]: invalid department")
	}

	if c.Role != "admin" && c.Role != "employee" {
		return errors.New("err[domain]: invalid role")
	}
	return nil
}
