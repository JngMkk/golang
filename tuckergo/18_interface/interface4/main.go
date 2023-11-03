// 서로 다른 인터페이스를 맞춰주기 위해 어댑터 패턴 사용 예시
package main

type Database interface {
	Get()
	Set()
}

type CDatabase struct {
}

func (c CDatabase) GetData() {

}

func (c CDatabase) SetData() {

}

// 어댑터
type Wrapper struct {
	cdb CDatabase
}

func (w Wrapper) Get() {
	w.cdb.GetData()
}

func (w Wrapper) Set() {
	w.cdb.SetData()
}

func main() {
	var cdatabase CDatabase
	var database Database
	database = Wrapper{cdatabase}
}
