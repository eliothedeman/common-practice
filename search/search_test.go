package search

import (
	"fmt"
	"testing"
)

func TestSplitData(t *testing.T) {
	d := "asdfgasdfasdkfjasldkjfal;sdjf;alsdkf"
	keys := SplitData(d)
	fmt.Println(keys)
}
