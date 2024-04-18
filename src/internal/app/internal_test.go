package app

import (
	"sm-box/src/pkg/core/env"
	"testing"
)

func Test_server_Ctx(t *testing.T) {
	var srv, err = New()

	if err != nil {
		t.Fatalf("Failed to create the '%s': '%s'. ", env.Vars.SystemName, err)
	}

	if ctx := srv.Ctx(); ctx == nil {
		t.Error("Context are nil. ")
	}
}

func Test_server_Components(t *testing.T) {
	var srv, err = New()

	if err != nil {
		t.Fatalf("Failed to create the '%s': '%s'. ", env.Vars.SystemName, err)
	}

	switch {
	case srv.Components() == nil:
		t.Error("Components is nil. ")
	case srv.Components().Logger() == nil:
		t.Error("Logger component is nil. ")
	}
}

func Test_server_Controllers(t *testing.T) {
	var srv, err = New()

	if err != nil {
		t.Fatalf("Failed to create the '%s': '%s'. ", env.Vars.SystemName, err)
	}

	switch {
	case srv.Controllers() == nil:
		t.Error("Controllers is nil. ")
	}
}
