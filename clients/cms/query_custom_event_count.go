package cms

import (
	"github.com/morlay/goaliyun"
)

type QueryCustomEventCountRequest struct {
	QueryJson string `position:"Query" name:"QueryJson"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *QueryCustomEventCountRequest) Invoke(client goaliyun.Client) (*QueryCustomEventCountResponse, error) {
	resp := &QueryCustomEventCountResponse{}
	err := client.Request("cms", "QueryCustomEventCount", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryCustomEventCountResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
	RequestId goaliyun.String
	Success   bool
}
