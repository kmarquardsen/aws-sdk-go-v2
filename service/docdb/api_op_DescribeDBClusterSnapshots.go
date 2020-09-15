// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about cluster snapshots. This API operation supports
// pagination.
func (c *Client) DescribeDBClusterSnapshots(ctx context.Context, params *DescribeDBClusterSnapshotsInput, optFns ...func(*Options)) (*DescribeDBClusterSnapshotsOutput, error) {
	stack := middleware.NewStack("DescribeDBClusterSnapshots", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeDBClusterSnapshotsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeDBClusterSnapshotsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterSnapshots(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DescribeDBClusterSnapshots",
			Err:           err,
		}
	}
	out := result.(*DescribeDBClusterSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DescribeDBClusterSnapshots ().
type DescribeDBClusterSnapshotsInput struct {
	// The ID of the cluster to retrieve the list of cluster snapshots for. This
	// parameter can't be used with the DBClusterSnapshotIdentifier parameter. This
	// parameter is not case sensitive. Constraints:
	//
	//     * If provided, must match the
	// identifier of an existing DBCluster.
	DBClusterIdentifier *string
	// A specific cluster snapshot identifier to describe. This parameter can't be used
	// with the DBClusterIdentifier parameter. This value is stored as a lowercase
	// string. Constraints:
	//
	//     * If provided, must match the identifier of an
	// existing DBClusterSnapshot.
	//
	//     * If this identifier is for an automated
	// snapshot, the SnapshotType parameter must also be specified.
	DBClusterSnapshotIdentifier *string
	// Set to true to include shared manual cluster snapshots from other AWS accounts
	// that this AWS account has been given permission to copy or restore, and
	// otherwise false. The default is false.
	IncludeShared *bool
	// The type of cluster snapshots to be returned. You can specify one of the
	// following values:
	//
	//     * automated - Return all cluster snapshots that Amazon
	// DocumentDB has automatically created for your AWS account.
	//
	//     * manual -
	// Return all cluster snapshots that you have manually created for your AWS
	// account.
	//
	//     * shared - Return all manual cluster snapshots that have been
	// shared to your AWS account.
	//
	//     * public - Return all cluster snapshots that
	// have been marked as public.
	//
	// If you don't specify a SnapshotType value, then
	// both automated and manual cluster snapshots are returned. You can include shared
	// cluster snapshots with these results by setting the IncludeShared parameter to
	// true. You can include public cluster snapshots with these results by setting the
	// IncludePublic parameter to true. The IncludeShared and IncludePublic parameters
	// don't apply for SnapshotType values of manual or automated. The IncludePublic
	// parameter doesn't apply when SnapshotType is set to shared. The IncludeShared
	// parameter doesn't apply when SnapshotType is set to public.
	SnapshotType *string
	// Set to true to include manual cluster snapshots that are public and can be
	// copied or restored by any AWS account, and otherwise false. The default is
	// false.
	IncludePublic *bool
	// This parameter is not currently supported.
	Filters []*types.Filter
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token (marker) is included in
	// the response so that the remaining results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string
}

// Represents the output of DescribeDBClusterSnapshots ().
type DescribeDBClusterSnapshotsOutput struct {
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string
	// Provides a list of cluster snapshots.
	DBClusterSnapshots []*types.DBClusterSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeDBClusterSnapshotsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterSnapshots{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterSnapshots{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDBClusterSnapshots(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusterSnapshots",
	}
}