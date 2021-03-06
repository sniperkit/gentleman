package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyUserGroupsRequest struct {
	ClusterId string                    `position:"Query" name:"ClusterId"`
	Users     *ModifyUserGroupsUserList `position:"Query" type:"Repeated" name:"User"`
	RegionId  string                    `position:"Query" name:"RegionId"`
}

func (req *ModifyUserGroupsRequest) Invoke(client goaliyun.Client) (*ModifyUserGroupsResponse, error) {
	resp := &ModifyUserGroupsResponse{}
	err := client.Request("ehpc", "ModifyUserGroups", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyUserGroupsUser struct {
	Name  string `name:"Name"`
	Group string `name:"Group"`
}

type ModifyUserGroupsResponse struct {
	RequestId goaliyun.String
}

type ModifyUserGroupsUserList []ModifyUserGroupsUser

func (list *ModifyUserGroupsUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyUserGroupsUser)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
