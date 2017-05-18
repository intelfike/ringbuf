package ringbuf

import (
	"fmt"
	"testing"
)

func TestRingbuf(t *testing.T) {
	rb, err := New(make([]int, 10))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\tWrite 0 to 4")
	for n := 0; n < 5; n++ {
		rb.Write(n)
	}
	t.Log("Get =", rb.Get())
	t.Log("IndexNew =", rb.IndexNew(0))
	t.Log("IndexOld =", rb.IndexOld(0))
	t.Log("Len =", rb.Len())
	t.Log("Cap =", rb.Cap())
	fmt.Println()
	t.Log("\tWrite 0 to 20")
	for n := 0; n < 20; n++ {
		rb.Write(n)
	}
	t.Log("Get =", rb.Get())
	t.Log("IndexNew =", rb.IndexNew(0))
	t.Log("IndexOld =", rb.IndexOld(0))
	t.Log("Len =", rb.Len())
	t.Log("Cap =", rb.Cap())

}
