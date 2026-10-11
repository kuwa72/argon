# argon

Experiment and verification of Go struct embedding, method inheritance behaviors, and reflection.

## Background Question

- **Q**: Can an embedded parent struct inspect the outer child struct type via reflection from within the parent's method?
- **A**: No, receiver methods on embedded structs receive only the parent type instance without dynamic virtual dispatch.
