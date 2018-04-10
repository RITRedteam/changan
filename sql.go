package main

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// BaseDB object for use in the as our db
var BaseDB *sqlx.DB

func GetAllDeviceInterfacesIP(db *sqlx.DB) []DeviceInterfaceIP {
	var interfaces []DeviceInterfaceIP
	sqlObj := sq.Select("device_id", "device_name", "team", "owner", "location", "ip", "mac",
		"interface_name").From("devices").Join("interfaces USING (device_id)").
		Join("ips USING (ip_id)").OrderBy("device_id")
	sql, args, err := sqlObj.ToSql()
	if err != nil {
		panic(err)
	}
	err = db.Select(&interfaces, sql, args...)
	if err != nil {
		panic(err)
	}
	return interfaces
}

// GetAllDevices returns a splice of all Device structs from the database
func GetAllDevices(db *sqlx.DB) []Device {
	var devices []Device
	err := db.Select(&devices, "SELECT * FROM devices")
	if err != nil {
		panic(err)
	}
	return devices
}

// GetDevice returns a Device struct from any non empty/default parts of a Device struct
//TODO
func GetDevice(db *sqlx.DB, device Device) Device {
	var returnDevice Device
	sqlObj := sq.Select("*").From("devices")
	returnDevice = device
	if device.DeviceID != 0 {
		db.Get(&returnDevice, "SELECT * FROM devices WHERE device_id=?", device.DeviceID)
		return returnDevice
	}
	if device.Name != "" {
		db.Get(&returnDevice, "SELECT * FROM devices WHERE name=?", device.Name)
		return returnDevice
	}
	if device.Team != "" {
		sqlObj = sqlObj.Where(sq.Eq{"team": device.Team})
	}
	if device.Owner != "" {
		sqlObj = sqlObj.Where(sq.Eq{"owner": device.Owner})
	}
	if device.Location != "" {
		sqlObj = sqlObj.Where(sq.Eq{"location": device.Location})
	}
	sql, args, err := sqlObj.ToSql()
	if err != nil {
		panic(err)
	}
	err = db.Get(&returnDevice, sql, args...)
	if err != nil {
		panic(err)
	}
	return returnDevice
}

// AddDevice adds a Device to the database from a Device struct
func AddDevice(db *sqlx.DB, device Device) {
	_, err := db.Exec("INSERT INTO devices (device_id, name, team, owner, location, mac, ip)"+
		"VALUES (?, ?, ?, ?, ?)", device.DeviceID, device.Name, device.Team, device.Owner,
		device.Location)
	if err != nil {
		panic(err)
	}
}

// DeleteDevice removes a Device from the database using the Devices ID
func DeleteDevice(db *sqlx.DB, device Device) {
	_, err := db.Exec("DELETE FROM devices WHERE device_id=?", device.DeviceID)
	if err != nil {
		panic(err)
	}
}

// EditDevice edits a Device in the database from a device struct using the device structs DeviceID
func EditDevice(db *sqlx.DB, device Device) {
	sqlObj := sq.Update("devices")
	if device.Name != "" {
		sqlObj = sqlObj.Set("name", device.Name)
	}
	if device.Team != "" {
		sqlObj = sqlObj.Set("team", device.Team)
	}
	if device.Owner != "" {
		sqlObj = sqlObj.Set("owner", device.Owner)
	}
	if device.Location != "" {
		sqlObj = sqlObj.Set("location", device.Location)
	}
	sql, args, err := sqlObj.Where(sq.Eq{"device_id": device.DeviceID}).ToSql()
	_, err = db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
}

// GetAllSubnets returns a splice of Subnet structs
func GetAllSubnets(db *sqlx.DB) []Subnet {
	var subnets []Subnet
	err := db.Select(&subnets, "SELECT * FROM subnets")
	if err != nil {
		panic(err)
	}
	return subnets
}

// GetSubnet returns a Subnet struct from any non empty/default parts of a Subnet struct
func GetSubnet(db *sqlx.DB, subnet Subnet) Subnet {
	var returnSubnet Subnet
	sqlObj := sq.Select("*").From("subnets")
	if subnet.SubnetID != 0 {
		sqlObj = sqlObj.Where(sq.Eq{"subnet_id": subnet.SubnetID})
	}
	if subnet.IP != "" {
		sqlObj = sqlObj.Where(sq.Eq{"ip": subnet.IP})
	}
	if subnet.Mask != "" {
		sqlObj = sqlObj.Where(sq.Eq{"mask": subnet.Mask})
	}
	if subnet.Name != "" {
		sqlObj = sqlObj.Where(sq.Eq{"name": subnet.Name})
	}
	sql, args, err := sqlObj.ToSql()
	err = db.Get(&returnSubnet, sql, args...)
	if err != nil {
		panic(err)
	}

	return returnSubnet
}

// AddSubnet adds a Subnet to the database from a subnet Struct
func AddSubnet(db *sqlx.DB, subnet Subnet) {
	_, err := db.Exec("INSERT INTO subnets (subnet_id, name, ip, mask) VALUES (?, ?, ?, ?)",
		subnet.SubnetID, subnet.Name, subnet.IP, subnet.Mask)
	if err != nil {
		panic(err)
	}
}

// DeleteSubnet removes a Subnet from the database using the subnets SubnetID
func DeleteSubnet(db *sqlx.DB, subnet Subnet) {
	_, err := db.Exec("DELETE FROM subnets WHERE subnet_id=?", subnet.SubnetID)
	if err != nil {
		panic(err)
	}
}

// EditSubnet edits a Subnet in the database from a Subnet struct using SubnetID as identifier
func EditSubnet(db *sqlx.DB, subnet Subnet) {
	sqlObj := sq.Update("subnets")
	if subnet.Name != "" {
		sqlObj = sqlObj.Set("name", subnet.Name)
	}
	if subnet.IP != "" {
		sqlObj = sqlObj.Set("ip", subnet.IP)
	}
	if subnet.Mask != "" {
		sqlObj = sqlObj.Set("mask", subnet.Mask)
	}
	sql, args, err := sqlObj.Where(sq.Eq{"subnet_id": subnet.SubnetID}).ToSql()
	_, err = db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
}

// GetAllIPs returns a splice of IP structs
func GetAllIPs(db *sqlx.DB) []IP {
	//TODO make this work with limits
	sqlObj := sq.Select("*").From("ips")
	var ips []IP
	sql, args, err := sqlObj.ToSql()
	err = db.Select(&ips, sql, args...) //THIS ACTUALLY WORKS
	if err != nil {
		panic(err)
	}
	return ips
}

// GetIP returns a single IP struct from the IP portions that are not default
func GetIP(db *sqlx.DB, ip IP) IP {
	var returnIP IP
	sqlObj := sq.Select("*").From("ips")
	if ip.IPID != 0 {
		sqlObj = sqlObj.Where(sq.Eq{"ip_id": ip.IPID}) // probably want ipid to happen alone as the id
		// is end all
	}
	if ip.IP != "" {
		sqlObj = sqlObj.Where(sq.Eq{"ip": ip.IP})
	}
	if ip.Name != "" {
		sqlObj = sqlObj.Where(sq.Eq{"name": ip.Name})
	}
	sql, args, err := sqlObj.ToSql()
	err = db.Get(&returnIP, sql, args...)
	if err != nil {
		panic(err)
	}
	return returnIP
}

// DeleteIP deletes an IP from the database using the ip_id
func DeleteIP(db *sqlx.DB, ip IP) {
	_, err := db.Exec("DELETE FROM ips WHERE ip_id=?", ip.IPID)
	if err != nil {
		panic(err)
	}
}

// AddIP adds an ip to the database
func AddIP(db *sqlx.DB, ip IP) {
	//TODO figure out returns also figure out subnets
	_, err := db.Exec("INSERT INTO ips (ip, name) VALUES (?, ?)", ip.IP, ip.Name)
	if err != nil {
		panic(err)
	}
}

// EditIP updates an ip in the database based off of the IPID
func EditIP(db *sqlx.DB, ip IP) {
	sqlObj := sq.Update("ips")
	if ip.Name != "" {
		sqlObj = sqlObj.Set("name", ip.Name)
	}
	if ip.IP != "" {
		sqlObj = sqlObj.Set("ip", ip.IP)
	}
	sql, args, err := sqlObj.Where(sq.Eq{"ip_id": ip.IPID}).ToSql()
	_, err = db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
}

func openDB(dsn string) *sql.DB {
	var db *sql.DB
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
