// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancerstatus"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
)

// LoadBalancerStatusDelete is the builder for deleting a LoadBalancerStatus entity.
type LoadBalancerStatusDelete struct {
	config
	hooks    []Hook
	mutation *LoadBalancerStatusMutation
}

// Where appends a list predicates to the LoadBalancerStatusDelete builder.
func (lbsd *LoadBalancerStatusDelete) Where(ps ...predicate.LoadBalancerStatus) *LoadBalancerStatusDelete {
	lbsd.mutation.Where(ps...)
	return lbsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lbsd *LoadBalancerStatusDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lbsd.sqlExec, lbsd.mutation, lbsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lbsd *LoadBalancerStatusDelete) ExecX(ctx context.Context) int {
	n, err := lbsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lbsd *LoadBalancerStatusDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(loadbalancerstatus.Table, sqlgraph.NewFieldSpec(loadbalancerstatus.FieldID, field.TypeString))
	if ps := lbsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lbsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lbsd.mutation.done = true
	return affected, err
}

// LoadBalancerStatusDeleteOne is the builder for deleting a single LoadBalancerStatus entity.
type LoadBalancerStatusDeleteOne struct {
	lbsd *LoadBalancerStatusDelete
}

// Where appends a list predicates to the LoadBalancerStatusDelete builder.
func (lbsdo *LoadBalancerStatusDeleteOne) Where(ps ...predicate.LoadBalancerStatus) *LoadBalancerStatusDeleteOne {
	lbsdo.lbsd.mutation.Where(ps...)
	return lbsdo
}

// Exec executes the deletion query.
func (lbsdo *LoadBalancerStatusDeleteOne) Exec(ctx context.Context) error {
	n, err := lbsdo.lbsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{loadbalancerstatus.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lbsdo *LoadBalancerStatusDeleteOne) ExecX(ctx context.Context) {
	if err := lbsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
