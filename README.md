# Test HTTP Client for Datetime Server Implemented in Go
Create an HTTP client in Go that consumes the datetime server APIs implemented in the previous project. This client will focus on building an HTTP client development and testing.

# Installation

To install the project use:

```golang
go get github.com/codescalersinternships/datetime-client-nabil
```

## Usage

1. Import the Package
```golang
import Client "github.com/codescalersinternships/datetime-client-nabil/pkg"
```

Using url string

2. Create a New Client Instance
```golang
myClient := Client.NewClient("http://localhost:8090", time.Duration(1*time.Second)
```
You can replace "http://localhost:8090" with the url your server listens to.

3. Get date and time
```golang
data, err := myClient.GetCurrentDate()
```

Or using Enviroment variables

2. Create a New Client Instance using Enviroment variables.
```golang
myClient, err := datetimeclient.NewClientUsingEnv("mybaseurlkey", time.Duration(1)*time.Second)
```

3. Get date and time
```golang
data, err := myClient.GetCurrentDate()
```
## How to Test

```golang
make test
```