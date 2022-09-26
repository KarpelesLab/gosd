package gosd

import (
	"net"
	"os"
)

type Handle struct {
	NotifySocket string
}

func New() *Handle {
	res := &Handle{
		NotifySocket: os.Getenv("NOTIFY_SOCKET"),
	}

	return res
}

func (h *Handle) Enabled() bool {
	return h.NotifySocket != ""
}

// Notify sends a message to the init daemon.
func (h *Handle) Notify(state string) error {
	if !h.Enabled() {
		return nil
	}

	socketAddr := &net.UnixAddr{
		Name: h.NotifySocket,
		Net:  "unixgram",
	}

	conn, err := net.DialUnix(socketAddr.Net, nil, socketAddr)
	if err != nil {
		return err
	}

	defer conn.Close()

	if _, err = conn.Write([]byte(state)); err != nil {
		return err
	}

	return nil
}

// Status passes a single-line UTF-8 status string back to the service manager that describes the service state
func (h *Handle) Status(status string) error {
	return h.Notify("STATUS=" + status)
}

// Ready should be called once the daemon initialization is done and it is ready to work
func (h *Handle) Ready() error {
	return h.Notify(ReadyState)
}
