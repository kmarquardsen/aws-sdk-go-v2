// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the themes in the current AWS account.
func (c *Client) ListThemes(ctx context.Context, params *ListThemesInput, optFns ...func(*Options)) (*ListThemesOutput, error) {
	if params == nil {
		params = &ListThemesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThemes", params, optFns, addOperationListThemesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThemesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThemesInput struct {

	// The ID of the AWS account that contains the themes that you're listing.
	//
	// This member is required.
	AwsAccountId *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The type of themes that you want to list. Valid options include the following:
	//
	//
	// * ALL (default)- Display all existing themes.
	//
	//     * CUSTOM - Display only the
	// themes created by people using Amazon QuickSight.
	//
	//     * QUICKSIGHT - Display
	// only the starting themes defined by QuickSight.
	Type types.ThemeType
}

type ListThemesOutput struct {

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// Information about the themes in the list.
	ThemeSummaryList []*types.ThemeSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListThemesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThemes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThemes{}, middleware.After)
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
	addOpListThemesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListThemes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListThemes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListThemes",
	}
}
