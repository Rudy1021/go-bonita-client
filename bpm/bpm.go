package bpm

import (
	"log"

	"github.com/go-resty/resty/v2"
)

var bc *BPMClient

type BPMClient struct {
	server   string
	token    string
	username string
	password string
	client   *resty.Client
}

func init() {
	const server_addr = "http://54.169.182.165:8080" + "/bonita/"
	// sources := fmt.Sprintf(server_addr,
	// 	// os.Getenv("BPM_SERVER_ADDR"),
	// 	os.Getenv("b.server"),
	// )
	bc = &BPMClient{
		server:   server_addr,
		token:    "",
		username: "",
		// password: "123456",
		client: resty.New(),
	}
}

// Login
// /bonita/loginservice
func Login(username string) {

	url := bc.server + "loginservice"
	resp, err := bc.client.R().
		SetFormData(map[string]string{
			"username": username,
			"password": "12345",
		}).
		Post(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "X-Bonita-API-Token" {
			bc.token = cookie.Value
		}
	}
}

// Start-Form
// /bonita/API/bpm/process/[ProcessId]/instantiation
// [ProcessId] == 表單編號
// return caseId
func StartForm(processID string, body string) string {

	url := bc.server + "API/bpm/process/" + processID + "/instantiation"

	resp, err := bc.client.R().
		SetHeaders(map[string]string{
			"Content-Type":       "application/json",
			"X-Bonita-API-Token": bc.token,
		}).
		SetBody(body).
		Post(url)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Body())
}
