package tester

import (
	"context"
	"testing"
	"time"

	"github.com/weisd/chromerunner"
)

// TestRunner TestRunner
func TestRunner(t *testing.T) {
	ctx := context.TODO()

	c, err := chromerunner.New(chromerunner.URL("http://www.baidu.com"))
	if err != nil {
		t.Fatal(err)
	}

	if err := c.Start(ctx); err != nil {
		t.Fatal(err)
	}

	t.Log("start done")

	time.Sleep(3 * time.Second)

	if err := c.Shutdown(ctx); err != nil {
		t.Fatal(err)
	}

	t.Log("Shutdown done")

	if err := c.Wait(); err != nil {
		t.Fatal(err)
	}

	t.Log("done")
}
