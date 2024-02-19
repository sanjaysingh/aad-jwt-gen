# AAD JWT Token Generator

## Overview
AAD JWT Token Generator is a web-based application designed to generate JWT tokens from Azure AD. It uses a client-server architecture with a Golang backend and a Vue.js frontend, styled with Tailwind CSS. The application allows users to input their Azure AD credentials (`clientId`, `tenantId`, `clientSecret`, and `scope`) and receive a JWT token in response.

## Features
- **Web Interface**: A clean and professional UI built with Vue.js and Tailwind CSS.
- **Flexible Token Generation**: Allows specifying various parameters for JWT token generation.
- **Headless Mode**: Can be run as a command-line tool for generating JWT tokens without the web interface.

## Installation
To install the application, follow these steps:

### Steps
1. Install the application using Go:
   ```bash
   go install github.com/sanjaysingh/aad-jwt-gen@latest

## Web Interface
1. Run the application:
```bash
aad-jwt-gen
```
2. Open a web browser and navigate to http://localhost:5555/.
3. Enter your Azure AD clientId, tenantId, clientSecret, and scope.
4. Click on "Generate Token" to receive the JWT token.

## Headless Mode
```bash
aad-jwt-gen -headless -clientId="your_client_id" -tenantId="your_tenant_id" -clientSecret="your_client_secret" -scope="your_scope"
```

## Run from Docker
```bash
docker run -p 5555:5555 sanjaysingh/aad-jwt-gen
```

