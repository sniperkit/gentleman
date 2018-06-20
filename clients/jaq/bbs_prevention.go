package jaq

import (
	"github.com/morlay/goaliyun"
)

type BbsPreventionRequest struct {
	CallerName      string `position:"Query" name:"CallerName"`
	Ip              string `position:"Query" name:"Ip"`
	ProtocolVersion string `position:"Query" name:"ProtocolVersion"`
	Source          int64  `position:"Query" name:"Source"`
	PhoneNumber     string `position:"Query" name:"PhoneNumber"`
	Email           string `position:"Query" name:"Email"`
	UserId          string `position:"Query" name:"UserId"`
	IdType          int64  `position:"Query" name:"IdType"`
	CurrentUrl      string `position:"Query" name:"CurrentUrl"`
	Agent           string `position:"Query" name:"Agent"`
	Cookie          string `position:"Query" name:"Cookie"`
	SessionId       string `position:"Query" name:"SessionId"`
	MacAddress      string `position:"Query" name:"MacAddress"`
	Referer         string `position:"Query" name:"Referer"`
	NickName        string `position:"Query" name:"NickName"`
	CompanyName     string `position:"Query" name:"CompanyName"`
	Address         string `position:"Query" name:"Address"`
	JsToken         string `position:"Query" name:"JsToken"`
	SDKToken        string `position:"Query" name:"SDKToken"`
	ExtendData      string `position:"Query" name:"ExtendData"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *BbsPreventionRequest) Invoke(client goaliyun.Client) (*BbsPreventionResponse, error) {
	resp := &BbsPreventionResponse{}
	err := client.Request("jaq", "BbsPrevention", "2016-11-23", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BbsPreventionResponse struct {
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Data      BbsPreventionData
}

type BbsPreventionData struct {
	FnalDecision     goaliyun.Integer
	EventId          goaliyun.String
	UserId           goaliyun.String
	FinalScore       goaliyun.Integer
	FinalDesc        goaliyun.String
	Detail           goaliyun.String
	CaptchaCheckData goaliyun.String
}
