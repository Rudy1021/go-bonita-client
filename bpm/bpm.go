package bpm

import (
	"log"

	"github.com/go-resty/resty/v2"
)

var Bc *BPMClient

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
	Bc = &BPMClient{
		server:   server_addr,
		token:    "",
		username: "",
		// password: "123456",
		client: resty.New(),
	}
}

// Login
// /bonita/loginservice
func (b *BPMClient) Login(username string) {

	url := b.server + "loginservice"
	resp, err := b.client.R().
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
			b.token = cookie.Value
		}
	}
}

// Start-Form
// /bonita/API/bpm/process/[ProcessId]/instantiation
// [ProcessId] == 表單編號
// return caseId
func (b *BPMClient) StartForm(processID string, body string) string {

	url := b.server + "API/bpm/process/" + processID + "/instantiation"

	resp, err := b.client.R().
		SetHeaders(map[string]string{
			"Content-Type":       "application/json",
			"X-Bonita-API-Token": b.token,
		}).
		SetBody(body).
		Post(url)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Body())
}

func (b *BPMClient) GetReadyCase(c string, state string, user_id string) string {

	url := b.server + "API/bpm/humanTask?c=" + c + "&f=state=" + state + "&f=user_id=" + user_id
	resp, err := b.client.R().
		Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Body())
}

func (b *BPMClient) GetDetailCase(case_id string) string {

	url := b.server + "API/bpm/humanTask?f=caseId=" + case_id
	resp, err := b.client.R().
		Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Body())
}

func (b *BPMClient) GetFinishCase(case_id string) string {
	//
	url := b.server + "API/bpm/archivedTask?f=caseId=" + case_id
	resp, err := b.client.R().
		Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Body())
}

func (b *BPMClient) GetFinishCaseState(sourceObjectId string) string {
	//
	url := b.server + "API/bpm/archivedHumanTask?f=sourceObjectId=" + sourceObjectId
	resp, err := b.client.R().
		Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Body())
}

func (b *BPMClient) GetAllProcessCase(c string, processId string) string {
	//
	url := b.server + "API/bpm/case?c=" + c + "&f=processDefinitionId=" + processId
	resp, err := b.client.R().
		Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Body())
}

func (b *BPMClient) ReviewCase(task_id string, body string) string {
	//
	url := b.server + "API/bpm/userTask/" + task_id + "/execution?assign=true"
	resp, err := b.client.R().
		SetHeaders(map[string]string{
			"Content-Type":       "application/json",
			"X-Bonita-API-Token": b.token,
		}).
		SetBody(body).
		Post(url)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Body())
}

//userTask/(TaskId)/execution?assign=true
