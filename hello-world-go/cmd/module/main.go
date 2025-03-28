package main

import (
	"helloworld"

	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
)

func main() {
	// ModularMain can take multiple APIModel arguments, if your module implements multiple models.
	module.ModularMain(resource.APIModel{camera.API, helloworld.HelloCamera}, resource.APIModel{sensor.API, helloworld.HelloSensor})
}
