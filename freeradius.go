package freeradius

import (
	"net"
	"time"
)

type LogType int

const (
	LogTypeAuth       LogType = 2
	LogTypeInfo       LogType = 3
	LogTypeError      LogType = 4
	LogTypeWarn       LogType = 5
	LogTypeProxy      LogType = 6
	LogTypeAccounting LogType = 7
)

type RlmCode int

const (
	RlmCodeReject RlmCode = iota
	RlmCodeFail
	RlmCodeOk
	RlmCodeHandled
	RlmCodeInvalid
	RlmCodeUserlock
	RlmCodeNotFound
	RlmCodeNoop
	RlmCodeUpdated
	RlmCodeNumCodes
	RlmCodeUnknown
)

type Log interface {
	Radlog(LogType, string, ...interface{}) int
	Info(format string, args ...interface{})
}

type Packet interface {
	AddValuePair(attribute, value string)
	Code() uint
	Id() int
	Timestamp() time.Time
	DestinationIp() net.IP
	SourceIp() net.IP
	DestinationPort() uint16
	SourcePort() uint16
}

type Request interface {
	Reply() Packet
	Packet() Packet
}

// Needs to be exported as CreateModule
type ModuleFunc func() Module

type Module interface {
	Init(logger Log) error
	Authorize(req Request) RlmCode
}
