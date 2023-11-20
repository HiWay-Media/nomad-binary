# nomad-binary
This is a test project for experimenting with the Nomad raw_exec driver. It includes a sample binary that can be used as a Nomad task to execute arbitrary commands.

## Getting Started
### Prerequisites
- Nomad installed on your system.
- A running Nomad server and client.

### Nomad Job Configuration
The nomad-job.hcl file defines a Nomad job using the raw_exec driver. It specifies the sample binary as a task to be executed.

```hcl
Copy code
job "sample-job" {
  datacenters = ["dc1"]

  group "sample-group" {
    task "sample-task" {
      driver = "raw_exec"
      config {
        command = "/path/to/sample-binary"
      }
    }
  }
}
```
Feel free to modify the job configuration and the sample binary to experiment with different commands and scenarios.

## Contributing
Contributions to this test project are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License
This test project is open-source software licensed under the MIT License.