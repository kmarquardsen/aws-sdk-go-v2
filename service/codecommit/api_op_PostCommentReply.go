// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Posts a comment in reply to an existing comment on a comparison between commits
// or a pull request.
func (c *Client) PostCommentReply(ctx context.Context, params *PostCommentReplyInput, optFns ...func(*Options)) (*PostCommentReplyOutput, error) {
	if params == nil {
		params = &PostCommentReplyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PostCommentReply", params, optFns, addOperationPostCommentReplyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PostCommentReplyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PostCommentReplyInput struct {

	// The contents of your reply to a comment.
	//
	// This member is required.
	Content *string

	// The system-generated ID of the comment to which you want to reply. To get this
	// ID, use GetCommentsForComparedCommit or GetCommentsForPullRequest.
	//
	// This member is required.
	InReplyTo *string

	// A unique, client-generated idempotency token that, when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request is
	// received with the same parameters and a token is included, the request returns
	// information about the initial request that used that token.
	ClientRequestToken *string
}

type PostCommentReplyOutput struct {

	// Information about the reply to a comment.
	Comment *types.Comment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPostCommentReplyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPostCommentReply{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPostCommentReply{}, middleware.After)
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
	addIdempotencyToken_opPostCommentReplyMiddleware(stack, options)
	addOpPostCommentReplyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPostCommentReply(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpPostCommentReply struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPostCommentReply) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPostCommentReply) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PostCommentReplyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PostCommentReplyInput ")
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
func addIdempotencyToken_opPostCommentReplyMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpPostCommentReply{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPostCommentReply(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "PostCommentReply",
	}
}
