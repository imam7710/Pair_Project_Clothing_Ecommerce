package usecase

import (
	"context"
	"time"

	"clean-archy/internal/domain"
	"clean-archy/internal/entity"
)

type reportUsecase struct {
	repo domain.ReportRepository
}

func NewReportUsecase(repo domain.ReportRepository) domain.ReportUsecase {

	return &reportUsecase{
		repo: repo,
	}
}

var timeout = 10 * time.Millisecond

func (u *reportUsecase) TotalGameSales(ctx context.Context) ([]entity.TotalGameSales, error) {

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return u.repo.GetTotalGameSales(ctx)
}
func (u *reportUsecase) MostPopularGame(ctx context.Context) ([]entity.MostPopularGame, error) {

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return u.repo.GetMostPopularGame(ctx)
}
func (u *reportUsecase) RevenuePerGame(ctx context.Context) ([]entity.RevenuePerGame, error) {

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return u.repo.GetRevenuePerGame(ctx)
}
func (u *reportUsecase) PlayerCountPerGame(ctx context.Context) ([]entity.PlayerCountPerGame, error) {

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return u.repo.GetPlayerCountPerGame(ctx)
}
