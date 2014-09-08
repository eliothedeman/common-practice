package storage

import (
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	r, _ := NewRedis(":6379", 10)
	r.Set("test", []byte("yup"))
	b, err := r.Get("test")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
