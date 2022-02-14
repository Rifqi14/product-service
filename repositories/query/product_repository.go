package query

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewQueryProductRepository(db *gorm.DB) query.IProductRepository {
	return &ProductRepository{DB: db}
}

func (repo ProductRepository) List(search, orderBy, sort, productName string, limit, offset, minPrice, maxPrice int64, brand []*uuid.UUID, product []*uuid.UUID, color []*uuid.UUID) (res []*models.Product, count int64, err error) {
	tx := repo.DB
	// search = strings.ToLower(search)
	productName = strings.ToLower(productName)

	if len(color) > 0 {
		tx = tx.Joins("left join product_colors as ps on ps.product_id = products.id").Where("ps.color_id in ?", color)
	}
	if len(product) > 0 {
		tx = tx.Where("id in ?", product)
	}
	if len(brand) > 0 {
		tx = tx.Where("brand_id in ?", color)
	}
	if minPrice > 0 && maxPrice > 0 {
		tx = tx.Where("final_price between ? and ?", minPrice, maxPrice)
	}

	err = tx.Where("lower(name) like ?", "%"+productName+"%").Preload(clause.Associations).Preload("Logs.Attachment").Preload("Logs.Verifier").Preload("Variants.Color").Preload("Images.Color").Preload("Images.Image").Preload("Brand." + clause.Associations).Preload("Categories." + clause.Associations).Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Where("lower(name) like ?", "%"+productName+"%").Preload(clause.Associations).Preload("Logs.Attachment").Preload("Variants.Color").Preload("Images.Color").Preload("Images.Image").Preload("Logs.Verifier").Preload("Brand." + clause.Associations).Order(orderBy + " " + sort).Preload("Categories." + clause.Associations).Find(&res).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	return res, count, nil
}

func (repo ProductRepository) Detail(productId uuid.UUID) (res *models.Product, err error) {
	tx := repo.DB

	err = tx.Preload(clause.Associations).Preload("Variants.Color").Preload("Images.Color").Preload("Images.Image").Preload("Logs.Attachment").Preload("Logs.Verifier").Preload("Brand."+clause.Associations).Preload("Categories."+clause.Associations).Find(&res, "id = ?", productId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (repo ProductRepository) FindBy(column, operator string, value interface{}) (res []*models.Product, err error) {
	tx := repo.DB

	err = tx.Preload(clause.Associations).Preload("Variants."+clause.Associations).Preload("Images."+clause.Associations).Preload("Logs."+clause.Associations).Preload("Brand."+clause.Associations).Preload("Categories."+clause.Associations).Find(&res, column+" "+operator+" ?", value).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
