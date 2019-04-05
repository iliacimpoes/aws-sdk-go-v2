// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilChangeSetCreateComplete uses the AWS CloudFormation API operation
// DescribeChangeSet to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFormation) WaitUntilChangeSetCreateComplete(ctx context.Context, input *DescribeChangeSetInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilChangeSetCreateComplete",
		MaxAttempts: 120,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Status",
				Expected: "CREATE_COMPLETE",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Status",
				Expected: "FAILED",
			},
			{
				State:    aws.FailureWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ValidationError",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeChangeSetInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeChangeSetRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilStackCreateComplete uses the AWS CloudFormation API operation
// DescribeStacks to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFormation) WaitUntilStackCreateComplete(ctx context.Context, input *DescribeStacksInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilStackCreateComplete",
		MaxAttempts: 120,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "CREATE_COMPLETE",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "CREATE_FAILED",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "DELETE_COMPLETE",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "DELETE_FAILED",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "ROLLBACK_FAILED",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "ROLLBACK_COMPLETE",
			},
			{
				State:    aws.FailureWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ValidationError",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeStacksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeStacksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilStackDeleteComplete uses the AWS CloudFormation API operation
// DescribeStacks to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFormation) WaitUntilStackDeleteComplete(ctx context.Context, input *DescribeStacksInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilStackDeleteComplete",
		MaxAttempts: 120,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "DELETE_COMPLETE",
			},
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ValidationError",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "DELETE_FAILED",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "CREATE_FAILED",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "ROLLBACK_FAILED",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "UPDATE_ROLLBACK_FAILED",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "UPDATE_ROLLBACK_IN_PROGRESS",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeStacksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeStacksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilStackExists uses the AWS CloudFormation API operation
// DescribeStacks to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFormation) WaitUntilStackExists(ctx context.Context, input *DescribeStacksInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilStackExists",
		MaxAttempts: 20,
		Delay:       aws.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ValidationError",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeStacksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeStacksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilStackUpdateComplete uses the AWS CloudFormation API operation
// DescribeStacks to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFormation) WaitUntilStackUpdateComplete(ctx context.Context, input *DescribeStacksInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilStackUpdateComplete",
		MaxAttempts: 120,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "UPDATE_COMPLETE",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "UPDATE_FAILED",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "UPDATE_ROLLBACK_FAILED",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Stacks[].StackStatus",
				Expected: "UPDATE_ROLLBACK_COMPLETE",
			},
			{
				State:    aws.FailureWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ValidationError",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeStacksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeStacksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
