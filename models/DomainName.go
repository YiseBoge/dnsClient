package models

import (
	"net/rpc"
	"time"
)

type DomainName struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name     string
	Address  string
	LastRead time.Time
}

func (DomainName) FindAllLocal(client *rpc.Client) ([]DomainName, error) {
	var results []DomainName
	err := client.Call("API.FindAll", "", &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (DomainName) Lookup(client *rpc.Client, name string) ([]DomainName, error) {
	var results []DomainName
	err := client.Call("API.Lookup", name, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (model *DomainName) Register(client *rpc.Client) error {
	var result bool
	err := client.Call("API.Register", model, &result)
	if err != nil {
		return err
	}
	return nil
}

func (model *DomainName) Remove(client *rpc.Client) error {
	var result bool
	err := client.Call("API.Remove", model, &result)
	if err != nil {
		return err
	}
	return nil
}
