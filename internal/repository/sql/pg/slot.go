package pg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nimoism/ad-rotator/internal/entity"
)

type SlotRepo struct {
	db *sql.DB
}

func NewSlotRepo(db *sql.DB) *SlotRepo {
	return &SlotRepo{db: db}
}

func (r *SlotRepo) Slots(ctx context.Context) ([]entity.Slot, error) {
	query := "SELECT id, name FROM slot ORDER BY id"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return r.fetchSlots(rows)
}

func (r *SlotRepo) SlotsByBanner(ctx context.Context, bannerID int) ([]entity.Slot, error) {
	query := "SELECT s.id, s.name FROM slot s WHERE id IN (SELECT slot_id FROM banner_slot WHERE banner_id = $1)"
	rows, err := r.db.QueryContext(ctx, query, bannerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.fetchSlots(rows)
}

func (r *SlotRepo) fetchSlots(rows *sql.Rows) ([]entity.Slot, error) {
	var err error
	slots := make([]entity.Slot, 0)
	for rows.Next() {
		slot := entity.Slot{}
		if err = rows.Scan(&slot.ID, &slot.Name); err != nil {
			return nil, fmt.Errorf("slot db mapping error: %w", err)
		}
		slots = append(slots, slot)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("getting slots error: %w", err)
	}
	return slots, nil
}

func (r *SlotRepo) CreateSlot(ctx context.Context, slot *entity.Slot) error {
	query := "INSERT INTO slot (name) VALUES ($1) RETURNING id"
	result := r.db.QueryRowContext(ctx, query, slot.Name)
	if err := result.Scan(&slot.ID); err != nil {
		return err
	}
	return result.Err()
}

func (r *SlotRepo) UpdateSlot(ctx context.Context, slot *entity.Slot) error {
	query := "UPDATE slot SET name = $2 WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, slot.ID, slot.Name)
	return err
}

func (r *SlotRepo) DeleteSlot(ctx context.Context, id int) error {
	query := "DELETE FROM slot WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
