package cores

import (
	"fmt"
	"strconv"
)

func ProPortion(dataTotal float64, data float64) string {
	res := data / dataTotal
	res *= 100
	strRes := strconv.FormatFloat(res, 'f', 2, 64)
	strRes = fmt.Sprintf("%s%%", strRes)
	return strRes

}
