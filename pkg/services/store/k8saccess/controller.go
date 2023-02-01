package k8saccess

import (
	"context"
	"fmt"
	"time"

	"github.com/grafana/grafana/pkg/apimachinery/bridge"
	"github.com/grafana/grafana/pkg/registry/corecrd"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/cache"
)

type DashboardController struct {
	bridgeService *bridge.Service
	reg           *corecrd.Registry
}

func ProvideDashboardController(bridgeService *bridge.Service, reg *corecrd.Registry) *DashboardController {
	return &DashboardController{
		bridgeService: bridgeService,
		reg:           reg,
	}
}

func (c *DashboardController) Run(ctx context.Context) error {
	dashboardCRD := c.reg.Dashboard()

	gvr := schema.GroupVersionResource{
		Group:    dashboardCRD.GVK().Group,
		Version:  dashboardCRD.GVK().Version,
		Resource: dashboardCRD.Schema.Name,
	}

	factory := dynamicinformer.NewDynamicSharedInformerFactory(c.bridgeService.ClientSet.Dynamic, time.Minute)
	dashboardInformer := factory.ForResource(gvr).Informer()

	dashboardInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Printf("dashboard added: %s \n", obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Printf("dashboard changed: %s \n", newObj)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Printf("dashboard deleted: %s \n", obj)
		},
	})

	stop := make(chan struct{})
	defer close(stop)

	factory.Start(stop)
	<-ctx.Done()
	return nil
}
