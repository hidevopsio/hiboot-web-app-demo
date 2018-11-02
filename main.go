// Copyright 2018 John Deng (hi.devops.io@gmail.com).
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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