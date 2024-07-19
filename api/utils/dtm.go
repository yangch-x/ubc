package utils

import (
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-gozero"
	"google.golang.org/protobuf/proto"
)

type MsgReq struct {
	Action     string
	Compensate string
	Message    proto.Message
}

// Msg msg类型
func Msg(dtmServer string, group []MsgReq) error {
	gid := dtmgrpc.MustGenGid(dtmServer)
	ms := dtmgrpc.NewMsgGrpc(dtmServer, gid)

	if len(group) == 0 {
		return nil
	}
	for _, g := range group {
		ms.Add(g.Action, g.Message)
	}
	return ms.Submit()
}

// Saga Saga 类型
func Saga(dtmServer string, group []MsgReq) error {
	gid := dtmgrpc.MustGenGid(dtmServer)
	s := dtmcli.NewSaga(dtmServer, gid)

	if len(group) == 0 {
		return nil
	}
	for _, g := range group {
		s.Add(g.Action, g.Compensate, g.Message)
	}
	return s.Submit()
}
