package filedirs

import (
	"errors"
	"io"
	"os"
)

func CheckErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func Operate() error {

	// 0755는 Chown명령과 비슷
	// ./ex_dir 디렉터리 생성
	err := os.Mkdir("ex_dir", os.FileMode(0755))
	CheckErr(err)

	// ./ex_dir로 이동
	err2 := os.Chdir("ex_dir")
	CheckErr(err2)

	f, err3 := os.Create("test.txt")
	CheckErr(err3)

	// 파일에 value를 쓰고 쓰기가 제대로 완료됐는지 검증
	value := []byte("hello")
	count, err4 := f.Write(value)
	CheckErr(err4)
	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}

	err5 := f.Close()
	CheckErr(err5)

	f, err3 = os.Open("test.txt")
	CheckErr(err3)

	io.Copy(os.Stdout, f)

	err6 := f.Close()
	CheckErr(err6)

	// /tmp로 이동
	err7 := os.Chdir("..")
	CheckErr(err7)

	// cleanup, os.RemoveAll은 잘못된 디렉터리를 가리키거나 사용자 입력을 사용하는 경우에 위험할 수 있음
	// 특히 root에서 실행하는 경우 위험할 수 있음
	err8 := os.RemoveAll("ex_dir")
	CheckErr(err8)

	return nil
}
