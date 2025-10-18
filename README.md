# Namespace SDK Examples

This repository contains examples demonstrating how to use the Namespace SDK across different programming languages.

## Repository Structure

```
examples/
├── go/          # Go SDK examples
└── README.md    # This file
```

## Go Examples

The Go examples are located in the `go/` directory. Each example demonstrates different aspects of the Namespace SDK:

### Available Examples

| Example | Description |
|---------|-------------|
| **build** | Building Docker images using BuildKit with authentication and secrets management |
| **createinstance** | Creating compute instances with specific configurations and running containers |
| **createpartnerinstance** | AWS identity federation with Namespace and creating instances with federated credentials |
| **ensuretenant** | Creating and managing tenants using AWS federation for identity management |
| **buildandrun** | Complete workflow: build image, create instance, run service, and connect via gRPC |
| **macrun** | Building and running Go applications on macOS/ARM64 instances |
| **sidecar** | Advanced instance configuration with sidecar containers and SSH access |

### Getting Started with Go Examples

1. **Prerequisites**
   - Go 1.21 or later
   - Namespace authentication credentials
   - For AWS federation examples: AWS credentials with appropriate permissions

2. **Running an Example**
   ```bash
   cd go/<example-name>
   go run .
   ```

3. **Common SDK Patterns**
   - **Authentication**: Token loading from workstations (`auth.LoadUserToken()`) or instances (`auth.LoadWorkloadToken()`)
   - **Compute API**: Creating instances, waiting for readiness, describing instances
   - **Builds API**: Building images with BuildKit, pushing to registry
   - **IAM API**: Tenant creation and management with external account federation
   - **gRPC Communication**: Client-server communication with TLS

### Example Details

#### build
Demonstrates using the Namespace Builds SDK with BuildKit to build Docker images from Dockerfiles.

```bash
cd go/build
go run build.go
```

#### createinstance
Shows how to create compute instances with specific CPU/memory/architecture configurations and run containers.

```bash
cd go/createinstance
go run createinstance.go
```

#### createpartnerinstance
Demonstrates AWS identity federation with Namespace and creating instances with federated credentials.

```bash
cd go/createpartnerinstance
go run createinstance.go
```

#### ensuretenant
Shows how to create and manage tenants using AWS federation for identity management.

```bash
cd go/ensuretenant
go run ensuretenant.go
```

#### buildandrun
Complete end-to-end example: build a Docker image, create an instance, run a gRPC service, and connect to it.

```bash
cd go/buildandrun
go run buildandrun.go
```

#### macrun
Demonstrates building Go binaries for macOS/ARM64 and running them on Apple Silicon instances.

```bash
cd go/macrun
go run macrun.go
```

#### sidecar
Advanced example showing instance configuration with sidecar containers, volume mounting, and SSH access.

```bash
cd go/sidecar
go run main.go
```

## Documentation

For more information about the Namespace SDK and APIs, please visit:
- [Namespace Documentation](https://namespace.so/docs)
- [SDK Reference](https://pkg.go.dev/namespacelabs.dev/integrations)

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

See LICENSE file for details.
