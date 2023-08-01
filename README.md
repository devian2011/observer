# Simple observer package

Async observer.

## How to

Use global

```go
package main

import (
	"fmt"
	"time"

	"github.com/devian2011/observer"
)

type OnAppStart struct {
	date   time.Time
	action string
}

func main() {


	observer.Register("onApplicationStart", func(data observer.EventData) {
		action, ok := data.(OnAppStart)
		if ok {
			fmt.Println(action.date, action.action)
		}
	})
	
	start := &OnAppStart{
		date: time.Now(),
		action: "boot",
    }

	observer.Notify("onApplicationStart", start) //Async notification
	observer.NotifySync("onApplicationStart", start) // Sync notification (in same goroutine)
}
```

Use by variable

```go
package main

import (
	"fmt"
	"time"

	"github.com/devian2011/observer"
)

type OnAppStart struct {
	date   time.Time
	action string
}

func main() {

    ob := observer.New()
	ob.Register("onApplicationStart", func(data observer.EventData) {
		action, ok := data.(OnAppStart)
		if ok {
			fmt.Println(action.date, action.action)
		}
	})
	
	start := &OnAppStart{
		date: time.Now(),
		action: "boot",
    }

	ob.Notify("onApplicationStart", start) //Async notification
	ob.NotifySync("onApplicationStart", start) // Sync notification (in same goroutine)
	ob.Wait()
}
```

