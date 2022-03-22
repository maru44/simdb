package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	db := tableAs{
		Data: map[uint]tableA{},
	}

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

}
