package main

import (
	"Pair_Project_Clothing_Ecommerce/config"
	"Pair_Project_Clothing_Ecommerce/internal/delivery/cli"
	"Pair_Project_Clothing_Ecommerce/internal/repository/mysql"
	"Pair_Project_Clothing_Ecommerce/internal/usecase"
	"fmt"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("go Kritis: Gagal terhubung ke MySQL:", err)
		return
	}
	defer db.Close()

	userRepo := mysql.NewUserRepository(db)
	prodRepo := mysql.NewProductRepository(db)
	orderRepo := mysql.NewOrderRepository(db)
	reportRepo := mysql.NewReportRepository(db)

	authUC := usecase.NewAuthUseCase(userRepo)
	userUC := usecase.NewUserUseCase(userRepo)
	prodUC := usecase.NewProductUseCase(prodRepo)
	orderUC := usecase.NewOrderUseCase(orderRepo)
	reportUC := usecase.NewReportUseCase(reportRepo)

	app := cli.NewApp(authUC, userUC, prodUC, orderUC, reportUC)
	app.Start()
}
