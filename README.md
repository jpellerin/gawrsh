gawrsh
======

Go Activated Websocket Redis Streaming Helper

Gawrsh plays in the same space as the old, lamented Juggernaut, which provided non-realtime applications an easy way to hook up a real-time UI using websockets, with redis as the bridge between worlds. Gawrsh does more or less the same thing, while also serving as a static file and proxy server, so that it can host an application without the need for nginx or apache (in some cases).

The Redis-powered streaming websockets feature works simply. Each request to gawrsh's subscription url will return a websocket connection that listens on the requested redis pubsub channel. The core application may then publish json blobs to that channel, and those json blobs will be sent to all connected subscribers. 
