package news

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Store struct {
	db bun.IDB
}

func NewStore(db bun.IDB) *Store {
	return &Store{
		db: db,
	}
}

// Create news record
func (s Store) Create(ctx context.Context, news Record) (createdNews Record, err error) {
	news.ID = uuid.New()
	err = s.db.NewInsert().Model(&news).Returning("*").Scan(ctx, &createdNews)
	if err != nil {
		return createdNews, NewCustomError(err, http.StatusInternalServerError)
	}

	return createdNews, nil
}

// FindByID finds a news record with the provided id
func (s Store) FindByID(ctx context.Context, id uuid.UUID) (news Record, err error) {
	err = s.db.NewSelect().Model(&news).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return news, NewCustomError(err, http.StatusNotFound)
		}
		return news, NewCustomError(err, http.StatusInternalServerError)
	}

	return news, nil
}

// FindAll returns all news store in the database
func (s Store) FindAll(ctx context.Context) (news []Record, err error) {
	err = s.db.NewSelect().Model(&news).Scan(ctx, &news)
	return news, err
}

// DeleteByID deletes a news by its ID
func (s Store) DeleteByID(ctx context.Context, id uuid.UUID) (err error) {
	_, err = s.db.NewDelete().Model(&Record{}).Where("id = ?", id).Returning("NULL").Exec(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return NewCustomError(err, http.StatusInternalServerError)
	}

	return nil
}

// UpdateByID update news by it's ID
func (s Store) UpdateByID(ctx context.Context, id uuid.UUID, news Record) (err error) {
	r, err := s.db.NewUpdate().Model(&news).Where("id = ?", id).Returning("NULL").Exec(ctx)
	if err != nil {
		return NewCustomError(err, http.StatusInternalServerError)
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return NewCustomError(err, http.StatusInternalServerError)
	}
	if rowsAffected == 0 {
		return NewCustomError(err, http.StatusNotFound)
	}

	return nil
}
