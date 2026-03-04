package traily

type State string

const (
    StateSucceeded State = "succeeded"
    StateFailed    State = "failed"
    StatePending   State = "pending"
    StateUnset     State = "unset"
)
