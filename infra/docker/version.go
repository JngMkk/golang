package docker

import (
	"encoding/json"
	"net/http"
	"time"
)

// 빌드 시점에 전달되는 값들을 저장
type Versioninfo struct {
	Version   string
	BuildDate time.Time
	Uptime    time.Duration
}

// 최신 버전 정보를 씀
func VersionHandler(v *Versioninfo) http.HandlerFunc {
	t := time.Now()
	return func(w http.ResponseWriter, r *http.Request) {
		v.Uptime = time.Since(t)
		vers, err := json.Marshal(v)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(vers)
	}
}
