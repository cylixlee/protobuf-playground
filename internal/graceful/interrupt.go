package graceful

import (
	"errors"
	"os"
	"os/signal"
	"reflect"
)

type Task struct {
	callable reflect.Value
	params   []reflect.Value
	defers   []*cleanupTask
}

type cleanupTask struct {
	callable reflect.Value
	params   []reflect.Value
}

func New(callable any) *Task {
	v := reflect.ValueOf(callable)
	if v.Kind() != reflect.Func {
		panic("Graceful task must be callable")
	}
	return &Task{callable: v}
}

func (t *Task) AddParams(params ...any) *Task {
	for _, param := range params {
		t.params = append(t.params, reflect.ValueOf(param))
	}
	return t
}

func (t *Task) Defer(cleanup any, params ...any) *Task {
	v := reflect.ValueOf(cleanup)
	if v.Kind() != reflect.Func {
		panic("Cleanup must be callable")
	}

	p := make([]reflect.Value, 0, len(params))
	for _, param := range params {
		p = append(p, reflect.ValueOf(param))
	}

	t.defers = append(t.defers, &cleanupTask{callable: v, params: p})
	return t
}

func (t *Task) Run() (err error) {
	interruptChan := make(chan os.Signal, 1)
	errChan := make(chan error)
	signal.Notify(interruptChan, os.Interrupt)

	go func() {
		r := t.callable.Call(t.params)
		for _, v := range r {
			if !v.CanInterface() {
				continue // discard non-error values
			}
			if e, ok := v.Interface().(error); ok {
				errChan <- e
			}
		}
		close(errChan)
	}()

	for {
		select {
		case e, ok := <-errChan:
			if !ok {
				return
			}
			err = errors.Join(err, e)
		case <-interruptChan:
			for i := len(t.defers) - 1; i >= 0; i-- {
				d := t.defers[i]
				v := d.callable.Call(d.params)
				for _, ret := range v {
					if !ret.CanInterface() {
						continue // discard non-error values
					}
					if e, ok := ret.Interface().(error); ok {
						err = errors.Join(err, e)
					}
				}
			}
			return
		}
	}
}
