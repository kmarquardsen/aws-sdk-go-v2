// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns metadata of your saved backup plan templates, including the template ID,
// name, and the creation and deletion dates.
func (c *Client) ListBackupPlanTemplates(ctx context.Context, params *ListBackupPlanTemplatesInput, optFns ...func(*Options)) (*ListBackupPlanTemplatesOutput, error) {
	if params == nil {
		params = &ListBackupPlanTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBackupPlanTemplates", params, optFns, addOperationListBackupPlanTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBackupPlanTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBackupPlanTemplatesInput struct {

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string
}

type ListBackupPlanTemplatesOutput struct {

	// An array of template list items containing metadata about your saved templates.
	BackupPlanTemplatesList []*types.BackupPlanTemplatesListMember

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBackupPlanTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBackupPlanTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackupPlanTemplates{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBackupPlanTemplates(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListBackupPlanTemplates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListBackupPlanTemplates",
	}
}
