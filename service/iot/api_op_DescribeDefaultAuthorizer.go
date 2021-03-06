// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the default authorizer.
func (c *Client) DescribeDefaultAuthorizer(ctx context.Context, params *DescribeDefaultAuthorizerInput, optFns ...func(*Options)) (*DescribeDefaultAuthorizerOutput, error) {
	if params == nil {
		params = &DescribeDefaultAuthorizerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDefaultAuthorizer", params, optFns, addOperationDescribeDefaultAuthorizerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDefaultAuthorizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDefaultAuthorizerInput struct {
}

type DescribeDefaultAuthorizerOutput struct {

	// The default authorizer's description.
	AuthorizerDescription *types.AuthorizerDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDefaultAuthorizerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDefaultAuthorizer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDefaultAuthorizer{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDefaultAuthorizer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDefaultAuthorizer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeDefaultAuthorizer",
	}
}
