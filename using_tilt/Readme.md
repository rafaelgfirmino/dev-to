# Go Development Has Never Been So Fast: Discover How Tilt Changes Everything!

Application development can be a challenging process, especially when it comes to complex environments and multiple services. For Go developers, finding ways to accelerate the workflow without compromising quality is a priority. Fortunately, tools like Tilt arise to transform this reality. In this article, we will explore how Tilt can boost your Go development process, making it more agile and efficient.

## What is Tilt?

Tilt is a development tool that facilitates the automation and integration of microservices workflows. It allows you to have a unified view of your development environment, providing agility to compile, test, and deploy applications quickly and dynamically.

Although it is more commonly associated with Kubernetes environments, Tilt can also be a powerful tool to speed up Go application development. It allows you to configure and observe changes in real-time, automatically updating your environment.

## How Does Tilt Accelerate Go Development?

### 1. Automatic Compilation and Deployment
Tilt eliminates the manual process of compiling and deploying code. Instead of executing several commands repeatedly, the tool watches for changes in files and performs compilation and deployment automatically. For Go developers, this means less time spent in the build cycle and more time focusing on the more important tasks.

### 2. Integration with Kubernetes and Docker
Tilt integrates seamlessly with Kubernetes and Docker, which makes it easier to run containers and services in local development environments. For those already working with containers in Go, Tilt offers a simplified way to configure and monitor services, optimizing the development workflow.

### 3. Real-Time Updates
With Tilt, whenever you make any modification to the code, it automatically detects these changes and applies them to the environment in real-time. This allows developers to immediately see the impact of their changes, without needing to restart the entire process manually. This significantly reduces the time between coding and verification.

### 4. Full Visibility and Control
Tilt provides an interactive interface that gives you a complete view of services, logs, and container status. This is especially useful in more complex environments where multiple services are running. You have full control over what’s happening, which helps to identify problems faster and optimize development.

## How to Configure Tilt for a Go Project?

Here’s a simple step-by-step guide to integrate Tilt into your Go project:

### Step 1: Install Tilt
First, you need to install [Tilt](https://docs.tilt.dev/install.html). In the terminal, use the following command (for Linux systems):

### Step 2: Tiltfile Configuration
In the root directory of your Go project, create a file called `Tiltfile`. This file will contain the configuration needed for Tilt to identify and monitor your code.

Basic example of Tiltfile configuration for Go:
This command generates mock files from mockery every time a file within the `./internal/port` folder is changed.
```bash
local_resource(
    name='mock update',
    labels=['test'],
    cmd='mockery',
    deps=['./internal/port'],
    ignore=['./internal/port/mock_*.go'],
)
```
We can also add a button with actions if you want to perform some actions manually:
```bash
local_resource(
    name='Run tests',
    labels=['test'],
    cmd='go test ./...',
    auto_init=False,
)
cmd_button('Run_tests',
    argv=['sh', '-c', 'go test ./...'], //command to be executed
    resource='Run tests', //the value here needs to be the same as the name of the local_resource
    icon_name='play_arrow', //we can use material design icons here
    text='Run tests', // button tooltip
)
```

Use Cases: Tilt in Action
Here are some examples of how Tilt can be useful in different Go development scenarios:

Microservices Development: For those working with a microservices architecture, Tilt simplifies the management of multiple containers and services. It allows you to quickly change a service and see the results in real-time, without interrupting the entire infrastructure.

Testing Environments: By automating the build and deploy cycle, Tilt helps quickly set up testing environments to ensure your Go code works as expected.

Continuous Integration: Tilt can easily be integrated with CI/CD systems to speed up the code integration and deployment process.

Conclusion
If you are a Go developer and are looking for ways to make your workflow more agile, Tilt is the tool that can transform your development approach. With the automation of the build cycle, deployment, and change monitoring, Tilt lets you focus on what truly matters: writing high-quality code.

Try Tilt and see how it can accelerate your Go development, taking your productivity to a new level!
![Exemplo](./example.gif)
