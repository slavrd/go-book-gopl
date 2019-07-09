// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 278.

// Package memo provides a concurrency-safe non-blocking memoization
// of a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package memo

//!+Func

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
	cancel   chan struct{} // cancel channel for the request
}

type Memo struct {
	requests chan request
	cancels  chan request
}

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request), cancels: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, cancel <-chan interface{}) (interface{}, error) {
	response := make(chan result)
	abort := make(chan struct{})
	req := request{key, response, abort}
	memo.requests <- req
	select {
	case res := <-response:
		return res.value, res.err
	case <-cancel:
		memo.cancels <- req
		return nil, nil
	}
}

func (memo *Memo) Close() { close(memo.requests) }

//!-get

//!+monitor

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for {

		// Handle cancel channels with priority instead of letting select choose randomly
	cloop:
		for {
			select {
			case req := <-memo.cancels:
				close(req.cancel)
				delete(cache, req.key)
			default:
				break cloop
			}
		}

		select {
		case req := <-memo.requests:
			e := cache[req.key]
			if e == nil {
				// This is the first request for this key.
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key) // call f(key)
			}
			go e.deliver(&req)
		default:
		}
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(req *request) {
	// Wait for the ready condition or request cancelation
	select {
	case <-e.ready:
		// Send the result to the client.
		req.response <- e.res
	case <-req.cancel:
	}
}

//!-monitor
