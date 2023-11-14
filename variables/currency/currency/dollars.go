package currency

import (
	"errors"
	"strconv"
	"strings"
)

func ConvertStringDollarsToPennies(amount string) (int64, error) {
	// amount 인자가 float로 변환 가능한지 확인
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	groups := strings.Split(amount, ".")

	res := groups[0]

	r := ""

	if len(groups) == 2 {
		if len(groups[1]) != 2 {
			return 0, errors.New("invalid cents")
		}
		r = groups[1]
	}

	for len(r) < 2 {
		r += "0"
	}

	res += r

	return strconv.ParseInt(res, 10, 64)
}
