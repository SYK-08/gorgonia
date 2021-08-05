package stdops

import (
	"context"
	"runtime/trace"

	"gorgonia.org/gorgonia/values"
)

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

// Lt is a tensor-tensor elementwise less-than.
type Lt struct {
	binop
	retSame bool
}

// String implements fmt.Stringer.
func (op Lt) String() string { return "<" }

// Do performs elementwise less-than.
func (op Lt) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	// Do the actual operation
	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.Lt(a, b, tensor.WithContext(ctx2), tensor.AsSame())
	} else {
		retVal, err = tensor.Lt(a, b, tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise less-than but with a preallocated return value.
// PreallocDo allows Lt to implement ops.PreallocOp.
func (op Lt) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.Lt(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2), tensor.AsSame())
	} else {
		retVal, err = tensor.Lt(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}

// LtVS is a tensor-scalar elementwise less-than.
type LtVS struct {
	binop
	retSame bool
}

// String implements fmt.Stringer.
func (op LtVS) String() string { return "<·" }

// Do performs elementwise less-than.
func (op LtVS) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	// Do the actual operation
	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.Lt(a, b, tensor.WithContext(ctx2), tensor.AsSame())
	} else {
		retVal, err = tensor.Lt(a, b, tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise less-than but with a preallocated return value.
// PreallocDo allows LtVS to implement ops.PreallocOp.
func (op LtVS) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.Lt(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2), tensor.AsSame())
	} else {
		retVal, err = tensor.Lt(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}

// LtSV is a scalar-tensor elementwise less-than.
type LtSV struct {
	binop
	retSame bool
}

// String implements fmt.Stringer.
func (op LtSV) String() string { return "·<" }

// Do performs elementwise less-than.
func (op LtSV) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	// Do the actual operation
	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.Lt(a, b, tensor.WithContext(ctx2), tensor.AsSame())
	} else {
		retVal, err = tensor.Lt(a, b, tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise less-than but with a preallocated return value.
// PreallocDo allows LtSV to implement ops.PreallocOp.
func (op LtSV) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (values.Value, error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.Lt(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2), tensor.AsSame())
	} else {
		retVal, err = tensor.Lt(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}
