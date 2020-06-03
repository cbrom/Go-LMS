package object

import "time"

//Interface defines model structs
type Interface interface {
	TableName() string
	GetID() string
	SetID(string)
	SetCreatedAt(t time.Time)
	SetUpdatedAt(t time.Time)
	SetArchivedAt(t *time.Time)
}

//Interfaces
type Interfaces interface {
	TableName() string
}

//Pagination defines list page pagination and sorting query
type Pagination interface {
	GetPage() uint16
	GetLimit() uint16
	GetHasNext() bool
	GetHasPrev() bool
	SetHasNext(bool)
	SetHasPrev(bool)
	GetPages() []string
	SetPages([]string)
	GetSortField() string
	GetSortOrder() string
}

//AdvancedSearch defines query for like and in query
type AdvancedSearch interface {
	GetSearchField() string
	SetSearchField(string)
	GetSearchQuery() string
	SetSearchQuery(string)
	GetSearchInField() string
	SetSearchInField(string)
	GetSearchInQuery() []interface{}
	SetSearchInQuery([]interface{})
}
