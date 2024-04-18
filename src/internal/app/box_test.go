package app

import (
	"sm-box/src/pkg/core"
	"sm-box/src/pkg/core/env"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var srv, err = New()

	if err != nil {
		t.Fatalf("Failed to create the '%s': '%s'. ", env.Vars.SystemName, err)
	}

	if srv == nil {
		t.Fatalf("'%s' instance is nil. ", env.Vars.SystemName)
	}

	// Проверки текущего состояния
	{
		if srv.State() != core.StateBooted {
			t.Fatalf("Invalid value of the '%s' state. ", env.Vars.SystemName)
		}
	}

	var completed = make(chan struct{})

	go func() {
		defer func() { completed <- struct{}{} }()

		if err = srv.Serve(); err != nil {
			t.Fatalf("Failed to launch '%s': '%s'. ", env.Vars.SystemName, err)
		}
	}()

	time.Sleep(time.Second)

	// Проверки текущего состояния
	{
		if srv.State() != core.StateServed {
			t.Fatalf("Invalid value of the '%s' state. ", env.Vars.SystemName)
		}
	}

	time.Sleep(time.Second * 3)

	if err = srv.Shutdown(); err != nil {
		t.Fatalf("'%s' could not be completed: '%s'. ", env.Vars.SystemName, err)
	}

	// Проверки текущего состояния
	{
		if srv.State() != core.StateOff {
			t.Fatalf("Invalid value of the '%s' state. ", env.Vars.SystemName)
		}
	}

	select {
	case <-completed:
		return
	case <-time.NewTimer(time.Second * 10).C:
		t.Fatal("The test execution time has exceeded the allowed value. ")
	}
}
