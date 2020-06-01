package postgresql

import (
	"fmt"
	"math"
	"strconv"
	"time"

	config "go-lms-of-pupilfirst/configs"
	"go-lms-of-pupilfirst/pkg/database/object"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //
)

type handler struct {
	client        *gorm.DB
	transactionDb *gorm.DB
	dialInfo      postgresqlInfo
	transaction   bool
}

// postgresqlInfo db connection info with postgresql
type postgresqlInfo struct {
	database string
	host     string
	port     string
	user     string
	password string
}

// NewHandler creates postgresql handler
func NewHandler(conf config.Storage) *handler {
	h := &handler{
		dialInfo: postgresqlInfo{
			database: conf.Database,
			host:     conf.Host,
			port:     conf.Port,
			user:     conf.Dbuser,
			password: conf.Dbpassword,
		},
	}

	return h
}

// Prepare prepares database for the system (migrate and other setup)
func (h *handler) Prepare() error {
	url := fmt.Sprintf("host=%s port%s user=%s password%s dbname%s sslmode=disabled", h.dialInfo.host, h.dialInfo.port, h.dialInfo.user, h.dialInfo.password, h.dialInfo.database)
	db, err := gorm.Open("postgres", url)
	if err != nil {
		return err
	}

	h.client = db
	h.transactionDb = nil
	h.transaction = false
	return nil
}

func (h *handler) Close() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("close err %s", r)
		}
	}()

	h.client.Close()
	return
}

func (h *handler) List(doc object.Interfaces, query interface{}) error {

	err := h.client.Where(query).Where("archived_at IS NULL").Find(doc)

	if err != nil {
		return err.Error
	}
	return nil
}

func (h *handler) ListPagination(doc object.Interfaces, query interface{}, pagination object.Pagination, advancedSearch object.AdvancedSearch) error {
	offest := pagination.GetPage() * pagination.GetLimit()
	//fetch doc
	dbQuery := h.client.Where(query).Where("archived_at IS NULL")
	if advancedSearch.GetSearchField() != "" {
		dbQuery = dbQuery.Where(advancedSearch.GetSearchField()+" LIKE ?", "%"+advancedSearch.GetSearchQuery()+"%")
	}
	if advancedSearch.GetSearchInField() != "" {
		dbQuery = dbQuery.Where(advancedSearch.GetSearchInField()+" IN (?)", advancedSearch.GetSearchInQuery())
	}
	if pagination.GetLimit() != 0 {
		dbQuery = dbQuery.Order(pagination.GetSortField() + " " + pagination.GetSortOrder()).Offset(offest).Limit(pagination.GetLimit())
	}
	err := dbQuery.Find(doc)
	//err := h.client.Where(query).Where("archived_at IS NULL").Order(pagination.GetSortField() + " " + pagination.GetSortOrder()).Offset(offest).Limit(pagination.GetLimit()).Find(doc)
	if err.Error != nil {
		return err.Error
	}
	//fetch count
	var totalRecord uint32
	if err := h.RecordCount(doc.GetNameSpace(), query, advancedSearch, &totalRecord); err != nil {
		return err
	}
	//calculate total number of pages
	totalPage := uint32(math.Ceil(float64(totalRecord) / float64(pagination.GetLimit())))
	pages := []string{}
	var i uint32
	for i = 1; i <= totalPage; i++ {
		pages = append(pages, string(strconv.FormatUint(uint64(i), 10)))
	}
	//set record pages
	pagination.SetPages(pages)
	//check if page has previous page
	if pagination.GetPage()+1 == 1 {
		pagination.SetHasPrev(false)
	} else {
		pagination.SetHasPrev(true)
	}
	//check if page has next page
	if pagination.GetPage()+1 >= uint16(len(pages)) {
		pagination.SetHasNext(false)
	} else {
		pagination.SetHasNext(true)
	}
	return nil
}

func (h *handler) FindOne(doc object.Interface, query interface{}) error {

	err := h.client.Where(query).Where("archived_at IS NULL").First(doc)
	if err.Error != nil {
		if gorm.IsRecordNotFoundError(err.Error) {
			return nil
		}
		return err.Error
	}

	return nil
}

func (h *handler) One(doc object.Interface) error {
	err := h.client.Where("id = ? AND archived_at IS NULL", doc.GetID()).First(doc)
	if err != nil {

		return err.Error
	}
	return nil
}

func (h *handler) Insert(doc object.Interface) error {
	doc.SetArchivedAt(nil)
	client := h.currentDb()
	err := client.Create(doc)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (h *handler) Update(doc object.Interface, query interface{}) error {

	doc.SetUpdatedAt(time.Now())
	client := h.currentDb()
	err := client.Table(doc.GetNameSpace()).Where(query).Where("archived_at IS NULL").Updates(doc)
	if err != nil {
		return err.Error
	}

	return nil
}

func (h *handler) Drop(doc object.Interface) error {
	client := h.currentDb()
	err := client.Unscoped().Delete(doc)
	if err != nil {
		return err.Error
	}

	return nil
}

func (h *handler) Remove(doc object.Interface) error {
	client := h.currentDb()
	result := client.Table(doc.GetNameSpace()).Where("id = ?", doc.GetID()).Where("archived_at IS NULL").Update("archived_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (h *handler) RemoveMany(doc object.Interface, query interface{}) error {
	client := h.currentDb()
	result := client.Table(doc.GetNameSpace()).Where(query).Where("archived_at IS NULL").Updates(map[string]interface{}{"archived_at": time.Now()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (h *handler) BeginTransaction() error {
	if h.transaction == false {
		//begin transaction
		h.transactionDb = h.client.Begin()
		h.transaction = true
	} else {
		return fmt.Errorf("Transaction have already begin")
	}
	return nil
}

func (h *handler) Rollback() error {
	if h.transaction == true {
		h.transactionDb.Rollback()
		h.transaction = false
		h.transactionDb = nil
	} else {
		return fmt.Errorf("Transaction was not started")
	}
	return nil
}

func (h *handler) CommitTransaction() error {
	if h.transaction == true {
		h.transactionDb.Commit()
		h.transaction = false
		h.transactionDb = nil
	} else {
		return fmt.Errorf("Transaction was not started")
	}
	return nil
}

func (h *handler) currentDb() *gorm.DB {
	if h.transaction == true {
		return h.transactionDb
	} else {
		return h.client
	}
}

func (h *handler) RecordCount(tableNameSpace string, query interface{}, advancedSearch object.AdvancedSearch, count *uint32) error {
	dbQuery := h.client.Table(tableNameSpace).Where(query).Where("archived_at IS NULL")
	if advancedSearch.GetSearchField() != "" {
		dbQuery = dbQuery.Where(advancedSearch.GetSearchField()+" LIKE ?", "%"+advancedSearch.GetSearchQuery()+"%")
	}
	if advancedSearch.GetSearchInField() != "" {
		dbQuery = dbQuery.Where(advancedSearch.GetSearchInField()+" IN (?)", advancedSearch.GetSearchInQuery())
	}

	if result := dbQuery.Count(count); result.Error != nil {
		return result.Error
	}
	return nil
}
