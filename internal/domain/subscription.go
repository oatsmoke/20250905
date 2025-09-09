package domain

import (
	"context"

	"github.com/oatsmoke/20250905/internal/model"
)

type Subscription interface {
	Create(ctx context.Context, subscription *model.Subscription) error
	Read(ctx context.Context, subscriptionId int64) (*model.Subscription, error)
	Update(ctx context.Context, subscription *model.Subscription) error
	Delete(ctx context.Context, subscriptionId int64) error
	List(ctx context.Context) ([]*model.Subscription, error)
	Total(ctx context.Context, subscription *model.Subscription) (int64, error)
}
