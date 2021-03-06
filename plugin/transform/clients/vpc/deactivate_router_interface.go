package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeactivateRouterInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeactivateRouterInterfaceRequest) Invoke(client goaliyun.Client) (*DeactivateRouterInterfaceResponse, error) {
	resp := &DeactivateRouterInterfaceResponse{}
	err := client.Request("vpc", "DeactivateRouterInterface", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeactivateRouterInterfaceResponse struct {
	RequestId goaliyun.String
}
