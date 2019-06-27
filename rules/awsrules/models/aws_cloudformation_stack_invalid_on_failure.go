// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudformationStackInvalidOnFailureRule checks the pattern is valid
type AwsCloudformationStackInvalidOnFailureRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCloudformationStackInvalidOnFailureRule returns new rule with default attributes
func NewAwsCloudformationStackInvalidOnFailureRule() *AwsCloudformationStackInvalidOnFailureRule {
	return &AwsCloudformationStackInvalidOnFailureRule{
		resourceType:  "aws_cloudformation_stack",
		attributeName: "on_failure",
		enum: []string{
			"DO_NOTHING",
			"ROLLBACK",
			"DELETE",
		},
	}
}

// Name returns the rule name
func (r *AwsCloudformationStackInvalidOnFailureRule) Name() string {
	return "aws_cloudformation_stack_invalid_on_failure"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudformationStackInvalidOnFailureRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCloudformationStackInvalidOnFailureRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudformationStackInvalidOnFailureRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudformationStackInvalidOnFailureRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`on_failure is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}