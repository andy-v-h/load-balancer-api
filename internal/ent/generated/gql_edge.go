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

	"github.com/99designs/gqlgen/graphql"
)

func (lb *LoadBalancer) Annotations(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *LoadBalancerAnnotationOrder, where *LoadBalancerAnnotationWhereInput,
) (*LoadBalancerAnnotationConnection, error) {
	opts := []LoadBalancerAnnotationPaginateOption{
		WithLoadBalancerAnnotationOrder(orderBy),
		WithLoadBalancerAnnotationFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := lb.Edges.totalCount[0][alias]
	if nodes, err := lb.NamedAnnotations(alias); err == nil || hasTotalCount {
		pager, err := newLoadBalancerAnnotationPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &LoadBalancerAnnotationConnection{Edges: []*LoadBalancerAnnotationEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return lb.QueryAnnotations().Paginate(ctx, after, first, before, last, opts...)
}

func (lb *LoadBalancer) Statuses(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *LoadBalancerStatusOrder, where *LoadBalancerStatusWhereInput,
) (*LoadBalancerStatusConnection, error) {
	opts := []LoadBalancerStatusPaginateOption{
		WithLoadBalancerStatusOrder(orderBy),
		WithLoadBalancerStatusFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := lb.Edges.totalCount[1][alias]
	if nodes, err := lb.NamedStatuses(alias); err == nil || hasTotalCount {
		pager, err := newLoadBalancerStatusPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &LoadBalancerStatusConnection{Edges: []*LoadBalancerStatusEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return lb.QueryStatuses().Paginate(ctx, after, first, before, last, opts...)
}

func (lb *LoadBalancer) Ports(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *LoadBalancerPortOrder, where *LoadBalancerPortWhereInput,
) (*LoadBalancerPortConnection, error) {
	opts := []LoadBalancerPortPaginateOption{
		WithLoadBalancerPortOrder(orderBy),
		WithLoadBalancerPortFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := lb.Edges.totalCount[2][alias]
	if nodes, err := lb.NamedPorts(alias); err == nil || hasTotalCount {
		pager, err := newLoadBalancerPortPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &LoadBalancerPortConnection{Edges: []*LoadBalancerPortEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return lb.QueryPorts().Paginate(ctx, after, first, before, last, opts...)
}

func (lb *LoadBalancer) Provider(ctx context.Context) (*Provider, error) {
	result, err := lb.Edges.ProviderOrErr()
	if IsNotLoaded(err) {
		result, err = lb.QueryProvider().Only(ctx)
	}
	return result, err
}

func (lba *LoadBalancerAnnotation) LoadBalancer(ctx context.Context) (*LoadBalancer, error) {
	result, err := lba.Edges.LoadBalancerOrErr()
	if IsNotLoaded(err) {
		result, err = lba.QueryLoadBalancer().Only(ctx)
	}
	return result, err
}

func (lbs *LoadBalancerStatus) LoadBalancer(ctx context.Context) (*LoadBalancer, error) {
	result, err := lbs.Edges.LoadBalancerOrErr()
	if IsNotLoaded(err) {
		result, err = lbs.QueryLoadBalancer().Only(ctx)
	}
	return result, err
}

func (o *Origin) Pool(ctx context.Context) (*Pool, error) {
	result, err := o.Edges.PoolOrErr()
	if IsNotLoaded(err) {
		result, err = o.QueryPool().Only(ctx)
	}
	return result, err
}

func (po *Pool) Ports(ctx context.Context) (result []*Port, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = po.NamedPorts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = po.Edges.PortsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = po.QueryPorts().All(ctx)
	}
	return result, err
}

func (po *Pool) Origins(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *LoadBalancerOriginOrder, where *LoadBalancerOriginWhereInput,
) (*LoadBalancerOriginConnection, error) {
	opts := []LoadBalancerOriginPaginateOption{
		WithLoadBalancerOriginOrder(orderBy),
		WithLoadBalancerOriginFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := po.Edges.totalCount[1][alias]
	if nodes, err := po.NamedOrigins(alias); err == nil || hasTotalCount {
		pager, err := newLoadBalancerOriginPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &LoadBalancerOriginConnection{Edges: []*LoadBalancerOriginEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return po.QueryOrigins().Paginate(ctx, after, first, before, last, opts...)
}

func (po *Port) Pools(ctx context.Context) (result []*Pool, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = po.NamedPools(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = po.Edges.PoolsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = po.QueryPools().All(ctx)
	}
	return result, err
}

func (po *Port) LoadBalancer(ctx context.Context) (*LoadBalancer, error) {
	result, err := po.Edges.LoadBalancerOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryLoadBalancer().Only(ctx)
	}
	return result, err
}

func (pr *Provider) LoadBalancers(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *LoadBalancerOrder, where *LoadBalancerWhereInput,
) (*LoadBalancerConnection, error) {
	opts := []LoadBalancerPaginateOption{
		WithLoadBalancerOrder(orderBy),
		WithLoadBalancerFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := pr.Edges.totalCount[0][alias]
	if nodes, err := pr.NamedLoadBalancers(alias); err == nil || hasTotalCount {
		pager, err := newLoadBalancerPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &LoadBalancerConnection{Edges: []*LoadBalancerEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return pr.QueryLoadBalancers().Paginate(ctx, after, first, before, last, opts...)
}
