// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

func (c *Client) CreateDocumentationVersion(ctx context.Context, params *CreateDocumentationVersionInput, optFns ...func(*Options)) (*CreateDocumentationVersionOutput, error) {
	if params == nil {
		params = &CreateDocumentationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDocumentationVersion", params, optFns, addOperationCreateDocumentationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDocumentationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new documentation version of a given API.
type CreateDocumentationVersionInput struct {

	// [Required] The version identifier of the new snapshot.
	//
	// This member is required.
	DocumentationVersion *string

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// A description about the new documentation snapshot.
	Description *string

	Name *string

	// The stage name to be associated with the new documentation snapshot.
	StageName *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// A snapshot of the documentation of an API. Publishing API documentation involves
// creating a documentation version associated with an API stage and exporting the
// versioned documentation to an external (e.g., OpenAPI) file. Documenting an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// DocumentationPart, DocumentationVersions
type CreateDocumentationVersionOutput struct {

	// The date when the API documentation snapshot is created.
	CreatedDate *time.Time

	// The description of the API documentation snapshot.
	Description *string

	// The version identifier of the API documentation snapshot.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDocumentationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDocumentationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDocumentationVersion{}, middleware.After)
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
	addOpCreateDocumentationVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDocumentationVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDocumentationVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateDocumentationVersion",
	}
}
