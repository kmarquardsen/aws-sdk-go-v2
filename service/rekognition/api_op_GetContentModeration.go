// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the unsafe content analysis results for a Amazon Rekognition Video analysis
// started by StartContentModeration. Unsafe content analysis of a video is an
// asynchronous operation. You start analysis by calling StartContentModeration
// which returns a job identifier (JobId). When analysis finishes, Amazon
// Rekognition Video publishes a completion status to the Amazon Simple
// Notification Service topic registered in the initial call to
// StartContentModeration. To get the results of the unsafe content analysis, first
// check that the status value published to the Amazon SNS topic is SUCCEEDED. If
// so, call GetContentModeration and pass the job identifier (JobId) from the
// initial call to StartContentModeration. For more information, see Working with
// Stored Videos in the Amazon Rekognition Devlopers Guide. GetContentModeration
// returns detected unsafe content labels, and the time they are detected, in an
// array, ModerationLabels, of ContentModerationDetection objects. By default, the
// moderated labels are returned sorted by time, in milliseconds from the start of
// the video. You can also sort them by moderated label by specifying NAME for the
// SortBy input parameter. Since video analysis can return a large number of
// results, use the MaxResults parameter to limit the number of labels returned in
// a single call to GetContentModeration. If there are more results than specified
// in MaxResults, the value of NextToken in the operation response contains a
// pagination token for getting the next set of results. To get the next page of
// results, call GetContentModeration and populate the NextToken request parameter
// with the value of NextToken returned from the previous call to
// GetContentModeration. For more information, see Detecting Unsafe Content in the
// Amazon Rekognition Developer Guide.
func (c *Client) GetContentModeration(ctx context.Context, params *GetContentModerationInput, optFns ...func(*Options)) (*GetContentModerationOutput, error) {
	if params == nil {
		params = &GetContentModerationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContentModeration", params, optFns, addOperationGetContentModerationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContentModerationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContentModerationInput struct {

	// The identifier for the unsafe content job. Use JobId to identify the job in a
	// subsequent call to GetContentModeration.
	//
	// This member is required.
	JobId *string

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Rekognition returns a pagination token in the response. You
	// can use this pagination token to retrieve the next set of unsafe content labels.
	NextToken *string

	// Sort to use for elements in the ModerationLabelDetections array. Use TIMESTAMP
	// to sort array elements by the time labels are detected. Use NAME to
	// alphabetically group elements for a label together. Within each label group, the
	// array element are sorted by detection confidence. The default sort is by
	// TIMESTAMP.
	SortBy types.ContentModerationSortBy
}

type GetContentModerationOutput struct {

	// The current status of the unsafe content analysis job.
	JobStatus types.VideoJobStatus

	// The detected unsafe content labels and the time(s) they were detected.
	ModerationLabels []*types.ContentModerationDetection

	// Version number of the moderation detection model that was used to detect unsafe
	// content.
	ModerationModelVersion *string

	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of unsafe content
	// labels.
	NextToken *string

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// Information about a video that Amazon Rekognition analyzed. Videometadata is
	// returned in every page of paginated responses from GetContentModeration.
	VideoMetadata *types.VideoMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetContentModerationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetContentModeration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetContentModeration{}, middleware.After)
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
	addOpGetContentModerationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetContentModeration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetContentModeration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetContentModeration",
	}
}
