package speedtest

import (
	"github.com/showwin/speedtest-go/speedtest"
	"log"
)

type TestResult struct {
	DownloadSpeed float64
	UploadSpeed   float64
}

// Test 테스트 수행 후 채널을 통해 결과를 전달
func Test(testChannel chan TestResult) {
	result := TestResult{}

	servers, err := speedtest.FetchServers()
	if err != nil {
		log.Fatalf("Error fetching server list: %v", err)
	}

	targets, err := servers.FindServer([]int{})
	if err != nil {
		log.Fatalf("Error finding server: %v", err)
	}

	for _, s := range targets {
		err := s.DownloadTest()
		if err != nil {
			result.DownloadSpeed = 0.0
		}

		result.DownloadSpeed = float64(s.DLSpeed)

		err = s.UploadTest()
		if err != nil {
			result.UploadSpeed = 0.0
		}

		result.UploadSpeed = float64(s.ULSpeed)
	}

	testChannel <- result
}
