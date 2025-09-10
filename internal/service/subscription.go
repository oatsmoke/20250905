package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/oatsmoke/20250905/internal/lib/err_msg"
	"github.com/oatsmoke/20250905/internal/lib/logger"
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

type SubscriptionService struct {
	subscriptionRepository Subscription
}

func New(subscriptionRepository Subscription) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepository: subscriptionRepository,
	}
}

func (s *SubscriptionService) Create(ctx context.Context, data *model.ExternalData) error {
	subscription, err := mapIn(data)
	if err != nil {
		return err
	}

	if subscription.EndDate != nil && subscription.StartDate.After(*subscription.EndDate) {
		return err_msg.LaterDate
	}

	return s.subscriptionRepository.Create(ctx, subscription)
}

func (s *SubscriptionService) Read(ctx context.Context, subscriptionId int64) (*model.ExternalData, error) {
	read, err := s.subscriptionRepository.Read(ctx, subscriptionId)
	if err != nil {
		return nil, err
	}

	return mapOut(read), nil
}

func (s *SubscriptionService) Update(ctx context.Context, subscriptionId int64, data *model.ExternalData) error {
	subscription, err := mapIn(data)
	if err != nil {
		return err
	}
	subscription.ID = subscriptionId

	return s.subscriptionRepository.Update(ctx, subscription)
}

func (s *SubscriptionService) Delete(ctx context.Context, subscriptionId int64) error {
	return s.subscriptionRepository.Delete(ctx, subscriptionId)
}

func (s *SubscriptionService) List(ctx context.Context) ([]*model.ExternalData, error) {
	subscriptions, err := s.subscriptionRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	list := make([]*model.ExternalData, len(subscriptions))
	for i, subscription := range subscriptions {
		list[i] = mapOut(subscription)
	}

	return list, nil
}

func (s *SubscriptionService) Total(ctx context.Context, data *model.ExternalData) (int64, error) {
	subscription, err := mapIn(data)
	if err != nil {
		return 0, err
	}

	if subscription.EndDate != nil && subscription.StartDate.After(*subscription.EndDate) {
		return 0, err_msg.LaterDate
	}

	total, err := s.subscriptionRepository.Total(ctx, subscription)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	logTotal := strconv.Itoa(int(total))
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		logTotal = err.Error()
	}

	logger.Info(fmt.Sprintf("user: %s, service: %s, from: %s, to: %s -> %s",
		data.UserId,
		data.ServiceName,
		data.StartDate,
		data.EndDate,
		logTotal,
	))

	return total, nil
}

func mapIn(data *model.ExternalData) (*model.Subscription, error) {
	var err error
	layout := "02-01-2006"
	subscription := &model.Subscription{
		ServiceName: data.ServiceName,
		Price:       data.Price,
		UserId:      data.UserId,
	}

	subscription.StartDate, err = time.Parse(layout, fmt.Sprintf("01-%s", data.StartDate))
	if err != nil {
		return nil, err
	}

	if data.EndDate != "" {
		subscription.EndDate = new(time.Time)
		*subscription.EndDate, err = time.Parse(layout, fmt.Sprintf("01-%s", data.EndDate))
		if err != nil {
			return nil, err
		}
	}

	return subscription, nil
}

func mapOut(subscription *model.Subscription) *model.ExternalData {
	layout := "01-2006"
	data := &model.ExternalData{
		ID:          subscription.ID,
		ServiceName: subscription.ServiceName,
		Price:       subscription.Price,
		UserId:      subscription.UserId,
	}

	data.StartDate = subscription.StartDate.Format(layout)

	if subscription.EndDate != nil {
		data.EndDate = subscription.EndDate.Format(layout)
	}

	return data
}
