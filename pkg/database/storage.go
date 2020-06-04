package database

import (
	"errors"
	"fmt"

	config "go-lms-of-pupilfirst/configs"
	"go-lms-of-pupilfirst/pkg/database/object"
	"go-lms-of-pupilfirst/pkg/database/postgresql"

	"github.com/jinzhu/gorm"
)

type (
	handlerName string

	Interface interface {
		Prepare() (*gorm.DB, error)
		Close() error
		//DropCollection(doc interface{}) error
		//DropDatabase(doc interface{}) error
		FindOne(doc object.Interface, query interface{}) error
		One(doc object.Interface) error
		List(doc object.Interfaces, query interface{}) error
		ListPagination(doc object.Interfaces, query interface{}, pagination object.Pagination, advancedSearch object.AdvancedSearch) error
		//ListParent(doc interface{}) error
		Insert(doc object.Interface) error
		Update(doc object.Interface, query interface{}) error
		Remove(doc object.Interface) error
		RemoveMany(doc object.Interface, query interface{}) error
		Drop(doc object.Interface) error
		BeginTransaction() error
		Rollback() error
		CommitTransaction() error
	}
)

var handler Interface

//New creates new db handler
func New(storageConfig config.Storage) error {
	switch storageConfig.HandlerName {
	case "postgresql":
		fmt.Println(storageConfig)
		handler = postgresql.NewHandler(storageConfig)
	default:
		return errors.New("Invalid storage handler `" + storageConfig.HandlerName + "`")
	}
	_, err := handler.Prepare()
	return err
}

//Handler returns storage handler
func Handler() Interface {

	return handler
}
