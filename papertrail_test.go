package papertrail

import (
	"testing"

	"github.com/apex/log"

	check "gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type MySuite struct{}

var _ = check.Suite(&MySuite{})

// func (s *MySuite) TestHandler(c *check.C) {
// 	endpoint := "localhost:udp"

// 	ch := make(chan string)
// 	ln, _ := net.Listen("udp", endpoint)
// 	go func() {
// 		conn, _ := ln.Accept()
// 		message, _ := bufio.NewReader(conn).ReadString('\n')
// 		ch <- message
// 	}()

// 	h, err := New(endpoint)
// 	c.Assert(h, check.NotNil)
// 	c.Assert(err, check.IsNil)

// 	log.SetHandler(h)
// 	ctx := log.WithField("key 1", "value 1")
// 	ctx.Info("this is an Info message")

// 	select {
// 	case <-time.After(1 * time.Second):
// 		c.Error("Timeout waiting for the log message")
// 	case msg := <-ch:
// 		c.Assert(strings.HasPrefix(msg, ""), check.Equals, true)
// 	}

// }

func (s *MySuite) TestRealEndpoint(c *check.C) {
	c.Skip("Only to test the real endpoint")
	/*
	  This is only to test the real endpoint.
	  Please remove the endpoint from the code after using this test
	*/
	endpoint := "ENDPOINT.papertrailapp.com:PORT"

	h, err := New(endpoint)
	c.Assert(h, check.NotNil)
	c.Assert(err, check.IsNil)

	log.SetHandler(h)
	ctx := log.WithField("key 5", "value 5")
	ctx.Info("this is an Info message")

	// After this point you should be able to see the entry on the web interface
	// Check page after registration on papertrail: https://papertrailapp.com/systems/setup?type=app&platform=unix

}
