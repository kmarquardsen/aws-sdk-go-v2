// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a new comment to the specified document version.
func (c *Client) CreateComment(ctx context.Context, params *CreateCommentInput, optFns ...func(*Options)) (*CreateCommentOutput, error) {
	if params == nil {
		params = &CreateCommentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateComment", params, optFns, addOperationCreateCommentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCommentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCommentInput struct {

	// The ID of the document.
	//
	// This member is required.
	DocumentId *string

	// The text of the comment.
	//
	// This member is required.
	Text *string

	// The ID of the document version.
	//
	// This member is required.
	VersionId *string

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// Set this parameter to TRUE to send an email out to the document collaborators
	// after the comment is created.
	NotifyCollaborators *bool

	// The ID of the parent comment.
	ParentId *string

	// The ID of the root comment in the thread.
	ThreadId *string

	// The visibility of the comment. Options are either PRIVATE, where the comment is
	// visible only to the comment author and document owner and co-owners, or PUBLIC,
	// where the comment is visible to document owners, co-owners, and contributors.
	Visibility types.CommentVisibilityType
}

type CreateCommentOutput struct {

	// The comment that has been created.
	Comment *types.Comment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCommentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateComment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateComment{}, middleware.After)
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
	addOpCreateCommentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateComment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateComment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "CreateComment",
	}
}
