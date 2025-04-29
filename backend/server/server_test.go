// backend/server/server_test.go
package server

import (
	"context"
	"testing"

	calculatorv1 "calculator/backend/gen/calculator/v1"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculatorOperations(t *testing.T) {
	server := &CalculatorServer{}
	ctx := context.Background()

	t.Run("Addition", func(t *testing.T) {
		tests := []struct {
			a, b, expected float64
		}{
			{2, 3, 5},
			{-1, 1, 0},
			{0.1, 0.2, 0.3},
		}

		for _, tt := range tests {
			resp, err := server.Add(ctx, connect.NewRequest(
				&calculatorv1.OperationRequest{A: tt.a, B: tt.b},
			))
			require.NoError(t, err)
			assert.InDelta(t, tt.expected, resp.Msg.Result, 0.0001)
		}
	})

	t.Run("Division", func(t *testing.T) {
		t.Run("ValidDivision", func(t *testing.T) {
			resp, err := server.Divide(ctx, connect.NewRequest(
				&calculatorv1.OperationRequest{A: 10, B: 2},
			))
			require.NoError(t, err)
			assert.Equal(t, 5.0, resp.Msg.Result)
		})

		t.Run("DivisionByZero", func(t *testing.T) {
			_, err := server.Divide(ctx, connect.NewRequest(
				&calculatorv1.OperationRequest{A: 5, B: 0},
			))
			require.Error(t, err)
			assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		})
	})

	// 类似编写 Subtract 和 Multiply 的测试
	t.Run("Subtraction", func(t *testing.T) {
		resp, err := server.Subtract(ctx, connect.NewRequest(
			&calculatorv1.OperationRequest{A: 5, B: 3},
		))
		require.NoError(t, err)
		assert.Equal(t, 2.0, resp.Msg.Result)
	})

	t.Run("Multiplication", func(t *testing.T) {
		resp, err := server.Multiply(ctx, connect.NewRequest(
			&calculatorv1.OperationRequest{A: 4, B: 2.5},
		))
		require.NoError(t, err)
		assert.Equal(t, 10.0, resp.Msg.Result)
	})
}
