package cores

func DataSorting(data [][]string) {
	for i := 0; i < len(data); i++ {
		if StringJudgment("三不同", data[i][12]) {
			data[i][12] = "三不同"
		}
		if StringJudgment("二同号", data[i][12]) {
			data[i][12] = "二同号"
		}
		if StringJudgment("和值 小", data[i][12]) {
			data[i][12] = "和值 大小单双"
		}
		if StringJudgment("和值 大", data[i][12]) {
			data[i][12] = "和值 大小单双"
		}
		if StringJudgment("和值 单", data[i][12]) {
			data[i][12] = "和值 大小单双"
		}
		if StringJudgment("和值 双", data[i][12]) {
			data[i][12] = "和值 大小单双"
		}
		if StringJudgment("豹子 1点", data[i][12]) {
			data[i][12] = "豹子 点数"
		}
		if StringJudgment("豹子 2点", data[i][12]) {
			data[i][12] = "豹子 点数"
		}
		if StringJudgment("豹子 3点", data[i][12]) {
			data[i][12] = "豹子 点数"
		}
		if StringJudgment("豹子 4点", data[i][12]) {
			data[i][12] = "豹子 点数"
		}
		if StringJudgment("豹子 5点", data[i][12]) {
			data[i][12] = "豹子 点数"
		}
		if StringJudgment("豹子 6点", data[i][12]) {
			data[i][12] = "豹子 点数"
		}
		if StringJudgment("三军 1点", data[i][12]) {
			data[i][12] = "三军"
		}
		if StringJudgment("三军 2点", data[i][12]) {
			data[i][12] = "三军"
		}
		if StringJudgment("三军 3点", data[i][12]) {
			data[i][12] = "三军"
		}
		if StringJudgment("三军 4点", data[i][12]) {
			data[i][12] = "三军"
		}
		if StringJudgment("三军 5点", data[i][12]) {
			data[i][12] = "三军"
		}
		if StringJudgment("三军 6点", data[i][12]) {
			data[i][12] = "三军"
		}
		if StringJudgment("短牌 1点", data[i][12]) {
			data[i][12] = "短牌"
		}
		if StringJudgment("短牌 2点", data[i][12]) {
			data[i][12] = "短牌"
		}
		if StringJudgment("短牌 3点", data[i][12]) {
			data[i][12] = "短牌"
		}
		if StringJudgment("短牌 4点", data[i][12]) {
			data[i][12] = "短牌"
		}
		if StringJudgment("短牌 5点", data[i][12]) {
			data[i][12] = "短牌"
		}
		if StringJudgment("短牌 6点", data[i][12]) {
			data[i][12] = "短牌"
		}
		if StringJudgment("第一球", data[i][12]) {
			data[i][12] = "定位胆"
		}
		if StringJudgment("第二球", data[i][12]) {
			data[i][12] = "定位胆"
		}
		if StringJudgment("第三球", data[i][12]) {
			data[i][12] = "定位胆"
		}
		if StringJudgment("第四球", data[i][12]) {
			data[i][12] = "定位胆"
		}
		if StringJudgment("第五球", data[i][12]) {
			data[i][12] = "定位胆"
		}
		if StringJudgment("第六球", data[i][12]) {
			data[i][12] = "第六球"
		}
		if StringJudgment("第七球", data[i][12]) {
			data[i][12] = "第七球"
		}
		if StringJudgment("第八球", data[i][12]) {
			data[i][12] = "第八球"
		}
		if StringJudgment("第九球", data[i][12]) {
			data[i][12] = "第九球"
		}

	}
}
