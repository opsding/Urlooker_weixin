package receiver

import (
	"time"

	"github.com/urlooker/alarm/judge"
	"github.com/urlooker/web/model"
)

type Alarm int

func (this *Alarm) Ping(req interface{}, reply *string) error {
	*reply = "ok"
	return nil
}

func (this *Alarm) Send(args []*model.ItemStatus, reply *string) error {
	// 把当前时间的计算放在最外层，是为了减少获取时间时的系统调用开销
	now := time.Now().Unix()

	for _, arg := range args {
		pk := arg.PK()
		judge.HistoryBigMap[pk[0:2]].PushFrontAndMaintain(pk, arg, 10, now)
	}
	*reply = ""
	return nil
}
