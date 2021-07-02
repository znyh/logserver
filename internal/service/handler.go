package service

import (
	"context"
	"encoding/json"

	"github.com/go-kratos/kratos/pkg/log"
	pb "github.com/znyh/middle-end/proto/logserver"
)

func (s *Service) Record(ctx context.Context, req *pb.RecordReq) (resp *pb.RecordRsp, err error) {
	resp = &pb.RecordRsp{}

	if len(req.Messages) == 0 {
		log.Errorc(ctx, "Record LogMsg is nil.")
		return
	}

	_data, err := json.Marshal(req.Messages)
	if err != nil {
		log.Errorc(ctx, "Record json Marshal is fail. err:%+v", err)
		return
	}

	if len(_data) == 0 {
		log.Errorc(ctx, "Record _data len is 0")
		return
	}

	err = s.dao.PubKafka(pb.LogServerTopic, _data)
	if err != nil {
		log.Errorc(ctx, "Record PubKafka is fail. err:%+v", err)
	}

	return
}
