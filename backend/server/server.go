package server

import (
	"context"
	"errors"

	calculatorv1 "calculator/backend/gen/calculator/v1"

	"calculator/backend/gen/calculator/v1/calculatorv1connect"

	connect_go "connectrpc.com/connect"
)

// 接口验证（必须放在包级别）
var _ calculatorv1connect.CalculatorServiceHandler = (*CalculatorServer)(nil)

type CalculatorServer struct{}

// Add 加法实现
func (s *CalculatorServer) Add(
	ctx context.Context,
	req *connect_go.Request[calculatorv1.OperationRequest],
) (*connect_go.Response[calculatorv1.OperationResponse], error) {
	a := req.Msg.A
	b := req.Msg.B
	return connect_go.NewResponse(&calculatorv1.OperationResponse{
		Result: a + b,
	}), nil
}

// Subtract 减法实现
func (s *CalculatorServer) Subtract(
	ctx context.Context,
	req *connect_go.Request[calculatorv1.OperationRequest],
) (*connect_go.Response[calculatorv1.OperationResponse], error) {
	return connect_go.NewResponse(&calculatorv1.OperationResponse{
		Result: req.Msg.A - req.Msg.B,
	}), nil
}

// Multiply 乘法实现
func (s *CalculatorServer) Multiply(
	ctx context.Context,
	req *connect_go.Request[calculatorv1.OperationRequest],
) (*connect_go.Response[calculatorv1.OperationResponse], error) {
	return connect_go.NewResponse(&calculatorv1.OperationResponse{
		Result: req.Msg.A * req.Msg.B,
	}), nil
}

// Divide 除法实现（带有错误处理）
func (s *CalculatorServer) Divide(
	ctx context.Context,
	req *connect_go.Request[calculatorv1.OperationRequest],
) (*connect_go.Response[calculatorv1.OperationResponse], error) {
	if req.Msg.B == 0 {
		return nil, connect_go.NewError(
			connect_go.CodeInvalidArgument, // 修正包引用
			errors.New("division by zero"),
		)
	}
	return connect_go.NewResponse(&calculatorv1.OperationResponse{
		Result: req.Msg.A / req.Msg.B,
	}), nil
}
