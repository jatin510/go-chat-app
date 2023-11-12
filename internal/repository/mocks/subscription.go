package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jatin510/go-chat-app/internal/models"
)

type SubscriptionRepoInterface interface {
	Create(subscription models.Subscription) (models.Subscription, error)
	Update(subscription models.Subscription) (models.Subscription, error)
	Delete(subscription uuid.UUID) error
	FindOne(filter map[string]any) (models.Subscription, error)
	FindAll(filter map[string]any) ([]models.Subscription, error)
	Count(filter map[string]any) (int, error)
}

type subscription struct {
	db *pgx.Conn
	l  models.Logger
}

func NewSubscriptionRepo(db *pgx.Conn, l models.Logger) SubscriptionRepoInterface {
	return &subscription{
		db: db,
		l:  l,
	}
}

func (r subscription) Create(subscription models.Subscription) (models.Subscription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBQueryTimeout)
	defer cancel()
	rows, err := r.db.Query(ctx, "INSERT INTO subscriptions VALUES ($1,$2,$3,$4,$5)", subscription.ID, subscription.RoomId, subscription.UserId, subscription.CreatedAt, subscription.UpdatedAt)
	if err != nil {
		r.l.Error("error in Create subscription query", err.Error())
		return models.Subscription{}, err
	}
	defer rows.Close()

	return subscription, nil
}

func (r subscription) Update(subscription models.Subscription) (models.Subscription, error) {
	return models.Subscription{}, nil
}

func (r subscription) Delete(id uuid.UUID) error {
	return nil
}

func (r subscription) FindOne(filter map[string]any) (models.Subscription, error) {
	return models.Subscription{}, nil
}

func (r subscription) FindAll(filter map[string]any) ([]models.Subscription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBQueryTimeout)
	defer cancel()

	rows, err := r.db.Query(ctx, "SELECT id, name, created_at, updated_at FROM subscriptions")
	if err != nil {
		r.l.Error("error in find all query", err.Error())
		return nil, err
	}
	defer rows.Close()

	var subscriptions []models.Subscription
	for rows.Next() {
		var subscription models.Subscription

		err := rows.Scan(&subscription.ID, &subscription.RoomId, &subscription.UserId, &subscription.CreatedAt, &subscription.UpdatedAt)
		if err != nil {
			return []models.Subscription{}, err
		}

		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions, nil
}

func (r subscription) FindAllSubscriptionsByUserId(userId uuid.UUID) ([]models.Subscription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBQueryTimeout)
	defer cancel()
	rows, err := r.db.Query(ctx, "SELECT * FROM subscriptions WHERE user_id = $1", userId)

	if err != nil {
		r.l.Error("error in FindAllsubscriptionsByUserId query", err.Error())
		return []models.Subscription{}, err
	}

	defer rows.Close()

	var subscriptions []models.Subscription
	for rows.Next() {
		var subscription models.Subscription

		err := rows.Scan(&subscription)
		if err != nil {
			return []models.Subscription{}, err
		}

		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions, nil
}

func (r subscription) Count(filter map[string]any) (int, error) {
	return 0, nil
}
