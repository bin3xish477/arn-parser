package main

import (
	"fmt"
	"strings"

	"github.com/alexflint/go-arg"
)

var args struct {
	Arn string `arg:"-r,--arn,required" help:"Amazon Resource Name (ARN) to parse"`
}

const (
	red    = "\u001b[91m"
	green  = "\u001b[32m"
	blue   = "\u001b[94m"
	yellow = "\u001b[33m"
	purple = "\u001b[35m"
	end    = "\u001b[0m"
)

// validateARN validates if an ARN is valid
func validateARN(arn string) (isValid bool) {
	isValid = true
	if strings.Count(arn, ":") != 5 {
		isValid = false
	}
	return
}

// parseARN parses an Amazon Resource Name (ARN).
// An example ARN: "arn:partition:service:region:account-id:resource"
func parseARN(arn string) (arnMap map[string]string, err error) {
	arnMap = make(map[string]string)
	err = nil

	if !validateARN(arn) {
		err = fmt.Errorf("`%s` is not a valid ARN", arn)
		return
	}

	arnSlice := strings.Split(arn, ":")

	partition := arnSlice[1]
	service := arnSlice[2]
	region := arnSlice[3]
	accountID := arnSlice[4]
	resource := arnSlice[5]

	arnMap[fmt.Sprintf("%s%s%s", red, "Partition", end)] = partition
	arnMap[fmt.Sprintf("%s%s%s", green, "Service", end)] = service
	arnMap[fmt.Sprintf("%s%s%s", blue, "Region", end)] = region
	arnMap[fmt.Sprintf("%s%s%s", yellow, "Account-Id", end)] = accountID
	arnMap[fmt.Sprintf("%s%s%s", purple, "Resource", end)] = resource

	return
}

func main() {
	arg.MustParse(&args)

	// exampleArn := "arn:aws:ec2:us-east-1:123456789012:vpc/vpc-0e9801d129EXAMPLE"
	arnMap, err := parseARN(args.Arn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for key, val := range arnMap {
		if val == "" {
			fmt.Printf("%s: N/A\n", key)
		} else {
			fmt.Printf("%s: %s\n", key, val)
		}
	}
}
