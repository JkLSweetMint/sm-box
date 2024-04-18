package closer

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test_Closer_CancelWithNilWg(t *testing.T) {
	var (
		wg        = new(sync.WaitGroup)
		parentCtx = context.Background()
		conf      = new(Config).Default()
		cl, _     = New(conf, parentCtx, wg)

		timeSecondLimit = 10
	)

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			time.Sleep(time.Second * time.Duration(rand.Intn(timeSecondLimit)))
		}()
	}

	var s = make(chan struct{}, 1)

	go func() {
		cl.Wait()
		s <- struct{}{}
	}()

	select {
	case <-s:
		return
	case <-time.NewTimer(time.Second * time.Duration(timeSecondLimit+1)).C:
		t.Error("The test execution time has exceeded the allowed value. ")
	}
}

func Test_Closer_CancelWithCloser(t *testing.T) {
	var (
		wg        = new(sync.WaitGroup)
		parentCtx = context.Background()
		conf      = new(Config).Default()
		cl, ctx   = New(conf, parentCtx, wg)

		timeSecondLimit = 5
	)

	go func() {
		time.Sleep(time.Second * 3)

		cl.Cancel()
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			case <-time.NewTimer(time.Minute).C:
				return
			}
		}()
	}

	var s = make(chan struct{}, 1)

	go func() {
		cl.Wait()
		s <- struct{}{}
	}()

	select {
	case <-s:
		return
	case <-time.NewTimer(time.Second * time.Duration(timeSecondLimit)).C:
		t.Error("The test execution time has exceeded the allowed value. ")
	}
}
