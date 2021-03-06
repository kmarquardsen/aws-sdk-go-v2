// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all provisioning artifacts (also known as versions) for the specified
// product.
func (c *Client) ListProvisioningArtifacts(ctx context.Context, params *ListProvisioningArtifactsInput, optFns ...func(*Options)) (*ListProvisioningArtifactsOutput, error) {
	if params == nil {
		params = &ListProvisioningArtifactsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProvisioningArtifacts", params, optFns, addOperationListProvisioningArtifactsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProvisioningArtifactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProvisioningArtifactsInput struct {

	// The product identifier.
	//
	// This member is required.
	ProductId *string

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
}

type ListProvisioningArtifactsOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Information about the provisioning artifacts.
	ProvisioningArtifactDetails []*types.ProvisioningArtifactDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListProvisioningArtifactsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProvisioningArtifacts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProvisioningArtifacts{}, middleware.After)
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
	addOpListProvisioningArtifactsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProvisioningArtifacts(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListProvisioningArtifacts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListProvisioningArtifacts",
	}
}
