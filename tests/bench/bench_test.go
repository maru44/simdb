package bench

import (
	"sync"
	"testing"

	pt "github.com/maru44/simdb/tests/pt"
)

func BenchmarkInsert(b *testing.B) {
	db := NewBenchs()
	value := bench{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = db.Insert(i, value)
	}
}

func BenchmarkInsertPt(b *testing.B) {
	db := pt.NewPts()
	value := pt.Pt{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = db.Insert(i, &value)
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
		err := db.Insert(i, value)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = db.Get(i)
	}
}

func BenchmarkLoad(b *testing.B) {
	db := NewBenchs()
	value := bench{}
	for i := 0; i < b.N; i++ {
		err := db.Insert(i, value)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = db.Load(i)
	}
}

func BenchmarkGetPt(b *testing.B) {
	db := pt.NewPts()
	value := &pt.Pt{}
	for i := 0; i < b.N; i++ {
		err := db.Insert(i, value)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = db.Get(i)
	}
}

func BenchmarkLoadPt(b *testing.B) {
	db := pt.NewPts()
	value := &pt.Pt{}
	for i := 0; i < b.N; i++ {
		err := db.Insert(i, value)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = db.Load(i)
	}
}

func BenchmarkLoad_SyncMap(b *testing.B) {
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

func BenchmarkDelete(b *testing.B) {
	db := NewBenchs()
	value := bench{}
	for i := 0; i < b.N; i++ {
		err := db.Insert(i, value)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Delete(i)
	}
}

func BenchmarkDeletePt(b *testing.B) {
	db := pt.NewPts()
	value := &pt.Pt{}
	for i := 0; i < b.N; i++ {
		err := db.Insert(i, value)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Delete(i)
	}
}

func BenchmarkDelete_SyncMap(b *testing.B) {
	db := sync.Map{}
	value := bench{}
	for i := 0; i < b.N; i++ {
		db.Store(i, value)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Delete(i)
	}
}

func BenchmarkUpdate(b *testing.B) {
	db := NewBenchs()
	value := bench{}
	for i := 0; i < b.N; i++ {
		err := db.Insert(i, value)
		if err != nil {
			b.Fatal(err)
		}
	}
	value2 := bench{Name: "a"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = db.Update(i, value2)
	}
}

func BenchmarkUpdatePt(b *testing.B) {
	db := pt.NewPts()
	value := &pt.Pt{}
	for i := 0; i < b.N; i++ {
		err := db.Insert(i, value)
		if err != nil {
			b.Fatal(err)
		}
	}
	value2 := &pt.Pt{Name: "a"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = db.Update(i, value2)
	}
}

func BenchmarkUpdateName(b *testing.B) {
	db := NewBenchs()
	value := bench{}
	for i := 0; i < b.N; i++ {
		err := db.Insert(i, value)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = db.UpdateName(i, "a")
	}
}

func BenchmarkUpdateNamePt(b *testing.B) {
	db := pt.NewPts()
	value := &pt.Pt{}
	for i := 0; i < b.N; i++ {
		err := db.Insert(i, value)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = db.UpdateName(i, "a")
	}
}
