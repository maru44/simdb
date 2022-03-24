// Code generated by "github.com/maru44/simdb/gen"; DO NOT EDIT.

package point

import (
	"fmt"
	"sync"
)

type (
	Pt struct {
		Name      string
		ExpiredAt int64
		IsExpired bool
	}

	Pts struct {
		data map[string]*Pt
		sync.RWMutex
	}
)

func NewPts() Pts {
	return Pts{
		data: map[string]*Pt{},
	}
}

func (t *Pts) List() map[string]*Pt {
	t.RLock()
	defer t.RUnlock()
	return t.data
}

func (t *Pts) Exists(id string) bool {
	t.RLock()
	defer t.RUnlock()
	_, ok := t.data[id]
	return ok
}

func (t *Pts) Get(id string) (*Pt, error) {
	t.RLock()
	defer t.RUnlock()
	v, ok := t.data[id]
	if !ok {
		return nil, fmt.Errorf("Not Exists: %v", id)
	}
	return v, nil
}

func (t *Pts) Load(id string) (*Pt, bool) {
	t.RLock()
	defer t.RUnlock()
	value, ok := t.data[id]
	return value, ok
}

func (t *Pts) Insert(id string, value *Pt) error {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.data[id]; ok {
		return fmt.Errorf("Duplicate Entry: %v", id)
	}
	t.data[id] = value
	return nil
}

func (t *Pts) BulkInsert(values map[string]*Pt) error {
	t.Lock()
	defer t.Unlock()
	for id := range values {
		if _, ok := t.data[id]; ok {
			return fmt.Errorf("Duplicate Entry: %v", id)
		}
	}
	for id, value := range values {
		t.data[id] = value
	}
	return nil
}

func (t *Pts) Update(id string, value *Pt) error {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.data[id]; !ok {
		return fmt.Errorf("Does not exists: %v", id)
	}
	t.data[id] = value
	return nil
}

func (t *Pts) UpdateName(id string, value string) error {
	t.Lock()
	defer t.Unlock()
	data, ok := t.data[id]
	if !ok {
		return fmt.Errorf("Does not exists: %v", id)
	}
	data.Name = value
	return nil
}

func (t *Pts) UpdateExpiredAt(id string, value int64) error {
	t.Lock()
	defer t.Unlock()
	data, ok := t.data[id]
	if !ok {
		return fmt.Errorf("Does not exists: %v", id)
	}
	data.ExpiredAt = value
	return nil
}

func (t *Pts) UpdateIsExpired(id string, value bool) error {
	t.Lock()
	defer t.Unlock()
	data, ok := t.data[id]
	if !ok {
		return fmt.Errorf("Does not exists: %v", id)
	}
	data.IsExpired = value
	return nil
}

func (t *Pts) Upsert(id string, value *Pt) {
	t.Lock()
	defer t.Unlock()
	t.data[id] = value
}

func (t *Pts) BulkUpsert(values map[string]*Pt) {
	t.Lock()
	defer t.Unlock()
	for id, value := range values {
		t.data[id] = value
	}
}

func (t *Pts) Delete(id string) {
	t.Lock()
	defer t.Unlock()
	delete(t.data, id)
}

func (t *Pts) BulkDelete(ids []string) {
	t.Lock()
	defer t.Unlock()
	for _, id := range ids {
		delete(t.data, id)
	}
}

func (t *Pts) Truncate() {
	t.Lock()
	defer t.Unlock()
	t.data = map[string]*Pt{}
}
