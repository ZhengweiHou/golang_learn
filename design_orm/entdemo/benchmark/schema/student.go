package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StorageKey("ID").Positive(),
		field.String("name").StorageKey("NAME"),
		field.Int("age").StorageKey("AGE"),
		field.Int32("version").StorageKey("VERSION").Default(0),
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}

func (Student) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "student"},
	}
}
