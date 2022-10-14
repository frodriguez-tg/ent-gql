// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"freg/ent/car"
	"freg/ent/group"
	"freg/ent/predicate"
	"freg/ent/user"
	"time"
)

// CarWhereInput represents a where input for filtering Car queries.
type CarWhereInput struct {
	Predicates []predicate.Car  `json:"-"`
	Not        *CarWhereInput   `json:"not,omitempty"`
	Or         []*CarWhereInput `json:"or,omitempty"`
	And        []*CarWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *string  `json:"id,omitempty"`
	IDNEQ   *string  `json:"idNEQ,omitempty"`
	IDIn    []string `json:"idIn,omitempty"`
	IDNotIn []string `json:"idNotIn,omitempty"`
	IDGT    *string  `json:"idGT,omitempty"`
	IDGTE   *string  `json:"idGTE,omitempty"`
	IDLT    *string  `json:"idLT,omitempty"`
	IDLTE   *string  `json:"idLTE,omitempty"`

	// "model" field predicates.
	Model             *string  `json:"model,omitempty"`
	ModelNEQ          *string  `json:"modelNEQ,omitempty"`
	ModelIn           []string `json:"modelIn,omitempty"`
	ModelNotIn        []string `json:"modelNotIn,omitempty"`
	ModelGT           *string  `json:"modelGT,omitempty"`
	ModelGTE          *string  `json:"modelGTE,omitempty"`
	ModelLT           *string  `json:"modelLT,omitempty"`
	ModelLTE          *string  `json:"modelLTE,omitempty"`
	ModelContains     *string  `json:"modelContains,omitempty"`
	ModelHasPrefix    *string  `json:"modelHasPrefix,omitempty"`
	ModelHasSuffix    *string  `json:"modelHasSuffix,omitempty"`
	ModelEqualFold    *string  `json:"modelEqualFold,omitempty"`
	ModelContainsFold *string  `json:"modelContainsFold,omitempty"`

	// "registered_at" field predicates.
	RegisteredAt      *time.Time  `json:"registeredAt,omitempty"`
	RegisteredAtNEQ   *time.Time  `json:"registeredAtNEQ,omitempty"`
	RegisteredAtIn    []time.Time `json:"registeredAtIn,omitempty"`
	RegisteredAtNotIn []time.Time `json:"registeredAtNotIn,omitempty"`
	RegisteredAtGT    *time.Time  `json:"registeredAtGT,omitempty"`
	RegisteredAtGTE   *time.Time  `json:"registeredAtGTE,omitempty"`
	RegisteredAtLT    *time.Time  `json:"registeredAtLT,omitempty"`
	RegisteredAtLTE   *time.Time  `json:"registeredAtLTE,omitempty"`

	// "owner" edge predicates.
	HasOwner     *bool             `json:"hasOwner,omitempty"`
	HasOwnerWith []*UserWhereInput `json:"hasOwnerWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *CarWhereInput) AddPredicates(predicates ...predicate.Car) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the CarWhereInput filter on the CarQuery builder.
func (i *CarWhereInput) Filter(q *CarQuery) (*CarQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyCarWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyCarWhereInput is returned in case the CarWhereInput is empty.
var ErrEmptyCarWhereInput = errors.New("ent: empty predicate CarWhereInput")

// P returns a predicate for filtering cars.
// An error is returned if the input is empty or invalid.
func (i *CarWhereInput) P() (predicate.Car, error) {
	var predicates []predicate.Car
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, car.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Car, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, car.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Car, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, car.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, car.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, car.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, car.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, car.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, car.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, car.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, car.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, car.IDLTE(*i.IDLTE))
	}
	if i.Model != nil {
		predicates = append(predicates, car.ModelEQ(*i.Model))
	}
	if i.ModelNEQ != nil {
		predicates = append(predicates, car.ModelNEQ(*i.ModelNEQ))
	}
	if len(i.ModelIn) > 0 {
		predicates = append(predicates, car.ModelIn(i.ModelIn...))
	}
	if len(i.ModelNotIn) > 0 {
		predicates = append(predicates, car.ModelNotIn(i.ModelNotIn...))
	}
	if i.ModelGT != nil {
		predicates = append(predicates, car.ModelGT(*i.ModelGT))
	}
	if i.ModelGTE != nil {
		predicates = append(predicates, car.ModelGTE(*i.ModelGTE))
	}
	if i.ModelLT != nil {
		predicates = append(predicates, car.ModelLT(*i.ModelLT))
	}
	if i.ModelLTE != nil {
		predicates = append(predicates, car.ModelLTE(*i.ModelLTE))
	}
	if i.ModelContains != nil {
		predicates = append(predicates, car.ModelContains(*i.ModelContains))
	}
	if i.ModelHasPrefix != nil {
		predicates = append(predicates, car.ModelHasPrefix(*i.ModelHasPrefix))
	}
	if i.ModelHasSuffix != nil {
		predicates = append(predicates, car.ModelHasSuffix(*i.ModelHasSuffix))
	}
	if i.ModelEqualFold != nil {
		predicates = append(predicates, car.ModelEqualFold(*i.ModelEqualFold))
	}
	if i.ModelContainsFold != nil {
		predicates = append(predicates, car.ModelContainsFold(*i.ModelContainsFold))
	}
	if i.RegisteredAt != nil {
		predicates = append(predicates, car.RegisteredAtEQ(*i.RegisteredAt))
	}
	if i.RegisteredAtNEQ != nil {
		predicates = append(predicates, car.RegisteredAtNEQ(*i.RegisteredAtNEQ))
	}
	if len(i.RegisteredAtIn) > 0 {
		predicates = append(predicates, car.RegisteredAtIn(i.RegisteredAtIn...))
	}
	if len(i.RegisteredAtNotIn) > 0 {
		predicates = append(predicates, car.RegisteredAtNotIn(i.RegisteredAtNotIn...))
	}
	if i.RegisteredAtGT != nil {
		predicates = append(predicates, car.RegisteredAtGT(*i.RegisteredAtGT))
	}
	if i.RegisteredAtGTE != nil {
		predicates = append(predicates, car.RegisteredAtGTE(*i.RegisteredAtGTE))
	}
	if i.RegisteredAtLT != nil {
		predicates = append(predicates, car.RegisteredAtLT(*i.RegisteredAtLT))
	}
	if i.RegisteredAtLTE != nil {
		predicates = append(predicates, car.RegisteredAtLTE(*i.RegisteredAtLTE))
	}

	if i.HasOwner != nil {
		p := car.HasOwner()
		if !*i.HasOwner {
			p = car.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOwnerWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasOwnerWith))
		for _, w := range i.HasOwnerWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasOwnerWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, car.HasOwnerWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyCarWhereInput
	case 1:
		return predicates[0], nil
	default:
		return car.And(predicates...), nil
	}
}

// GroupWhereInput represents a where input for filtering Group queries.
type GroupWhereInput struct {
	Predicates []predicate.Group  `json:"-"`
	Not        *GroupWhereInput   `json:"not,omitempty"`
	Or         []*GroupWhereInput `json:"or,omitempty"`
	And        []*GroupWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *string  `json:"id,omitempty"`
	IDNEQ   *string  `json:"idNEQ,omitempty"`
	IDIn    []string `json:"idIn,omitempty"`
	IDNotIn []string `json:"idNotIn,omitempty"`
	IDGT    *string  `json:"idGT,omitempty"`
	IDGTE   *string  `json:"idGTE,omitempty"`
	IDLT    *string  `json:"idLT,omitempty"`
	IDLTE   *string  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "users" edge predicates.
	HasUsers     *bool             `json:"hasUsers,omitempty"`
	HasUsersWith []*UserWhereInput `json:"hasUsersWith,omitempty"`

	// "parent" edge predicates.
	HasParent     *bool              `json:"hasParent,omitempty"`
	HasParentWith []*GroupWhereInput `json:"hasParentWith,omitempty"`

	// "children" edge predicates.
	HasChildren     *bool              `json:"hasChildren,omitempty"`
	HasChildrenWith []*GroupWhereInput `json:"hasChildrenWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *GroupWhereInput) AddPredicates(predicates ...predicate.Group) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the GroupWhereInput filter on the GroupQuery builder.
func (i *GroupWhereInput) Filter(q *GroupQuery) (*GroupQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyGroupWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyGroupWhereInput is returned in case the GroupWhereInput is empty.
var ErrEmptyGroupWhereInput = errors.New("ent: empty predicate GroupWhereInput")

// P returns a predicate for filtering groups.
// An error is returned if the input is empty or invalid.
func (i *GroupWhereInput) P() (predicate.Group, error) {
	var predicates []predicate.Group
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, group.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Group, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, group.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Group, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, group.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, group.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, group.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, group.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, group.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, group.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, group.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, group.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, group.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, group.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, group.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, group.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, group.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, group.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, group.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, group.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, group.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, group.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, group.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, group.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, group.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, group.NameContainsFold(*i.NameContainsFold))
	}

	if i.HasUsers != nil {
		p := group.HasUsers()
		if !*i.HasUsers {
			p = group.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasUsersWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasUsersWith))
		for _, w := range i.HasUsersWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasUsersWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, group.HasUsersWith(with...))
	}
	if i.HasParent != nil {
		p := group.HasParent()
		if !*i.HasParent {
			p = group.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasParentWith) > 0 {
		with := make([]predicate.Group, 0, len(i.HasParentWith))
		for _, w := range i.HasParentWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasParentWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, group.HasParentWith(with...))
	}
	if i.HasChildren != nil {
		p := group.HasChildren()
		if !*i.HasChildren {
			p = group.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasChildrenWith) > 0 {
		with := make([]predicate.Group, 0, len(i.HasChildrenWith))
		for _, w := range i.HasChildrenWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasChildrenWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, group.HasChildrenWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyGroupWhereInput
	case 1:
		return predicates[0], nil
	default:
		return group.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Predicates []predicate.User  `json:"-"`
	Not        *UserWhereInput   `json:"not,omitempty"`
	Or         []*UserWhereInput `json:"or,omitempty"`
	And        []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *string  `json:"id,omitempty"`
	IDNEQ   *string  `json:"idNEQ,omitempty"`
	IDIn    []string `json:"idIn,omitempty"`
	IDNotIn []string `json:"idNotIn,omitempty"`
	IDGT    *string  `json:"idGT,omitempty"`
	IDGTE   *string  `json:"idGTE,omitempty"`
	IDLT    *string  `json:"idLT,omitempty"`
	IDLTE   *string  `json:"idLTE,omitempty"`

	// "age" field predicates.
	Age      *int  `json:"age,omitempty"`
	AgeNEQ   *int  `json:"ageNEQ,omitempty"`
	AgeIn    []int `json:"ageIn,omitempty"`
	AgeNotIn []int `json:"ageNotIn,omitempty"`
	AgeGT    *int  `json:"ageGT,omitempty"`
	AgeGTE   *int  `json:"ageGTE,omitempty"`
	AgeLT    *int  `json:"ageLT,omitempty"`
	AgeLTE   *int  `json:"ageLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "cars" edge predicates.
	HasCars     *bool            `json:"hasCars,omitempty"`
	HasCarsWith []*CarWhereInput `json:"hasCarsWith,omitempty"`

	// "groups" edge predicates.
	HasGroups     *bool              `json:"hasGroups,omitempty"`
	HasGroupsWith []*GroupWhereInput `json:"hasGroupsWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserWhereInput) AddPredicates(predicates ...predicate.User) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUserWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUserWhereInput is returned in case the UserWhereInput is empty.
var ErrEmptyUserWhereInput = errors.New("ent: empty predicate UserWhereInput")

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.Age != nil {
		predicates = append(predicates, user.AgeEQ(*i.Age))
	}
	if i.AgeNEQ != nil {
		predicates = append(predicates, user.AgeNEQ(*i.AgeNEQ))
	}
	if len(i.AgeIn) > 0 {
		predicates = append(predicates, user.AgeIn(i.AgeIn...))
	}
	if len(i.AgeNotIn) > 0 {
		predicates = append(predicates, user.AgeNotIn(i.AgeNotIn...))
	}
	if i.AgeGT != nil {
		predicates = append(predicates, user.AgeGT(*i.AgeGT))
	}
	if i.AgeGTE != nil {
		predicates = append(predicates, user.AgeGTE(*i.AgeGTE))
	}
	if i.AgeLT != nil {
		predicates = append(predicates, user.AgeLT(*i.AgeLT))
	}
	if i.AgeLTE != nil {
		predicates = append(predicates, user.AgeLTE(*i.AgeLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, user.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, user.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, user.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, user.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, user.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, user.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, user.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, user.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, user.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, user.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, user.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, user.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, user.NameContainsFold(*i.NameContainsFold))
	}

	if i.HasCars != nil {
		p := user.HasCars()
		if !*i.HasCars {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasCarsWith) > 0 {
		with := make([]predicate.Car, 0, len(i.HasCarsWith))
		for _, w := range i.HasCarsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasCarsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasCarsWith(with...))
	}
	if i.HasGroups != nil {
		p := user.HasGroups()
		if !*i.HasGroups {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasGroupsWith) > 0 {
		with := make([]predicate.Group, 0, len(i.HasGroupsWith))
		for _, w := range i.HasGroupsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasGroupsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasGroupsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserWhereInput
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}
