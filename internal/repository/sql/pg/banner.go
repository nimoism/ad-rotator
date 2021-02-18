package pg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nimoism/ad-rotator/internal/entity"
)

type BannerRepo struct {
	db *sql.DB
}

func NewBannerRepo(db *sql.DB) *BannerRepo {
	return &BannerRepo{db: db}
}

func (r *BannerRepo) BannerStats(ctx context.Context, slotID, userGroupID int) ([]entity.BannerStat, error) {
	query := "" +
		"SELECT " +
		"	b.id, b.name, " +
		"	coalesce(sum(se.count), 0) show_count, " +
		"	coalesce(sum(ce.count), 0) click_count " +
		"FROM banner b " +
		"LEFT JOIN banner_slot bs ON b.id = bs.banner_id " +
		"LEFT JOIN show_banner_event_stat se ON b.id = se.banner_id AND se.user_group_id = $2 " +
		"LEFT JOIN click_banner_event_stat ce ON b.id = ce.banner_id AND ce.user_group_id = $2 " +
		"WHERE bs.slot_id = $1 " +
		"GROUP BY b.id, b.name"
	rows, err := r.db.QueryContext(ctx, query, slotID, userGroupID)
	if err != nil {
		return nil, fmt.Errorf("banners stats querying error: %w", err)
	}
	defer rows.Close()
	stats := make([]entity.BannerStat, 0)
	for rows.Next() {
		banner := entity.Banner{}
		var showCount int
		var clickCount int
		if err = rows.Scan(&banner.ID, &banner.Name, &showCount, &clickCount); err != nil {
			return nil, fmt.Errorf("banner stat mapping error: %w", err)
		}
		stats = append(stats, entity.BannerStat{Banner: banner, ShowCount: showCount, ClickCount: clickCount})
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("banners stats getting error: %w", err)
	}
	return stats, nil
}

func (r *BannerRepo) CreateBannerClick(ctx context.Context, click entity.ClickEvent) error {
	updateQuery := "" +
		"WITH updated AS (" +
		"	UPDATE click_banner_event_stat " +
		"	SET count = count + 1 " +
		"	WHERE banner_id = $1 AND slot_id = $2 AND user_group_id = $3 " +
		"	RETURNING id" +
		") SELECT count(id) FROM updated"
	insertQuery := "INSERT INTO click_banner_event_stat " +
		"(banner_id, slot_id, user_group_id, count) VALUES " +
		"($1, $2, $3, $4)"
	return r.upsertEvent(ctx, insertQuery, updateQuery, click.BannerID, click.SlotID, click.UserGroupID)
}

func (r *BannerRepo) CreateBannerShow(ctx context.Context, show entity.ShowEvent) error {
	updateQuery := "" +
		"WITH updated AS (" +
		"	UPDATE show_banner_event_stat " +
		"	SET count = count + 1 " +
		"	WHERE banner_id = $1 AND slot_id = $2 AND user_group_id = $3 " +
		"	RETURNING id" +
		") SELECT count(id) FROM updated"
	insertQuery := "" +
		"INSERT INTO show_banner_event_stat " +
		"(banner_id, slot_id, user_group_id, count) VALUES " +
		"($1, $2, $3, $4)"
	return r.upsertEvent(ctx, insertQuery, updateQuery, show.BannerID, show.SlotID, show.UserGroupID)
}

func (r *BannerRepo) upsertEvent(ctx context.Context, insQ, updQ string, bannerID, slotID, ugID int) error {
	res := r.db.QueryRowContext(ctx, updQ, bannerID, slotID, ugID)
	var count int
	if err := res.Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	if _, insertErr := r.db.ExecContext(ctx, insQ, bannerID, slotID, ugID, 1); insertErr != nil {
		res = r.db.QueryRowContext(ctx, updQ, bannerID, slotID, ugID)
		if err := res.Scan(&count); err != nil {
			return err
		}
		if count == 0 {
			return insertErr
		}
	}
	return nil
}

func (r *BannerRepo) Banners(ctx context.Context) ([]entity.Banner, error) {
	query := "SELECT id, name FROM banner ORDER BY id"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	banners := make([]entity.Banner, 0)
	for rows.Next() {
		banner := entity.Banner{}
		if err = rows.Scan(&banner.ID, &banner.Name); err != nil {
			return nil, fmt.Errorf("banner DB mapping error: %w", err)
		}
		banners = append(banners, banner)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("banners getting error: %w", err)
	}
	return banners, nil
}

func (r *BannerRepo) CreateBanner(ctx context.Context, banner *entity.Banner) error {
	query := "INSERT INTO banner (name) VALUES ($1) RETURNING id"
	result := r.db.QueryRowContext(ctx, query, banner.Name)
	if err := result.Scan(&banner.ID); err != nil {
		return err
	}
	return result.Err()
}

func (r *BannerRepo) UpdateBanner(ctx context.Context, banner *entity.Banner) error {
	query := "UPDATE banner SET name = $2 WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, banner.ID, banner.Name)
	return err
}

func (r *BannerRepo) DeleteBanner(ctx context.Context, id int) error {
	query := "DELETE FROM banner WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *BannerRepo) CreateBannerSlot(ctx context.Context, bannerID, slotID int) error {
	query := "INSERT INTO banner_slot (banner_id, slot_id) VALUES ($1, $2)"
	_, err := r.db.ExecContext(ctx, query, bannerID, slotID)
	return err
}

func (r *BannerRepo) RemoveBannerSlot(ctx context.Context, bannerID, slotID int) error {
	query := "DELETE FROM banner_slot WHERE banner_id = $1 AND slot_id = $2"
	_, err := r.db.ExecContext(ctx, query, bannerID, slotID)
	return err
}
