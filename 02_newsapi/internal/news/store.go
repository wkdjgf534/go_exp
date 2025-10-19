package news

import (
	"context"

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
		return createdNews, err
	}

	return createdNews, nil
}

// FindByID finds a news record with the provided id
func (s Store) FindByID(ctx context.Context, id uuid.UUID) (news Record, err error) {
	err = s.db.NewSelect().Model(&news).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return news, err
	}

	return news, nil
}

// FindAll returns all news store in the database
func (s Store) FindAll(ctx context.Context) (news []Record, err error) {
	err = s.db.NewSelect().Model(&news).Scan(ctx, &news)
	if err != nil {
		return news, err
	}

	return news, nil
}

// DeleteByID deletes a news by its ID
func (s Store) DeleteByID(ctx context.Context, id uuid.UUID) (err error) {
	_, err = s.db.NewDelete().Model(&Record{}).Where("id = ?", id).Returning("NULL").Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// UpdateByID update news by it's ID
func (s Store) UpdateByID(ctx context.Context, id uuid.UUID, news Record) (err error) {
	_, err = s.db.NewUpdate().Model(&news).Where("id = ?", id).Returning("NULL").Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
