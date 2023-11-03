package math

import "math/big"

// 피보나치 수열의 저장을 위한 글로벌 변수
var memoize map[int]*big.Int

func init() {
	// 맵 초기화
	memoize = make(map[int]*big.Int)
}

func Fib(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}

	if n < 2 {
		memoize[n] = big.NewInt(1)
	}

	if val, ok := memoize[n]; ok {
		return val
	}

	memoize[n] = big.NewInt(0)
	memoize[n].Add(memoize[n], Fib(n-1))
	memoize[n].Add(memoize[n], Fib(n-2))

	return memoize[n]
}
