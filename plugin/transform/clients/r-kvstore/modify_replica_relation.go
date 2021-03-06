package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifyReplicaRelationRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyReplicaRelationRequest) Invoke(client goaliyun.Client) (*ModifyReplicaRelationResponse, error) {
	resp := &ModifyReplicaRelationResponse{}
	err := client.Request("r-kvstore", "ModifyReplicaRelation", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyReplicaRelationResponse struct {
	RequestId goaliyun.String
}
