// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package gensql

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Dummy struct {
	ID   uuid.UUID
	Data json.RawMessage
}
