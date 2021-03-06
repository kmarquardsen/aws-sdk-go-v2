// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the versions of a core definition.
func (c *Client) ListCoreDefinitionVersions(ctx context.Context, params *ListCoreDefinitionVersionsInput, optFns ...func(*Options)) (*ListCoreDefinitionVersionsOutput, error) {
	if params == nil {
		params = &ListCoreDefinitionVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCoreDefinitionVersions", params, optFns, addOperationListCoreDefinitionVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCoreDefinitionVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCoreDefinitionVersionsInput struct {

	// The ID of the core definition.
	//
	// This member is required.
	CoreDefinitionId *string

	// The maximum number of results to be returned per request.
	MaxResults *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
}

type ListCoreDefinitionVersionsOutput struct {

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	// Information about a version.
	Versions []*types.VersionInformation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCoreDefinitionVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCoreDefinitionVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCoreDefinitionVersions{}, middleware.After)
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
	addOpListCoreDefinitionVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCoreDefinitionVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListCoreDefinitionVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ListCoreDefinitionVersions",
	}
}
