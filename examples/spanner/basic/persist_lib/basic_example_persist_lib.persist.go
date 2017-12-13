// This file is generated by protoc-gen-persist
// Source File: examples/spanner/basic/basic_example.proto
// DO NOT EDIT
package persist_lib

import (
	"cloud.google.com/go/spanner"
	"golang.org/x/net/context"
)

type MySpannerPersistHelper struct {
	Handlers MySpannerHandlers
}
type MySpannerHandlers interface {
	GetUniaryInsertHandler() func(context.Context, *MySpannerUniaryInsertInput, func(*spanner.Row)) error
	GetUniarySelectHandler() func(context.Context, *MySpannerUniarySelectInput, func(*spanner.Row)) error
	GetUniarySelectWithDirectivesHandler() func(context.Context, *MySpannerUniarySelectWithDirectivesInput, func(*spanner.Row)) error
	GetUniaryUpdateHandler() func(context.Context, *MySpannerUniaryUpdateInput, func(*spanner.Row)) error
	GetUniaryDeleteRangeHandler() func(context.Context, *MySpannerUniaryDeleteRangeInput, func(*spanner.Row)) error
	GetUniaryDeleteSingleHandler() func(context.Context, *MySpannerUniaryDeleteSingleInput, func(*spanner.Row)) error
	GetNoArgsHandler() func(context.Context, *MySpannerNoArgsInput, func(*spanner.Row)) error
	GetServerStreamHandler() func(context.Context, *MySpannerServerStreamInput, func(*spanner.Row)) error
	GetClientStreamInsertHandler() func(context.Context) (func(*MySpannerClientStreamInsertInput), func() (*spanner.Row, error))
	GetClientStreamDeleteHandler() func(context.Context) (func(*MySpannerClientStreamDeleteInput), func() (*spanner.Row, error))
	GetClientStreamUpdateHandler() func(context.Context) (func(*MySpannerClientStreamUpdateInput), func() (*spanner.Row, error))
	GetUniaryInsertWithHooksHandler() func(context.Context, *MySpannerUniaryInsertWithHooksInput, func(*spanner.Row)) error
	GetUniarySelectWithHooksHandler() func(context.Context, *MySpannerUniarySelectWithHooksInput, func(*spanner.Row)) error
	GetUniaryUpdateWithHooksHandler() func(context.Context, *MySpannerUniaryUpdateWithHooksInput, func(*spanner.Row)) error
	GetUniaryDeleteWithHooksHandler() func(context.Context, *MySpannerUniaryDeleteWithHooksInput, func(*spanner.Row)) error
	GetServerStreamWithHooksHandler() func(context.Context, *MySpannerServerStreamWithHooksInput, func(*spanner.Row)) error
	GetClientStreamUpdateWithHooksHandler() func(context.Context) (func(*MySpannerClientStreamUpdateWithHooksInput), func() (*spanner.Row, error))
}

func (p *MySpannerPersistHelper) UniaryInsert(ctx context.Context, params *MySpannerUniaryInsertInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetUniaryInsertHandler()(ctx, params, fn)
}
func (p *MySpannerPersistHelper) UniarySelect(ctx context.Context, params *MySpannerUniarySelectInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetUniarySelectHandler()(ctx, params, fn)
}
func (p *MySpannerPersistHelper) UniarySelectWithDirectives(ctx context.Context, params *MySpannerUniarySelectWithDirectivesInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetUniarySelectWithDirectivesHandler()(ctx, params, fn)
}
func (p *MySpannerPersistHelper) UniaryUpdate(ctx context.Context, params *MySpannerUniaryUpdateInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetUniaryUpdateHandler()(ctx, params, fn)
}
func (p *MySpannerPersistHelper) UniaryDeleteRange(ctx context.Context, params *MySpannerUniaryDeleteRangeInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetUniaryDeleteRangeHandler()(ctx, params, fn)
}
func (p *MySpannerPersistHelper) UniaryDeleteSingle(ctx context.Context, params *MySpannerUniaryDeleteSingleInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetUniaryDeleteSingleHandler()(ctx, params, fn)
}
func (p *MySpannerPersistHelper) NoArgs(ctx context.Context, params *MySpannerNoArgsInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetNoArgsHandler()(ctx, params, fn)
}
func (p *MySpannerPersistHelper) ServerStream(ctx context.Context, params *MySpannerServerStreamInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetServerStreamHandler()(ctx, params, fn)
}

// given a context, returns two functions.  (feed, stop)
// feed will be called once for every row recieved by the handler
// stop will be called when the client is done streaming it expects some sort of results to be returned
// that can be marshalled into a response
func (p *MySpannerPersistHelper) ClientStreamInsert(ctx context.Context) (func(*MySpannerClientStreamInsertInput), func() (*spanner.Row, error)) {
	return p.Handlers.GetClientStreamInsertHandler()(ctx)
}

// given a context, returns two functions.  (feed, stop)
// feed will be called once for every row recieved by the handler
// stop will be called when the client is done streaming it expects some sort of results to be returned
// that can be marshalled into a response
func (p *MySpannerPersistHelper) ClientStreamDelete(ctx context.Context) (func(*MySpannerClientStreamDeleteInput), func() (*spanner.Row, error)) {
	return p.Handlers.GetClientStreamDeleteHandler()(ctx)
}

// given a context, returns two functions.  (feed, stop)
// feed will be called once for every row recieved by the handler
// stop will be called when the client is done streaming it expects some sort of results to be returned
// that can be marshalled into a response
func (p *MySpannerPersistHelper) ClientStreamUpdate(ctx context.Context) (func(*MySpannerClientStreamUpdateInput), func() (*spanner.Row, error)) {
	return p.Handlers.GetClientStreamUpdateHandler()(ctx)
}
func (p *MySpannerPersistHelper) UniaryInsertWithHooks(ctx context.Context, params *MySpannerUniaryInsertWithHooksInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetUniaryInsertWithHooksHandler()(ctx, params, fn)
}
func (p *MySpannerPersistHelper) UniarySelectWithHooks(ctx context.Context, params *MySpannerUniarySelectWithHooksInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetUniarySelectWithHooksHandler()(ctx, params, fn)
}
func (p *MySpannerPersistHelper) UniaryUpdateWithHooks(ctx context.Context, params *MySpannerUniaryUpdateWithHooksInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetUniaryUpdateWithHooksHandler()(ctx, params, fn)
}
func (p *MySpannerPersistHelper) UniaryDeleteWithHooks(ctx context.Context, params *MySpannerUniaryDeleteWithHooksInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetUniaryDeleteWithHooksHandler()(ctx, params, fn)
}
func (p *MySpannerPersistHelper) ServerStreamWithHooks(ctx context.Context, params *MySpannerServerStreamWithHooksInput, fn func(row *spanner.Row)) error {
	return p.Handlers.GetServerStreamWithHooksHandler()(ctx, params, fn)
}

// given a context, returns two functions.  (feed, stop)
// feed will be called once for every row recieved by the handler
// stop will be called when the client is done streaming it expects some sort of results to be returned
// that can be marshalled into a response
func (p *MySpannerPersistHelper) ClientStreamUpdateWithHooks(ctx context.Context) (func(*MySpannerClientStreamUpdateWithHooksInput), func() (*spanner.Row, error)) {
	return p.Handlers.GetClientStreamUpdateWithHooksHandler()(ctx)
}

// input type definitions
type MySpannerUniaryInsertInput struct {
	Id        int64
	Name      string
	StartTime interface{}
}
type MySpannerUniarySelectInput struct {
	Id   int64
	Name string
}
type MySpannerUniarySelectWithDirectivesInput struct {
	Id   int64
	Name string
}
type MySpannerUniaryUpdateInput struct {
	Id        int64
	Name      string
	StartTime interface{}
}
type MySpannerUniaryDeleteRangeInput struct {
	EndId   int64
	StartId int64
}
type MySpannerUniaryDeleteSingleInput struct {
	Id int64
}
type MySpannerNoArgsInput struct {
}
type MySpannerServerStreamInput struct {
}
type MySpannerClientStreamInsertInput struct {
	Id        int64
	Name      string
	StartTime interface{}
}
type MySpannerClientStreamDeleteInput struct {
	Id int64
}
type MySpannerClientStreamUpdateInput struct {
	Id        int64
	Name      string
	StartTime interface{}
}
type MySpannerUniaryInsertWithHooksInput struct {
	Id        int64
	Name      string
	StartTime interface{}
}
type MySpannerUniarySelectWithHooksInput struct {
	Id int64
}
type MySpannerUniaryUpdateWithHooksInput struct {
	Id        int64
	Name      string
	StartTime interface{}
}
type MySpannerUniaryDeleteWithHooksInput struct {
	EndId   int64
	StartId int64
}
type MySpannerServerStreamWithHooksInput struct {
}
type MySpannerClientStreamUpdateWithHooksInput struct {
	Id   int64
	Name string
}

func ExampleTableForUniaryInsert(req *MySpannerUniaryInsertInput) *spanner.Mutation {
	return spanner.InsertMap("example_table", map[string]interface{}{
		"id":         req.Id,
		"start_time": req.StartTime,
		"name":       "bananas",
	})
}
func ExampleTableForUniarySelect(req *MySpannerUniarySelectInput) spanner.Statement {
	return spanner.Statement{
		SQL: "SELECT * from example_table Where id=@id AND name=@name",
		Params: map[string]interface{}{
			"@id":   req.Id,
			"@name": req.Name,
		},
	}
}
func ExampleTableForUniarySelectWithDirectives(req *MySpannerUniarySelectWithDirectivesInput) spanner.Statement {
	return spanner.Statement{
		SQL: "SELECT * from example_table@{FORCE_INDEX=index} Where id=@id AND name=@name",
		Params: map[string]interface{}{
			"@id":   req.Id,
			"@name": req.Name,
		},
	}
}
func ExampleTableForUniaryUpdate(req *MySpannerUniaryUpdateInput) *spanner.Mutation {
	return spanner.UpdateMap("example_table", map[string]interface{}{
		"start_time": req.StartTime,
		"name":       "oranges",
		"id":         req.Id,
	})
}
func ExampleTableRangeForUniaryDeleteRange(req *MySpannerUniaryDeleteRangeInput) *spanner.Mutation {
	return spanner.Delete("example_table", spanner.KeyRange{
		Start: spanner.Key{
			req.StartId,
		},
		End: spanner.Key{
			req.EndId,
		},
		Kind: spanner.ClosedOpen,
	})
}
func ExampleTableForUniaryDeleteSingle(req *MySpannerUniaryDeleteSingleInput) *spanner.Mutation {
	return spanner.Delete("example_table", spanner.Key{
		"abc",
		123,
		req.Id,
	})
}
func ExampleTableForNoArgs(req *MySpannerNoArgsInput) spanner.Statement {
	return spanner.Statement{
		SQL:    "select * from example_table limit 1",
		Params: map[string]interface{}{},
	}
}
func NameForServerStream(req *MySpannerServerStreamInput) spanner.Statement {
	return spanner.Statement{
		SQL:    "SELECT * FROM example_table",
		Params: map[string]interface{}{},
	}
}
func ExampleTableForClientStreamInsert(req *MySpannerClientStreamInsertInput) *spanner.Mutation {
	return spanner.InsertMap("example_table", map[string]interface{}{
		"id":         req.Id,
		"start_time": req.StartTime,
		"name":       3,
	})
}
func ExampleTableForClientStreamDelete(req *MySpannerClientStreamDeleteInput) *spanner.Mutation {
	return spanner.Delete("example_table", spanner.Key{
		req.Id,
	})
}
func ExampleTableForClientStreamUpdate(req *MySpannerClientStreamUpdateInput) *spanner.Mutation {
	return spanner.UpdateMap("example_table", map[string]interface{}{
		"start_time": req.StartTime,
		"name":       req.Name,
		"id":         req.Id,
	})
}
func ExampleTableForUniaryInsertWithHooks(req *MySpannerUniaryInsertWithHooksInput) *spanner.Mutation {
	return spanner.InsertMap("example_table", map[string]interface{}{
		"id":         req.Id,
		"start_time": req.StartTime,
		"name":       "bananas",
	})
}
func ExampleTableForUniarySelectWithHooks(req *MySpannerUniarySelectWithHooksInput) spanner.Statement {
	return spanner.Statement{
		SQL: "SELECT * from example_table Where id=@id",
		Params: map[string]interface{}{
			"@id": req.Id,
		},
	}
}
func ExampleTableForUniaryUpdateWithHooks(req *MySpannerUniaryUpdateWithHooksInput) *spanner.Mutation {
	return spanner.UpdateMap("example_table", map[string]interface{}{
		"start_time": req.StartTime,
		"name":       "oranges",
		"id":         req.Id,
	})
}
func ExampleTableRangeForUniaryDeleteWithHooks(req *MySpannerUniaryDeleteWithHooksInput) *spanner.Mutation {
	return spanner.Delete("example_table", spanner.KeyRange{
		Start: spanner.Key{
			req.StartId,
		},
		End: spanner.Key{
			req.EndId,
		},
		Kind: spanner.ClosedOpen,
	})
}
func NameForServerStreamWithHooks(req *MySpannerServerStreamWithHooksInput) spanner.Statement {
	return spanner.Statement{
		SQL:    "SELECT * FROM example_table",
		Params: map[string]interface{}{},
	}
}
func ExampleTableForClientStreamUpdateWithHooks(req *MySpannerClientStreamUpdateWithHooksInput) *spanner.Mutation {
	return spanner.UpdateMap("example_table", map[string]interface{}{
		"name": "asdf",
		"id":   req.Id,
	})
}

// Default method implementations
func DefaultUniaryInsertHandler(cli *spanner.Client) func(context.Context, *MySpannerUniaryInsertInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerUniaryInsertInput, next func(*spanner.Row)) error {
		if _, err := cli.Apply(ctx, []*spanner.Mutation{ExampleTableForUniaryInsert(req)}); err != nil {
			return err
		}
		next(nil) // this is an apply, it has no result
		return nil
	}
}

func DefaultUniarySelectHandler(cli *spanner.Client) func(context.Context, *MySpannerUniarySelectInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerUniarySelectInput, next func(*spanner.Row)) error {
		iter := cli.Single().Query(ctx, ExampleTableForUniarySelect(req))
		if err := iter.Do(func(r *spanner.Row) error {
			next(r)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}
}

func DefaultUniarySelectWithDirectivesHandler(cli *spanner.Client) func(context.Context, *MySpannerUniarySelectWithDirectivesInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerUniarySelectWithDirectivesInput, next func(*spanner.Row)) error {
		iter := cli.Single().Query(ctx, ExampleTableForUniarySelectWithDirectives(req))
		if err := iter.Do(func(r *spanner.Row) error {
			next(r)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}
}

func DefaultUniaryUpdateHandler(cli *spanner.Client) func(context.Context, *MySpannerUniaryUpdateInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerUniaryUpdateInput, next func(*spanner.Row)) error {
		if _, err := cli.Apply(ctx, []*spanner.Mutation{ExampleTableForUniaryUpdate(req)}); err != nil {
			return err
		}
		next(nil) // this is an apply, it has no result
		return nil
	}
}

func DefaultUniaryDeleteRangeHandler(cli *spanner.Client) func(context.Context, *MySpannerUniaryDeleteRangeInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerUniaryDeleteRangeInput, next func(*spanner.Row)) error {
		if _, err := cli.Apply(ctx, []*spanner.Mutation{ExampleTableRangeForUniaryDeleteRange(req)}); err != nil {
			return err
		}
		next(nil) // this is an apply, it has no result
		return nil
	}
}

func DefaultUniaryDeleteSingleHandler(cli *spanner.Client) func(context.Context, *MySpannerUniaryDeleteSingleInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerUniaryDeleteSingleInput, next func(*spanner.Row)) error {
		if _, err := cli.Apply(ctx, []*spanner.Mutation{ExampleTableForUniaryDeleteSingle(req)}); err != nil {
			return err
		}
		next(nil) // this is an apply, it has no result
		return nil
	}
}

func DefaultNoArgsHandler(cli *spanner.Client) func(context.Context, *MySpannerNoArgsInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerNoArgsInput, next func(*spanner.Row)) error {
		iter := cli.Single().Query(ctx, ExampleTableForNoArgs(req))
		if err := iter.Do(func(r *spanner.Row) error {
			next(r)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}
}

func DefaultServerStreamHandler(cli *spanner.Client) func(context.Context, *MySpannerServerStreamInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerServerStreamInput, next func(*spanner.Row)) error {
		iter := cli.Single().Query(ctx, NameForServerStream(req))
		if err := iter.Do(func(r *spanner.Row) error {
			next(r)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}
}

func DefaultClientStreamInsertHandler(cli *spanner.Client) func(context.Context) (func(*MySpannerClientStreamInsertInput), func() (*spanner.Row, error)) {
	return func(ctx context.Context) (feed func(*MySpannerClientStreamInsertInput), done func() (*spanner.Row, error)) {
		var muts []*spanner.Mutation
		feed = func(req *MySpannerClientStreamInsertInput) {
			muts = append(muts, ExampleTableForClientStreamInsert(req))
		}
		done = func() (*spanner.Row, error) {
			if _, err := cli.Apply(ctx, muts); err != nil {
				return nil, err
			}
			return nil, nil // we dont have a row, because we are an apply
		}
		return feed, done
	}
}

func DefaultClientStreamDeleteHandler(cli *spanner.Client) func(context.Context) (func(*MySpannerClientStreamDeleteInput), func() (*spanner.Row, error)) {
	return func(ctx context.Context) (feed func(*MySpannerClientStreamDeleteInput), done func() (*spanner.Row, error)) {
		var muts []*spanner.Mutation
		feed = func(req *MySpannerClientStreamDeleteInput) {
			muts = append(muts, ExampleTableForClientStreamDelete(req))
		}
		done = func() (*spanner.Row, error) {
			if _, err := cli.Apply(ctx, muts); err != nil {
				return nil, err
			}
			return nil, nil // we dont have a row, because we are an apply
		}
		return feed, done
	}
}

func DefaultClientStreamUpdateHandler(cli *spanner.Client) func(context.Context) (func(*MySpannerClientStreamUpdateInput), func() (*spanner.Row, error)) {
	return func(ctx context.Context) (feed func(*MySpannerClientStreamUpdateInput), done func() (*spanner.Row, error)) {
		var muts []*spanner.Mutation
		feed = func(req *MySpannerClientStreamUpdateInput) {
			muts = append(muts, ExampleTableForClientStreamUpdate(req))
		}
		done = func() (*spanner.Row, error) {
			if _, err := cli.Apply(ctx, muts); err != nil {
				return nil, err
			}
			return nil, nil // we dont have a row, because we are an apply
		}
		return feed, done
	}
}

func DefaultUniaryInsertWithHooksHandler(cli *spanner.Client) func(context.Context, *MySpannerUniaryInsertWithHooksInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerUniaryInsertWithHooksInput, next func(*spanner.Row)) error {
		if _, err := cli.Apply(ctx, []*spanner.Mutation{ExampleTableForUniaryInsertWithHooks(req)}); err != nil {
			return err
		}
		next(nil) // this is an apply, it has no result
		return nil
	}
}

func DefaultUniarySelectWithHooksHandler(cli *spanner.Client) func(context.Context, *MySpannerUniarySelectWithHooksInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerUniarySelectWithHooksInput, next func(*spanner.Row)) error {
		iter := cli.Single().Query(ctx, ExampleTableForUniarySelectWithHooks(req))
		if err := iter.Do(func(r *spanner.Row) error {
			next(r)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}
}

func DefaultUniaryUpdateWithHooksHandler(cli *spanner.Client) func(context.Context, *MySpannerUniaryUpdateWithHooksInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerUniaryUpdateWithHooksInput, next func(*spanner.Row)) error {
		if _, err := cli.Apply(ctx, []*spanner.Mutation{ExampleTableForUniaryUpdateWithHooks(req)}); err != nil {
			return err
		}
		next(nil) // this is an apply, it has no result
		return nil
	}
}

func DefaultUniaryDeleteWithHooksHandler(cli *spanner.Client) func(context.Context, *MySpannerUniaryDeleteWithHooksInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerUniaryDeleteWithHooksInput, next func(*spanner.Row)) error {
		if _, err := cli.Apply(ctx, []*spanner.Mutation{ExampleTableRangeForUniaryDeleteWithHooks(req)}); err != nil {
			return err
		}
		next(nil) // this is an apply, it has no result
		return nil
	}
}

func DefaultServerStreamWithHooksHandler(cli *spanner.Client) func(context.Context, *MySpannerServerStreamWithHooksInput, func(*spanner.Row)) error {
	return func(ctx context.Context, req *MySpannerServerStreamWithHooksInput, next func(*spanner.Row)) error {
		iter := cli.Single().Query(ctx, NameForServerStreamWithHooks(req))
		if err := iter.Do(func(r *spanner.Row) error {
			next(r)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}
}

func DefaultClientStreamUpdateWithHooksHandler(cli *spanner.Client) func(context.Context) (func(*MySpannerClientStreamUpdateWithHooksInput), func() (*spanner.Row, error)) {
	return func(ctx context.Context) (feed func(*MySpannerClientStreamUpdateWithHooksInput), done func() (*spanner.Row, error)) {
		var muts []*spanner.Mutation
		feed = func(req *MySpannerClientStreamUpdateWithHooksInput) {
			muts = append(muts, ExampleTableForClientStreamUpdateWithHooks(req))
		}
		done = func() (*spanner.Row, error) {
			if _, err := cli.Apply(ctx, muts); err != nil {
				return nil, err
			}
			return nil, nil // we dont have a row, because we are an apply
		}
		return feed, done
	}
}
