package Billing

import "github.com/aws/aws-sdk-go/service/costexplorer"

type Billinger interface {
	GetCostByTag(svc *costexplorer.CostExplorer)
	BillingCostSetTag(groupByTag string)
	BillingCostSetDimensionsValue(dimensionsValue string)
}

var (
	factoryByName = make(map[string]func() Billinger)
)

func Register(name string, factory func() Billinger) {
	factoryByName[name] = factory
}

func Create(name string) Billinger {
	if f, ok := factoryByName[name]; ok {
		return f()
	} else {
		panic("name not found")
	}
}
