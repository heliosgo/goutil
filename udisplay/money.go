package udisplay

import (
	"errors"
	"goutil/umath"
	"goutil/utype"
	"regexp"
	"strconv"
	"strings"
)

func UpperRMB[T utype.Number | string](amount T) (string, error) {
	var str string
	switch a := any(amount).(type) {
	case int:
		str = strconv.FormatInt(int64(a), 10)
	case int8:
		str = strconv.FormatInt(int64(a), 10)
	case int16:
		str = strconv.FormatInt(int64(a), 10)
	case int32:
		str = strconv.FormatInt(int64(a), 10)
	case int64:
		str = strconv.FormatInt(int64(a), 10)
	case uint:
		str = strconv.FormatUint(uint64(a), 10)
	case uint8:
		str = strconv.FormatUint(uint64(a), 10)
	case uint16:
		str = strconv.FormatUint(uint64(a), 10)
	case uint32:
		str = strconv.FormatUint(uint64(a), 10)
	case uint64:
		str = strconv.FormatUint(uint64(a), 10)
	case float32:
		str = strconv.FormatFloat(float64(a), 'f', 2, 64)
	case float64:
		str = strconv.FormatFloat(float64(a), 'f', 2, 64)
	case string:
		str = a
	}

	return upperRMB(str)
}

func formatInt[T utype.Int](x T) string {
	return strconv.FormatInt(int64(x), 10)
}

func upperRMB(amount string) (string, error) {
	matched, err := regexp.MatchString(
		`^¥?([0-9]+|[0-9]{1,3}(,[0-9]{3})*)(.[0-9]+)?$`,
		amount,
	)
	if err != nil {
		return amount, err
	}
	if !matched {
		return amount, errors.New("invliad amount")
	}

	if len(amount) >= 2 && amount[:2] == "¥" {
		amount = amount[2:]
	}
	amount = strings.ReplaceAll(amount, ",", "")
	sli := strings.Split(amount, ".")
	num := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	unit := []string{"仟", "佰", "拾", ""}
	levelUnit := []string{"万", "亿", "万亿"}
	suffix := func(str string) bool {
		for _, v := range levelUnit {
			if strings.HasSuffix(str, v) {
				return true
			}
		}

		return false
	}

	var dfs func(string, int) string
	dfs = func(str string, level int) string {
		if len(str) == 0 {
			return ""
		}
		index := umath.Max(0, len(str)-len(unit))
		var res string
		curStr := str[index:]
		for i := 0; i < len(curStr); i++ {
			if curStr[i] == '0' {
				for i < len(curStr) && curStr[i] == '0' {
					i++
				}
				if i < len(curStr) {
					res += num[0]
				}
			}
			if i < len(curStr) {
				res += num[curStr[i]-'0'] + unit[len(unit)-len(curStr):][i]
			}
		}

		pre := dfs(str[:index], level+1)
		if pre != "" && !suffix(pre) && level < len(levelUnit) {
			pre += levelUnit[level]
		}

		return pre + res
	}
	res := dfs(sli[0], 0)
	if res == "" {
		res = num[0]
	}
	res += "元"
	if len(sli) < 2 {
		return res + "整", nil
	}
	if len(sli[1]) < 2 {
		sli[1] += "0"
	}
	if len(sli[1]) > 2 {
		sli[1] = sli[1][:2]
	}
	if sli[1] == "" || sli[1] == "0" || sli[1] == "00" {
		return res + "整", nil
	}

	res += num[sli[1][0]-'0'] + "角"
	if len(sli[1]) > 1 {
		res += num[sli[1][1]-'0'] + "分"
	}

	return res, nil
}
