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
	gameMethod := data[0][12]

	if strings.ContainsAny(gameName, "百家乐") {
		return models.Baccarat[gameMethod]
	}

	if strings.ContainsAny(gameName, "赛车") {
		return models.RacingTen[gameMethod]
	}
	if strings.ContainsAny(gameName, "飞艇") {
		return models.RacingTen[gameMethod]
	}

	if strings.ContainsAny(gameName, "澳洲幸运10") {
		return models.RacingTen[gameMethod]
	}

	if strings.ContainsAny(gameName, "时时彩") {
		return models.ColorfulFive[gameMethod]
	}

	if strings.ContainsAny(gameName, "快三") {
		for k, v := range models.FastThree {
			if StringJudgment(gameMethod, k) {
				return v
			}
		}
	}
	return 0

}

// 判断两个字符str1是否都存在于str2里面
func StringJudgment(str1 string, str2 string) bool {

	for _, v := range str1 {
		found := false
		for _, v2 := range str2 {
			if v == v2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
