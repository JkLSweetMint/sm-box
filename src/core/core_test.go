package core

import (
	"context"
	"fmt"
	"sm-box/src/core/tools/task_scheduler"
	"testing"
	"time"
)

type TestServer struct {
}

func (srv *TestServer) Serve(ctx context.Context) (err error) {
	fmt.Println("Server serve")
	return
}

func (srv *TestServer) Shutdown(ctx context.Context) (err error) {
	fmt.Println("Server shutdown")
	return
}

func Test(t *testing.T) {
	var (
		cr, err = New()
		srv     = new(TestServer)
	)

	if err != nil {
		t.Fatalf("Failed to create the system core: '%s'. ", err)
	}

	if cr == nil {
		t.Fatal("Core instance is nil. ")
	}

	if err = cr.Tools().TaskScheduler().Register(task_scheduler.Task{
		Name: "Start test server",
		Type: task_scheduler.TaskServe,
		Func: srv.Serve,
	}); err != nil {
		t.Fatalf("Failed to register a task in task scheduler: '%s'. ", err)
	}

	if err = cr.Tools().TaskScheduler().Register(task_scheduler.Task{
		Name: "Shutdown test server",
		Type: task_scheduler.TaskShutdown,
		Func: srv.Shutdown,
	}); err != nil {
		t.Fatalf("Failed to register a task in task scheduler: '%s'. ", err)
	}

	// Проверки текущего состояния
	{
		switch {
		case cr.State() != StateNew:
			t.Fatal("Invalid value of the system core state. ")
		default:
			{
				if e := cr.Shutdown(); e == nil {
					t.Fatal("Invalid behavior of the system core in the 'New' state. ")
				}
				if e := cr.Serve(); e == nil {
					t.Fatal("Invalid behavior of the system core in the 'New' state. ")
				}
			}
		}
	}

	if err = cr.Boot(); err != nil {
		t.Fatalf("An error occurred while loading the core: '%s'. ", err)
	}

	// Проверки текущего состояния
	{
		switch {
		case cr.State() != StateBooted:
			t.Fatal("Invalid value of the system core state. ")
		default:
			{
				if e := cr.Boot(); e == nil {
					t.Fatal("Invalid behavior of the system core in the 'New' state. ")
				}
				if e := cr.Shutdown(); e == nil {
					t.Fatal("Invalid behavior of the system core in the 'New' state. ")
				}
			}
		}
	}

	var completed = make(chan struct{})

	go func() {
		defer func() { completed <- struct{}{} }()

		if err = cr.Serve(); err != nil {
			t.Fatalf("An error occurred during the start of system maintenance by the core: '%s'. ", err)
		}
	}()

	time.Sleep(time.Second)

	// Проверки текущего состояния
	{
		switch {
		case cr.State() != StateServed:
			t.Fatal("Invalid value of the system core state. ")
		default:
			{
				if e := cr.Boot(); e == nil {
					t.Fatal("Invalid behavior of the system core in the 'New' state. ")
				}
				if e := cr.Serve(); e == nil {
					t.Fatal("Invalid behavior of the system core in the 'New' state. ")
				}
			}
		}
	}

	time.Sleep(time.Second * 3)

	if err = cr.Shutdown(); err != nil {
		t.Fatalf("An error occurred during the completion of maintenance by the core: '%s'. ", err)
	}

	// Проверки текущего состояния
	{
		switch {
		case cr.State() != StateOff:
			t.Fatal("Invalid value of the system core state. ")
		default:
			{
				if e := cr.Boot(); e == nil {
					t.Fatal("Invalid behavior of the system core in the 'New' state. ")
				}
				if e := cr.Serve(); e == nil {
					t.Fatal("Invalid behavior of the system core in the 'New' state. ")
				}
				if e := cr.Shutdown(); e == nil {
					t.Fatal("Invalid behavior of the system core in the 'New' state. ")
				}
			}
		}
	}

	select {
	case <-completed:
		return
	case <-time.NewTimer(time.Second * 10).C:
		t.Fatal("The test execution time has exceeded the allowed value. ")
	}
}
