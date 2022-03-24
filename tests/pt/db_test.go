package point

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	db := NewPts()

	err := db.BulkInsert(map[string]*Pt{
		"1": &Pt{Name: "foo"},
		"2": &Pt{Name: "bar"},
	})
	require.NoError(t, err)

	d1, err := db.Get("1")
	require.NoError(t, err)
	d1.Name = "foo2"

	tests := []struct {
		name          string
		id            string
		wantVal       *Pt
		wantIsNoError bool
	}{
		{
			name:          "success",
			id:            "1",
			wantVal:       &Pt{Name: "foo2"},
			wantIsNoError: true,
		},
		{
			name:          "failed",
			id:            "3",
			wantVal:       nil,
			wantIsNoError: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.Get(tt.id)
			assert.Equal(t, tt.wantVal, got)
			assert.Equal(t, tt.wantIsNoError, err == nil)
		})
	}
}
