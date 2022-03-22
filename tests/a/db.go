// Code generated by "github.com/maru44/simdb/gen"; DO NOT EDIT.

package main

import (
	"fmt"
	"sync"
)

type (
	tableA struct {
		Name      int32
		ExpiredAt int64
		IsExpired bool
	}

	tableAs struct {
		Data map[uint]tableA
		sync.RWMutex
	}
)

func (t *tableAs) Get(id uint) (*tableA, error) {
	t.RLock()
	defer t.RUnlock()
	v, ok := t.Data[id]
	if !ok {
		return nil, fmt.Errorf("Not Exists: %v", id)
	}
	return &v, nil
}

func (t *tableAs) Insert(id uint, value tableA) error {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.Data[id]; ok {
		return fmt.Errorf("Duplicate Entry: %v", id)
	}
	t.Data[id] = value
	return nil
}

func (t *tableAs) BulkInsert(values map[uint]tableA) error {
	t.Lock()
	defer t.Unlock()
	for id, value := range values {
		if _, ok := t.Data[id]; ok {
			return fmt.Errorf("Duplicate Entry: %v", id)
		}
		t.Data[id] = value
	}
	return nil
}

func (t *tableAs) Update(id uint, value tableA) error {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.Data[id]; !ok {
		return fmt.Errorf("Does not exists: %v", id)
	}
	t.Data[id] = value
	return nil
}

func (t *tableAs) Upsert(id uint, value tableA) {
	t.Lock()
	defer t.Unlock()
	t.Data[id] = value
}

func (t *tableAs) BulkUpsert(values map[uint]tableA) {
	t.Lock()
	defer t.Unlock()
	for id, value := range values {
		t.Data[id] = value
	}
}

func (t *tableAs) Delete(id uint) {
	t.Lock()
	defer t.Unlock()
	delete(t.Data, id)
}

func (t *tableAs) Truncate() {
	t.Lock()
	defer t.Unlock()
	t.Data = map[uint]tableA{}
}
