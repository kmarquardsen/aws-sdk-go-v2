// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a peer node in a member.
func (c *Client) CreateNode(ctx context.Context, params *CreateNodeInput, optFns ...func(*Options)) (*CreateNodeOutput, error) {
	if params == nil {
		params = &CreateNodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNode", params, optFns, addOperationCreateNodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNodeInput struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the operation. An idempotent operation completes no more than one time. This
	// identifier is required only if you make a service request directly using an HTTP
	// client. It is generated automatically if you use an AWS SDK or the AWS CLI.
	//
	// This member is required.
	ClientRequestToken *string

	// The unique identifier of the member that owns this node.
	//
	// This member is required.
	MemberId *string

	// The unique identifier of the network in which this node runs.
	//
	// This member is required.
	NetworkId *string

	// The properties of a node configuration.
	//
	// This member is required.
	NodeConfiguration *types.NodeConfiguration
}

type CreateNodeOutput struct {

	// The unique identifier of the node.
	NodeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateNodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateNode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateNode{}, middleware.After)
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
	addIdempotencyToken_opCreateNodeMiddleware(stack, options)
	addOpCreateNodeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNode(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateNode struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNode) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNode) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNodeInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNodeInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateNodeMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateNode{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNode(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "CreateNode",
	}
}
