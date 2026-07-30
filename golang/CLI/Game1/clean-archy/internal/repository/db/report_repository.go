package db

import (
	"context"
	"database/sql"

	"clean-archy/internal/domain"
	"clean-archy/internal/entity"
)

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) domain.ReportRepository {

	return &reportRepository{
		db: db,
	}
}
func (r *reportRepository) GetTotalGameSales(ctx context.Context) ([]entity.TotalGameSales, error) {
	query := `
		SELECT
		g.game_name,
		COUNT(s.sale_id) AS total_sales
		FROM Games g
		JOIN Sales s
		ON g.game_id = s.game_id
		GROUP BY g.game_id,g.game_name
		ORDER BY total_sales DESC;
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reports []entity.TotalGameSales

	for rows.Next() {

		var report entity.TotalGameSales

		err := rows.Scan(
			&report.GameName,
			&report.TotalSales,
		)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *reportRepository) GetMostPopularGame(ctx context.Context) ([]entity.MostPopularGame, error) {
	query := `
		SELECT
		g.game_name,
		COUNT(s.sale_id) AS total_sales
		FROM Games g
		JOIN Sales s
		ON g.game_id=s.game_id
		GROUP BY g.game_id,g.game_name
		ORDER BY total_sales DESC
		LIMIT 1;
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []entity.MostPopularGame

	for rows.Next() {

		var report entity.MostPopularGame

		err := rows.Scan(
			&report.GameName,
			&report.TotalSales,
		)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *reportRepository) GetRevenuePerGame(ctx context.Context) ([]entity.RevenuePerGame, error) {
	query := `
		SELECT
		g.game_name,
		SUM(s.total_price) AS revenue
		FROM Games g
		JOIN Sales s
		ON g.game_id=s.game_id
		GROUP BY g.game_id,g.game_name
		ORDER BY revenue DESC;
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []entity.RevenuePerGame

	for rows.Next() {

		var report entity.RevenuePerGame

		err := rows.Scan(
			&report.GameName,
			&report.Revenue,
		)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *reportRepository) GetPlayerCountPerGame(ctx context.Context) ([]entity.PlayerCountPerGame, error) {
	query := `
		SELECT
		g.game_name,
		COUNT(DISTINCT s.player_id) AS total_players
		FROM Games g
		JOIN Sales s
		ON g.game_id=s.game_id
		GROUP BY g.game_id,g.game_name
		ORDER BY total_players DESC;
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []entity.PlayerCountPerGame

	for rows.Next() {

		var report entity.PlayerCountPerGame

		err := rows.Scan(
			&report.GameName,
			&report.TotalPlayers,
		)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil

}
