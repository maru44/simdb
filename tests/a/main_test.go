package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	db := NewTableAs()

	tests := []struct {
		name       string
		insertID   uint
		insertItem tableA
		wantNoErr  bool
		wantItems  map[uint]tableA
	}{
		{
			name:     "success: first",
			insertID: uint(1),
			insertItem: tableA{
				Name:      777777,
				ExpiredAt: time.Now().Add(2 * time.Hour).Unix(),
				IsExpired: false,
			},
			wantNoErr: true,
			wantItems: map[uint]tableA{
				1: {
					Name:      777777,
					ExpiredAt: time.Now().Add(2 * time.Hour).Unix(),
					IsExpired: false,
				},
			},
		},
		{
			name:     "failed: duplicate entry",
			insertID: uint(1),
			insertItem: tableA{
				Name:      777777,
				ExpiredAt: time.Now().Add(2 * time.Hour).Unix(),
				IsExpired: false,
			},
			wantNoErr: false,
			wantItems: map[uint]tableA{
				1: {
					Name:      777777,
					ExpiredAt: time.Now().Add(2 * time.Hour).Unix(),
					IsExpired: false,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := db.Insert(tt.insertID, tt.insertItem)
			assert.Equal(t, tt.wantNoErr, err == nil)
			assert.Equal(t, tt.wantItems, db.List())
		})
	}
}

func TestUpdate(t *testing.T) {
	db := tableAs{
		data: map[uint]tableA{
			1: {
				Name:      777777,
				ExpiredAt: time.Now().Add(2 * time.Hour).Unix(),
				IsExpired: false,
			},
		},
	}

	tim := time.Now().Add(3 * time.Hour).Unix()

	tests := []struct {
		name       string
		updateID   uint
		updateItem tableA
		wantNoErr  bool
		wantItems  map[uint]tableA
	}{
		{
			name:     "success: first",
			updateID: uint(1),
			updateItem: tableA{
				Name:      888888,
				ExpiredAt: tim,
				IsExpired: false,
			},
			wantNoErr: true,
			wantItems: map[uint]tableA{
				1: {
					Name:      888888,
					ExpiredAt: tim,
					IsExpired: false,
				},
			},
		},
		{
			name:     "failed: key does not exist",
			updateID: uint(3),
			updateItem: tableA{
				Name:      777777,
				ExpiredAt: time.Now().Add(2 * time.Hour).Unix(),
				IsExpired: false,
			},
			wantNoErr: false,
			wantItems: map[uint]tableA{
				1: {
					Name:      888888,
					ExpiredAt: tim,
					IsExpired: false,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := db.Update(tt.updateID, tt.updateItem)
			assert.Equal(t, tt.wantNoErr, err == nil)

			assert.Equal(t, tt.wantItems, db.List())
		})
	}
}

func TestGetAndLoadAndExists(t *testing.T) {
	timeBefore := time.Now().Add(-2 * time.Hour).Unix()
	timeAfter := time.Now().Add(2 * time.Hour).Unix()

	d1 := tableA{
		Name:      777777,
		ExpiredAt: timeAfter,
		IsExpired: false,
	}
	d2 := tableA{
		Name:      900000,
		ExpiredAt: timeBefore,
		IsExpired: true,
	}

	db := tableAs{
		data: map[uint]tableA{
			1: d1,
			2: d2,
		},
	}

	tests := []struct {
		name          string
		id            uint
		wantItem      *tableA
		wantIsNoError bool
		wantLoadItem  *tableA
		wantLoadOK    bool
		wantExists    bool
	}{
		{
			name:          "success: first",
			id:            1,
			wantItem:      &d1,
			wantIsNoError: true,
			wantLoadItem:  &d1,
			wantLoadOK:    true,
			wantExists:    true,
		},
		{
			name:          "success: second",
			id:            2,
			wantItem:      &d2,
			wantIsNoError: true,
			wantLoadItem:  &d2,
			wantLoadOK:    true,
			wantExists:    true,
		},
		{
			name:          "failed: not ex",
			id:            3,
			wantItem:      nil,
			wantIsNoError: false,
			wantLoadItem:  nil,
			wantLoadOK:    false,
			wantExists:    false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.Get(tt.id)
			assert.Equal(t, tt.wantIsNoError, err == nil)
			assert.Equal(t, tt.wantItem, got)
			gotLoad, gotOK := db.Load(tt.id)
			assert.Equal(t, tt.wantLoadItem, gotLoad)
			assert.Equal(t, tt.wantLoadOK, gotOK)
			gotEx := db.Exists(tt.id)
			assert.Equal(t, tt.wantExists, gotEx)
		})
	}
}

func TestBulkInsert(t *testing.T) {
	db := NewTableAs()

	tests := []struct {
		name          string
		inserts       map[uint]tableA
		wantItems     map[uint]tableA
		wantIsNoError bool
	}{
		{
			name: "success: ",
			inserts: map[uint]tableA{
				1: {Name: 1},
				2: {Name: 2},
				3: {Name: 3},
			},
			wantIsNoError: true,
			wantItems: map[uint]tableA{
				1: {Name: 1},
				2: {Name: 2},
				3: {Name: 3},
			},
		},
		{
			name: "failed: duplicate",
			inserts: map[uint]tableA{
				4: {Name: 4},
				3: {Name: 2},
				5: {Name: 3},
			},
			wantIsNoError: false,
			wantItems: map[uint]tableA{
				1: {Name: 1},
				2: {Name: 2},
				3: {Name: 3},
			},
		},
		{
			name: "success: second",
			inserts: map[uint]tableA{
				4: {Name: 4},
				5: {Name: 2},
				6: {Name: 3},
			},
			wantIsNoError: true,
			wantItems: map[uint]tableA{
				1: {Name: 1},
				2: {Name: 2},
				3: {Name: 3},
				4: {Name: 4},
				5: {Name: 2},
				6: {Name: 3},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := db.BulkInsert(tt.inserts)
			assert.Equal(t, tt.wantIsNoError, err == nil)
			assert.Equal(t, tt.wantItems, db.List())
		})
	}
}

func TestUpsert(t *testing.T) {
	db := NewTableAs()

	tests := []struct {
		name      string
		inputs    map[uint]tableA
		wantItems map[uint]tableA
	}{
		{
			name: "success: ",
			inputs: map[uint]tableA{
				1: {Name: 1},
				2: {Name: 2},
				3: {Name: 3},
			},
			wantItems: map[uint]tableA{
				1: {Name: 1},
				2: {Name: 2},
				3: {Name: 3},
			},
		},
		{
			name: "success: even if duplicate entry",
			inputs: map[uint]tableA{
				4: {Name: 4},
				3: {Name: 2},
				5: {Name: 3},
			},
			wantItems: map[uint]tableA{
				1: {Name: 1},
				2: {Name: 2},
				4: {Name: 4},
				3: {Name: 2},
				5: {Name: 3},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db.BulkUpsert(tt.inputs)
			assert.Equal(t, tt.wantItems, db.List())
		})
	}
}

func TestDelete(t *testing.T) {
	db := tableAs{
		data: map[uint]tableA{
			1: {Name: 1},
			2: {Name: 2},
		},
	}

	tests := []struct {
		name      string
		id        uint
		wantItems map[uint]tableA
	}{
		{
			name: "success: ",
			id:   1,
			wantItems: map[uint]tableA{
				2: {Name: 2},
			},
		},
		{
			name: "success: even if key does not ex",
			id:   3,
			wantItems: map[uint]tableA{
				2: {Name: 2},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db.Delete(tt.id)
			assert.Equal(t, tt.wantItems, db.List())
		})
	}
}
