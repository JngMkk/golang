package main

type Attacker interface {
	Attack()
}

// 런타임 에러
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x457c16]
// 인터페이스의 기본값은 nil이다.
func main() {
	var att Attacker
	att.Attack()
}
