package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Hzw holds the schema definition for the Hzw entity.
type Hzw struct {
	ent.Schema
}

// Fields of the Hzw.
func (Hzw) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("HZW_ID").StructTag(`json:"hzwId,omitempty"`),
		field.String("stuNo").StorageKey("STU_NO"),
		field.String("name").StorageKey("NAME"),
		field.Int("age").StorageKey("AGE"),
		field.Int32("version").StorageKey("VERSION").Default(0),
	}
}

// Edges of the Hzw.
func (Hzw) Edges() []ent.Edge {
	return nil
}

func (Hzw) Annotations() []schema.Annotation {
	return []schema.Annotation{
		//		field.ID("stuNo", "name"),
		entsql.Annotation{Table: "hzw"},
		//		entsql.Skip()
	}
}
