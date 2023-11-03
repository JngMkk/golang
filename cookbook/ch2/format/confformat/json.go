package confformat

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSON 구조체 태그를 갖는 일반적인 데이터 구조체
type JSONData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// JSONData 구조체를 JSON 포맷의 bytes.Buffer로 반환
func (j *JSONData) ToJSON() (*bytes.Buffer, error) {
	d, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)
	return b, nil
}

// JSONData 구조체 디코딩
func (j *JSONData) Decode(data []byte) error {
	return json.Unmarshal(data, j)
}

func OtherJSONExamples() error {
	res := make(map[string]string)
	err := json.Unmarshal([]byte(`{"key": "value"}`), &res)
	if err != nil {
		return err
	}

	fmt.Println("We can unmarshal into a map instead of a struct :", res)

	b := bytes.NewReader([]byte(`{"key2": "value2"}`))
	decoder := json.NewDecoder(b)
	if err := decoder.Decode(&res); err != nil {
		return err
	}

	fmt.Println("we can also use decoders/encoders to work with streams :", res)

	return nil
}
