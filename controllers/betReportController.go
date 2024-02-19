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
	b.Content = fmt.Sprintf("%s：\n", b.BetGames.GameName)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintf("(1)注单数：%v（占总注单数%s）", b.BetGames.BettingOrderQuantity, b.BetGames.BetOrderPro)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintf("(2)有效投注：%v（占总有效金额%s）", b.BetGames.BettingAmount, b.BetGames.BetAmountPro)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintf("(3)游戏输赢：%v（占总输赢%s）", b.BetGames.WinOrLose, b.BetGames.WinOrLosePro)
	cores.UpDataReport(b.Content, b.TxtPath)
}
