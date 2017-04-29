package participant

import (
	"fmt"
	"sync"
)

type Object struct {
	Key       string
	Value     string
	lock      *sync.RWMutex
	cond      *sync.Cond
	running   bool
	currTrans int32
}

func (o *Object) start() {
	o.lock.Lock()
	// o.running = true
	o.lock.Unlock()
}

func (o *Object) stop() {
	fmt.Println("Called stop!")
	o.lock.Lock()
	o.running = false
	o.currTrans = 0
	o.cond.Broadcast()
	o.lock.Unlock()
}

func (o *Object) resetKey(value string, trans int32) {
	o.lock.Lock()
	for o.running && trans != o.currTrans {
		o.cond.Wait()
	}
	fmt.Printf("In resetKey: %v->%v\n", o.Value, value)
	o.Value = value
	o.currTrans = 0
	o.lock.Unlock()
}

func (o *Object) setKey(key string, value string, trans int32) {
	// o.lock.Lock()
	// for o.running && trans != o.currTrans {
	// 	o.cond.Wait()
	// }

	if _, ok := self.held[key]; ok {
	  self.held[key].lock.Lock()
		for self.held[key].holding {
			self.held[key].cond.Wait()
		}
	}
	fmt.Printf("In setKey: %v is value\n", value)
	o.Value = value
	self.held[key].lock.Unlock()
	// o.running = true
	// o.currTrans = trans
	// o.lock.Unlock()
	fmt.Println(o)
}

func (o *Object) getKey() string {
	o.lock.RLock()
	res := o.Value
	o.lock.RUnlock()
	return res
}

func NewObject(key string, value string, trans int32) *Object {
  if _, ok := self.held[key]; ok {
    self.held[key].lock.Lock()
		for self.held[key].holding {
			self.held[key].cond.Wait()
		}
	}
	self.held[key].lock.Unlock()
	m := &sync.RWMutex{}
	c := sync.NewCond(m)
	return &Object{key, value, m, c, true, trans}
}
