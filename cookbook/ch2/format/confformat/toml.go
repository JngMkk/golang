package confformat

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

// TOML 구조체 태그를 갖는 일반적인 데이터 구조체
type TOMLData struct {
	Name string `toml:"name"`
	Age  int    `toml:"age"`
}

// TOMLData 구조체를 TOML 포맷의 bytes.Buffer로 반환
func (t *TOMLData) ToTOML() (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	encoder := toml.NewEncoder(b)
	if err := encoder.Encode(t); err != nil {
		return nil, err
	}

	return b, nil
}

// TOMLData 구조체로 디코딩
func (t *TOMLData) Decode(data []byte) (toml.MetaData, error) {
	return toml.Decode(string(data), t)
}
