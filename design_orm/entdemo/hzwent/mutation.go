// Code generated by ent, DO NOT EDIT.

package hzwent

import (
	"context"
	"entdemo/hzwent/hzw"
	"entdemo/hzwent/predicate"
	"entdemo/hzwent/student"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeHzw     = "Hzw"
	TypeStudent = "Student"
)

// HzwMutation represents an operation that mutates the Hzw nodes in the graph.
type HzwMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	stuNo         *string
	name          *string
	age           *int
	addage        *int
	version       *int32
	addversion    *int32
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Hzw, error)
	predicates    []predicate.Hzw
}

var _ ent.Mutation = (*HzwMutation)(nil)

// hzwOption allows management of the mutation configuration using functional options.
type hzwOption func(*HzwMutation)

// newHzwMutation creates new mutation for the Hzw entity.
func newHzwMutation(c config, op Op, opts ...hzwOption) *HzwMutation {
	m := &HzwMutation{
		config:        c,
		op:            op,
		typ:           TypeHzw,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withHzwID sets the ID field of the mutation.
func withHzwID(id uuid.UUID) hzwOption {
	return func(m *HzwMutation) {
		var (
			err   error
			once  sync.Once
			value *Hzw
		)
		m.oldValue = func(ctx context.Context) (*Hzw, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Hzw.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withHzw sets the old Hzw of the mutation.
func withHzw(node *Hzw) hzwOption {
	return func(m *HzwMutation) {
		m.oldValue = func(context.Context) (*Hzw, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m HzwMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m HzwMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("hzwent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Hzw entities.
func (m *HzwMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *HzwMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *HzwMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Hzw.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetStuNo sets the "stuNo" field.
func (m *HzwMutation) SetStuNo(s string) {
	m.stuNo = &s
}

// StuNo returns the value of the "stuNo" field in the mutation.
func (m *HzwMutation) StuNo() (r string, exists bool) {
	v := m.stuNo
	if v == nil {
		return
	}
	return *v, true
}

// OldStuNo returns the old "stuNo" field's value of the Hzw entity.
// If the Hzw object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *HzwMutation) OldStuNo(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStuNo is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStuNo requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStuNo: %w", err)
	}
	return oldValue.StuNo, nil
}

// ResetStuNo resets all changes to the "stuNo" field.
func (m *HzwMutation) ResetStuNo() {
	m.stuNo = nil
}

// SetName sets the "name" field.
func (m *HzwMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *HzwMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Hzw entity.
// If the Hzw object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *HzwMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *HzwMutation) ResetName() {
	m.name = nil
}

// SetAge sets the "age" field.
func (m *HzwMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *HzwMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the Hzw entity.
// If the Hzw object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *HzwMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *HzwMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *HzwMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *HzwMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetVersion sets the "version" field.
func (m *HzwMutation) SetVersion(i int32) {
	m.version = &i
	m.addversion = nil
}

// Version returns the value of the "version" field in the mutation.
func (m *HzwMutation) Version() (r int32, exists bool) {
	v := m.version
	if v == nil {
		return
	}
	return *v, true
}

// OldVersion returns the old "version" field's value of the Hzw entity.
// If the Hzw object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *HzwMutation) OldVersion(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldVersion is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldVersion requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldVersion: %w", err)
	}
	return oldValue.Version, nil
}

// AddVersion adds i to the "version" field.
func (m *HzwMutation) AddVersion(i int32) {
	if m.addversion != nil {
		*m.addversion += i
	} else {
		m.addversion = &i
	}
}

// AddedVersion returns the value that was added to the "version" field in this mutation.
func (m *HzwMutation) AddedVersion() (r int32, exists bool) {
	v := m.addversion
	if v == nil {
		return
	}
	return *v, true
}

// ResetVersion resets all changes to the "version" field.
func (m *HzwMutation) ResetVersion() {
	m.version = nil
	m.addversion = nil
}

// Where appends a list predicates to the HzwMutation builder.
func (m *HzwMutation) Where(ps ...predicate.Hzw) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the HzwMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *HzwMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Hzw, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *HzwMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *HzwMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Hzw).
func (m *HzwMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *HzwMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.stuNo != nil {
		fields = append(fields, hzw.FieldStuNo)
	}
	if m.name != nil {
		fields = append(fields, hzw.FieldName)
	}
	if m.age != nil {
		fields = append(fields, hzw.FieldAge)
	}
	if m.version != nil {
		fields = append(fields, hzw.FieldVersion)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *HzwMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case hzw.FieldStuNo:
		return m.StuNo()
	case hzw.FieldName:
		return m.Name()
	case hzw.FieldAge:
		return m.Age()
	case hzw.FieldVersion:
		return m.Version()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *HzwMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case hzw.FieldStuNo:
		return m.OldStuNo(ctx)
	case hzw.FieldName:
		return m.OldName(ctx)
	case hzw.FieldAge:
		return m.OldAge(ctx)
	case hzw.FieldVersion:
		return m.OldVersion(ctx)
	}
	return nil, fmt.Errorf("unknown Hzw field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *HzwMutation) SetField(name string, value ent.Value) error {
	switch name {
	case hzw.FieldStuNo:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStuNo(v)
		return nil
	case hzw.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case hzw.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case hzw.FieldVersion:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetVersion(v)
		return nil
	}
	return fmt.Errorf("unknown Hzw field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *HzwMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, hzw.FieldAge)
	}
	if m.addversion != nil {
		fields = append(fields, hzw.FieldVersion)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *HzwMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case hzw.FieldAge:
		return m.AddedAge()
	case hzw.FieldVersion:
		return m.AddedVersion()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *HzwMutation) AddField(name string, value ent.Value) error {
	switch name {
	case hzw.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	case hzw.FieldVersion:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddVersion(v)
		return nil
	}
	return fmt.Errorf("unknown Hzw numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *HzwMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *HzwMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *HzwMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Hzw nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *HzwMutation) ResetField(name string) error {
	switch name {
	case hzw.FieldStuNo:
		m.ResetStuNo()
		return nil
	case hzw.FieldName:
		m.ResetName()
		return nil
	case hzw.FieldAge:
		m.ResetAge()
		return nil
	case hzw.FieldVersion:
		m.ResetVersion()
		return nil
	}
	return fmt.Errorf("unknown Hzw field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *HzwMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *HzwMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *HzwMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *HzwMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *HzwMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *HzwMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *HzwMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Hzw unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *HzwMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Hzw edge %s", name)
}

// StudentMutation represents an operation that mutates the Student nodes in the graph.
type StudentMutation struct {
	config
	op            Op
	typ           string
	id            *int
	stuNo         *string
	name          *string
	age           *int
	addage        *int
	version       *int32
	addversion    *int32
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Student, error)
	predicates    []predicate.Student
}

var _ ent.Mutation = (*StudentMutation)(nil)

// studentOption allows management of the mutation configuration using functional options.
type studentOption func(*StudentMutation)

// newStudentMutation creates new mutation for the Student entity.
func newStudentMutation(c config, op Op, opts ...studentOption) *StudentMutation {
	m := &StudentMutation{
		config:        c,
		op:            op,
		typ:           TypeStudent,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withStudentID sets the ID field of the mutation.
func withStudentID(id int) studentOption {
	return func(m *StudentMutation) {
		var (
			err   error
			once  sync.Once
			value *Student
		)
		m.oldValue = func(ctx context.Context) (*Student, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Student.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withStudent sets the old Student of the mutation.
func withStudent(node *Student) studentOption {
	return func(m *StudentMutation) {
		m.oldValue = func(context.Context) (*Student, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m StudentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m StudentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("hzwent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Student entities.
func (m *StudentMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *StudentMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *StudentMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Student.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetStuNo sets the "stuNo" field.
func (m *StudentMutation) SetStuNo(s string) {
	m.stuNo = &s
}

// StuNo returns the value of the "stuNo" field in the mutation.
func (m *StudentMutation) StuNo() (r string, exists bool) {
	v := m.stuNo
	if v == nil {
		return
	}
	return *v, true
}

// OldStuNo returns the old "stuNo" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldStuNo(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStuNo is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStuNo requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStuNo: %w", err)
	}
	return oldValue.StuNo, nil
}

// ResetStuNo resets all changes to the "stuNo" field.
func (m *StudentMutation) ResetStuNo() {
	m.stuNo = nil
}

// SetName sets the "name" field.
func (m *StudentMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *StudentMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *StudentMutation) ResetName() {
	m.name = nil
}

// SetAge sets the "age" field.
func (m *StudentMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *StudentMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *StudentMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *StudentMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *StudentMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetVersion sets the "version" field.
func (m *StudentMutation) SetVersion(i int32) {
	m.version = &i
	m.addversion = nil
}

// Version returns the value of the "version" field in the mutation.
func (m *StudentMutation) Version() (r int32, exists bool) {
	v := m.version
	if v == nil {
		return
	}
	return *v, true
}

// OldVersion returns the old "version" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldVersion(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldVersion is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldVersion requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldVersion: %w", err)
	}
	return oldValue.Version, nil
}

// AddVersion adds i to the "version" field.
func (m *StudentMutation) AddVersion(i int32) {
	if m.addversion != nil {
		*m.addversion += i
	} else {
		m.addversion = &i
	}
}

// AddedVersion returns the value that was added to the "version" field in this mutation.
func (m *StudentMutation) AddedVersion() (r int32, exists bool) {
	v := m.addversion
	if v == nil {
		return
	}
	return *v, true
}

// ResetVersion resets all changes to the "version" field.
func (m *StudentMutation) ResetVersion() {
	m.version = nil
	m.addversion = nil
}

// Where appends a list predicates to the StudentMutation builder.
func (m *StudentMutation) Where(ps ...predicate.Student) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the StudentMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *StudentMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Student, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *StudentMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *StudentMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Student).
func (m *StudentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *StudentMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.stuNo != nil {
		fields = append(fields, student.FieldStuNo)
	}
	if m.name != nil {
		fields = append(fields, student.FieldName)
	}
	if m.age != nil {
		fields = append(fields, student.FieldAge)
	}
	if m.version != nil {
		fields = append(fields, student.FieldVersion)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *StudentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case student.FieldStuNo:
		return m.StuNo()
	case student.FieldName:
		return m.Name()
	case student.FieldAge:
		return m.Age()
	case student.FieldVersion:
		return m.Version()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *StudentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case student.FieldStuNo:
		return m.OldStuNo(ctx)
	case student.FieldName:
		return m.OldName(ctx)
	case student.FieldAge:
		return m.OldAge(ctx)
	case student.FieldVersion:
		return m.OldVersion(ctx)
	}
	return nil, fmt.Errorf("unknown Student field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StudentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case student.FieldStuNo:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStuNo(v)
		return nil
	case student.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case student.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case student.FieldVersion:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetVersion(v)
		return nil
	}
	return fmt.Errorf("unknown Student field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *StudentMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, student.FieldAge)
	}
	if m.addversion != nil {
		fields = append(fields, student.FieldVersion)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *StudentMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case student.FieldAge:
		return m.AddedAge()
	case student.FieldVersion:
		return m.AddedVersion()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StudentMutation) AddField(name string, value ent.Value) error {
	switch name {
	case student.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	case student.FieldVersion:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddVersion(v)
		return nil
	}
	return fmt.Errorf("unknown Student numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *StudentMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *StudentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *StudentMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Student nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *StudentMutation) ResetField(name string) error {
	switch name {
	case student.FieldStuNo:
		m.ResetStuNo()
		return nil
	case student.FieldName:
		m.ResetName()
		return nil
	case student.FieldAge:
		m.ResetAge()
		return nil
	case student.FieldVersion:
		m.ResetVersion()
		return nil
	}
	return fmt.Errorf("unknown Student field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *StudentMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *StudentMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *StudentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *StudentMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *StudentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *StudentMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *StudentMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Student unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *StudentMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Student edge %s", name)
}
