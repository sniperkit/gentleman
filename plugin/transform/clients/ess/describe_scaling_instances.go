package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeScalingInstancesRequest struct {
	InstanceId10           string `position:"Query" name:"InstanceId.10"`
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId12           string `position:"Query" name:"InstanceId.12"`
	InstanceId11           string `position:"Query" name:"InstanceId.11"`
	ScalingGroupId         string `position:"Query" name:"ScalingGroupId"`
	LifecycleState         string `position:"Query" name:"LifecycleState"`
	CreationType           string `position:"Query" name:"CreationType"`
	PageNumber             int64  `position:"Query" name:"PageNumber"`
	PageSize               int64  `position:"Query" name:"PageSize"`
	InstanceId20           string `position:"Query" name:"InstanceId.20"`
	InstanceId1            string `position:"Query" name:"InstanceId.1"`
	InstanceId3            string `position:"Query" name:"InstanceId.3"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceId2            string `position:"Query" name:"InstanceId.2"`
	InstanceId5            string `position:"Query" name:"InstanceId.5"`
	InstanceId4            string `position:"Query" name:"InstanceId.4"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	InstanceId7            string `position:"Query" name:"InstanceId.7"`
	InstanceId6            string `position:"Query" name:"InstanceId.6"`
	InstanceId9            string `position:"Query" name:"InstanceId.9"`
	InstanceId8            string `position:"Query" name:"InstanceId.8"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	ScalingConfigurationId string `position:"Query" name:"ScalingConfigurationId"`
	HealthStatus           string `position:"Query" name:"HealthStatus"`
	InstanceId18           string `position:"Query" name:"InstanceId.18"`
	InstanceId17           string `position:"Query" name:"InstanceId.17"`
	InstanceId19           string `position:"Query" name:"InstanceId.19"`
	InstanceId14           string `position:"Query" name:"InstanceId.14"`
	InstanceId13           string `position:"Query" name:"InstanceId.13"`
	InstanceId16           string `position:"Query" name:"InstanceId.16"`
	InstanceId15           string `position:"Query" name:"InstanceId.15"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *DescribeScalingInstancesRequest) Invoke(client goaliyun.Client) (*DescribeScalingInstancesResponse, error) {
	resp := &DescribeScalingInstancesResponse{}
	err := client.Request("ess", "DescribeScalingInstances", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeScalingInstancesResponse struct {
	TotalCount       goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageSize         goaliyun.Integer
	RequestId        goaliyun.String
	ScalingInstances DescribeScalingInstancesScalingInstanceList
}

type DescribeScalingInstancesScalingInstance struct {
	InstanceId             goaliyun.String
	ScalingConfigurationId goaliyun.String
	ScalingGroupId         goaliyun.String
	HealthStatus           goaliyun.String
	LoadBalancerWeight     goaliyun.Integer
	LifecycleState         goaliyun.String
	CreationTime           goaliyun.String
	CreationType           goaliyun.String
}

type DescribeScalingInstancesScalingInstanceList []DescribeScalingInstancesScalingInstance

func (list *DescribeScalingInstancesScalingInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingInstancesScalingInstance)
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
