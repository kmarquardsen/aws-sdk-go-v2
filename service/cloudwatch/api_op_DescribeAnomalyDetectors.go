// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the anomaly detection models that you have created in your account. You
// can list all models in your account or filter the results to only the models
// that are related to a certain namespace, metric name, or metric dimension.
func (c *Client) DescribeAnomalyDetectors(ctx context.Context, params *DescribeAnomalyDetectorsInput, optFns ...func(*Options)) (*DescribeAnomalyDetectorsOutput, error) {
	if params == nil {
		params = &DescribeAnomalyDetectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAnomalyDetectors", params, optFns, addOperationDescribeAnomalyDetectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAnomalyDetectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAnomalyDetectorsInput struct {

	// Limits the results to only the anomaly detection models that are associated with
	// the specified metric dimensions. If there are multiple metrics that have these
	// dimensions and have anomaly detection models associated, they're all returned.
	Dimensions []*types.Dimension

	// The maximum number of results to return in one operation. The maximum value that
	// you can specify is 100. To retrieve the remaining results, make another call
	// with the returned NextToken value.
	MaxResults *int32

	// Limits the results to only the anomaly detection models that are associated with
	// the specified metric name. If there are multiple metrics with this name in
	// different namespaces that have anomaly detection models, they're all returned.
	MetricName *string

	// Limits the results to only the anomaly detection models that are associated with
	// the specified namespace.
	Namespace *string

	// Use the token returned by the previous operation to request the next page of
	// results.
	NextToken *string
}

type DescribeAnomalyDetectorsOutput struct {

	// The list of anomaly detection models returned by the operation.
	AnomalyDetectors []*types.AnomalyDetector

	// A token that you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAnomalyDetectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAnomalyDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAnomalyDetectors{}, middleware.After)
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
	addOpDescribeAnomalyDetectorsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAnomalyDetectors(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeAnomalyDetectors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "DescribeAnomalyDetectors",
	}
}
