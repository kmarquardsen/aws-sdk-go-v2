// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a specified resolver endpoint, such as whether it's an
// inbound or an outbound resolver endpoint, and the current status of the
// endpoint.
func (c *Client) GetResolverEndpoint(ctx context.Context, params *GetResolverEndpointInput, optFns ...func(*Options)) (*GetResolverEndpointOutput, error) {
	stack := middleware.NewStack("GetResolverEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetResolverEndpointMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetResolverEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetResolverEndpoint(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetResolverEndpoint",
			Err:           err,
		}
	}
	out := result.(*GetResolverEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResolverEndpointInput struct {
	// The ID of the resolver endpoint that you want to get information about.
	ResolverEndpointId *string
}

type GetResolverEndpointOutput struct {
	// Information about the resolver endpoint that you specified in a
	// GetResolverEndpoint request.
	ResolverEndpoint *types.ResolverEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetResolverEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetResolverEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResolverEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetResolverEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "GetResolverEndpoint",
	}
}