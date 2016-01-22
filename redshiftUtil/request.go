package redshiftUtil

import (
	"fmt"

	"github.com/anxiousmodernman/stats/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/redshift"
)

func GetRedShiftCreateParams(cfg config.RedShiftConfig) *redshift.CreateClusterInput {

	params := &redshift.CreateClusterInput{
		ClusterIdentifier:                aws.String(cfg.ClusterIdentifier),
		MasterUserPassword:               aws.String(cfg.MasterUserPassword),
		MasterUsername:                   aws.String(cfg.MasterUsername),
		NodeType:                         aws.String(cfg.NodeType),
		AllowVersionUpgrade:              aws.Boolean(true),
		AutomatedSnapshotRetentionPeriod: aws.Long(1),
		AvailabilityZone:                 aws.String(cfg.AvailabilityZone),
		ClusterParameterGroupName:        aws.String(cfg.ClusterParameterGroupName),
		ClusterSubnetGroupName:           aws.String(cfg.ClusterSubnetGroupName),
		DBName:                           aws.String(cfg.DbName),
		Encrypted:                        aws.Boolean(cfg.Encrypted),
		NumberOfNodes:                    aws.Long(int64(cfg.NumberOfNodes)),
		Port:                             aws.Long(int64(cfg.Port)),
		PubliclyAccessible:               aws.Boolean(cfg.PubliclyAccessible),
		//		Tags: []*redshift.Tag{
		//			{ // Required
		//				Key:   aws.String("String"),
		//				Value: aws.String("String"),
		//			},
		//			// More values...
		//		},
		VPCSecurityGroupIDs: []*string{
			aws.String("sg-6f82510b"),
		},
	}
	return params
}

func CreateCluster(cfg config.RedShiftConfig) error {
	_ = "breakpoint"
	svc := redshift.New(nil)
	params := GetRedShiftCreateParams(cfg)
	_, err := svc.CreateCluster(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}

		return err
	}
	return nil
}
