package main

import (
	"docker"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// 빌드 시점에 설정되는 변수들
var (
	version   string
	buildDate string
)

var versionInfo docker.VersionInfo

func init() {
	// 빌드 시간에 설정
	versionInfo.Version = version
	intDate, err := strconv.ParseInt(buildDate, 10, 64)
	if err != nil {
		panic(err)
	}

	unixTime := time.Unix(intDate, 0)
	versionInfo.BuildDate = unixTime
}

func main() {
	http.HandleFunc("/version", docker.VersionHandler(&versionInfo))
	fmt.Println(http.ListenAndServe(":4000", nil))
}
