package confformat

import (
	"bytes"

	"github.com/go-yaml/yaml"
)

// YAML 구조체 태그를 갖는 일반적인 데이터 구조체
type YAMLData struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

// YAMLData 구조체를 YAML 포맷의 bytes.Buffer로 반환
func (y *YAMLData) ToYAML() (*bytes.Buffer, error) {
	d, err := yaml.Marshal(y)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)
	return b, nil
}

// YAMLData 구조체로 디코딩
func (y *YAMLData) Decode(data []byte) error {
	return yaml.Unmarshal(data, y)
}
