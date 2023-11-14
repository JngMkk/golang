package tmpFiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 임시 파일 및 디렉터리 작업을 하는 기본적인 패턴을 보여줌
func WorkWithTemp() error {

	// 예를 들어 template1-10.thml과 동일한 이름을 가진 파일을 저장할 임시 공간이 필요한 경우 임시 디렉터리가 좋은 방법
	// 첫 번째 인자는 공백인데, 이는 os.Tempdir()이 반환한 위치에 디렉터리를 생성한다는 것을 의미함
	// os.MkdirTemp("", "tmp")
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}
	// 함수가 종료될 때 임시 파일의 모든 내용을 삭제함
	// 따라서 반드시 디렉터리 이름을 호출 함수에 반환해야 함
	defer os.RemoveAll(t)

	// 임시 파일을 생성하기 위해서는 디렉터리가 반드시 존재해야 함
	// t는 *os.File 객체임
	// os.CreateTemp(t, "tmp")
	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}

	fmt.Println(tf.Name())

	// 일반적으로 여기에서 임시 파일을 삭제함. 하지만 임시 디렉터리에 파일을 저장했기 때문에 앞서 호출된 defer 구문에 의해 정리됨

	return nil
}
