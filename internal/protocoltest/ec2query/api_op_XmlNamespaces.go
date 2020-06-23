// Code generated by smithy-go-codegen DO NOT EDIT.
package ec2query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/ec2query/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) XmlNamespaces(ctx context.Context, params *XmlNamespacesInput, optFns ...func(*Options)) (*XmlNamespacesOutput, error) {
	stack := middleware.NewStack("XmlNamespaces", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Initialize.Add(awsmiddleware.RegisterServiceMetadata{
		Region:         options.Region,
		ServiceName:    "EC2 Protocol",
		ServiceID:      "ec2protocol",
		EndpointPrefix: "ec2protocol",
		OperationName:  "XmlNamespaces",
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
			OperationName: "XmlNamespaces",
			Err:           err,
		}
	}
	out := result.(*XmlNamespacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlNamespacesInput struct {
}

type XmlNamespacesOutput struct {
	Nested *types.XmlNamespaceNested

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
