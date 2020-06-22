// Code generated by smithy-go-codegen DO NOT EDIT.
package ec2query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The xmlName trait on the output structure is ignored in AWS Query. The wrapping
// element is always operation name + "Response".
func (c *Client) IgnoresWrappingXmlName(ctx context.Context, params *IgnoresWrappingXmlNameInput, optFns ...func(*Options)) (*IgnoresWrappingXmlNameOutput, error) {
	stack := middleware.NewStack("IgnoresWrappingXmlName", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "IgnoresWrappingXmlName",
			Err:           err,
		}
	}
	out := result.(*IgnoresWrappingXmlNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IgnoresWrappingXmlNameInput struct {
}

type IgnoresWrappingXmlNameOutput struct {
	Foo *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
