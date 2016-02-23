package maplock

import (
	"sync"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	m := New()
	m.Lock("foo")
	m.Unlock("foo")
}

func Test2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(3)

	m := New()
	n := 0
	m.Lock("foo")
	n++
	m.Unlock("foo")
	if n != 1 {
		t.Fatal("not 1")
	}
	wg.Done()

	go func() {
		m.Lock("foo")
		n++
		m.Unlock("foo")
		if n != 2 {
			t.Fatal("not 2")
		}
		wg.Done()
	}()
	go func() {
		time.Sleep(1 * time.Millisecond)
		m.Lock("foo")
		n++
		m.Unlock("foo")
		if n != 3 {
			t.Fatal("not 3")
		}
		wg.Done()
	}()
	wg.Wait()
}
