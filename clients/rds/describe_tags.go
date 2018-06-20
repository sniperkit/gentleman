package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeTagsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ProxyId              string `position:"Query" name:"ProxyId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeTagsRequest) Invoke(client goaliyun.Client) (*DescribeTagsResponse, error) {
	resp := &DescribeTagsResponse{}
	err := client.Request("rds", "DescribeTags", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTagsResponse struct {
	RequestId goaliyun.String
	Items     DescribeTagsTagInfosList
}

type DescribeTagsTagInfos struct {
	TagKey        goaliyun.String
	TagValue      goaliyun.String
	DBInstanceIds DescribeTagsDBInstanceIdList
}

type DescribeTagsTagInfosList []DescribeTagsTagInfos

func (list *DescribeTagsTagInfosList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTagsTagInfos)
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

type DescribeTagsDBInstanceIdList []goaliyun.String

func (list *DescribeTagsDBInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
