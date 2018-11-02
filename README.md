# Hiboot web application demo

This is the Hiboot web application demo, we are going to demonstrate Hiboot we application programming step by step.

## Step 1, Writing the first Hiboot web application

To write Hiboot application, as we know, the executable commands must always use package main, so we need to create the main package first.

See [Effective GO](https://golang.org/doc/effective_go.html#names) to learn more about Go's naming conventions.

## Step 2, adding some starters

Here we are going to add starter [actuator](https://hidevops.io/hiboot/tree/master/pkg/starter/actuator) and [logging](https://hidevops.io/hiboot/tree/master/pkg/starter/logging).

## Step 3, adding Hiboot controller

### Writing the code

```go

package main

import (
	"hidevops.io/hiboot/pkg/app/web"
	"hidevops.io/hiboot/pkg/app"
	"hidevops.io/hiboot/pkg/starter/actuator"
	"hidevops.io/hiboot/pkg/starter/logging"
	"hidevops.io/hiboot/pkg/at"
)

// Controller Rest Controller with path /
// RESTful Controller, derived from web.Controller. The context mapping of this controller is '/' by default
type Controller struct {
	// at.RestController or web.Controller must be embedded here
	at.RestController
}

// Get GET /
// Get method, the context mapping of this method is '/' by default
// the Method name Get means that the http request method is GET
func (c *Controller) Get() string {
	// response
	return "My first Hiboot web application"
}

func main()  {
	web.NewApplication(new(Controller)).
		SetProperty(app.ProfilesInclude, actuator.Profile, logging.Profile).
		Run()
}

```

### Run the web application

As shown above, we just start the application, we did not write any business logic yet. But it's runnable, let's take a look,

```bash
go run main.go
```

The output will be,

```bash
___  / / /__(_)__  /_______________  /_
__  /_/ /__  /__  __ \  __ \  __ \  __/
_  __  / _  / _  /_/ / /_/ / /_/ / /_     Hiboot Application Framework
/_/ /_/  /_/  /_.___/\____/\____/\__/     https://hidevops.io/hiboot

[INFO] 2018/11/03 07:47 Starting Hiboot web application hiboot-app on localhost with PID 4092
[INFO] 2018/11/03 07:47 Working directory: /Users/johnd/.gvm/pkgsets/go1.10/hidevops/src/hidevops.io/hiboot-web-app-demo
[INFO] 2018/11/03 07:47 The following profiles are active: local, [actuator logging web]
[INFO] 2018/11/03 07:47 Initializing Hiboot Application
[INFO] 2018/11/03 07:47 Auto configure web starter
[INFO] 2018/11/03 07:47 Auto configure actuator starter
[INFO] 2018/11/03 07:47 Auto configure logging starter
[INFO] 2018/11/03 07:47 Resolving dependencies
[INFO] 2018/11/03 07:47 Injecting dependencies
[INFO] 2018/11/03 07:47 Injected dependencies
[INFO] 2018/11/03 07:47 Mapped "/" onto main.Controller.Get()
[INFO] 2018/11/03 07:47 Mapped "/health" onto actuator.healthController.Get()
[INFO] 2018/11/03 07:47 Hiboot started on port(s) http://localhost:8080
[INFO] 2018/11/03 07:47 Started hiboot-app in 0.003143 seconds
```

As you can see above, the starter actuator and logging is auto configured.

Now we have health check endpoint GET /health, let's use [httpie](https://httpie.org/) to check this endpoint.

```bash
http :8080/health
```

We can see the output of above request, my first Hiboot application is up and running. Yeah!

```bash
HTTP/1.1 200 OK
Content-Length: 15
Content-Type: application/json; charset=UTF-8
Date: Fri, 26 Oct 2018 10:49:12 GMT

{
    "status": "UP"
}
```

The application log also prints the request info.

```bash

[INFO] 2018/10/26 18:49 200 215.239µs ::1 GET /health HTTPie/0.9.3

```
