package controllers

import (
	"BettingAnalysis/cores"
	"BettingAnalysis/models"
	"fmt"
)

type BetReportController struct {
	BettingReport *models.BettingReport
	BetGames      *models.BettingGame
	GamesClass    [][][]string
	Content       string
	TxtPath       string
}

func (b *BetReportController) BetReportFormatContent() {
	b.Content = cores.GetPlayerName(b.BettingReport.PlayerName)
	b.Content = fmt.Sprintf("会员账号：%s\n", b.Content)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintf("查询投注时间段：%s至%s\n", b.BettingReport.BettingTimeStart, b.BettingReport.BettingTimeEnd)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintf("总投注注单数：%v\n", b.BettingReport.BettingOrderQuantity)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintf("总有效投注金额:%v\n", b.BettingReport.BettingAmount)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintf("总输赢:%v\n", b.BettingReport.WinOrLose)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintln("主要投注游戏如下：")
	cores.UpDataReport(b.Content, b.TxtPath)

}

func (b *BetReportController) BetGamesFormatContent() {
}
