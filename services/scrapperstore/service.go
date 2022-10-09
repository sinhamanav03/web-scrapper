package scrapperstore

import (
	"context"

	"github.com/sinhamanav03/web-scrapper/services/scrapperstore/model"
	"github.com/sinhamanav03/web-scrapper/services/scrapperstore/repo"
)

type ScrapperStoreService interface {
	Add(ctx context.Context, item *model.Item) (int, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, item *model.Item) error
	Get(ctx context.Context, id int) (*model.Item, error)
}

type scrapperService struct {
	repo repo.ScrapperStoreRepo
}

func NewScrapperStoreService(repo repo.ScrapperStoreRepo) ScrapperStoreService {
	return &scrapperService{
		repo: repo,
	}
}

func (svc *scrapperService) Add(ctx context.Context, item *model.Item) (int, error) {
	id, err := svc.repo.Add(ctx, item)
	if err != nil {
		return 0, err
	}
	return id, err
}
func (svc *scrapperService) Delete(ctx context.Context, id int) error {
	err := svc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (svc *scrapperService) Update(ctx context.Context, id int, item *model.Item) error {
	err := svc.repo.Update(ctx, id, item)
	if err != nil {
		return err
	}
	return nil
}
func (svc *scrapperService) Get(ctx context.Context, id int) (*model.Item, error) {
	item, err := svc.repo.Get(ctx, id)

	if err != nil {
		return nil, err
	}
	return item, nil
}
