package core

import (
	"testing"
)

func Test_core_Ctx(t *testing.T) {
	var cr, _ = New()

	if ctx := cr.Ctx(); ctx == nil {
		t.Error("Context are nil. ")
	}
}

func Test_core_Components(t *testing.T) {
	var (
		cr, _      = New()
		components = cr.Components()
	)

	switch {
	case components == nil:
		t.Error("Components is nil. ")
	case components.Logger() == nil:
		t.Error("Logger component is nil. ")
	}
}

func Test_core_Tools(t *testing.T) {
	var (
		cr, _ = New()
		tools = cr.Tools()
	)

	switch {
	case tools == nil:
		t.Error("Tools is nil. ")
	case tools.TaskScheduler() == nil:
		t.Error("Task scheduler tool is nil. ")
	}
}
