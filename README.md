# grocket

A dynamic task scheduler service

Grocket is based on the abstraction of an Event.

And Event is characterized by:

* an endpoint that the event is pushed to (a REST API call)
* a time at which the event should fire
* a payload, which is a string (could be JSON, could be base64-encoded binary, etc.)
* a REST verb to use when calling the endpoint (defaults to PUT)
* a timeout after which the event should be discarded (optionally, defaults to infinite)
* a max-attempts number (optional, defaults to infinite)

````
PUT /events {time: <long/timestamp>,
             payload: <string>,
             expiry: <long/timestamp>,
             endpoint: "https://service.domain.tld/collection",
             max-attempts: 1,
             verb: <PUT|POST|DELETE|GET>}
````
