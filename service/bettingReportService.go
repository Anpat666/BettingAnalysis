package service

import (
	"BettingAnalysis/controllers"
	"BettingAnalysis/cores"
	"BettingAnalysis/models"

	"github.com/xuri/excelize/v2"
)

type BetReportService struct {
	BetReportController *controllers.BetReportController
	BetDataTotal        [][]string
	F                   *excelize.File
}

func NewBetReportService() *BetReportService {
	return &BetReportService{
		BetReportController: &controllers.BetReportController{
			BettingReport: &models.BettingReport{},
			BetGames:      &models.BettingGame{},
			TxtPath:       "document/投注评估.txt",
		},

		F: cores.OpenExcel("document/投注报表.xlsx"),
	}
}

func (b *BetReportService) GetBetReportService() {
	cores.ClearDocument(b.BetReportController.TxtPath)
	b.BetDataTotal = cores.GetExcelRows(b.F, "代理端-报表管理-投注详情_0")
	b.BetReportController.BettingReport.PlayerName = cores.GetUniqueValuesInColumn(b.BetDataTotal, 3)
	b.BetReportController.BettingReport.BettingTimeStart = b.BetDataTotal[len(b.BetDataTotal)-1][17]
	b.BetReportController.BettingReport.BettingTimeEnd = b.BetDataTotal[1][17]
	b.BetReportController.BettingReport.BettingOrderQuantity = float64(len(b.BetDataTotal) - 1)
	b.BetReportController.BettingReport.BettingAmount = cores.SumColumnValues(b.BetDataTotal, 11)
	b.BetReportController.BettingReport.WinOrLose = cores.SumColumnValues(b.BetDataTotal, 16)
	b.BetReportController.BettingReport.GamesName = cores.GetUniqueValuesInColumn(b.BetDataTotal, 4)
	b.BetReportController.BetReportFormatContent()
	b.BetReportController.GamesClass = make([][][]string, len(b.BetReportController.BettingReport.GamesName))
	b.GetBetGamesClass()
	for i := 0; i < len(b.BetReportController.GamesClass); i++ {
		b.BetReportController.BetGames.GameName = b.BetReportController.GamesClass[i][i][4]
		b.BetReportController.BetGames.BettingOrderQuantity = float64(len(b.BetReportController.GamesClass[i]) - 1)
		b.BetReportController.BetGames.BetOrderPro = cores.ProPortion(b.BetReportController.BettingReport.BettingOrderQuantity, b.BetReportController.BetGames.BettingOrderQuantity)
		b.BetReportController.BetGames.BettingAmount = cores.SumColumnValues(b.BetReportController.GamesClass[i], 11)
		b.BetReportController.BetGames.BetAmountPro = cores.ProPortion(b.BetReportController.BettingReport.BettingAmount, b.BetReportController.BetGames.BettingAmount)
		b.BetReportController.BetGames.WinOrLose = cores.SumColumnValues(b.BetReportController.GamesClass[i], 16)
		b.BetReportController.BetGames.WinOrLosePro = cores.ProPortion(b.BetReportController.BettingReport.WinOrLose, b.BetReportController.BetGames.WinOrLose)

	}

}

// func (b *BetReportService) GetGamesDataService(num int) {
// 	b.BetReportController.BetGames.GameName =
// }

func (b *BetReportService) GetBetGamesClass() {
	for i := 0; i < len(b.BetReportController.BettingReport.GamesName); i++ {
		b.BetReportController.GamesClass[i] = make([][]string, 0)
		for k := range b.BetDataTotal {
			if b.BetDataTotal[k][4] == b.BetReportController.BettingReport.GamesName[i] {
				b.BetReportController.GamesClass[i] = append(b.BetReportController.GamesClass[i], b.BetDataTotal[k])
			}
		}
	}

}
