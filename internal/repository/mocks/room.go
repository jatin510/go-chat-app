package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jatin510/go-chat-app/internal/models"
)

type RoomRepoInterface interface {
	Create(room models.Room) (models.Room, error)
	Update(room models.Room) (models.Room, error)
	Delete(room uuid.UUID) error
	FindOne(filter map[string]any) (models.Room, error)
	FindAll(filter map[string]any) ([]models.Room, error)
	FindAllRoomsByUserId(userId uuid.UUID) ([]models.Room, error)
	Count(filter map[string]any) (int, error)
}

type room struct {
	db *pgx.Conn
	l  models.Logger
}

func NewRoomRepo(db *pgx.Conn, l models.Logger) RoomRepoInterface {
	return &room{
		db: db,
		l:  l,
	}
}

func (r room) Create(room models.Room) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBQueryTimeout)
	defer cancel()
	rows, err := r.db.Query(ctx, "INSERT INTO rooms VALUES ($1,$2,$3,$4)", room.ID, room.Name, room.CreatedAt, room.UpdatedAt)
	if err != nil {
		r.l.Error("error in Create room query", err.Error())
		return models.Room{}, err
	}
	defer rows.Close()

	return room, nil
}

func (r room) Update(room models.Room) (models.Room, error) {
	return models.Room{}, nil
}

func (r room) Delete(id uuid.UUID) error {
	return nil
}

func (r room) FindOne(filter map[string]any) (models.Room, error) {
	return models.Room{}, nil
}

func (r room) FindAll(filter map[string]any) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBQueryTimeout)
	defer cancel()

	rows, err := r.db.Query(ctx, "SELECT id, name, created_at, updated_at FROM rooms")
	if err != nil {
		r.l.Error("error in find all query", err.Error())
		return nil, err
	}
	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		var room models.Room

		err := rows.Scan(&room.ID, &room.Name, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return []models.Room{}, err
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r room) FindAllRoomsByUserId(userId uuid.UUID) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBQueryTimeout)
	defer cancel()
	rows, err := r.db.Query(ctx, "SELECT DISTINCT r.id AS id FROM rooms AS r, subscriptions AS s WHERE s.user_id = $1 AND r.id = s.room_id", userId)

	if err != nil {
		r.l.Error("error in FindAllRoomsByUserId query", err.Error())
		return []models.Room{}, err
	}

	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		var room models.Room

		err := rows.Scan(&room.ID)
		if err != nil {
			return []models.Room{}, err
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r room) Count(filter map[string]any) (int, error) {
	return 0, nil
}
