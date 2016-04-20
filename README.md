# maplock

A map of locks

```go
m := maplock.New()
m.Lock("foo")
m.Lock("bar")
//do stuff with foo and bar
m.Unlock("foo")
m.Unlock("bar")
```

---

**Update** April 2016

Replaced the underlying implementation with Docker's [locker](https://github.com/docker/docker/tree/master/pkg/locker) implementation since it has exactly the same API though also adds auto-cleanup and doesn't panic on missing locks.

Locker
=====

locker provides a mechanism for creating finer-grained locking to help
free up more global locks to handle other tasks.

The implementation looks close to a sync.Mutex, however the user must provide a
reference to use to refer to the underlying lock when locking and unlocking,
and unlock may generate an error.

If a lock with a given name does not exist when `Lock` is called, one is
created.
Lock references are automatically cleaned up on `Unlock` if nothing else is
waiting for the lock.


## Usage

```go
package important

import (
	"sync"
	"time"

	"github.com/docker/docker/pkg/locker"
)

type important struct {
	locks *locker.Locker
	data  map[string]interface{}
	mu    sync.Mutex
}

func (i *important) Get(name string) interface{} {
	i.locks.Lock(name)
	defer i.locks.Unlock(name)
	return data[name]
}

func (i *important) Create(name string, data interface{}) {
	i.locks.Lock(name)
	defer i.locks.Unlock(name)

	i.createImportant(data)

	s.mu.Lock()
	i.data[name] = data
	s.mu.Unlock()
}

func (i *important) createImportant(data interface{}) {
	time.Sleep(10 * time.Second)
}
```

For functions dealing with a given name, always lock at the beginning of the
function (or before doing anything with the underlying state), this ensures any
other function that is dealing with the same name will block.

When needing to modify the underlying data, use the global lock to ensure nothing
else is modfying it at the same time.
Since name lock is already in place, no reads will occur while the modification
is being performed.
