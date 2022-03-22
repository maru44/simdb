package main

func main() {
	// dbdb := NewTableA()
	dbdb := TableAs{
		Data: map[uint]TableA{},
	}
	dbdb.Insert(88, TableA{})
}
