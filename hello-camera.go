package helloworld

import (
	"context"
	"errors"
	"os"
	"reflect"

	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/gostream"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/pointcloud"
	"go.viam.com/rdk/resource"
	"go.viam.com/utils/rpc"
)

var (
	HelloCamera      = resource.NewModel("jessamy", "hello-world", "hello-camera")
	errUnimplemented = errors.New("unimplemented")
	imagePath        = ""
)

func init() {
	resource.RegisterComponent(camera.API, HelloCamera,
		resource.Registration[camera.Camera, *Config]{
			Constructor: newHelloWorldHelloCamera,
		},
	)
}

type Config struct {
	resource.AlwaysRebuild
	ImagePath string `json:"image_path"`
}

// Validate ensures all parts of the config are valid and important fields exist.
// Returns implicit dependencies based on the config.
// The path is the JSON path in your robot's config (not the `Config` struct) to the
// resource being validated; e.g. "components.0".
func (cfg *Config) Validate(path string) ([]string, error) {
	var deps []string
	if cfg.ImagePath == "" {
		return nil, resource.NewConfigValidationFieldRequiredError(path, "image_path")
	}
	if reflect.TypeOf(cfg.ImagePath).Kind() != reflect.String {
		return nil, errors.New("image_path must be a string")
	}
	imagePath = cfg.ImagePath
	return deps, nil
}

type helloWorldHelloCamera struct {
	resource.AlwaysRebuild

	name resource.Name

	logger logging.Logger
	cfg    *Config

	cancelCtx  context.Context
	cancelFunc func()
}

func newHelloWorldHelloCamera(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (camera.Camera, error) {
	conf, err := resource.NativeConfig[*Config](rawConf)
	if err != nil {
		return nil, err
	}

	return NewHelloCamera(ctx, deps, rawConf.ResourceName(), conf, logger)

}

func NewHelloCamera(ctx context.Context, deps resource.Dependencies, name resource.Name, conf *Config, logger logging.Logger) (camera.Camera, error) {

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	s := &helloWorldHelloCamera{
		name:       name,
		logger:     logger,
		cfg:        conf,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
	}
	return s, nil
}

func (s *helloWorldHelloCamera) Name() resource.Name {
	return s.name
}

func (s *helloWorldHelloCamera) Stream(ctx context.Context, errHandlers ...gostream.ErrorHandler) (gostream.VideoStream, error) {
	return nil, errors.New("not implemented")
}

func (s *helloWorldHelloCamera) Image(ctx context.Context, mimeType string, extra map[string]interface{}) ([]byte, camera.ImageMetadata, error) {
	imgFile, err := os.Open(imagePath)
	if err != nil {
		return nil, camera.ImageMetadata{}, errors.New("error opening image")
	}
	defer imgFile.Close()
	imgByte, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, camera.ImageMetadata{}, errors.New("error reading image")
	}
	return imgByte, camera.ImageMetadata{}, nil
}

func (s *helloWorldHelloCamera) NewClientFromConn(ctx context.Context, conn rpc.ClientConn, remoteName string, name resource.Name, logger logging.Logger) (camera.Camera, error) {
	return nil, errors.New("not implemented")
}

func (s *helloWorldHelloCamera) Images(ctx context.Context) ([]camera.NamedImage, resource.ResponseMetadata, error) {
	return []camera.NamedImage{}, resource.ResponseMetadata{}, errors.New("not implemented")
}

func (s *helloWorldHelloCamera) NextPointCloud(ctx context.Context) (pointcloud.PointCloud, error) {
	return nil, errors.New("not implemented")
}

func (s *helloWorldHelloCamera) Properties(ctx context.Context) (camera.Properties, error) {
	return camera.Properties{}, errors.New("not implemented")
}

func (s *helloWorldHelloCamera) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{}, errors.New("not implemented")
}

func (s *helloWorldHelloCamera) Close(context.Context) error {
	// Put close code here
	s.cancelFunc()
	return nil
}
