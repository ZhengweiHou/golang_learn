// Code generated by ent, DO NOT EDIT.

package hzwent

import (
	"entdemo/hzwent/hzw"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Hzw is the model entity for the Hzw schema.
type Hzw struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"hzwId,omitempty"`
	// StuNo holds the value of the "stuNo" field.
	StuNo string `json:"stuNo,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Version holds the value of the "version" field.
	Version      int32 `json:"version,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Hzw) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hzw.FieldAge, hzw.FieldVersion:
			values[i] = new(sql.NullInt64)
		case hzw.FieldStuNo, hzw.FieldName:
			values[i] = new(sql.NullString)
		case hzw.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Hzw fields.
func (h *Hzw) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hzw.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				h.ID = *value
			}
		case hzw.FieldStuNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stuNo", values[i])
			} else if value.Valid {
				h.StuNo = value.String
			}
		case hzw.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				h.Name = value.String
			}
		case hzw.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				h.Age = int(value.Int64)
			}
		case hzw.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				h.Version = int32(value.Int64)
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Hzw.
// This includes values selected through modifiers, order, etc.
func (h *Hzw) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// Update returns a builder for updating this Hzw.
// Note that you need to call Hzw.Unwrap() before calling this method if this Hzw
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Hzw) Update() *HzwUpdateOne {
	return NewHzwClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the Hzw entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Hzw) Unwrap() *Hzw {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("hzwent: Hzw is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Hzw) String() string {
	var builder strings.Builder
	builder.WriteString("Hzw(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("stuNo=")
	builder.WriteString(h.StuNo)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(h.Name)
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", h.Age))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", h.Version))
	builder.WriteByte(')')
	return builder.String()
}

// Hzws is a parsable slice of Hzw.
type Hzws []*Hzw
