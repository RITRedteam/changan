package main

// Device struct is a translation from devices in the database to a struct
type Device struct {
	DeviceID   int    `db:"device_id" json:"device_id"`
	Name       string `db:"device_name" json:"device_name"`
	Team       string `db:"team" json:"team"`
	Owner      string `db:"owner" json:"owner"`
	Location   string `db:"location" json:"location"`
	Interfaces map[string]map[string][]string
}

type DeviceInterfaceIP struct {
	DeviceID   int    `db:"device_id" json:"device_id"`
	DeviceName string `db:"device_name" json:"device_name"`
	Team       string `db:"team" json:"team"`
	Owner      string `db:"owner" json:"owner"`
	Location   string `db:"location" json:"location"`
	IP         string `db:"ip"`
	MAC        string `db:"mac"`
	Name       string `db:"interface_name"`
}

type InterfaceIP struct {
	DeviceID int    `db:"device_id"`
	IP       string `db:"ip"`
	MAC      string `db:"mac"`
	Name     string `db:"interface_name"`
}

// Interface struct is a translation from interfaces in the database to a struct
type Interface struct {
	DeviceID int    `db:"device_id"`
	IPID     int    `db:"ip_id"`
	MAC      string `db:"mac"`
}

// Subnet struct is a translation from subnets in the database to a struct
type Subnet struct {
	SubnetID int    `db:"subnet_id" json:"subnet_id"`
	Name     string `db:"subnet_name" json:"subnet_name"`
	IP       string `db:"ip" json:"ip"`
	Mask     string `db:"mask" json:"mask"`
}

// IP struct is a translation from ips in the database to a stuct
type IP struct {
	IPID int    `db:"ip_id" json:"ip_id"`
	Name string `db:"name" json:"name"`
	IP   string `db:"ip" json:"ip"`
}

type User struct {
	UserID   int    `db:"user_id"`
	GroupID  int    `db:"group_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	APIKey   string `db:"api_key"`
}

type Group struct {
	GroupID int    `db:"group_id"`
	Name    string `db:"name"`
}

type Implant struct {
	ImplantID   int    `db:"implant_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Active      bool   `db:"dsecription"`
}

type Report struct {
	DeviceID   int    `db:"device_id"`
	Report     string `db:"report"`
	LastWriter int    `db:"last_user_id"`
}
