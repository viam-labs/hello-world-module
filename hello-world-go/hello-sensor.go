package helloworld

import (
	"context"
	"fmt"
	"math/rand"

	sensor "go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

var (
	HelloSensor      = resource.NewModel("naomi", "hello-world", "hello-sensor")
)

func init() {
	resource.RegisterComponent(sensor.API, HelloSensor,
		resource.Registration[sensor.Sensor, *sensorConfig]{
			Constructor: newHelloWorldHelloSensor,
		},
	)
}

type sensorConfig struct {
	/*
		Put config attributes here. There should be public/exported fields
		with a `json` parameter at the end of each attribute.

		Example config struct:
			type Config struct {
				Pin   string `json:"pin"`
				Board string `json:"board"`
				MinDeg *float64 `json:"min_angle_deg,omitempty"`
			}

		If your model does not need a config, replace *sensorConfig in the init
		function with resource.NoNativeConfig
	*/
}

// Validate ensures all parts of the config are valid and important fields exist.
// Returns implicit required (first return) and optional (second return) dependencies based on the config.
// The path is the JSON path in your robot's config (not the `Config` struct) to the
// resource being validated; e.g. "components.0".
func (cfg *sensorConfig) Validate(path string) ([]string, []string, error) {
	// Add config validation code here
	return nil, nil, nil
}

type helloWorldHelloSensor struct {
	resource.AlwaysRebuild

	name resource.Name

	logger logging.Logger
	cfg    *sensorConfig

	cancelCtx  context.Context
	cancelFunc func()
}

func newHelloWorldHelloSensor(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (sensor.Sensor, error) {
	conf, err := resource.NativeConfig[*sensorConfig](rawConf)
	if err != nil {
		return nil, err
	}

	return NewHelloSensor(ctx, deps, rawConf.ResourceName(), conf, logger)

}

func NewHelloSensor(ctx context.Context, deps resource.Dependencies, name resource.Name, conf *sensorConfig, logger logging.Logger) (sensor.Sensor, error) {

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	s := &helloWorldHelloSensor{
		name:       name,
		logger:     logger,
		cfg:        conf,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
	}
	return s, nil
}

func (s *helloWorldHelloSensor) Name() resource.Name {
	return s.name
}

func (s *helloWorldHelloSensor) Readings(ctx context.Context, extra map[string]interface{}) (map[string]interface{}, error) {
    number := rand.Float64()
    return map[string]interface{}{
        "random_number": number,
    }, nil
}

func (s *helloWorldHelloSensor) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *helloWorldHelloSensor) Close(context.Context) error {
	// Put close code here
	s.cancelFunc()
	return nil
}
