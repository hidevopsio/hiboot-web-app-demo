# Hiboot web application demo

This is the Hiboot web application demo, we are going to demonstrate Hiboot we application programming step by step.

## Step 1, Writing the first Hiboot web application

To write Hiboot application, as we know, the executable commands must always use package main, so we need to create the main package first.

See [Effective GO](https://golang.org/doc/effective_go.html#names) to learn more about Go's naming conventions.

### Writing the code

```go
package main

import (
	"github.com/hidevopsio/hiboot/pkg/app/web"
)

func main()  {
	web.NewApplication().Run()
}
```

### Run the web application

As shown above, we just start the application, we did not write any business logic yet. But it's runnable, let's take a look,

```bash
go run main.go
```

The output will be,

```bash
______  ____________             _____
___  / / /__(_)__  /_______________  /_
__  /_/ /__  /__  __ \  __ \  __ \  __/
_  __  / _  / _  /_/ / /_/ / /_/ / /_     Hiboot Application Framework
/_/ /_/  /_/  /_.___/\____/\____/\__/     https://github.com/hidevopsio/hiboot

[INFO] 2018/10/26 18:29 Starting Hiboot web application hiboot-app on localhost with PID 15988
[INFO] 2018/10/26 18:29 Working directory: /Users/johnd/.gvm/pkgsets/go1.10/hidevops/src/github.com/hidevopsio/hiboot-web-app-demo
[INFO] 2018/10/26 18:29 The following profiles are active: default, []
Now listening on: http://localhost:8080
Application started. Press CMD+C to shut down.

```