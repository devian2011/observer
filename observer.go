package observer

import (
	"reflect"
	"runtime"
	"sync"
)

type EventCode string
type EventData any
type EventHandler func(data EventData)

type Observer struct {
	wg       *sync.WaitGroup
	handlers map[EventCode][]EventHandler
}

var observer = &Observer{
	handlers: make(map[EventCode][]EventHandler),
}

func (o *Observer) Notify(code EventCode, data EventData) {
	if handlers, exists := o.handlers[code]; exists {
		for _, h := range handlers {
			o.wg.Add(1)
			go func(h EventHandler, wg *sync.WaitGroup) {
				h(data)
				wg.Done()
			}(h, o.wg)
		}
	}
}

func (o *Observer) Register(code EventCode, handlers ...EventHandler) {
	if _, exists := o.handlers[code]; !exists {
		o.handlers[code] = make([]EventHandler, 0)
	}
	o.handlers[code] = append(o.handlers[code], handlers...)
}

func (o *Observer) EventCodes() []EventCode {
	result := make([]EventCode, len(o.handlers))
	for code := range o.handlers {
		result = append(result, code)
	}

	return result
}

func (o *Observer) GetFunctionsForEvent(code EventCode) []string {
	if handlers, exists := o.handlers[code]; exists {
		result := make([]string, len(handlers))
		for _, fn := range handlers {
			result = append(result, runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())
		}
	}

	return []string{}
}

func (o *Observer) GetFunctions() map[EventCode][]string {
	result := make(map[EventCode][]string, len(o.handlers))
	for code := range o.handlers {
		result[code] = o.GetFunctionsForEvent(code)
	}

	return result
}

func (o *Observer) Wait() {
	o.wg.Wait()
}

func Notify(code EventCode, data EventData) {
	observer.Notify(code, data)
}

func Register(code EventCode, handler EventHandler) {
	observer.Register(code, handler)
}

func EventCodes() []EventCode {
	return observer.EventCodes()
}

func GetFunctionsForEvent(code EventCode) []string {
	return observer.GetFunctionsForEvent(code)
}

func GetFunctions() map[EventCode][]string {
	return observer.GetFunctions()
}

func Wait() {
	observer.Wait()
}
