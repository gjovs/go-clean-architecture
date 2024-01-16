# Go Clean Architecture Repository 🚀

Welcome to the Clean Architecture At Go repository! This project follows the principles of Clean Architecture in Go, providing a scalable and maintainable foundation for building robust applications. 🛠️

## Table of Contents 📚

- [Go Clean Architecture Repository 🚀](#go-clean-architecture-repository-)
  - [Table of Contents 📚](#table-of-contents-)
  - [Introduction 🌐](#introduction-)
  - [Getting Started 🚦](#getting-started-)
  - [Project Structure 🏗️](#project-structure-️)
  - [Dependencies 📦](#dependencies-)
  - [Usage 🚀](#usage-)
  - [Testing 🧪](#testing-)
  - [Contributing 🤝](#contributing-)
  - [License 📄](#license-)

## Introduction 🌐

This project aims to demonstrate the implementation of Clean Architecture in Go, separating concerns into distinct layers to achieve better testability, maintainability, and flexibility. The architecture consists of the following layers:

- **Entities**: Represents the core business logic.
- **Use Cases**: Contains application-specific business rules.
- **Interfaces**: Defines the boundaries of the application.
- **Adapters**: Implements the interfaces, connecting the application to external frameworks or tools.

## Getting Started 🚦

To get started with this project, make sure you have Go installed on your machine. Clone the repository and run:

```bash
go build -o ./bin/ source.go
```

This will build and execute the application.

## Project Structure 🏗️

The project structure follows the Clean Architecture principles, separating concerns into different packages. Here's an overview:

```
/api
/configs
/cmd
    /app
/internal
    /entities
    /usecases
    /event
    /infra
/pkg
  /types
```

- **cmd**: Contains the application's entry point and configuration.
- **internal**: Holds the core application code.
- **pkg**: Houses reusable packages.
  
## Dependencies 📦

This project uses default golang dependency managing tool  for managing external dependencies. To install the required dependencies, run:

```bash
go install 
```

## Usage 🚀

Describe how to use your application or library here. Include code snippets or examples to guide users.

## Testing 🧪

Ensure that all unit tests and integration tests pass before submitting a pull request. Run the tests using:

```bash
go test ./...
```

## Contributing 🤝

Contributions are welcome! Please follow our [Contribution Guidelines](CONTRIBUTING.md) when submitting pull requests.

## License 📄

This project is licensed under the [MIT License](LICENSE.md) - see the [LICENSE.md](LICENSE.md) file for details.
