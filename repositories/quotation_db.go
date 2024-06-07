package repositories

import (
	"github.com/ratchav/go-mssql/entities"
	"gorm.io/gorm"
)

type quotationRepositoryDB struct {
	db *gorm.DB
}

func NewQuotationRepositoryDB(db *gorm.DB) QuotationRepository {
	db.AutoMigrate(&entities.Quotation{})
	db.AutoMigrate(&entities.QuotationItem{})
	return quotationRepositoryDB{db}
}

func (q quotationRepositoryDB) GetQuotation(int) (quotation entities.Quotation, err error) {
	err = q.db.Preload("Items").Find(&quotation).Error
	return quotation, err
}

func (q quotationRepositoryDB) UpdateQuotation(entities.Quotation) error {
	panic("unimplemented")
}

// Create implements QuotationRepository.
func (q quotationRepositoryDB) Create(quotation *entities.Quotation) error {
	return q.db.Create(quotation).Error
}
