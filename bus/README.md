# Event Bus

## Ben Matthews 

### 2020-01-18

#### Template for an event bus - event service for microservice patterns

Must have an API to receive events (HTTP POST) (w. validation?)

Get requests (HTTP GET) with search function to return events relevant to that service (searching on Consumed_by)

Connection to MongoDB instance to archive old events (or ones without an expiry date)

Uses gorllia mux

Messages must always follow the following pattern:

``` JSON
{
    "_ID": "",
    "Event": "event_id", // this is what the services are subscribed to
    "Posted_by": "",
    "Posted_on": "",
    "Expiry":,
    "Payload": { // Must be as concise as possible - only PK for new information

        "Event_identifier": "BSM01"
    },
    "Consumed_by": {
        // list of services that have consumed that service (value is number of times)
        }
}
```

Go struct representation of this pattern:

``` Go
type Event struct {
    ID        string
    PostedBy   string
    PostedOn   time.Time
    Expiry     int64
    Payload    map[string]string
    ConsumedBy map[string]int
}
```

Example:

``` JSON
{
    "_ID": "xxxx",
    "Event": "new-user",
    "Posted_by": "service_id",
    "Posted_on": "2020-01-01 00:01:02.33",
    "Expiry": 3600, // seconds
    "Payload": {
        "Event": "new-user",
        "Event_identifier": "BSM01"
    },
    "Consumed_by": {
        "<service_id1>": 1,
        "<service_id2>": 3,
        "<service_id3>": 0
        }
}
```

The api to query this looks like:

``` HTTP
/services/{event}/consumption/{n_consumed}
```

`{event}` Refers to the event ID that a handler will subscribe to

`{n_consumed}` Refers to the maximum number of times that an event can be 

This will return all events in the buffer

``` HTTP
/services/
```

This will return all events with nil consumption (orphaned events)

``` HTTP
/services/{event}/consumption/0
```
