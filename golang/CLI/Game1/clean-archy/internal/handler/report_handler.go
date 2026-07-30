package handlers

import (
	"context"
	"fmt"

	"clean-archy/internal/domain"
)

type ReportHandler struct {
	usecase domain.ReportUsecase
}

func NewReportHandler(u domain.ReportUsecase) *ReportHandler {

	return &ReportHandler{
		usecase: u,
	}
}

func (h *ReportHandler) ShowMenu() {

	fmt.Println("Select report to generate:")
	fmt.Println("1. Total Game Sales Report")
	fmt.Println("2. Most Popular Game Report")
	fmt.Println("3. Total Revenue Per Game Report")
	fmt.Println("4. Player Count Per Game Report")
	fmt.Println("0. Exit")

	var choice int

	fmt.Print("Choose : ")
	fmt.Scan(&choice)

	switch choice {

	case 1:

		data, _ := h.usecase.TotalGameSales(context.Background())

		fmt.Println()

		fmt.Println("Total Game Sales Report")

		for _, d := range data {

			fmt.Printf(
				"Game : %-20s Total Sales : %d\n",
				d.GameName,
				d.TotalSales,
			)
		}

	case 2:

		data, _ := h.usecase.MostPopularGame(context.Background())

		fmt.Println()

		fmt.Println("Most Popular Game")

		for _, d := range data {

			fmt.Printf(
				"%s (%d Sales)\n",
				d.GameName,
				d.TotalSales,
			)
		}

	case 3:

		data, _ := h.usecase.RevenuePerGame(context.Background())

		fmt.Println()

		fmt.Println("Revenue Per Game")

		for _, d := range data {

			fmt.Printf(
				"%-20s Rp %.0f\n",
				d.GameName,
				d.Revenue,
			)
		}

	case 4:

		data, _ := h.usecase.PlayerCountPerGame(context.Background())

		fmt.Println()

		fmt.Println("Player Count")

		for _, d := range data {

			fmt.Printf(
				"%-20s %d Player\n",
				d.GameName,
				d.TotalPlayers,
			)
		}
	case 0:
		fmt.Println("Thank you for using Game Report System")
		return

	default:

		fmt.Println("Invalid Menu")
	}
}
