package filedirs

import (
	"bytes"
	"io"
	"os"
	"strings"
)

// 파일열 열고 내용을 읽은 다음 읽은 내용을 두 번째 파일에 씀
func Capitalizer(f1, f2 *os.File) error {
	_, err := f1.Seek(0, io.SeekStart)
	CheckErr(err)

	var tmp = new(bytes.Buffer)
	_, err2 := io.Copy(tmp, f1)
	CheckErr(err2)

	s := strings.ToUpper(tmp.String())
	_, err3 := io.Copy(f2, strings.NewReader(s))
	CheckErr(err3)

	err4 := f1.Close()
	CheckErr(err4)

	err5 := f2.Close()
	CheckErr(err5)

	return nil
}

// 두 파일을 생성하고 두 파일 모두 Capitalizer() 함수를 호출
func CaplitalizerExample() error {
	f1, err := os.Create("file1.txt")
	CheckErr(err)

	_, err2 := f1.Write([]byte(`this file contains a number of words and
	new lines`))
	CheckErr(err2)

	f2, err3 := os.Create("file2.txt")
	CheckErr(err3)

	err4 := Capitalizer(f1, f2)
	CheckErr(err4)

	err5 := os.Remove("file1.txt")
	CheckErr(err5)

	err6 := os.Remove("file2.txt")
	CheckErr(err6)

	return nil
}
