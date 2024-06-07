package repositories

import (
	"github.com/ratchav/go-mssql/entities"
)

type QuotationRepository interface {
	GetQuotation(int) (entities.Quotation, error)
	UpdateQuotation(entities.Quotation) error
	Create(*entities.Quotation) error
}
