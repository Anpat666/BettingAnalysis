package cores

func GameMethodLevelJudgment(profit float64, WinRateCoe float64, maxSingleWin float64, totalBet float64) string {
	var level string

	switch {
	case profit < 0:
		level = "安全"
		if WinRateCoe >= 0.8 {
			level = "正常"
		}
	case profit < 3000:
		level = "正常"
		if WinRateCoe >= 1.0 {
			level = "持续观察"
		}
	case profit >= 3000 && profit < 10000:
		level = "持续观察"
		if maxSingleWin < 0.2*totalBet && profit < 0.15*totalBet && WinRateCoe < 1.0 {
			level = "正常"
		} else if WinRateCoe >= 1 {
			level = "警告"
		}
	case profit >= 10000 && profit < 30000:
		level = "警告"
		if maxSingleWin < 0.15*totalBet && profit < 0.1*totalBet && WinRateCoe < 1.0 {
			level = "持续观察"
		} else if WinRateCoe >= 1 {
			level = "危险"
		}
	case profit >= 30000:
		level = "危险"
		if maxSingleWin < 0.1*totalBet && profit < 0.05*totalBet && WinRateCoe < 1.0 {
			level = "警告"
		}
	}

	return level
}

func AccountLevelJudgment(profit, winRateCoefficient float64) string {
	var level string

	switch {
	case profit < 0:
		level = "安全"
		if winRateCoefficient > 0.7 {
			level = "正常"
		}
	case profit < 3000:
		level = "正常"
		if winRateCoefficient > 1.5 {
			level = "持续观察"
		}
	case profit >= 3000 && profit < 10000:
		level = "持续观察"
		if winRateCoefficient > 1.2 {
			level = "警告"
		}
	case profit >= 10000 && profit < 30000:
		level = "警告"
		if winRateCoefficient > 1.1 {
			level = "危险"
		}
	case profit >= 30000:
		level = "危险"
		if winRateCoefficient < 1.0 {
			level = "警告"
		}
	}

	return level
}
