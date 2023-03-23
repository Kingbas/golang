package proto

import (
	"context"
	"fmt"
)

type Arith struct {}

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func (t *Arith) Add(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

type DivideByZeroError struct{}

func (e *DivideByZeroError) Error() string {
	return fmt.Sprintln("The divisor cannot be zero!")
}

func (t *Arith) Div(ctx context.Context, args *Args, reply *Reply) error {
	if args.B == 0 {
		return &DivideByZeroError{}
	}
	reply.C = args.A / args.B
	return nil
}
