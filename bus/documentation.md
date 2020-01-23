# Event Queue Documentation

## v0.1

## copyright 2020 Ben Matthews

The JSON requests must conform to the following standard

GET requests will return:

``` JSON
{
 "UID": "a uid auto created by the queue service",
 "EventType": "new-user",
 "Posted_by": "service_id",
 "Posted_on": "<ISO 8601>",
 "Published_on": "<ISO 8601>",
 "Expiry": 3600, // Time in seconds between Published_on and Archiving
 "Payload": { // UID for the consuming service to identify important information without race conditions
  "Identifier": "BSM01"
 },
 "Consumed_by": { // List of services that have consumed that service n times
  "service_id1": 0,
  "service_id2": 1,
  "service_id3": 2
  ...
 }
}
```

POST requests for *new events* must be in the following format:

They will return a status 200 (or 400) if not parsed correctly

``` JSON
{
 "EventType": "new-user",
 "Posted_by": "service_id",
 "Posted_on": "ISO 8601",
 "Expiry": 3600, // Time in seconds between Published_on and Archiving
 "Payload": { // UID for the consuming service to identify important information without race conditions
  "Identifier": "BSM01"
 }
}
```

POST requests for *consumed events* are not transferred with JSON - but are in the URL of the request. These will increment the value of the ```Consumed_by``` key by N.
