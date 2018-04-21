// Coder: koalatea
// Email: koalateac@gmail.com

package main

import (
	"net/http"
)

// Route is a struct to represent a route for the router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Auth        []string
	API         bool
}

// Routes is a splice of multiple Route for use by the router
type Routes []Route

func (app *App) GenerateRoutes() []Route {
	var routes = Routes{
		Route{"Index", "GET", "/", app.Home, nil, false},
		Route{"Devices", "GET", "/devices", app.viewDevices, []string{"user"}, false},
		Route{"NewDevice", "GET", "/devices/new", app.newDevice, []string{"user"}, false},
		Route{"CreateDevice", "POST", "/devices/new", app.createDevice, []string{"user"}, false},
		Route{"ViewDevice", "GET", "/devices/{deviceID}", app.viewDevice, []string{"user"}, false},
		Route{"Subnets", "GET", "/subnets", app.viewSubnets, nil, false},
		Route{"ViewSubnet", "GET", "/subnets/{subnetID}", app.viewSubnet, nil, false},
		Route{"NewReport", "GET", "/reports/new", app.newReport, []string{"user"}, false},
		Route{"CreateReport", "POST", "/reports/new", app.createReport, []string{"user"}, false},
		Route{"ViewReport", "GET", "/reports/{reportID}", app.viewReport, []string{"user"}, false},
		//Route{"IPs", "GET", "/ips", viewIPs, nil, false},

		// users/auth
		Route{"User Signup Page", "GET", "/user/signup", app.SignupUser, nil, false},
		Route{"Signup User", "POST", "/user/signup", app.CreateUser, nil, false},
		Route{"User Login Page", "GET", "/user/login", app.LoginUser, nil, false},
		Route{"Login User", "POST", "/user/login", app.VerifyUser, nil, false},
		Route{"Logout User", "POST", "/user/logout", app.LogoutUser, []string{"user"}, false},
		Route{"Review Users", "GET", "/user/review", app.ReviewUsers, []string{"user"}, false},
		Route{"Activate Users", "POST", "/user/active", app.ActivateUser, []string{"user"}, false},
		Route{"View User", "GET", "/user/{userID}", app.ViewUser, []string{"user"}, false},

		// API Handlers
		// Device API Handlers
		Route{"APIDevices", "GET", "/api/v1/devices", app.apiViewDevices, nil, true},
		Route{"APIDevicesAdd", "PUT", "/api/v1/devices", app.apiAddDevices, nil, true},
		Route{"APIDevicesDelete", "DELETE", "/api/v1/devices", app.apiDeleteDevices, nil, true},
		Route{"APIDevicesEdit", "POST", "/api/v1/devices", app.apiEditDevices, nil, true},
		Route{"APIGetDeviceByName", "GET", "/api/v1/device", app.apiViewDevice, nil, true},
		// Subnet API Handlers
		Route{"APISubnets", "GET", "/api/v1/subnets", app.apiViewSubnets, nil, true},
		Route{"APISubnetsAdd", "PUT", "/api/v1/subnets", app.apiAddSubnets, nil, true},
		Route{"APISubnetsDelete", "DELETE", "/api/v1/subnets", app.apiDeleteSubnet, nil, true},
		Route{"APISubnetsEdit", "POST", "/api/v1/subnets", app.apiEditSubnet, nil, true},
		// Report API Handlers
		Route{"APIReports", "GET", "/api/v1/reports", app.apiViewReports, nil, true},
		Route{"APIReportAdd", "PUT", "/api/v1/reports", app.apiAddReport, nil, true},
		/*
			// IP API Handlers
			Route{"APIIPs", "GET", "/api/v1/ips", apiViewIPs, nil, true},
			Route{"APIIPsDelete", "DELETE", "/api/v1/ips", apiDeleteIPs, nil, true},
			Route{"APIIPsAdd", "PUT", "/api/v1/ips", apiAddIPs, nil, true},
			Route{"APIIPsEdit", "POST", "/api/v1/ips", apiEditIPs, nil, true},
		*/
	}

	return routes
}
