package utils

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

// Base serves as a base model for other models
type Base struct {
	ID         string     `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()" validate:"omitempty,uuid,required"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	ArchivedAt *time.Time `sql:"index" json:"deleted_at"`
}

// BeforeCreate generates a uuid into the id before saving the model
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewRandom().String()
	return scope.SetColumn("ID", uuid)
}

// Error defines error type to be returned by handlers
type Error struct {
	Errors    map[string][]string
	Code      int
	Message   string
	Timestamp time.Time
}

// GenerateError generates a new error
func GenerateError(errors map[string][]string, code int, message string) *Error {
	return &Error{
		Errors:    errors,
		Code:      code,
		Message:   message,
		Timestamp: time.Now(),
	}
}

func (e *Error) Error() string {
	return e.Message
}

// PaginationQuery defines pagination and sorting query
type PaginationQuery struct {
	Page        uint16
	Limit       uint16
	HasPrevious bool
	HasNext     bool
	Pages       []string
	SortField   string
	SortOrder   string
}

// GetPage fetches page number
func (p *PaginationQuery) GetPage() uint16 {
	return p.Page
}

// GetLimit fetches limit of records per page number
func (p *PaginationQuery) GetLimit() uint16 {
	return p.Limit
}

// GetHasNext fetches if record has next page
func (p *PaginationQuery) GetHasNext() bool {
	return p.HasNext
}

// GetHasPrev fetches if record has previous page
func (p *PaginationQuery) GetHasPrev() bool {
	return p.HasPrevious
}

// SetHasNext sets if record has next page
func (p *PaginationQuery) SetHasNext(next bool) {
	p.HasNext = next
}

// SetHasPrev sets if record has previous page
func (p *PaginationQuery) SetHasPrev(prev bool) {
	p.HasPrevious = prev
}

// GetPages fetches record pages
func (p *PaginationQuery) GetPages() []string {
	return p.Pages
}

// SetPages sets record pages
func (p *PaginationQuery) SetPages(pages []string) {
	p.Pages = pages
}

// GetSortField fetches sort field
func (p *PaginationQuery) GetSortField() string {
	return p.SortField
}

// GetSortOrder fetches sort order
func (p *PaginationQuery) GetSortOrder() string {
	return p.SortOrder
}

// AdvancedSearch defines advanced search query
type AdvancedSearch struct {
	SearchField   string
	SearchQuery   string
	SearchInField string
	SearchInQuery []interface{}
}

// GetSearchField retrieves search field
func (as *AdvancedSearch) GetSearchField() string {
	return as.SearchField
}

// SetSearchField sets search field
func (as *AdvancedSearch) SetSearchField(field string) {
	as.SearchField = field
}

// GetSearchQuery retrieves search query
func (as *AdvancedSearch) GetSearchQuery() string {
	return as.SearchQuery
}

// SetSearchQuery sets search query
func (as *AdvancedSearch) SetSearchQuery(query string) {
	as.SearchQuery = query
}

// GetSearchInField retrieves search in field
func (as *AdvancedSearch) GetSearchInField() string {
	return as.SearchInField
}

// SetSearchInField sets search in field
func (as *AdvancedSearch) SetSearchInField(field string) {
	as.SearchInField = field
}

// GetSearchInQuery retrieves search in query
func (as *AdvancedSearch) GetSearchInQuery() []interface{} {
	return as.SearchInQuery
}

// SetSearchInQuery sets search in query
func (as *AdvancedSearch) SetSearchInQuery(query []interface{}) {
	as.SearchInQuery = query
}
