package monitor

import "k8s.io/client-go/tools/remotecommand"

type SizeQueue struct {
	ResizeChan   chan remotecommand.TerminalSize
	StopResizing chan struct{}
}

func (s *SizeQueue) Next() *remotecommand.TerminalSize {
	size, ok := <-s.ResizeChan
	if !ok {
		return nil
	}
	return &size
}

func (s *SizeQueue) Stop() {
	close(s.StopResizing)
}
