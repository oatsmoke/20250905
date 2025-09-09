package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oatsmoke/20250905/internal/lib/err_msg"
	"github.com/oatsmoke/20250905/internal/lib/logger"
	"github.com/oatsmoke/20250905/internal/model"
)

type SubscriptionRepository struct {
	postgresDB *pgxpool.Pool
}

func New(postgresDB *pgxpool.Pool) *SubscriptionRepository {
	return &SubscriptionRepository{
		postgresDB: postgresDB,
	}
}

func (r *SubscriptionRepository) Create(ctx context.Context, subscription *model.Subscription) error {
	var id int64
	const query = `
		INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;`

	if err := r.postgresDB.QueryRow(
		ctx,
		query,
		subscription.ServiceName,
		subscription.Price,
		subscription.UserId,
		subscription.StartDate,
		subscription.EndDate,
	).Scan(&id); err != nil {
		return err
	}

	if id == 0 {
		return err_msg.NoRowsAffected
	}

	logger.Info(fmt.Sprintf("subscription with id %d created", id))
	return nil
}

func (r *SubscriptionRepository) Read(ctx context.Context, subscriptionId int64) (*model.Subscription, error) {
	subscription := new(model.Subscription)
	const query = `
		SELECT id, service_name, price, user_id, start_date, end_date
		FROM subscriptions
		WHERE id = $1;`

	if err := r.postgresDB.QueryRow(ctx, query, subscriptionId).Scan(
		&subscription.ID,
		&subscription.ServiceName,
		&subscription.Price,
		&subscription.UserId,
		&subscription.StartDate,
		&subscription.EndDate,
	); err != nil {
		return nil, err
	}

	logger.Info(fmt.Sprintf("subscription with id %d read", subscriptionId))
	return subscription, nil
}

func (r *SubscriptionRepository) Update(ctx context.Context, subscription *model.Subscription) error {
	const query = `
		UPDATE subscriptions
		SET service_name = $2, price = $3, user_id = $4, start_date = $5, end_date = $6
		WHERE id = $1;`

	tag, err := r.postgresDB.Exec(
		ctx,
		query,
		subscription.ID,
		subscription.ServiceName,
		subscription.Price,
		subscription.UserId,
		subscription.StartDate,
		subscription.EndDate,
	)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return err_msg.NoRowsAffected
	}

	logger.Info(fmt.Sprintf("subscription with id %d updated", subscription.ID))
	return nil
}

func (r *SubscriptionRepository) Delete(ctx context.Context, subscriptionId int64) error {
	const query = `
		DELETE FROM subscriptions
		WHERE id = $1;`

	tag, err := r.postgresDB.Exec(ctx, query, subscriptionId)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return err_msg.NoRowsAffected
	}

	logger.Info(fmt.Sprintf("subscription with id %d deleted", subscriptionId))
	return nil
}

func (r *SubscriptionRepository) List(ctx context.Context) ([]*model.Subscription, error) {
	var subscriptions []*model.Subscription
	const query = `
		SELECT id, service_name, price, user_id, start_date, end_date
		FROM subscriptions;`

	rows, err := r.postgresDB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		subscription := new(model.Subscription)
		if err := rows.Scan(
			&subscription.ID,
			&subscription.ServiceName,
			&subscription.Price,
			&subscription.UserId,
			&subscription.StartDate,
			&subscription.EndDate,
		); err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	logger.Info(fmt.Sprintf("%d subscriptions listed\n", len(subscriptions)))
	return subscriptions, nil
}

func (r *SubscriptionRepository) Total(ctx context.Context, subscription *model.Subscription) (int64, error) {
	var total int64
	const query = `
		SELECT SUM(price) 
		FROM subscriptions
		WHERE start_date BETWEEN $3 AND $4
		GROUP BY user_id, service_name
		HAVING user_id = $1 AND service_name = $2;`

	if err := r.postgresDB.QueryRow(
		ctx,
		query,
		subscription.UserId,
		subscription.ServiceName,
		subscription.StartDate,
		subscription.EndDate,
	).Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
