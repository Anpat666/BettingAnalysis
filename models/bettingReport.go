package models

type BettingReport struct {
	PlayerName           []string //会员账号
	BettingTimeStart     string   //投注开始时间
	BettingTimeEnd       string   //投注结束时间
	BettingOrderQuantity float64  //投注订单数量
	BettingAmount        float64  //总有效投注金额
	WinOrLose            float64  //总输赢
	BigProfitBet         int      //大额盈利注单数量
	CoeTotal             float64  //账号总体系数
	AccountLevel         string   //账号等级评估
	GamesName            []string //投注过的游戏
}

type BettingGame struct {
	GameName             string   //游戏名称
	BettingOrderQuantity float64  //投注订单数量
	BetOrderPro          string   //注单数占比
	BettingAmount        float64  //总有效投注金额
	BetAmountPro         string   //有效投注占比
	WinOrLose            float64  //总输赢
	WinOrLosePro         string   //总输赢占比
	MethodsTotal         []string //投注过的玩法
}

type BettingMethod struct {
	MethodName           string  //玩法名称
	BettingOrderQuantity float64 //投注订单数量
	BettingAmount        float64 //总有效投注金额
	WinOrLose            float64 //总输赢
	MaxSingleWin         float64 //单注最高盈利金额
	WinningRate          string  //胜率
	WinRateCoe           float64 //胜率系数
	GamePlayLevel        string  //玩法等级评估
	CoePrO               float64 //系数占比
}
