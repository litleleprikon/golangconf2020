package ex03

//go:generate $GOPATH/bin/mockgen -source ./log.go // HL0001
type Logger interface {
	Print(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}
