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
			assert.Equal(t, tt.wantItems, db.Data)
		})
	}
}

func TestUpdate(t *testing.T) {
	db := tableAs{
		Data: map[uint]tableA{
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

			assert.Equal(t, tt.wantItems, db.Data)
		})
	}
}

func TestGet(t *testing.T) {
	timeBefore := time.Now().Add(-2 * time.Hour).Unix()
	timeAfter := time.Now().Add(2 * time.Hour).Unix()

	db := tableAs{
		Data: map[uint]tableA{
			1: {
				Name:      777777,
				ExpiredAt: timeAfter,
				IsExpired: false,
			},
			2: {
				Name:      900000,
				ExpiredAt: timeBefore,
				IsExpired: true,
			},
		},
	}

	tests := []struct {
		name          string
		id            uint
		wantItem      tableA
		wantIsNoError bool
	}{
		{
			name:          "success: first",
			id:            1,
			wantItem:      db.Data[1],
			wantIsNoError: true,
		},
		{
			name:          "success: second",
			id:            2,
			wantItem:      db.Data[2],
			wantIsNoError: true,
		},
		{
			name:          "failed: not ex",
			id:            3,
			wantItem:      tableA{},
			wantIsNoError: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.Get(tt.id)
			assert.Equal(t, tt.wantIsNoError, err == nil)
			assert.Equal(t, tt.wantItem, got)
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
				1: {
					Name: 1,
				},
				2: {
					Name: 2,
				},
				3: {
					Name: 3,
				},
			},
			wantIsNoError: true,
			wantItems: map[uint]tableA{
				1: {
					Name: 1,
				},
				2: {
					Name: 2,
				},
				3: {
					Name: 3,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := db.BulkInsert(tt.inserts)
			assert.Equal(t, tt.wantIsNoError, err == nil)
			assert.Equal(t, tt.wantItems, db.Data)
		})
	}
}
