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
		field.String("stuNo").StorageKey("STU_NO"),
		//field.String("stuNo"), // StorageKey默认为转为蛇形命名即stu_no
		field.String("name").StorageKey("NAME"),
		field.Int("age").StorageKey("AGE"),
		field.Int32("version").StorageKey("VERSION").Default(0),
		//		field.String("msg1").MaxLen(20),
		//		field.Time("date"),
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
