package main

import (
	Billing "awsbilling/Billing"
	session "awsbilling/Session"

	"github.com/aws/aws-sdk-go/service/costexplorer"
)

func main() {
	svc := costexplorer.New(session.GetSession())
	awsProduct := []string{"Amazon Simple Storage Service", "AWS Lambda"}
	c := Billing.Create("BillingCost")
	//c.BillingCostSetDimensionsValue("AWS Lambda")
	//c.GetCostByTag()
	for _, awsproduct := range awsProduct {
		c.BillingCostSetDimensionsValue(awsproduct)
		c.BillingCostSetTag("Name")
		c.GetCostByTag(svc)
	}

	//go c.GetCostByTag(svc, CostInputChan)
	//c.GetCostByTag(svc)

}
