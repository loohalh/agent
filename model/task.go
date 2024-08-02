package model

const (
	_ = iota
	TaskTypeHTTPGet
	TaskTypeICMPPing
	TaskTypeTCPPing
	TaskTypeCommand
	TaskTypeTerminal
	TaskTypeUpgrade
	TaskTypeKeepalive
	TaskTypeTerminalGRPC
	TaskTypeNAT
	TaskTypeReportHostInfo
)

type TerminalTask struct {
	StreamID string
}

type TaskNAT struct {
	StreamID string
	Host     string
}
