# marlin-eventbus

This repository contains the implementation of `marlin-eventbus`, a
lightweight synchronous event bus used across Silverpine's service mesh
adapters.

## Status

Released - v0.9.x is the current supported line.

## Example

```go
bus := eventbus.New()
bus.Subscribe("order.created", handler)
bus.Publish("order.created", payload)
```
