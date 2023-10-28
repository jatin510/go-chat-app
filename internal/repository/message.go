package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jatin510/go-chat-app/internal/models"
)

type MessageRepoInterface interface {
	Create(models.Message) (models.Message, error)
	Update(models.Message) (models.Message, error)
	Delete(uuid.UUID) error
	FindOne(any) (models.Message, error)
	FindAll(any) ([]models.Message, error)
	Count(any) (int, error)
}

type message struct {
	db *pgx.Conn
	l  models.Logger
}

func NewMessageRepo(db *pgx.Conn, l models.Logger) MessageRepoInterface {
	return &message{
		db: db,
		l:  l,
	}
}

func (m message) Create(message models.Message) (models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBQueryTimeout)
	defer cancel()
	m.l.Info("msg models", message)
	rows, err := m.db.Query(ctx, "INSERT INTO messages VALUES ($1,$2,$3,$4,$5,$6)", message.ID, message.Msg, message.RoomId, message.UserId, message.UpdatedAt, message.CreatedAt)
	if err != nil {
		m.l.Error("error in Create message query", err.Error())
		return models.Message{}, err
	}
	defer rows.Close()
	m.l.Info("rows", rows)
	return message, nil
}

func (m message) Update(message models.Message) (models.Message, error) {
	return models.Message{}, nil
}

func (m message) Delete(id uuid.UUID) error {
	return nil
}

func (m message) FindOne(filter any) (models.Message, error) {
	return models.Message{}, nil
}

func (m message) FindAll(filter any) ([]models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBQueryTimeout)
	defer cancel()

	rows, err := m.db.Query(ctx, "SELECT id, msg, room_id, user_id, created_at, updated_at FROM messages")
	if err != nil {
		m.l.Error("error in find all query", err.Error())
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message

		err := rows.Scan(&message.ID, &message.Msg, &message.RoomId, &message.UserId, &message.CreatedAt, &message.UpdatedAt)
		if err != nil {
			return []models.Message{}, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func (m message) Count(filter any) (int, error) {
	return 0, nil
}
