# Module hello-world

This repo contains example code for a Viam module that provides an example camera and sensor resource.
To build it yourself, see the [Create a Hello World module guide](https://docs.viam.com/operate/get-started/other-hardware/hello-world-module/).

Note that the example code in this repo uses the `jessamy` namespace, but you would use your own namespace when authoring your own module.

## Model jessamy:hello-world:hello-camera

A camera that returns a static image.

### Configuration

The following attribute template can be used to configure this model:

```json
{
  "image_path": <string>
}
```

#### Attributes

The following attributes are available for this model:

| Name         | Type   | Inclusion | Description                      |
|--------------|--------|-----------|----------------------------------|
| `image_path` | string | Required  | The path to the image to return. |

#### Example Configuration

```json
{
  "image_path": "path/to/image.jpg"
}
```

## Model jessamy:hello-world:hello-sensor

A sensor that returns a random number.

### Configuration

This sensor does not have any configurable attributes.
