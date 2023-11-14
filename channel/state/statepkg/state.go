package statepkg

type op string

const (
	Add op = "add"
	Sub op = "sub"
	Mul op = "mul"
	Div op = "div"
)

// 두 값에 대한 연산을 처리함
type WorkRequest struct {
	Operation op
	Value1    int64
	Value2    int64
}

// 연산의 결과와 발생한 오류 반환
type WorkResponse struct {
	Wr     *WorkRequest
	Result int64
	Err    error
}
