// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/login-door/pkg/db/ent/user"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetSalt sets the "salt" field.
func (uc *UserCreate) SetSalt(s string) *UserCreate {
	uc.mutation.SetSalt(s)
	return uc
}

// SetDisplayName sets the "display_name" field.
func (uc *UserCreate) SetDisplayName(s string) *UserCreate {
	uc.mutation.SetDisplayName(s)
	return uc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (uc *UserCreate) SetNillableDisplayName(s *string) *UserCreate {
	if s != nil {
		uc.SetDisplayName(*s)
	}
	return uc
}

// SetPhoneNumber sets the "phone_number" field.
func (uc *UserCreate) SetPhoneNumber(s string) *UserCreate {
	uc.mutation.SetPhoneNumber(s)
	return uc
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (uc *UserCreate) SetNillablePhoneNumber(s *string) *UserCreate {
	if s != nil {
		uc.SetPhoneNumber(*s)
	}
	return uc
}

// SetEmailAddress sets the "email_address" field.
func (uc *UserCreate) SetEmailAddress(s string) *UserCreate {
	uc.mutation.SetEmailAddress(s)
	return uc
}

// SetNillableEmailAddress sets the "email_address" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmailAddress(s *string) *UserCreate {
	if s != nil {
		uc.SetEmailAddress(*s)
	}
	return uc
}

// SetLoginTimes sets the "login_times" field.
func (uc *UserCreate) SetLoginTimes(u uint32) *UserCreate {
	uc.mutation.SetLoginTimes(u)
	return uc
}

// SetNillableLoginTimes sets the "login_times" field if the given value is not nil.
func (uc *UserCreate) SetNillableLoginTimes(u *uint32) *UserCreate {
	if u != nil {
		uc.SetLoginTimes(*u)
	}
	return uc
}

// SetKycVerify sets the "kyc_verify" field.
func (uc *UserCreate) SetKycVerify(b bool) *UserCreate {
	uc.mutation.SetKycVerify(b)
	return uc
}

// SetNillableKycVerify sets the "kyc_verify" field if the given value is not nil.
func (uc *UserCreate) SetNillableKycVerify(b *bool) *UserCreate {
	if b != nil {
		uc.SetKycVerify(*b)
	}
	return uc
}

// SetGaVerify sets the "ga_verify" field.
func (uc *UserCreate) SetGaVerify(b bool) *UserCreate {
	uc.mutation.SetGaVerify(b)
	return uc
}

// SetNillableGaVerify sets the "ga_verify" field if the given value is not nil.
func (uc *UserCreate) SetNillableGaVerify(b *bool) *UserCreate {
	if b != nil {
		uc.SetGaVerify(*b)
	}
	return uc
}

// SetSignupMethod sets the "signup_method" field.
func (uc *UserCreate) SetSignupMethod(s string) *UserCreate {
	uc.mutation.SetSignupMethod(s)
	return uc
}

// SetCreateAt sets the "create_at" field.
func (uc *UserCreate) SetCreateAt(u uint32) *UserCreate {
	uc.mutation.SetCreateAt(u)
	return uc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateAt(u *uint32) *UserCreate {
	if u != nil {
		uc.SetCreateAt(*u)
	}
	return uc
}

// SetUpdateAt sets the "update_at" field.
func (uc *UserCreate) SetUpdateAt(u uint32) *UserCreate {
	uc.mutation.SetUpdateAt(u)
	return uc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdateAt(u *uint32) *UserCreate {
	if u != nil {
		uc.SetUpdateAt(*u)
	}
	return uc
}

// SetDeleteAt sets the "delete_at" field.
func (uc *UserCreate) SetDeleteAt(u uint32) *UserCreate {
	uc.mutation.SetDeleteAt(u)
	return uc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeleteAt(u *uint32) *UserCreate {
	if u != nil {
		uc.SetDeleteAt(*u)
	}
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UserCreate) SetAvatar(s string) *UserCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatar(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatar(*s)
	}
	return uc
}

// SetRegion sets the "region" field.
func (uc *UserCreate) SetRegion(s string) *UserCreate {
	uc.mutation.SetRegion(s)
	return uc
}

// SetNillableRegion sets the "region" field if the given value is not nil.
func (uc *UserCreate) SetNillableRegion(s *string) *UserCreate {
	if s != nil {
		uc.SetRegion(*s)
	}
	return uc
}

// SetAge sets the "age" field.
func (uc *UserCreate) SetAge(u uint32) *UserCreate {
	uc.mutation.SetAge(u)
	return uc
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uc *UserCreate) SetNillableAge(u *uint32) *UserCreate {
	if u != nil {
		uc.SetAge(*u)
	}
	return uc
}

// SetGender sets the "gender" field.
func (uc *UserCreate) SetGender(s string) *UserCreate {
	uc.mutation.SetGender(s)
	return uc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uc *UserCreate) SetNillableGender(s *string) *UserCreate {
	if s != nil {
		uc.SetGender(*s)
	}
	return uc
}

// SetBirthday sets the "birthday" field.
func (uc *UserCreate) SetBirthday(s string) *UserCreate {
	uc.mutation.SetBirthday(s)
	return uc
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (uc *UserCreate) SetNillableBirthday(s *string) *UserCreate {
	if s != nil {
		uc.SetBirthday(*s)
	}
	return uc
}

// SetCountry sets the "country" field.
func (uc *UserCreate) SetCountry(s string) *UserCreate {
	uc.mutation.SetCountry(s)
	return uc
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (uc *UserCreate) SetNillableCountry(s *string) *UserCreate {
	if s != nil {
		uc.SetCountry(*s)
	}
	return uc
}

// SetProvince sets the "province" field.
func (uc *UserCreate) SetProvince(s string) *UserCreate {
	uc.mutation.SetProvince(s)
	return uc
}

// SetNillableProvince sets the "province" field if the given value is not nil.
func (uc *UserCreate) SetNillableProvince(s *string) *UserCreate {
	if s != nil {
		uc.SetProvince(*s)
	}
	return uc
}

// SetCity sets the "city" field.
func (uc *UserCreate) SetCity(s string) *UserCreate {
	uc.mutation.SetCity(s)
	return uc
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (uc *UserCreate) SetNillableCity(s *string) *UserCreate {
	if s != nil {
		uc.SetCity(*s)
	}
	return uc
}

// SetCareer sets the "career" field.
func (uc *UserCreate) SetCareer(s string) *UserCreate {
	uc.mutation.SetCareer(s)
	return uc
}

// SetNillableCareer sets the "career" field if the given value is not nil.
func (uc *UserCreate) SetNillableCareer(s *string) *UserCreate {
	if s != nil {
		uc.SetCareer(*s)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.DisplayName(); !ok {
		v := user.DefaultDisplayName
		uc.mutation.SetDisplayName(v)
	}
	if _, ok := uc.mutation.LoginTimes(); !ok {
		v := user.DefaultLoginTimes
		uc.mutation.SetLoginTimes(v)
	}
	if _, ok := uc.mutation.KycVerify(); !ok {
		v := user.DefaultKycVerify
		uc.mutation.SetKycVerify(v)
	}
	if _, ok := uc.mutation.GaVerify(); !ok {
		v := user.DefaultGaVerify
		uc.mutation.SetGaVerify(v)
	}
	if _, ok := uc.mutation.CreateAt(); !ok {
		v := user.DefaultCreateAt()
		uc.mutation.SetCreateAt(v)
	}
	if _, ok := uc.mutation.UpdateAt(); !ok {
		v := user.DefaultUpdateAt()
		uc.mutation.SetUpdateAt(v)
	}
	if _, ok := uc.mutation.DeleteAt(); !ok {
		v := user.DefaultDeleteAt()
		uc.mutation.SetDeleteAt(v)
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		v := user.DefaultAvatar
		uc.mutation.SetAvatar(v)
	}
	if _, ok := uc.mutation.Region(); !ok {
		v := user.DefaultRegion
		uc.mutation.SetRegion(v)
	}
	if _, ok := uc.mutation.Age(); !ok {
		v := user.DefaultAge
		uc.mutation.SetAge(v)
	}
	if _, ok := uc.mutation.Gender(); !ok {
		v := user.DefaultGender
		uc.mutation.SetGender(v)
	}
	if _, ok := uc.mutation.Birthday(); !ok {
		v := user.DefaultBirthday
		uc.mutation.SetBirthday(v)
	}
	if _, ok := uc.mutation.Country(); !ok {
		v := user.DefaultCountry
		uc.mutation.SetCountry(v)
	}
	if _, ok := uc.mutation.Province(); !ok {
		v := user.DefaultProvince
		uc.mutation.SetProvince(v)
	}
	if _, ok := uc.mutation.City(); !ok {
		v := user.DefaultCity
		uc.mutation.SetCity(v)
	}
	if _, ok := uc.mutation.Career(); !ok {
		v := user.DefaultCareer
		uc.mutation.SetCareer(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "username"`)}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "password"`)}
	}
	if _, ok := uc.mutation.Salt(); !ok {
		return &ValidationError{Name: "salt", err: errors.New(`ent: missing required field "salt"`)}
	}
	if _, ok := uc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`ent: missing required field "display_name"`)}
	}
	if _, ok := uc.mutation.LoginTimes(); !ok {
		return &ValidationError{Name: "login_times", err: errors.New(`ent: missing required field "login_times"`)}
	}
	if _, ok := uc.mutation.KycVerify(); !ok {
		return &ValidationError{Name: "kyc_verify", err: errors.New(`ent: missing required field "kyc_verify"`)}
	}
	if _, ok := uc.mutation.GaVerify(); !ok {
		return &ValidationError{Name: "ga_verify", err: errors.New(`ent: missing required field "ga_verify"`)}
	}
	if _, ok := uc.mutation.SignupMethod(); !ok {
		return &ValidationError{Name: "signup_method", err: errors.New(`ent: missing required field "signup_method"`)}
	}
	if _, ok := uc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := uc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := uc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`ent: missing required field "avatar"`)}
	}
	if _, ok := uc.mutation.Region(); !ok {
		return &ValidationError{Name: "region", err: errors.New(`ent: missing required field "region"`)}
	}
	if _, ok := uc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "age"`)}
	}
	if _, ok := uc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "gender"`)}
	}
	if _, ok := uc.mutation.Birthday(); !ok {
		return &ValidationError{Name: "birthday", err: errors.New(`ent: missing required field "birthday"`)}
	}
	if _, ok := uc.mutation.Country(); !ok {
		return &ValidationError{Name: "country", err: errors.New(`ent: missing required field "country"`)}
	}
	if _, ok := uc.mutation.Province(); !ok {
		return &ValidationError{Name: "province", err: errors.New(`ent: missing required field "province"`)}
	}
	if _, ok := uc.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New(`ent: missing required field "city"`)}
	}
	if _, ok := uc.mutation.Career(); !ok {
		return &ValidationError{Name: "career", err: errors.New(`ent: missing required field "career"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := uc.mutation.Salt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSalt,
		})
		_node.Salt = value
	}
	if value, ok := uc.mutation.DisplayName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDisplayName,
		})
		_node.DisplayName = value
	}
	if value, ok := uc.mutation.PhoneNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhoneNumber,
		})
		_node.PhoneNumber = value
	}
	if value, ok := uc.mutation.EmailAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmailAddress,
		})
		_node.EmailAddress = value
	}
	if value, ok := uc.mutation.LoginTimes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldLoginTimes,
		})
		_node.LoginTimes = value
	}
	if value, ok := uc.mutation.KycVerify(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldKycVerify,
		})
		_node.KycVerify = value
	}
	if value, ok := uc.mutation.GaVerify(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldGaVerify,
		})
		_node.GaVerify = value
	}
	if value, ok := uc.mutation.SignupMethod(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSignupMethod,
		})
		_node.SignupMethod = value
	}
	if value, ok := uc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := uc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := uc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatar,
		})
		_node.Avatar = value
	}
	if value, ok := uc.mutation.Region(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRegion,
		})
		_node.Region = value
	}
	if value, ok := uc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldAge,
		})
		_node.Age = value
	}
	if value, ok := uc.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldGender,
		})
		_node.Gender = value
	}
	if value, ok := uc.mutation.Birthday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldBirthday,
		})
		_node.Birthday = value
	}
	if value, ok := uc.mutation.Country(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldCountry,
		})
		_node.Country = value
	}
	if value, ok := uc.mutation.Province(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldProvince,
		})
		_node.Province = value
	}
	if value, ok := uc.mutation.City(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldCity,
		})
		_node.City = value
	}
	if value, ok := uc.mutation.Career(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldCareer,
		})
		_node.Career = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
