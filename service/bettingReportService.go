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
	BetMethodClass      [][][]string
	GamesClass          [][][]string
	F                   *excelize.File
}

func NewBetReportService() *BetReportService {
	return &BetReportService{
		BetReportController: &controllers.BetReportController{
			BettingReport:  &models.BettingReport{},
			BetGames:       &models.BettingGame{},
			BettingMethods: &models.BettingMethod{},
			TxtPath:        "document/投注评估.txt",
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
	b.GamesClass = make([][][]string, len(b.BetReportController.BettingReport.GamesName))
	b.GetBetGamesClass()

	for i := 0; i < len(b.GamesClass); i++ {
		b.BetReportController.BetGames.GameName = b.GamesClass[i][0][4]
		b.BetReportController.BetGames.BettingOrderQuantity = float64(len(b.GamesClass[i]))
		b.BetReportController.BetGames.BetOrderPro = cores.ProPortion(b.BetReportController.BettingReport.BettingOrderQuantity, b.BetReportController.BetGames.BettingOrderQuantity)
		b.BetReportController.BetGames.BettingAmount = cores.SumColumnValues(b.GamesClass[i], 11)
		b.BetReportController.BetGames.BetAmountPro = cores.ProPortion(b.BetReportController.BettingReport.BettingAmount, b.BetReportController.BetGames.BettingAmount)
		b.BetReportController.BetGames.WinOrLose = cores.SumColumnValues(b.GamesClass[i], 16)
		b.BetReportController.BetGames.WinOrLosePro = cores.ProPortion(b.BetReportController.BettingReport.WinOrLose, b.BetReportController.BetGames.WinOrLose)
		b.BetReportController.BetGamesFormatContent(i)

		b.BetReportController.BetGames.MethodsTotal = cores.GetUniqueValuesInColumn(b.GamesClass[i], 12)
		b.BetMethodClass = make([][][]string, len(b.BetReportController.BetGames.MethodsTotal))
		b.GetBetGameMethodsClass(i)
		for k := 0; k < len(b.BetMethodClass); k++ {
			b.BetReportController.BettingMethods.MethodName = b.BetMethodClass[k][0][12]
			b.BetReportController.BettingMethods.BettingOrderQuantity = float64(len(b.BetMethodClass[k]))
			b.BetReportController.BettingMethods.BettingAmount = cores.SumColumnValues(b.BetMethodClass[k], 11)
			b.BetReportController.BettingMethods.WinOrLose = cores.SumColumnValues(b.BetMethodClass[k], 16)
			b.BetReportController.BettingMethods.WinningRate = cores.GetWinningRate(b.BetMethodClass[k])
			b.BetReportController.BettingMethods.WinRateCoe = cores.GetWinRateCoe(b.BetMethodClass[k])
			b.BetReportController.BettingMethods.MaxSingleWin = cores.GetMaxWin(b.BetMethodClass[k])
			b.BetReportController.BettingMethods.GamePlayLevel = cores.GameMethodLevelJudgment(
				b.BetReportController.BettingMethods.WinOrLose,
				b.BetReportController.BettingMethods.WinRateCoe,
				b.BetReportController.BettingMethods.MaxSingleWin,
				b.BetReportController.BettingMethods.BettingAmount,
			)
			b.BetReportController.BettingMethods.CoePrO = b.BetReportController.BettingMethods.WinRateCoe * (b.BetReportController.BettingMethods.BettingAmount / b.BetReportController.BettingReport.BettingAmount)
			b.BetReportController.BettingReport.CoeTotal += b.BetReportController.BettingMethods.CoePrO
			b.BetReportController.BetGameMethodFormatContent(k)
		}
	}
	b.BetReportController.BettingReport.AccountLevel = cores.AccountLevelJudgment(b.BetReportController.BettingReport.WinOrLose, b.BetReportController.BettingReport.CoeTotal)
	b.BetReportController.AccountLevelFormatContent()

}

func (b *BetReportService) GetBetGamesClass() {
	for i := 0; i < len(b.BetReportController.BettingReport.GamesName); i++ {
		b.GamesClass[i] = make([][]string, 0)
		for k := range b.BetDataTotal {
			if b.BetDataTotal[k][4] == b.BetReportController.BettingReport.GamesName[i] {
				b.GamesClass[i] = append(b.GamesClass[i], b.BetDataTotal[k])
			}
		}
	}

}

func (b *BetReportService) GetBetGameMethodsClass(num int) {
	for i := 0; i < len(b.BetReportController.BetGames.MethodsTotal); i++ {
		b.BetMethodClass[i] = make([][]string, 0)
		for k := range b.GamesClass[num] {
			if b.GamesClass[num][k][12] == b.BetReportController.BetGames.MethodsTotal[i] {
				b.BetMethodClass[i] = append(b.BetMethodClass[i], b.GamesClass[num][k])
			}
		}
	}
}
