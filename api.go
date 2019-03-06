/**
Simple publish/subscribe storage.
*/

package pubsub

// Publish/subscribe storage
type Storage struct {
	// Map topic name to set of subscribers names
	topics map[string]map[string]struct{}
	// Map topic name and subscriber name to message storage
	subs map[SubID]*Fifo
}

// Subscription unique key
type SubID struct {
	// Topic name
	Topic string
	// Subscriber name
	SubName string
}

// NewStorage creates new instance of publish/subscribe
// in-memory storage.
func NewStorage() *Storage {
	return &Storage{
		topics: map[string]map[string]struct{}{},
		subs:   map[SubID]*Fifo{},
	}
}

// Publish document for all subscribers of given topic.
// To keep every message in the memory only once(*) one should
// consider to publish pointers instead of raw values.
//
// * of course, we still need mem to store the pointer for each subscriber.
func (s *Storage) Publish(topicName string, doc interface{}) {
	if subscribers, ok := s.topics[topicName]; ok {
		var id = SubID{Topic: topicName}
		for subscriberName := range subscribers {
			id.SubName = subscriberName
			fifo, ok := s.subs[id]
			if !ok {
				fifo = &Fifo{}
				s.subs[id] = fifo
			}
			fifo.Append(doc)
		}
	}
}

// Subscribe to JSON documents in given topic.
func (s *Storage) Subscribe(topicName, subscriberName string) {
	subscribers, ok := s.topics[topicName]
	if !ok {
		subscribers = map[string]struct{}{subscriberName: struct{}{}}
		s.topics[topicName] = subscribers
	} else if _, ok := subscribers[subscriberName]; !ok {
		subscribers[subscriberName] = struct{}{}
	}
}

// Unsubscribe discards existing subscription, if any.
func (s *Storage) Unsubscribe(topicName, subscriberName string) {
	if subscribers, ok := s.topics[topicName]; ok {
		delete(subscribers, subscriberName)
		delete(s.subs, SubID{topicName, subscriberName})
	}
}

// Poll returns either next unseen message or no message
// if all are seen or error if no subscription found.
// FIXME: there is no way to distinguish absence of messages
// and published nil value.
func (s *Storage) Poll(topicName, subscriberName string) (interface{}, error) {
	if fifo, ok := s.subs[SubID{topicName, subscriberName}]; ok {
		return fifo.Extract(), nil
	} else if _, ok := s.topics[topicName]; !ok {
		return nil, NoSuchTopicError
	}
	// subscription exists but message queue is still empty
	return nil, nil
}
