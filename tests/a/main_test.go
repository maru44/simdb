package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	db := TableAs{
		Data: map[uint]TableA{},
	}

	tests := []struct {
		name       string
		insertID   uint
		insertItem TableA
		wantNoErr  bool
		wantItems  map[uint]TableA
	}{
		{
			name:     "first success",
			insertID: uint(1),
			insertItem: TableA{
				Name:      777777,
				ExpiredAt: time.Now().Add(2 * time.Hour).Unix(),
				IsExpired: false,
			},
			wantNoErr: true,
			wantItems: map[uint]TableA{
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
			insertItem: TableA{
				Name:      777777,
				ExpiredAt: time.Now().Add(2 * time.Hour).Unix(),
				IsExpired: false,
			},
			wantNoErr: false,
			wantItems: map[uint]TableA{
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
