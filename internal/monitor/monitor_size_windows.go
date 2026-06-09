//go:build windows
// +build windows

package monitor

import (
	"os"

	"golang.org/x/term"
	"k8s.io/client-go/tools/remotecommand"
)

func (s *SizeQueue) MonitorSize() {
	sigCh := make(chan os.Signal, 1)
	// Need to fix this to get it working on windows
	//signal.Notify(sigCh, syscall.SIGWINCH)

	for {
		select {
		case <-sigCh:
			width, height, err := term.GetSize(int(os.Stdout.Fd()))
			if err == nil {
				select {
				case s.ResizeChan <- remotecommand.TerminalSize{Width: uint16(width), Height: uint16(height)}:
				default:
				}
			}
		case <-s.StopResizing:
			close(s.ResizeChan)
			return
		}
	}
}
