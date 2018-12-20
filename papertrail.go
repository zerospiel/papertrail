package papertrail

import (
	"encoding/json"
	"net"
	"sync"

	"github.com/apex/log"
)

// Handler implementation struct for apex/log
type Handler struct {
	mu   sync.Mutex
	enc  *json.Encoder
	conn net.Conn
}

// New creates new Handler pointer to use it within apex/log logger
func New(endpoint string) (h *Handler, err error) {
	dialer := &net.Dialer{
		LocalAddr: &net.UDPAddr{
			Port: 31337,
		},
	}
	conn, err := dialer.Dial("udp", endpoint)
	if err != nil {
		return nil, err
	}

	return &Handler{
		conn: conn,
		enc:  json.NewEncoder(conn),
	}, nil
}

// HandleLog implements standart apex/log Handlers func
func (h *Handler) HandleLog(e *log.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	err := h.enc.Encode(e)
	if err != nil {
		return err
	}

	_, err = h.conn.Write([]byte("\n"))
	return err
}
