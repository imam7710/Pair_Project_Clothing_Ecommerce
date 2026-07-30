package main

import (
	"log"

	"clean-archy/internal/config"
	handlers "clean-archy/internal/handler"
	dbrepo "clean-archy/internal/repository/db"
	"clean-archy/internal/usecase"
)

func main() {

	db, err := config.InitDB()

	if err != nil {

		log.Fatal(err)
	}

	defer db.Close()

	reportRepo := dbrepo.NewReportRepository(db)

	reportUsecase := usecase.NewReportUsecase(reportRepo)

	reportHandler := handlers.NewReportHandler(reportUsecase)

	reportHandler.ShowMenu()
}
