package traily

import (
    "encoding/json"
    "time"

    "github.com/google/uuid"
)

type Entry struct {
    ID         uuid.UUID        `json:"id"`
    ResourceID uuid.UUID        `json:"resource_id"`
    ActorType  string           `json:"actor_type"`
    ActorID    *uuid.UUID       `json:"actor_id"`
    Action     string           `json:"action"`
    State      State            `json:"state"`
    Metadata   *json.RawMessage `json:"metadata,omitempty"`
    CreatedAt  time.Time        `json:"created_at"`
}
