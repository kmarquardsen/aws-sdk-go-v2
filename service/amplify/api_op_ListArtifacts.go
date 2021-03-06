// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of artifacts for a specified app, branch, and job.
func (c *Client) ListArtifacts(ctx context.Context, params *ListArtifactsInput, optFns ...func(*Options)) (*ListArtifactsOutput, error) {
	if params == nil {
		params = &ListArtifactsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListArtifacts", params, optFns, addOperationListArtifactsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListArtifactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Describes the request structure for the list artifacts request.
type ListArtifactsInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name of a branch that is part of an Amplify app.
	//
	// This member is required.
	BranchName *string

	// The unique ID for a job.
	//
	// This member is required.
	JobId *string

	// The maximum number of records to list in a single response.
	MaxResults *int32

	// A pagination token. Set to null to start listing artifacts from start. If a
	// non-null pagination token is returned in a result, pass its value in here to
	// list more artifacts.
	NextToken *string
}

// The result structure for the list artifacts request.
type ListArtifactsOutput struct {

	// A list of artifacts.
	//
	// This member is required.
	Artifacts []*types.Artifact

	// A pagination token. If a non-null pagination token is returned in a result, pass
	// its value in another request to retrieve more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListArtifactsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListArtifacts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListArtifacts{}, middleware.After)
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
	addOpListArtifactsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListArtifacts(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListArtifacts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "ListArtifacts",
	}
}
