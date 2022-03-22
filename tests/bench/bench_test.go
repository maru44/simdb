package bench

import (
	"sync"
	"testing"
)

func BenchmarkInsert(b *testing.B) {
	db := NewBenchs()
	value := bench{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Insert(i, value)
	}
}

func BenchmarkInsert_SyncMap(b *testing.B) {
	db := sync.Map{}
	value := bench{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Store(i, value)
	}
}

func BenchmarkGet(b *testing.B) {
	db := NewBenchs()
	value := bench{}
	for i := 0; i < b.N; i++ {
		db.Insert(i, value)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Get(i)
	}
}

func BenchmarkGet_SyncMap(b *testing.B) {
	db := sync.Map{}
	value := bench{}
	for i := 0; i < b.N; i++ {
		db.Store(i, value)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, _ := db.Load(i)
		_ = val.(bench)
	}
}
