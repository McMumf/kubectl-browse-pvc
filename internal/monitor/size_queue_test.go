package monitor

import (
	"testing"

	"k8s.io/client-go/tools/remotecommand"
)

// Verifies that Next() reads a TerminalSize value from ResizeChan.
func TestNextReturnsSize(t *testing.T) {
	q := &SizeQueue{
		ResizeChan:   make(chan remotecommand.TerminalSize, 1),
		StopResizing: make(chan struct{}),
	}

	expected := remotecommand.TerminalSize{
		Width:  80,
		Height: 24,
	}

	q.ResizeChan <- expected

	actual := q.Next()
	if actual == nil {
		t.Fatal("expected size, got nil")
	}

	if actual.Width != expected.Width {
		t.Fatalf("expected width %d, got %d", expected.Width, actual.Width)
	}

	if actual.Height != expected.Height {
		t.Fatalf("expected height %d, got %d", expected.Height, actual.Height)
	}
}

// Verifies that Next() returns nil when ResizeChan has been closed.
func TestNextWhenChannelClosed(t *testing.T) {
	q := &SizeQueue{
		ResizeChan:   make(chan remotecommand.TerminalSize),
		StopResizing: make(chan struct{}),
	}

	close(q.ResizeChan)

	actual := q.Next()
	if actual != nil {
		t.Fatalf("expected nil, got %+v", actual)
	}
}

// Verifies that Stop() closes the StopResizing channel
func TestStopClosesChannel(t *testing.T) {
	q := &SizeQueue{
		ResizeChan:   make(chan remotecommand.TerminalSize),
		StopResizing: make(chan struct{}),
	}

	q.Stop()

	select {
	case <-q.StopResizing:
		// expected: channel is closed
	default:
		t.Fatal("StopResizing channel was not closed")
	}
}
