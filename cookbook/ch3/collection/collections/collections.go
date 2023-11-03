package collections

type WorkWith struct {
	Data    string
	Version int
}

// WorkWith의 리스트와 불리언을 리턴하는 WorkWith 함수를 입력받음
func Filter(ws []WorkWith, f func(w WorkWith) bool) []WorkWith {
	result := make([]WorkWith, 0)
	for _, w := range ws {
		if f(w) {
			result = append(result, w)
		}
	}
	return result
}

// WorkWith의 리스트와 WorkWith을 입력받아 수정된 WorkWith을 반환하는 함수를 인자로 받음
func Map(ws []WorkWith, f func(w WorkWith) WorkWith) []WorkWith {
	result := make([]WorkWith, len(ws))

	for pos, w := range ws {
		newW := f(w)
		result[pos] = newW
	}
	return result
}
