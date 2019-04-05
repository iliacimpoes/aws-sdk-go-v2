// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilTableExists uses the DynamoDB API operation
// DescribeTable to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DynamoDB) WaitUntilTableExists(ctx context.Context, input *DescribeTableInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilTableExists",
		MaxAttempts: 25,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Table.TableStatus",
				Expected: "ACTIVE",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeTableInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeTableRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilTableNotExists uses the DynamoDB API operation
// DescribeTable to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DynamoDB) WaitUntilTableNotExists(ctx context.Context, input *DescribeTableInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilTableNotExists",
		MaxAttempts: 25,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeTableInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeTableRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
