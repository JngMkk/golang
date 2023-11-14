package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// Movie 구조체는 파싱된 csv를 저장하는 데 사용됨
type Movie struct {
	Title    string
	Director string
	Year     int
}

// io.Reader에 전달된 csv를 처리하는 예제
func ReadCSV(b io.Reader) ([]Movie, error) {
	r := csv.NewReader(b)

	// 설정 옵션을 설정하는 선택적 작업
	r.Comma = ';'
	r.Comment = '-'
	var movies []Movie

	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}
		m := Movie{record[0], record[1], int(year)}
		movies = append(movies, m)
	}
	return movies, nil
}

// csv 파서 사용
// csv 패키지를 사용해 문자열을 가져온 후 버퍼로 변환해 읽는 예
func AddMoviesFromText() error {
	in := `
- first our headers
movie title;director;year released

- then some data
Guardians of the Galaxy Vol. 2;James Gumm;2017
Star Wars: Episode VIII;Rian Johnson;2017
`

	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", m)
	return nil
}
