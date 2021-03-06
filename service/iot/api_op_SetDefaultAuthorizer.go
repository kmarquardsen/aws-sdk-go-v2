// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the default authorizer. This will be used if a websocket connection is made
// without specifying an authorizer.
func (c *Client) SetDefaultAuthorizer(ctx context.Context, params *SetDefaultAuthorizerInput, optFns ...func(*Options)) (*SetDefaultAuthorizerOutput, error) {
	if params == nil {
		params = &SetDefaultAuthorizerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetDefaultAuthorizer", params, optFns, addOperationSetDefaultAuthorizerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetDefaultAuthorizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetDefaultAuthorizerInput struct {

	// The authorizer name.
	//
	// This member is required.
	AuthorizerName *string
}

type SetDefaultAuthorizerOutput struct {

	// The authorizer ARN.
	AuthorizerArn *string

	// The authorizer name.
	AuthorizerName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetDefaultAuthorizerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSetDefaultAuthorizer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSetDefaultAuthorizer{}, middleware.After)
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
	addOpSetDefaultAuthorizerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetDefaultAuthorizer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetDefaultAuthorizer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "SetDefaultAuthorizer",
	}
}
