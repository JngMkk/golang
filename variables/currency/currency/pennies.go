package currency

import "strconv"

func ConvertPenniestoDollarString(amount int64) string {
	res := strconv.FormatInt(amount, 10)

	negative := false
	if res[0] == '-' {
		res = res[1:]
		negative = true
	}

	for len(res) < 3 {
		res = "0" + res
	}

	l := len(res)
	res = res[0:l-2] + "." + res[l-2:]

	if negative {
		res = "-" + res
	}

	return res
}
