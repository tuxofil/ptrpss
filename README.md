# Publish/Subscribe Storage

## Summary

Provides API for communicating using publish/subscribe pattern.
Once a message written to a topic, it will be received only once
by every topic subscriber existing at the moment of publishing.

## Example

```
import pubsub

storage := pubsub.NewStorage()

// Consuming entity: subscribe to a topic
var (
    topic = "topic1"
    subscriber = "subscriber1"
)
storage.Subscribe(topic, subscriber)

// Producing entity: send some data to the topic
storage.Publish(topic, []byte("byte slice"))

// Consuming entity: read data from the topic
data, err := storage.Poll(topic, subscriber)
```

## Testing

```
make test
```

## TODO

At the moment it's not ready for production for obvious reasons:

* there is no memory usage limitation;
* it is not thread safe;
* it doesn't use advantages of multicore systems.

And those are the main targets for further improvements.
