package pubsub

import "testing"

type PublishTask struct {
	Topic string
	Value interface{}
}

type SubscribeTask struct {
	Topic      string
	Subscriber string
}

type UnsubscribeTask struct {
	Topic      string
	Subscriber string
}

type PollTask struct {
	Topic       string
	Subscriber  string
	ExpectError bool
	ExpectValue interface{}
}

func TestStorage(t *testing.T) {
	testset := []interface{}{
		PollTask{"topic1", "sub1", true, nil},
		PollTask{"topic2", "sub1", true, nil},
		SubscribeTask{"topic1", "sub1"},
		PollTask{"topic1", "sub1", false, nil},
		PollTask{"topic2", "sub1", true, nil},
		PublishTask{"topic2", 1},
		PollTask{"topic1", "sub1", false, nil},
		PollTask{"topic2", "sub1", true, nil},
		PublishTask{"topic1", 2},
		PollTask{"topic1", "sub1", false, 2},
		PollTask{"topic1", "sub1", false, nil},
		PollTask{"topic1", "sub2", false, nil},
		PollTask{"topic2", "sub2", true, nil},
		PublishTask{"topic1", 3},
		SubscribeTask{"topic1", "sub2"},
		PublishTask{"topic1", "4"},
		PublishTask{"topic1", "five"},
		PollTask{"topic1", "sub1", false, 3},
		PollTask{"topic1", "sub2", false, "4"},
		PollTask{"topic1", "sub1", false, "4"},
		PollTask{"topic1", "sub2", false, "five"},
		PollTask{"topic1", "sub1", false, "five"},
		PollTask{"topic1", "sub2", false, nil},
		PollTask{"topic1", "sub1", false, nil},
		PollTask{"topic1", "sub2", false, nil},
		PublishTask{"topic1", 6},
		PublishTask{"topic1", 7},
		PollTask{"topic1", "sub1", false, 6},
		PollTask{"topic1", "sub2", false, 6},
		UnsubscribeTask{"topic1", "sub1"},
		PublishTask{"topic1", 8},
		PollTask{"topic1", "sub1", false, nil},
		PollTask{"topic1", "sub2", false, 7},
		PollTask{"topic1", "sub1", false, nil},
		PollTask{"topic1", "sub2", false, 8},
		PollTask{"topic1", "sub1", false, nil},
		PollTask{"topic1", "sub2", false, nil},
	}
	s := NewStorage()
	for n, task0 := range testset {
		if task, ok := task0.(PublishTask); ok {
			s.Publish(task.Topic, task.Value)
		} else if task, ok := task0.(SubscribeTask); ok {
			s.Subscribe(task.Topic, task.Subscriber)
		} else if task, ok := task0.(UnsubscribeTask); ok {
			s.Unsubscribe(task.Topic, task.Subscriber)
		} else if task, ok := task0.(PollTask); ok {
			value, err := s.Poll(task.Topic, task.Subscriber)
			if task.ExpectError && err == nil {
				t.Fatalf("#%d> %#v> unexpected poll success: %#v", n, task, value)
			} else if !task.ExpectError && err != nil {
				t.Fatalf("#%d> %#v> unexpected poll error: %s", n, task, err)
			} else if !task.ExpectError && value != task.ExpectValue {
				t.Fatalf("#%d> %#v> unexpected poll result: %#v", n, task, value)
			}
		} else {
			t.Fatalf("#%d> unknown task type: %#v", n, task0)
		}
		t.Logf("#%03d> task OK %#v", n, task0)
	}
}
