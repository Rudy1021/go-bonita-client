package main

import (
	"fmt"

	bpm "github.com/Rudy1021/go-bonita-client/bpm"
)

func main() {
	client := bpm.Bc
	client.Login("isabelle_wu")
	tm := []int{44}
	body := client.StartB2Form(44, tm)
	fmt.Println(body)

	// res := client.GetReadyCase("50", "ready", "38")
	// fmt.Println(res)

	// res = client.GetDetailCase("3075")
	// fmt.Println(res)

	//res = client.GetFinishCase("3075")
	//fmt.Println(res)

	//res = client.GetFinishCaseState("3075")
	//fmt.Println(res)

	/*res = client.ReviewCase("", `{
		"modelInput":{
			"gmApprovalStatus":最高主管審核狀態 : boolean
		}
	}`)
	fmt.Println(res)

	res = client.ReviewCase("", `{
	  "modelInput":{
	  	"dmApprovalStatus":部門主管審核狀態 : boolean
	  }
	}`)
	fmt.Println(res)
	*/
	// res = client.GetAllProcessCase("20", "8759976868088592450")
	// fmt.Println(res)
}
