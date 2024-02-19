package controllers

import (
	"BettingAnalysis/cores"
	"BettingAnalysis/models"
	"fmt"
)

type BetReportController struct {
	BettingReport  *models.BettingReport
	BetGames       *models.BettingGame
	BettingMethods *models.BettingMethod
	Content        string
	TxtPath        string
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

func (b *BetReportController) BetGamesFormatContent(num int) {
	b.Content = fmt.Sprintf("%v、%s：\n", num+1, b.BetGames.GameName)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintf("   (1)注单数：%v（占总注单数%s）\n", b.BetGames.BettingOrderQuantity, b.BetGames.BetOrderPro)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintf("   (2)有效投注：%v（占总有效金额%s）\n", b.BetGames.BettingAmount, b.BetGames.BetAmountPro)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintf("   (3)游戏输赢：%v（占总输赢%s）\n", b.BetGames.WinOrLose, b.BetGames.WinOrLosePro)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintln("   (4)主要投注玩法：")
	cores.UpDataReport(b.Content, b.TxtPath)
}

func (b *BetReportController) BetGameMethodFormatContent(num int) {
	b.Content = fmt.Sprintf("       [%v]%s 注单数%v,有效投注%v,输赢%v,胜率%s,胜率系数%.2f,玩法等级评估:%s\n",
		num+1,
		b.BettingMethods.MethodName,
		b.BettingMethods.BettingOrderQuantity,
		b.BettingMethods.BettingAmount,
		b.BettingMethods.WinOrLose,
		b.BettingMethods.WinningRate,
		b.BettingMethods.WinRateCoe,
		b.BettingMethods.GamePlayLevel,
	)
	cores.UpDataReport(b.Content, b.TxtPath)
}

func (b *BetReportController) AccountLevelFormatContent() {
	b.Content = fmt.Sprintf("该会员查询时间段存在大额单注盈利情况（%v注，请自行到投注报表查询）\n", b.BettingReport.BigProfitBet)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintf("账总总体胜率系数：%.2f\n", b.BettingReport.CoeTotal)
	cores.UpDataReport(b.Content, b.TxtPath)

	b.Content = fmt.Sprintf("账号系统综合等级评估：%s\n", b.BettingReport.AccountLevel)
	cores.UpDataReport(b.Content, b.TxtPath)
}
