package relationDB

import (
	"context"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/stores"
	"gorm.io/gorm"
)

type StockMoveRepo struct {
	db *gorm.DB
}

type ProductFilter struct {
	DeviceType   int64
	ProductName  string
	ProductIDs   []string
	ProductNames []string
	Tags         map[string]string
}

func NewStockMoveRepo(in any) *StockMoveRepo {
	return &StockMoveRepo{db: stores.GetCommonConn(in)}
}

func (p StockMoveRepo) fmtFilter(ctx context.Context, f ProductFilter) *gorm.DB {
	db := p.db.WithContext(ctx)
	if f.DeviceType != 0 {
		db = db.Where("device_type=?", f.DeviceType)
	}
	if f.ProductName != "" {
		db = db.Where("product_name like ?", "%"+f.ProductName+"%")
	}
	if len(f.ProductIDs) != 0 {
		db = db.Where("product_id = ?", f.ProductIDs)
	}
	if len(f.ProductNames) != 0 {
		db = db.Where("product_name = ?", f.ProductNames)
	}
	if f.Tags != nil {
		for k, v := range f.Tags {
			db = db.Where("JSON_CONTAINS(tags, JSON_OBJECT(?,?))",
				k, v)
		}
	}
	return db
}

func (p StockMoveRepo) Insert(ctx context.Context, data *StockMove) error {
	result := p.db.WithContext(ctx).Create(data)
	return stores.ErrFmt(result.Error)
}

func (p StockMoveRepo) FindOneByFilter(ctx context.Context, f ProductFilter) (*StockMove, error) {
	var result StockMove
	db := p.fmtFilter(ctx, f)
	err := db.First(&result).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return &result, nil
}

func (p StockMoveRepo) Update(ctx context.Context, data *StockMove) error {
	err := p.db.WithContext(ctx).Where("product_id = ?", data.ProductID).Save(data).Error
	return stores.ErrFmt(err)
}

func (p StockMoveRepo) DeleteByFilter(ctx context.Context, f ProductFilter) error {
	db := p.fmtFilter(ctx, f)
	err := db.Delete(&StockMove{}).Error
	return stores.ErrFmt(err)
}

func (p StockMoveRepo) FindByFilter(ctx context.Context, f ProductFilter, page *def.PageInfo) ([]*StockMove, error) {
	var results []*StockMove
	db := p.fmtFilter(ctx, f).Model(&StockMove{})
	db = page.ToGorm(db)
	err := db.Find(&results).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return results, nil
}

func (p StockMoveRepo) CountByFilter(ctx context.Context, f ProductFilter) (size int64, err error) {
	db := p.fmtFilter(ctx, f).Model(&StockMove{})
	err = db.Count(&size).Error
	return size, stores.ErrFmt(err)
}
