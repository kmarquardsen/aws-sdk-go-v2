// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a new pool of dedicated IP addresses. A pool can include one or more
// dedicated IP addresses that are associated with your AWS account. You can
// associate a pool with a configuration set. When you send an email that uses that
// configuration set, the message is sent from one of the addresses in the
// associated pool.
func (c *Client) CreateDedicatedIpPool(ctx context.Context, params *CreateDedicatedIpPoolInput, optFns ...func(*Options)) (*CreateDedicatedIpPoolOutput, error) {
	if params == nil {
		params = &CreateDedicatedIpPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDedicatedIpPool", params, optFns, addOperationCreateDedicatedIpPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDedicatedIpPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to create a new dedicated IP pool.
type CreateDedicatedIpPoolInput struct {

	// The name of the dedicated IP pool.
	//
	// This member is required.
	PoolName *string

	// An object that defines the tags (keys and values) that you want to associate
	// with the pool.
	Tags []*types.Tag
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type CreateDedicatedIpPoolOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDedicatedIpPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDedicatedIpPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDedicatedIpPool{}, middleware.After)
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
	addOpCreateDedicatedIpPoolValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDedicatedIpPool(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDedicatedIpPool(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateDedicatedIpPool",
	}
}
