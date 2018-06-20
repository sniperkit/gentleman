package rds

import (
	"github.com/morlay/goaliyun"
)

type PurgeDBInstanceLogRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *PurgeDBInstanceLogRequest) Invoke(client goaliyun.Client) (*PurgeDBInstanceLogResponse, error) {
	resp := &PurgeDBInstanceLogResponse{}
	err := client.Request("rds", "PurgeDBInstanceLog", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PurgeDBInstanceLogResponse struct {
	RequestId goaliyun.String
}
