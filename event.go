package zlog

type Event interface {
	Enabled() bool
	Msg(msg string)
	Str(key string, val string) Event
	Int(key string, val int) Event
	Error(e error) Event
	Timestamp() Event
}
