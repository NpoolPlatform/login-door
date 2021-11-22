package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LoginRecord holds the schema definition for the LoginRecord entity.
type LoginRecord struct {
	ent.Schema
}

// Fields of the LoginRecord.
func (LoginRecord) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("app_id", uuid.UUID{}),
		field.Uint32("login_time").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.String("ip"),
		field.String("location"),
		field.Float("lat"),
		field.Float("lon"),
		field.String("timezone"),
	}
}

// Edges of the LoginRecord.
func (LoginRecord) Edges() []ent.Edge {
	return nil
}
