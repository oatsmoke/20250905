package service

import (
	"context"

	"github.com/oatsmoke/20250905/internal/domain"
	"github.com/oatsmoke/20250905/internal/model"
)

type SubscriptionService struct {
	subscriptionRepository domain.Subscription
}

func New(subscriptionRepository domain.Subscription) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepository: subscriptionRepository,
	}
}

func (s *SubscriptionService) Create(ctx context.Context, subscription *model.Subscription) error {
	return s.subscriptionRepository.Create(ctx, subscription)
}

func (s *SubscriptionService) Read(ctx context.Context, subscriptionId int64) (*model.Subscription, error) {
	return s.subscriptionRepository.Read(ctx, subscriptionId)
}

func (s *SubscriptionService) Update(ctx context.Context, subscription *model.Subscription) error {
	return s.subscriptionRepository.Update(ctx, subscription)
}

func (s *SubscriptionService) Delete(ctx context.Context, subscriptionId int64) error {
	return s.subscriptionRepository.Delete(ctx, subscriptionId)
}

func (s *SubscriptionService) List(ctx context.Context) ([]*model.Subscription, error) {
	return s.subscriptionRepository.List(ctx)
}

func (s *SubscriptionService) Total(ctx context.Context, subscription *model.Subscription) (int64, error) {
	return s.subscriptionRepository.Total(ctx, subscription)
}
