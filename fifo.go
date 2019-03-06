/**
Simple FIFO queue based on linked list.
*/

package pubsub

// Linked list, used as FIFO queue.
type Fifo struct {
	// Pointer to the oldest entry
	first *FifoEntry
	// Pointer to the newest entry
	last *FifoEntry
}

// Linked list element.
type FifoEntry struct {
	// The value itself
	value interface{}
	// Pointer to the next (more recent) entry
	newer *FifoEntry
}

// Append a value to the end of the list.
func (a *Fifo) Append(v interface{}) {
	entry := &FifoEntry{value: v}
	if a.first == nil {
		a.first = entry
		a.last = a.first
		return
	}
	a.last.newer = entry
	a.last = entry
}

// Get the oldest value from the list, if any.
func (a *Fifo) Extract() interface{} {
	if a.first == nil {
		return nil
	}
	value := a.first.value
	a.first = a.first.newer
	return value
}
