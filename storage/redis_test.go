package storage

import "testing"

// TestSetGet must have a redis server running on localhost:6379
func TestSetGet(t *testing.T) {
	expected := "test"
	r, _ := NewRedis(":6379", 10)
	r.Set("test", []byte(expected))
	b, err := r.Get("test")
	if err != nil {
		t.Fatal(err)
	}

	// test that b comes back with the correct value
	if string(b) != expected {
		t.Fatal("Expected: " + expected + " got: " + string(b))
	}
}

// BenchmarkSet benchmark the SET command to the redis server
func BenchmarkSetPoolsize10(b *testing.B) {
	setValue := []byte("set-test-value")
	r, _ := NewRedis(":6379", 100)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r.Set("set", setValue)
		}
	})

}
