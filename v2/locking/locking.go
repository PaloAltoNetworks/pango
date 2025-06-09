package locking

import (
	"sync"
)

type LockCategory string

const (
	ImportFileLockCategory LockCategory = "import"
	XpathLockCategory      LockCategory = "xpath"
)

type categoryLocks struct {
	mutex *sync.Mutex
	locks map[string]*sync.Mutex
}

func newCategoryLocks() *categoryLocks {
	return &categoryLocks{
		mutex: &sync.Mutex{},
		locks: make(map[string]*sync.Mutex),
	}
}

func (o *categoryLocks) getMutex(path string) *sync.Mutex {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	mutex, found := o.locks[path]
	if !found {
		mutex = &sync.Mutex{}
		o.locks[path] = mutex
	}

	return mutex
}

type locksManager struct {
	mutex      *sync.Mutex
	categories map[LockCategory]*categoryLocks
}

func newLocksManager() *locksManager {
	return &locksManager{
		mutex:      &sync.Mutex{},
		categories: make(map[LockCategory]*categoryLocks),
	}
}

var manager = newLocksManager()

func GetMutex(category LockCategory, path string) *sync.Mutex {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	categoryLocks, found := manager.categories[category]
	if !found {
		categoryLocks = newCategoryLocks()
		manager.categories[category] = categoryLocks
	}

	return categoryLocks.getMutex(path)
}
