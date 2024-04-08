package core

import "testing"

func Test(t *testing.T) {
	var (
		core Core
		err  error
	)

	// Создание и проверки экземпляра ядра
	{
		if core, err = New(); err != nil {
			t.Fatalf("An error occurred while creating an instance of the system core: '%s'. ", err)
		}

		switch {
		case core == nil:
			t.Fatal("Core instance is nil. ")
			//case core.State() != StateNew:
			//	t.Fatal("Core instance state not 'New'. ")
		}
	}
}
