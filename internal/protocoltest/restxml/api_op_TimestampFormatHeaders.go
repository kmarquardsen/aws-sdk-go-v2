// Code generated by smithy-go-codegen DO NOT EDIT.
package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The example tests how timestamp request and response headers are serialized.
func (c *Client) TimestampFormatHeaders(ctx context.Context, params *TimestampFormatHeadersInput, optFns ...func(*Options)) (*TimestampFormatHeadersOutput, error) {
	stack := middleware.NewStack("TimestampFormatHeaders", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Initialize.Add(awsmiddleware.RegisterServiceMetadata{
		Region:         options.Region,
		ServiceName:    "Rest Xml Protocol",
		ServiceID:      "restxmlprotocol",
		EndpointPrefix: "restxmlprotocol",
		OperationName:  "TimestampFormatHeaders",
	}, middleware.Before)
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
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
			OperationName: "TimestampFormatHeaders",
			Err:           err,
		}
	}
	out := result.(*TimestampFormatHeadersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TimestampFormatHeadersInput struct {
	DefaultFormat      *time.Time
	MemberDateTime     *time.Time
	MemberEpochSeconds *time.Time
	MemberHttpDate     *time.Time
	TargetDateTime     *time.Time
	TargetEpochSeconds *time.Time
	TargetHttpDate     *time.Time
}

type TimestampFormatHeadersOutput struct {
	DefaultFormat      *time.Time
	MemberDateTime     *time.Time
	MemberEpochSeconds *time.Time
	MemberHttpDate     *time.Time
	TargetDateTime     *time.Time
	TargetEpochSeconds *time.Time
	TargetHttpDate     *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
