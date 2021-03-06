// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of runs for a machine learning transform. Machine learning task runs
// are asynchronous tasks that AWS Glue runs on your behalf as part of various
// machine learning workflows. You can get a sortable, filterable list of machine
// learning task runs by calling GetMLTaskRuns with their parent transform's
// TransformID and other optional parameters as documented in this section. This
// operation returns a list of historic runs and must be paginated.
func (c *Client) GetMLTaskRuns(ctx context.Context, params *GetMLTaskRunsInput, optFns ...func(*Options)) (*GetMLTaskRunsOutput, error) {
	if params == nil {
		params = &GetMLTaskRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMLTaskRuns", params, optFns, addOperationGetMLTaskRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMLTaskRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMLTaskRunsInput struct {

	// The unique identifier of the machine learning transform.
	//
	// This member is required.
	TransformId *string

	// The filter criteria, in the TaskRunFilterCriteria structure, for the task run.
	Filter *types.TaskRunFilterCriteria

	// The maximum number of results to return.
	MaxResults *int32

	// A token for pagination of the results. The default is empty.
	NextToken *string

	// The sorting criteria, in the TaskRunSortCriteria structure, for the task run.
	Sort *types.TaskRunSortCriteria
}

type GetMLTaskRunsOutput struct {

	// A pagination token, if more results are available.
	NextToken *string

	// A list of task runs that are associated with the transform.
	TaskRuns []*types.TaskRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMLTaskRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMLTaskRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMLTaskRuns{}, middleware.After)
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
	addOpGetMLTaskRunsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMLTaskRuns(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetMLTaskRuns(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetMLTaskRuns",
	}
}
