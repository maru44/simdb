// Code generated by "github.com/maru44/simdb/gen"; DO NOT EDIT.

package main_test

import (
	"fmt"
	"sync"
)

type (
	bench struct {
		Name    string
		Email   string
		age     uint
		IsValid bool
	}

	benchs struct {
		Data map[string]bench
		sync.RWMutex
	}
)

func (t *benchs) Get(id string) (bench, error) {
	t.RLock()
	defer t.RUnlock()
	v, ok := t.Data[id]
	if !ok {
		return bench{}, fmt.Errorf("Not Exists: %v", id)
	}
	return v, nil
}

func (t *benchs) Insert(id string, value bench) error {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.Data[id]; ok {
		return fmt.Errorf("Duplicate Entry: %v", id)
	}
	t.Data[id] = value
	return nil
}

func (t *benchs) BulkInsert(values map[string]bench) error {
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

func (t *benchs) Update(id string, value bench) error {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.Data[id]; !ok {
		return fmt.Errorf("Does not exists: %v", id)
	}
	t.Data[id] = value
	return nil
}

func (t *benchs) Upsert(id string, value bench) {
	t.Lock()
	defer t.Unlock()
	t.Data[id] = value
}

func (t *benchs) BulkUpsert(values map[string]bench) {
	t.Lock()
	defer t.Unlock()
	for id, value := range values {
		t.Data[id] = value
	}
}

func (t *benchs) Delete(id string) {
	t.Lock()
	defer t.Unlock()
	delete(t.Data, id)
}

func (t *benchs) Truncate() {
	t.Lock()
	defer t.Unlock()
	t.Data = map[string]bench{}
}
