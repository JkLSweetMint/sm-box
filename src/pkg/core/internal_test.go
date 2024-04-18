package core

import (
	"testing"
)

func Test_core_Ctx(t *testing.T) {
	var cr, err = New()

	if err != nil {
		t.Fatalf("Failed to create the system core: '%s'. ", err)
	}

	if ctx := cr.Ctx(); ctx == nil {
		t.Error("Context are nil. ")
	}
}

func Test_core_Components(t *testing.T) {
	var cr, err = New()

	if err != nil {
		t.Fatalf("Failed to create the system core: '%s'. ", err)
	}

	switch {
	case cr.Components() == nil:
		t.Error("Components is nil. ")
	case cr.Components().Logger() == nil:
		t.Error("Logger component is nil. ")
	}
}

func Test_core_Tools(t *testing.T) {
	var cr, err = New()

	if err != nil {
		t.Fatalf("Failed to create the system core: '%s'. ", err)
	}

	switch {
	case cr.Tools() == nil:
		t.Error("Tools is nil. ")
	case cr.Tools().TaskScheduler() == nil:
		t.Error("Task scheduler tool is nil. ")
	}
}
