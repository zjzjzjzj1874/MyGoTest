package golover

import (
	"context"
	"testing"
)

func setup() (srv DataService, ctx context.Context) {
	return NewDataService(), context.Background()
}

func TestDataService_Status(t *testing.T) {
	srv, ctx := setup()

	s, err := srv.Status(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing status
	ok := s == "ok"
	if !ok {
		t.Errorf("expected service to be ok")
	}
}

func TestClosure(t *testing.T) {
	closure()
}

func TestDefer(t *testing.T) {
	deferWithParam()
}

func TestRunTime(t *testing.T) {
	calRuntime()
}

func TestLearnChannel(t *testing.T) {
	//learnChannel()
	//personChan()
	quitChannel()
}
