# AAD JWT Token Generator

## Overview
AAD JWT Token Generator is a web-based application designed to generate JWT tokens from Azure AD. It uses a client-server architecture with a Golang backend and a Vue.js frontend, styled with Tailwind CSS. The application allows users to input their Azure AD credentials (`clientId`, `tenantId`, `clientSecret`, and `scope`) and receive a JWT token in response.

## Features
- **Web Interface**: A clean and professional UI built with Vue.js and Tailwind CSS
- **Flexible Token Generation**: Allows specifying various parameters for JWT token generation
- **Headless Mode**: Can be run as a command-line tool for generating JWT tokens without the web interface
- **Docker Support**: Easy deployment using Docker containers
- **Embedded Static Files**: All frontend assets are embedded in the binary for easy distribution

## Prerequisites
- Go 1.21.0 or later
- Docker (optional, for containerized deployment)

## Installation

### Using Go
```bash
go install github.com/sanjaysingh/aad-jwt-gen@latest
```

### Using Docker
```bash
docker pull sanjaysingh/aad-jwt-gen
```

## Usage

### Web Interface
1. Run the application:
   ```bash
   aad-jwt-gen
   ```
2. Open a web browser and navigate to http://localhost:5555/
3. Enter your Azure AD credentials:
   - Client ID
   - Tenant ID
   - Client Secret
   - Scope (optional, defaults to ".default")
4. Click on "Generate Token" to receive the JWT token

### Headless Mode
```bash
aad-jwt-gen -headless \
  -clientId="your_client_id" \
  -tenantId="your_tenant_id" \
  -clientSecret="your_client_secret" \
  -scope="your_scope"
```

### Docker Deployment
```bash
docker run -p 5555:5555 sanjaysingh/aad-jwt-gen
```

## Development

### Building from Source
1. Clone the repository:
   ```bash
   git clone https://github.com/sanjaysingh/aad-jwt-gen.git
   cd aad-jwt-gen
   ```
2. Build the application:
   ```bash
   go build
   ```

### Project Structure
- `cmd/`: Command-line interface implementation
- `handlers/`: HTTP request handlers
- `services/`: Core business logic
- `static/`: Frontend assets (Vue.js + Tailwind CSS)
- `main.go`: Application entry point

## License
This project is licensed under the terms specified in the LICENSE file.

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

