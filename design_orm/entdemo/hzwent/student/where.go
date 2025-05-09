// Code generated by ent, DO NOT EDIT.

package student

import (
	"entdemo/hzwent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldID, id))
}

// StuNo applies equality check predicate on the "stuNo" field. It's identical to StuNoEQ.
func StuNo(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldStuNo, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldName, v))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAge, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int32) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldVersion, v))
}

// StuNoEQ applies the EQ predicate on the "stuNo" field.
func StuNoEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldStuNo, v))
}

// StuNoNEQ applies the NEQ predicate on the "stuNo" field.
func StuNoNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldStuNo, v))
}

// StuNoIn applies the In predicate on the "stuNo" field.
func StuNoIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldStuNo, vs...))
}

// StuNoNotIn applies the NotIn predicate on the "stuNo" field.
func StuNoNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldStuNo, vs...))
}

// StuNoGT applies the GT predicate on the "stuNo" field.
func StuNoGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldStuNo, v))
}

// StuNoGTE applies the GTE predicate on the "stuNo" field.
func StuNoGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldStuNo, v))
}

// StuNoLT applies the LT predicate on the "stuNo" field.
func StuNoLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldStuNo, v))
}

// StuNoLTE applies the LTE predicate on the "stuNo" field.
func StuNoLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldStuNo, v))
}

// StuNoContains applies the Contains predicate on the "stuNo" field.
func StuNoContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldStuNo, v))
}

// StuNoHasPrefix applies the HasPrefix predicate on the "stuNo" field.
func StuNoHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldStuNo, v))
}

// StuNoHasSuffix applies the HasSuffix predicate on the "stuNo" field.
func StuNoHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldStuNo, v))
}

// StuNoEqualFold applies the EqualFold predicate on the "stuNo" field.
func StuNoEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldStuNo, v))
}

// StuNoContainsFold applies the ContainsFold predicate on the "stuNo" field.
func StuNoContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldStuNo, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldName, v))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldAge, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int32) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int32) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int32) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int32) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int32) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int32) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int32) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int32) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldVersion, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Student) predicate.Student {
	return predicate.Student(sql.NotPredicates(p))
}
