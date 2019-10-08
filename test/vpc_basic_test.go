package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)
//	"github.com/stretchr/testify/assert"

func TestVPCBasic(t *testing.T) {
	t.Parallel()

	awsRegion := "ap-northeast-2"
	vpcName := fmt.Sprintf("terratest-vpc")
	expectedCidr := "10.10.0.0/16"

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "fixtures/default",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"aws_region": awsRegion,
			"cidr_block": expectedCidr,
			"vpc_name": vpcName,
		},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	terraformOutputs := terraform.OutputAll(t, terraformOptions)
	fmt.Println("Output:", terraformOutputs)

	// Retreive outputs from the terraform output
	vpcId := terraform.Output(t, terraformOptions, "vpc_id")
	fmt.Println("VpcId:", vpcId)

	// Verify we're getting back the outputs we expect
}
