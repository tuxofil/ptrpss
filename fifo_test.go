package pubsub

import "testing"

func TestFifo(t *testing.T) {
	var fifo Fifo
	if v := fifo.Extract(); v != nil {
		t.Fatalf("unexpected value extracted: %#v", v)
	}
	for i := 0; i < 5; i++ {
		fifo.Append(i)
	}
	for i := 0; i < 3; i++ {
		v, ok := fifo.Extract().(int)
		if !ok || v != i {
			t.Fatalf("expected %d but found: %#v", i, v)
		}
	}
	for _, v := range []string{"a", "b", "c"} {
		fifo.Append(v)
	}
	for _, v1 := range []interface{}{3, 4, "a", "b", "c", nil} {
		if v2 := fifo.Extract(); v2 != v1 {
			t.Fatalf("expected %#v but found: %#v", v1, v2)
		}
	}
}
