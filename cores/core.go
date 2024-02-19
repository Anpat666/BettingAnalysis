package cores

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// 提取列数据
func GetExcelCols(f *excelize.File, tableName string) [][]string {
	rows, err := f.GetCols(tableName)
	if err != nil {
		fmt.Println("表格获取列元素失败：", tableName, err)
	}
	return rows
}

// 提取行数据
func GetExcelRows(f *excelize.File, tableName string) [][]string {
	rows, err := f.GetRows(tableName)
	if err != nil {
		fmt.Println("表格获取列元素失败：", tableName, err)
	}
	return rows
}

// 获取二数组中去重后的元素
func GetUniqueValuesInColumn(data [][]string, num int) []string {
	names := make([]string, 0)
	for i := 1; i < len(data); i++ {
		name := data[i][num]
		found := false
		for _, value := range names {
			if value == name {
				found = true
				break
			}
		}
		if !found {
			names = append(names, name)
		}
	}
	return names
}

// 二维数组中一维相同下标元素的和值
func SumColumnValues(data [][]string, num int) float64 {
	var amount float64
	for i := 0; i < len(data); i++ {
		bet, _ := strconv.ParseFloat(data[i][num], 64)
		amount += bet
	}
	return amount
}

func ColumeClass(data [][]string, gameNames []string) [][]string {
	gameClass := make([][]string, 0)
	for _, value := range gameNames {
		for i := 0; i < len(data); i++ {
			if data[i][4] == value {
				gameClass = append(gameClass, data[i])
			}
		}
	}
	return gameClass
}

func GetPlayerName(data []string) string {
	if len(data) == 1 {
		return data[0]
	}
	res := data[0]
	for i := 1; i < len(data); i++ {
		res = res + "、" + data[i]
	}
	return res
}

func GetThreeArrayAmount(data [][][]string, name string) int {
	for i := 0; i < len(data); i++ {
		for k := 0; k < len(data); k++ {
			if data[i][k][4] == name {
				return len(data[i][k]) - 1
			}
		}
	}
	fmt.Println("游戏名字错误")
	os.Exit(0)
	return -1
}

// 获取二维数组中相同一维下标的最大值
func GetMaxWin(data [][]string) float64 {
	var MaxWin float64 = 0
	for i := 0; i < len(data); i++ {
		valueF, _ := strconv.ParseFloat(data[i][16], 64)
		if MaxWin < valueF {
			MaxWin = valueF
		}
	}
	return MaxWin
}
