// Code generated by "github.com/maru44/simdb/gen"; DO NOT EDIT.

package sample

import (
	"pkg/errors"
	"sync"
)

type (
	TableSample struct {
		Column1 int
		Column2 int64
		Column3 uint64
	}

	TableSamples struct {
		Data map[int]TableSample
		mux  sync.Mutex
	}
)

func NewTableSample() TableSamples {
	return TableSamples{
		Data: map[int]TableSample{},
	}
}

func (t TableSamples) Get(id int) (TableSample, error) {
	v, ok := t.Data[id]
	if !ok {
		return TableSample{}, errors.New("Not Exists")
	}
	return v, nil
}

func (t TableSamples) Insert(id int, value TableSample) error {
	if _, ok := t.Data[id]; ok {
		return errors.New("Already Exists")
	}
	t.mux.Lock()
	defer t.mux.Unlock()
	t.Data[id] = value
	return nil
}

func (t TableSamples) BulkInsert(values TableSamples) error {
	t.mux.Lock()
	defer t.mux.Unlock()
	for id, value := range values {
		if _, ok := t.Data[id]; ok {
			return errors.New("Already Exists")
		}
		t.Data[id] = value
	}
	return nil
}

func (t TableSamples) Update(id int, value TableSample) error {
	if _, ok := t.Data[id]; !ok {
		return errors.New("Not Exists")
	}
	t.mux.Lock()
	defer t.mux.Unlock()
	t.Data[id] = value
	return nil
}

func (t TableSamples) Upsert(id int, value TableSample) {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.Data[id] = value
}

func (t TableSamples) BulkUpsert(values TableSamples) {
	t.mux.Lock()
	defer t.mux.Unlock()
	for id, value := range values {
		t.Data[id] = value
	}
}

func (t TableSamples) Delete(id int) {
	t.mux.Lock()
	defer t.mux.Unlock()
	delete(t, id)
}

func (t TableSamples) Truncate() {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.Data = map[int]TableSample{}
}
