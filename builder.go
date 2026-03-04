package traily

import (
    "encoding/json"
    "time"

    "github.com/google/uuid"
)

type Builder struct {
    resourceID uuid.UUID
    actorType  string
    actorID    *uuid.UUID
    current    Entry
    entries    []Entry
    metadata   map[string]any
}

func Start(resourceID uuid.UUID, actorType string, actorID *uuid.UUID) *Builder {
    return &Builder{
        resourceID: resourceID,
        actorType:  actorType,
        actorID:    actorID,
        entries:    make([]Entry, 0),
        current:    Entry{State: StateUnset},
    }
}

func (b *Builder) Action(a string) *Builder    { b.current.Action = a; return b }
func (b *Builder) Actor(a string) *Builder     { b.actorType = a; b.current.ActorType = a; return b }
func (b *Builder) State(s State) *Builder      { b.current.State = s; return b }
func (b *Builder) ActorType() string           { return b.current.ActorType }

func (b *Builder) Meta(key string, value any) *Builder {
    if b.metadata == nil {
        b.metadata = make(map[string]any)
    }
    b.metadata[key] = value
    return b
}

func (b *Builder) Fail(err error, reason string) error {
    b.State(StateFailed).Meta("reason", reason)
    return err
}

func (b *Builder) Entries() []Entry { return b.entries }

func (b *Builder) Emit() {
    if b.current.Action == "" || b.resourceID == uuid.Nil {
        return
    }

    b.current.ID = uuid.New()
    b.current.ResourceID = b.resourceID
    b.current.ActorID = b.actorID
    b.current.CreatedAt = time.Now()

    if b.current.State == "" {
        b.current.State = StateUnset
    }
    if b.current.ActorType == "" {
        b.current.ActorType = b.actorType
    }
    if b.metadata != nil {
        if blob, err := json.Marshal(b.metadata); err == nil {
            raw := json.RawMessage(blob)
            b.current.Metadata = &raw
        }
    }

    b.entries = append(b.entries, b.current)
    b.metadata = nil
    b.current = Entry{State: StateUnset}
}
