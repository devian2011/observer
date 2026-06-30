# Go Observer Pattern Package

A lightweight, concurrent, and flexible Event Observer package for Go. It supports both asynchronous (goroutine-backed) and synchronous event execution, reflection-based handler inspection, and built-in `sync.WaitGroup` tracking for safe application shutdowns.

## Features

* **Async & Sync Notifications**: Choose between concurrent `Notify()` or blocking `NotifySync()`.
* **Dynamic Registrations**: Bind multiple handlers to a single event code.
* **Wait Group Tracking**: Built-in `Wait()` mechanism ensures all background goroutines finish processing before your app exits.
* **Introspection Tools**: Inspect registered handlers by retrieving Go function names via runtime reflection.
* **Global & Local Instances**: Use the ready-to-go package-level global functions or instantiate separate `New()` observers.

## Installation

```bash
go get github.com/devian2011/observer
```

## Quick Start

### 1. Using Global Functions (Asynchronous)

```go
package main

import (
	"fmt"
	"github.com/devian2011/observer"
)

func main() {
	// Define an event
	const UserRegistered observer.EventCode = "USER_REGISTERED"

	// Register handlers
	observer.Register(UserRegistered, func(data observer.EventData) {
		email := data.(string)
		fmt.Printf("[Handler 1] Sending welcome email to %s\n", email)
	})

	observer.Register(UserRegistered, func(data observer.EventData) {
		fmt.Println("[Handler 2] Updating analytics database...")
	})

	// Fire event asynchronously
	observer.Notify(UserRegistered, "john.doe@example.com")

	// Wait for all async tasks to finish before exiting
	observer.Wait()
	fmt.Println("All events processed.")
}
```

### 2. Using Synchronous Notifications

If your handlers must execute sequentially on the main thread, use `NotifySync`.

```go
// This blocks until all registered handlers finish executing one after another
observer.NotifySync(UserRegistered, "jane.doe@example.com")
```

### 3. Creating Isolated Instances

If you want to avoid global state, instantiate an isolated observer pool.

```go
package main

import "github.com/devian2011/observer"

func main() {
    customObserver := observer.New()
    
    customObserver.Register("LOG_EVENT", func(data observer.EventData) {
        // handle log
    })
    
    customObserver.Notify("LOG_EVENT", "System started")
    customObserver.Wait()
}
```

## Introspection API

You can programmatically inspect your registered events and look up the underlying Go runtime function names for debugging or logging purposes.

```go
// Get all registered event codes
codes := observer.EventCodes() 

// Get function names bound to a specific event
funcs := observer.GetFunctionsForEvent("USER_REGISTERED") 

// Get the full map of events and their functions
allFuncs := observer.GetFunctions() 
```

## API Reference

| Function / Method | Scope | Description |
| :--- | :--- | :--- |
| `New()` | Global | Creates a new isolated `Observer` instance. |
| `Register(code, ...handlers)` | Both | Binds one or multiple handlers to an event code. |
| `Notify(code, data)` | Both | Triggers all handlers concurrently in separate goroutines. |
| `NotifySync(code, data)` | Both | Triggers handlers sequentially on the current thread. |
| `Wait()` | Both | Blocks until all background goroutine handlers finish. |
| `EventCodes()` | Both | Returns a slice of all registered event codes. |
| `GetFunctionsForEvent(code)` | Both | Returns runtime string names of functions registered to an event. |
| `GetFunctions()` | Both | Returns a complete map of event codes to their function name slices. |
