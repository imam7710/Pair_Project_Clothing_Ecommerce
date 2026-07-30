package entity

type TotalGameSales struct {
	GameName   string
	TotalSales int
}

type MostPopularGame struct {
	GameName   string
	TotalSales int
}

type RevenuePerGame struct {
	GameName string
	Revenue  float64
}

type PlayerCountPerGame struct {
	GameName     string
	TotalPlayers int
}
