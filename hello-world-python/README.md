# Module hello-world

This repo contains example code for a Viam module that provides an example camera and sensor resource.
To build it yourself, see the [Support additional hardware and software](https://docs.viam.com/operate/modules/support-hardware/).

## Model naomi:hello-world:hello-camera

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

| Name          | Type   | Inclusion | Description                      |
|---------------|--------|-----------|----------------------------------|
| `image_path`  | string | Required  | The path to the image to return. |

#### Example Configuration

```json
{
  "image_path": "path/to/image.jpg"
}
```

## Model naomi:hello-world:hello-sensor

A sensor that returns a random number.

### Configuration

This sensor does not have any configurable attributes.
