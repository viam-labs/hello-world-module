package helloworld

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"os"

	camera "go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/components/camera/rtppassthrough"
	"go.viam.com/rdk/gostream"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/pointcloud"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/spatialmath"
)

var (
	HelloCamera      = resource.NewModel("naomi", "hello-world", "hello-camera")
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
// Returns implicit required (first return) and optional (second return) dependencies based on the config.
// The path is the JSON path in your robot's config (not the `Config` struct) to the
// resource being validated; e.g. "components.0".
func (cfg *Config) Validate(path string) ([]string, []string, error) {
    var deps []string
    if cfg.ImagePath == "" {
        return nil, nil, resource.NewConfigValidationFieldRequiredError(path, "image_path")
    }
    if reflect.TypeOf(cfg.ImagePath).Kind() != reflect.String {
        return nil, nil, errors.New("image_path must be a string.")
    }
    imagePath = cfg.ImagePath
    return deps, []string{}, nil
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
	var videoStreamRetVal gostream.VideoStream

	return videoStreamRetVal, fmt.Errorf("not implemented")
}

// Image returns a byte slice representing an image that tries to adhere to the MIME type hint.
// Image also may return metadata about the frame.
func (s *helloWorldHelloCamera) Image(ctx context.Context, mimeType string, extra map[string]interface{}) ([]byte, camera.ImageMetadata, error) {
    imgFile, err := os.Open(imagePath)
    if err != nil {
       return nil, camera.ImageMetadata{}, errors.New("Error opening image.")
    }
    defer imgFile.Close()
    imgByte, err := os.ReadFile(imagePath)
    return imgByte, camera.ImageMetadata{}, nil
}

// Images is used for getting simultaneous images from different imagers,
// along with associated metadata (just timestamp for now). It's not for getting a time series of images from the same imager.
// The extra parameter can be used to pass additional options to the camera resource. The filterSourceNames parameter can be used to filter
// only the images from the specified source names. When unspecified, all images are returned.
func (s *helloWorldHelloCamera) Images(ctx context.Context, filterSourceNames []string, extra map[string]interface{}) ([]camera.NamedImage, resource.ResponseMetadata, error) {
    var responseMetadataRetVal resource.ResponseMetadata

    imgFile, err := os.Open(imagePath)
    if err != nil {
        return nil, responseMetadataRetVal, errors.New("Error opening image.")
    }
    defer imgFile.Close()

    imgByte, err := os.ReadFile(imagePath)
    if err != nil {
        return nil, responseMetadataRetVal, err
    }

	named, err := camera.NamedImageFromBytes(imgByte, "default", "image/png")
    if err != nil {
        return nil, responseMetadataRetVal, err
    }

    return []camera.NamedImage{named}, responseMetadataRetVal, nil
}

// NextPointCloud returns the next immediately available point cloud, not necessarily one
// a part of a sequence. In the future, there could be streaming of point clouds.
func (s *helloWorldHelloCamera) NextPointCloud(ctx context.Context, extra map[string]interface{}) (pointcloud.PointCloud, error) {
	var pointCloudRetVal pointcloud.PointCloud

	return pointCloudRetVal, fmt.Errorf("not implemented")
}

// Properties returns properties that are intrinsic to the particular
// implementation of a camera.
func (s *helloWorldHelloCamera) Properties(ctx context.Context) (camera.Properties, error) {
	var propertiesRetVal camera.Properties

	return propertiesRetVal, fmt.Errorf("not implemented")
}

func (s *helloWorldHelloCamera) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *helloWorldHelloCamera) Geometries(ctx context.Context, extra map[string]interface{}) ([]spatialmath.Geometry, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *helloWorldHelloCamera) SubscribeRTP(ctx context.Context, bufferSize int, packetsCB rtppassthrough.PacketCallback) (rtppassthrough.Subscription, error) {
	var subscriptionRetVal rtppassthrough.Subscription

	return subscriptionRetVal, fmt.Errorf("not implemented")
}

func (s *helloWorldHelloCamera) Unsubscribe(ctx context.Context, id rtppassthrough.SubscriptionID) error {
	return fmt.Errorf("not implemented")
}

func (s *helloWorldHelloCamera) Close(context.Context) error {
	// Put close code here
	s.cancelFunc()
	return nil
}
