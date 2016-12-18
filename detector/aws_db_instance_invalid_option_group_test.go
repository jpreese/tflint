package detector

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/golang/mock/gomock"
	"github.com/wata727/tflint/awsmock"
	"github.com/wata727/tflint/config"
	"github.com/wata727/tflint/issue"
)

func TestDetectAwsDBInstanceInvalidOptionGroup(t *testing.T) {
	cases := []struct {
		Name     string
		Src      string
		Response []*rds.OptionGroup
		Issues   []*issue.Issue
	}{
		{
			Name: "option_group is invalid",
			Src: `
resource "aws_db_instance" "mysql" {
    option_group_name = "app-server"
}`,
			Response: []*rds.OptionGroup{
				&rds.OptionGroup{
					OptionGroupName: aws.String("app-server1"),
				},
				&rds.OptionGroup{
					OptionGroupName: aws.String("app-server2"),
				},
			},
			Issues: []*issue.Issue{
				&issue.Issue{
					Type:    "ERROR",
					Message: "\"app-server\" is invalid option group name.",
					Line:    3,
					File:    "test.tf",
				},
			},
		},
		{
			Name: "option_group is valid",
			Src: `
resource "aws_db_instance" "mysql" {
    option_group_name = "app-server"
}`,
			Response: []*rds.OptionGroup{
				&rds.OptionGroup{
					OptionGroupName: aws.String("app-server1"),
				},
				&rds.OptionGroup{
					OptionGroupName: aws.String("app-server2"),
				},
				&rds.OptionGroup{
					OptionGroupName: aws.String("app-server"),
				},
			},
			Issues: []*issue.Issue{},
		},
	}

	for _, tc := range cases {
		c := config.Init()
		c.DeepCheck = true

		awsClient := c.NewAwsClient()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		rdsmock := awsmock.NewMockRDSAPI(ctrl)
		rdsmock.EXPECT().DescribeOptionGroups(&rds.DescribeOptionGroupsInput{}).Return(&rds.DescribeOptionGroupsOutput{
			OptionGroupsList: tc.Response,
		}, nil)
		awsClient.Rds = rdsmock

		var issues = []*issue.Issue{}
		TestDetectByCreatorName(
			"CreateAwsDBInstanceInvalidOptionGroupDetector",
			tc.Src,
			c,
			awsClient,
			&issues,
		)

		if !reflect.DeepEqual(issues, tc.Issues) {
			t.Fatalf("Bad: %s\nExpected: %s\n\ntestcase: %s", issues, tc.Issues, tc.Name)
		}
	}
}
