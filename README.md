# Go-dummy

This is a Go application built with the Gin framework. It includes several endpoints for various functionalities.

## Prerequisites

- Go (version 1.20 or higher) installed
- AWS credentials properly configured if using the `/` endpoint

## Endpoints

The application provides the following endpoints:

- `/`: Returns the AWS caller identity information in JSON format.
- `/counter`: Returns the current count of requests made to the server.
- `/ping`: Returns "pong" as a response.
- `/switch`: Toggles the switch status between true and false, returning the updated status as a JSON response.
- `/liveness`: Returns a response indicating the liveness probe status (200 OK if the switch is true, otherwise 503 Service Unavailable).
- `/readiness`: Returns a response indicating the readiness probe status (200 OK if the switch is true, otherwise 503 Service Unavailable).

### Response
+ `/`
```=json
{
  "Account": "111111111",
  "Arn": "arn:aws:iam::111111111:user/foo.var",
  "UserId": "C4NTH1SB3AR34LUS3R1D"
}
```

## Configuration

No additional configuration is required to run the application. However, if you want to use the / endpoint to retrieve AWS caller identity information, ensure that your AWS credentials are properly configured.

