// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/webmin7761/go-school/homework/final/internal/data/ent/fare"
	"github.com/webmin7761/go-school/homework/final/internal/data/ent/predicate"
)

// FareUpdate is the builder for updating Fare entities.
type FareUpdate struct {
	config
	hooks    []Hook
	mutation *FareMutation
}

// Where adds a new predicate for the FareUpdate builder.
func (fu *FareUpdate) Where(ps ...predicate.Fare) *FareUpdate {
	fu.mutation.predicates = append(fu.mutation.predicates, ps...)
	return fu
}

// SetOrgAirport sets the "org_airport" field.
func (fu *FareUpdate) SetOrgAirport(s string) *FareUpdate {
	fu.mutation.SetOrgAirport(s)
	return fu
}

// SetArrAirport sets the "arr_airport" field.
func (fu *FareUpdate) SetArrAirport(s string) *FareUpdate {
	fu.mutation.SetArrAirport(s)
	return fu
}

// SetPassageType sets the "passage_type" field.
func (fu *FareUpdate) SetPassageType(s string) *FareUpdate {
	fu.mutation.SetPassageType(s)
	return fu
}

// SetFirstTravelDate sets the "first_travel_date" field.
func (fu *FareUpdate) SetFirstTravelDate(t time.Time) *FareUpdate {
	fu.mutation.SetFirstTravelDate(t)
	return fu
}

// SetNillableFirstTravelDate sets the "first_travel_date" field if the given value is not nil.
func (fu *FareUpdate) SetNillableFirstTravelDate(t *time.Time) *FareUpdate {
	if t != nil {
		fu.SetFirstTravelDate(*t)
	}
	return fu
}

// SetLastTravelDate sets the "last_travel_date" field.
func (fu *FareUpdate) SetLastTravelDate(t time.Time) *FareUpdate {
	fu.mutation.SetLastTravelDate(t)
	return fu
}

// SetNillableLastTravelDate sets the "last_travel_date" field if the given value is not nil.
func (fu *FareUpdate) SetNillableLastTravelDate(t *time.Time) *FareUpdate {
	if t != nil {
		fu.SetLastTravelDate(*t)
	}
	return fu
}

// SetAmount sets the "amount" field.
func (fu *FareUpdate) SetAmount(f float64) *FareUpdate {
	fu.mutation.ResetAmount()
	fu.mutation.SetAmount(f)
	return fu
}

// AddAmount adds f to the "amount" field.
func (fu *FareUpdate) AddAmount(f float64) *FareUpdate {
	fu.mutation.AddAmount(f)
	return fu
}

// SetCreatedAt sets the "created_at" field.
func (fu *FareUpdate) SetCreatedAt(t time.Time) *FareUpdate {
	fu.mutation.SetCreatedAt(t)
	return fu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fu *FareUpdate) SetNillableCreatedAt(t *time.Time) *FareUpdate {
	if t != nil {
		fu.SetCreatedAt(*t)
	}
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FareUpdate) SetUpdatedAt(t time.Time) *FareUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fu *FareUpdate) SetNillableUpdatedAt(t *time.Time) *FareUpdate {
	if t != nil {
		fu.SetUpdatedAt(*t)
	}
	return fu
}

// Mutation returns the FareMutation object of the builder.
func (fu *FareUpdate) Mutation() *FareMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FareUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FareUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FareUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FareUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FareUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fare.Table,
			Columns: fare.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: fare.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.OrgAirport(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fare.FieldOrgAirport,
		})
	}
	if value, ok := fu.mutation.ArrAirport(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fare.FieldArrAirport,
		})
	}
	if value, ok := fu.mutation.PassageType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fare.FieldPassageType,
		})
	}
	if value, ok := fu.mutation.FirstTravelDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fare.FieldFirstTravelDate,
		})
	}
	if value, ok := fu.mutation.LastTravelDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fare.FieldLastTravelDate,
		})
	}
	if value, ok := fu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fare.FieldAmount,
		})
	}
	if value, ok := fu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fare.FieldAmount,
		})
	}
	if value, ok := fu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fare.FieldCreatedAt,
		})
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fare.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fare.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FareUpdateOne is the builder for updating a single Fare entity.
type FareUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FareMutation
}

// SetOrgAirport sets the "org_airport" field.
func (fuo *FareUpdateOne) SetOrgAirport(s string) *FareUpdateOne {
	fuo.mutation.SetOrgAirport(s)
	return fuo
}

// SetArrAirport sets the "arr_airport" field.
func (fuo *FareUpdateOne) SetArrAirport(s string) *FareUpdateOne {
	fuo.mutation.SetArrAirport(s)
	return fuo
}

// SetPassageType sets the "passage_type" field.
func (fuo *FareUpdateOne) SetPassageType(s string) *FareUpdateOne {
	fuo.mutation.SetPassageType(s)
	return fuo
}

// SetFirstTravelDate sets the "first_travel_date" field.
func (fuo *FareUpdateOne) SetFirstTravelDate(t time.Time) *FareUpdateOne {
	fuo.mutation.SetFirstTravelDate(t)
	return fuo
}

// SetNillableFirstTravelDate sets the "first_travel_date" field if the given value is not nil.
func (fuo *FareUpdateOne) SetNillableFirstTravelDate(t *time.Time) *FareUpdateOne {
	if t != nil {
		fuo.SetFirstTravelDate(*t)
	}
	return fuo
}

// SetLastTravelDate sets the "last_travel_date" field.
func (fuo *FareUpdateOne) SetLastTravelDate(t time.Time) *FareUpdateOne {
	fuo.mutation.SetLastTravelDate(t)
	return fuo
}

// SetNillableLastTravelDate sets the "last_travel_date" field if the given value is not nil.
func (fuo *FareUpdateOne) SetNillableLastTravelDate(t *time.Time) *FareUpdateOne {
	if t != nil {
		fuo.SetLastTravelDate(*t)
	}
	return fuo
}

// SetAmount sets the "amount" field.
func (fuo *FareUpdateOne) SetAmount(f float64) *FareUpdateOne {
	fuo.mutation.ResetAmount()
	fuo.mutation.SetAmount(f)
	return fuo
}

// AddAmount adds f to the "amount" field.
func (fuo *FareUpdateOne) AddAmount(f float64) *FareUpdateOne {
	fuo.mutation.AddAmount(f)
	return fuo
}

// SetCreatedAt sets the "created_at" field.
func (fuo *FareUpdateOne) SetCreatedAt(t time.Time) *FareUpdateOne {
	fuo.mutation.SetCreatedAt(t)
	return fuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fuo *FareUpdateOne) SetNillableCreatedAt(t *time.Time) *FareUpdateOne {
	if t != nil {
		fuo.SetCreatedAt(*t)
	}
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FareUpdateOne) SetUpdatedAt(t time.Time) *FareUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fuo *FareUpdateOne) SetNillableUpdatedAt(t *time.Time) *FareUpdateOne {
	if t != nil {
		fuo.SetUpdatedAt(*t)
	}
	return fuo
}

// Mutation returns the FareMutation object of the builder.
func (fuo *FareUpdateOne) Mutation() *FareMutation {
	return fuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FareUpdateOne) Select(field string, fields ...string) *FareUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Fare entity.
func (fuo *FareUpdateOne) Save(ctx context.Context) (*Fare, error) {
	var (
		err  error
		node *Fare
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FareUpdateOne) SaveX(ctx context.Context) *Fare {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FareUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FareUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FareUpdateOne) sqlSave(ctx context.Context) (_node *Fare, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fare.Table,
			Columns: fare.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: fare.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Fare.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fare.FieldID)
		for _, f := range fields {
			if !fare.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fare.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.OrgAirport(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fare.FieldOrgAirport,
		})
	}
	if value, ok := fuo.mutation.ArrAirport(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fare.FieldArrAirport,
		})
	}
	if value, ok := fuo.mutation.PassageType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fare.FieldPassageType,
		})
	}
	if value, ok := fuo.mutation.FirstTravelDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fare.FieldFirstTravelDate,
		})
	}
	if value, ok := fuo.mutation.LastTravelDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fare.FieldLastTravelDate,
		})
	}
	if value, ok := fuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fare.FieldAmount,
		})
	}
	if value, ok := fuo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fare.FieldAmount,
		})
	}
	if value, ok := fuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fare.FieldCreatedAt,
		})
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fare.FieldUpdatedAt,
		})
	}
	_node = &Fare{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fare.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
