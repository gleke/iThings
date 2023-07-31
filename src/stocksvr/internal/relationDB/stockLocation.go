package relationDB

import (
	"context"
	"github.com/i-Things/things/shared/stores"
	"gorm.io/gorm"
)

type StockLocationRepo struct {
	db *gorm.DB
}

func NewStockLocationRepo(in any) *StockLocationRepo {
	return &StockLocationRepo{db: stores.GetCommonConn(in)}
}

func (p StockLocationRepo) Insert(ctx context.Context, data *StockLocation) error {
	result := p.db.WithContext(ctx).Create(data)
	return stores.ErrFmt(result.Error)
}

func (p StockLocationRepo) FindOneByProductID(ctx context.Context, productID string) (*StockLocation, error) {
	var result StockLocation
	err := p.db.WithContext(ctx).Where("product_id = ?", productID).First(&result).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return &result, nil
}

func (p StockLocationRepo) Update(ctx context.Context, data *StockLocation) error {
	err := p.db.WithContext(ctx).Save(data).Error
	return stores.ErrFmt(err)
}
