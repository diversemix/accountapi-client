package accountclient

import (
	"testing"
)

type goodClientLogger struct{}

func (c *goodClientLogger) Fatalln(v ...interface{}) {
}

func (c *goodClientLogger) Println(v ...interface{}) {
}

func (c *goodClientLogger) Printf(fmt string, v ...interface{}) {
}

func isClientLogger(c *goodClientLogger) ClientLogger {
	return c
}

func TestClientLogger(t *testing.T) {

	t.Run("interface has required members", func(t *testing.T) {
		x := goodClientLogger{}
		if isClientLogger(&x) == nil {
			t.Error("Expecting ClientLogger")
		}
		return
	})
}
