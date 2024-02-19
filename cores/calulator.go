package cores

import (
	"BettingAnalysis/models"
	"fmt"
	"strconv"
	"strings"
)

func ProPortion(dataTotal float64, data float64) string {
	res := data / dataTotal
	res *= 100
	strRes := strconv.FormatFloat(res, 'f', 2, 64)
	strRes = fmt.Sprintf("%s%%", strRes)
	return strRes

}

// 胜率
func GetWinningRate(data [][]string) string {
	var win float64 = 0
	var lose float64 = 0
	for i := 0; i < len(data); i++ {
		if data[i][14] == "赢" {
			win += 1
		}
		if data[i][14] == "输" {
			lose += 1
		}
	}
	res := ProPortion(win+lose, win)
	return res

}

// 胜率系数
func GetWinRateCoe(data [][]string) float64 {
	var win float64 = 0
	var lose float64 = 0
	for i := 0; i < len(data); i++ {
		if data[i][14] == "赢" {
			win += 1
		}
		if data[i][14] == "输" {
			lose += 1
		}
	}
	rate := win / (win + lose)
	res := rate / GetGamesByRate(data)
	return res

}

// 游戏判定
func GetGamesByRate(data [][]string) float64 {
	gameName := data[0][4]
	baccarat := "百家乐"
	gameMethod := data[0][12]
	if strings.ContainsAny(gameName, baccarat) {
		return models.Baccarat[gameMethod]
	}
	return 0

}
