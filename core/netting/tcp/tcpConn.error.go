package tcp

import (
	"errors"
)

var (
	ConnErr_ConnIsNotWorking      = errors.New("connect is not working ")
	ConnErr_CloseSendTimeout      = errors.New("CloseSendTimeout")
	ConnErr_CloseShutdown         = errors.New("CloseShutdown")
	ConnErr_CloseHeartLoopTimeout = errors.New("CloseHeartLoopTimeout")
	ConnErr_CloseNormal           = errors.New("CloseNormal")
)
