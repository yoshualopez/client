// Auto-generated to Go types and interfaces using avdl-compiler v1.4.6 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/notify_perflogevents.avdl

package keybase1

import (
	"fmt"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
	"time"
)

type PerfLogEventType int

const (
	PerfLogEventType_LOAD_USER      PerfLogEventType = 0
	PerfLogEventType_LOAD_TEAM      PerfLogEventType = 1
	PerfLogEventType_LOAD_TEAM_FAST PerfLogEventType = 2
	PerfLogEventType_AUDIT_TEAM     PerfLogEventType = 3
	PerfLogEventType_BOX_AUDIT      PerfLogEventType = 4
)

func (o PerfLogEventType) DeepCopy() PerfLogEventType { return o }

var PerfLogEventTypeMap = map[string]PerfLogEventType{
	"LOAD_USER":      0,
	"LOAD_TEAM":      1,
	"LOAD_TEAM_FAST": 2,
	"AUDIT_TEAM":     3,
	"BOX_AUDIT":      4,
}

var PerfLogEventTypeRevMap = map[PerfLogEventType]string{
	0: "LOAD_USER",
	1: "LOAD_TEAM",
	2: "LOAD_TEAM_FAST",
	3: "AUDIT_TEAM",
	4: "BOX_AUDIT",
}

func (e PerfLogEventType) String() string {
	if v, ok := PerfLogEventTypeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type PerfLogEvent struct {
	EventType PerfLogEventType `codec:"eventType" json:"eventType"`
	Message   string           `codec:"message" json:"message"`
	Ctime     Time             `codec:"ctime" json:"ctime"`
}

func (o PerfLogEvent) DeepCopy() PerfLogEvent {
	return PerfLogEvent{
		EventType: o.EventType.DeepCopy(),
		Message:   o.Message,
		Ctime:     o.Ctime.DeepCopy(),
	}
}

type PerfLogEventUpdateArg struct {
	Event PerfLogEvent `codec:"event" json:"event"`
}

type NotifyPerfLogEventsInterface interface {
	PerfLogEventUpdate(context.Context, PerfLogEvent) error
}

func NotifyPerfLogEventsProtocol(i NotifyPerfLogEventsInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.NotifyPerfLogEvents",
		Methods: map[string]rpc.ServeHandlerDescription{
			"perfLogEventUpdate": {
				MakeArg: func() interface{} {
					var ret [1]PerfLogEventUpdateArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]PerfLogEventUpdateArg)
					if !ok {
						err = rpc.NewTypeError((*[1]PerfLogEventUpdateArg)(nil), args)
						return
					}
					err = i.PerfLogEventUpdate(ctx, typedArgs[0].Event)
					return
				},
			},
		},
	}
}

type NotifyPerfLogEventsClient struct {
	Cli rpc.GenericClient
}

func (c NotifyPerfLogEventsClient) PerfLogEventUpdate(ctx context.Context, event PerfLogEvent) (err error) {
	__arg := PerfLogEventUpdateArg{Event: event}
	err = c.Cli.Notify(ctx, "keybase.1.NotifyPerfLogEvents.perfLogEventUpdate", []interface{}{__arg}, 0*time.Millisecond)
	return
}
