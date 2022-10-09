package repo

import (
	"context"
	"log"

	"github.com/sinhamanav03/web-scrapper/services/scrapperstore/model"
	repoModel "github.com/sinhamanav03/web-scrapper/services/scrapperstore/repo/model"
	"gorm.io/gorm"
)

type ScrapperStoreRepo interface {
	Add(ctx context.Context, item *model.Item) (int, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, item *model.Item) error
	Get(ctx context.Context, id int) (*model.Item, error)
}

type scrapperStoreRepo struct {
	db *gorm.DB
}

func NewScrapperStoreServiceRepo(db *gorm.DB) ScrapperStoreRepo {

	db.AutoMigrate(&repoModel.Item{})
	return &scrapperStoreRepo{
		db: db,
	}
}

func (repo *scrapperStoreRepo) Add(ctx context.Context, item *model.Item) (int, error) {

	newItem := repoModel.Item{
		Name:        item.Name,
		ImageURL:    item.ImageURL,
		Description: item.Description,
		Price:       item.Price,
		Reviews:     item.Reviews,
	}

	db := repo.db.Select("name", "image_url", "description", "price", "reviews").Create(&newItem)
	if db.Error != nil {
		log.Println(db.Error)
		return 0, db.Error
	}
	return int(newItem.ID), nil

}
func (repo *scrapperStoreRepo) Delete(ctx context.Context, id int) error {
	db := repo.db.WithContext(ctx).Where("id = ?", id).Delete(&repoModel.Item{})
	if db.Error != nil {
		return db.Error
	}
	return nil

}
func (repo *scrapperStoreRepo) Update(ctx context.Context, id int, item *model.Item) error {
	updatedUser := repoModel.Item{
		Name:        item.Name,
		ImageURL:    item.ImageURL,
		Description: item.Description,
		Price:       item.Price,
		Reviews:     item.Reviews,
	}

	db := repo.db.WithContext(ctx).Model(&repoModel.Item{}).Where("id = ?", id).Updates(updatedUser)

	if db.Error != nil {
		return db.Error
	}
	return nil
}
func (repo *scrapperStoreRepo) Get(ctx context.Context, id int) (*model.Item, error) {

	item := repoModel.Item{}

	db := repo.db.WithContext(ctx).First(&item, id)
	if db.Error != nil {
		return nil, db.Error
	}

	resp := model.Item{
		Name:        item.Name,
		ImageURL:    item.ImageURL,
		Description: item.Description,
		Price:       item.Price,
		Reviews:     item.Reviews,
	}

	return &resp, nil

}
