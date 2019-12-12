package Billing

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/costexplorer"
)

type BillingCost struct {
	groupByTag      string
	dimensionsValue string
}

func (arg *BillingCost) BillingCostSetTag(groupByTag string) {
	arg.groupByTag = groupByTag
}

func (arg *BillingCost) BillingCostSetDimensionsValue(dimensionsValue string) {
	arg.dimensionsValue = dimensionsValue
}
func (arg *BillingCost) GetCostByTag(svc *costexplorer.CostExplorer) {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	queryTime := &costexplorer.DateInterval{}
	queryTime.SetEnd(fmt.Sprintf("%d-%d-%d", nTime.Year(), nTime.Month(), nTime.Day()))
	queryTime.SetStart(fmt.Sprintf("%d-%d-%d", yesTime.Year(), yesTime.Month(), yesTime.Day()))
	getCostInput := &costexplorer.GetCostAndUsageInput{}
	getCostInput.SetTimePeriod(queryTime)
	getCostInput.SetGranularity("DAILY")
	groupDefinition := &costexplorer.GroupDefinition{}
	groupDefinition.SetKey(arg.groupByTag)
	groupDefinition.SetType("TAG")
	metrics := "BLENDED_COST"
	groupby := []*costexplorer.GroupDefinition{groupDefinition}
	getCostInput.SetMetrics([]*string{&metrics})
	getCostInput.SetGroupBy(groupby)

	filter := &costexplorer.Expression{}
	dimensions := &costexplorer.DimensionValues{}
	dimensions.SetKey("SERVICE")

	dimensions.SetValues([]*string{&arg.dimensionsValue})
	filter.SetDimensions(dimensions)
	getCostInput.SetFilter(filter)
	log.Println(svc.GetCostAndUsage(getCostInput))
	//CostInputChan := make(chan *costexplorer.GetCostAndUsageInput)
	//CostInputChan <- getCostInput
	//log.Println(len(CostInputChan), cap(CostInputChan))
}

func BillingCostReg() Billinger {
	return &BillingCost{groupByTag: "Name"}
}

func init() {
	Register("BillingCost", BillingCostReg)
}
