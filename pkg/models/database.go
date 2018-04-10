package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Database struct {
	*mgo.Database
}

func (db Database) GetAllDevices() ([]Device, error) {
	var devices []Device
	err := db.C("devices").Find(bson.M{}).All(&devices)
	return devices, err
}

func (db Database) GetDevice(device *Device) (*Device, error) {
	err := db.C("devices").FindId(device.ID).One(device)
	return device, err
}

func (db Database) GetDeviceByName(device *Device) (*Device, error) {
	err := db.C("devices").Find(bson.M{"device_name": device.Name}).One(device)
	return device, err
}

func (db Database) AddDevice(device *Device) error {
	err := db.C("devices").Insert(device)
	return err
}

func (db Database) DeleteDevice(device *Device) error {
	err := db.C("devices").Remove(device) // By id?
	return err
}

func (db Database) EditDevice(device *Device) error {
	err := db.C("devices").UpdateId(device.ID, device)
	return err
}

func (db Database) GetDevicesForSubnet(subnet *Subnet) ([]Device, error) {
	var devices []Device
	err := db.C("devices").Find(bson.M{"interfaces.ips.subnet_id": subnet.ID}).
		Select(bson.M{"_id": 1, "device_name": 1, "interfaces.ips.ip": 1,
			"interfaces.ips.subnet_id": 1}).All(&devices)
	return devices, err
}

func (db Database) GetAllSubnets() ([]Subnet, error) {
	var subnets []Subnet
	err := db.C("subnets").Find(bson.M{}).All(&subnets)
	return subnets, err
}

func (db Database) GetSubnet(subnet *Subnet) (*Subnet, error) {
	err := db.C("subnets").FindId(subnet.ID).One(subnet)
	return subnet, err
}

func (db Database) GetSubnetByName(subnet *Subnet) (*Subnet, error) {
	err := db.C("subnets").Find(bson.M{"subnet_name": subnet.Name}).One(subnet)
	return subnet, err
}

// TODO
func (db Database) AddSubnet(subnet *Subnet) error {
	err := db.C("subnets").Insert(subnet)
	return err
}

// TODO
func (db Database) DeleteSubnet(subnet *Subnet) error {
	err := db.C("subnets").Remove(subnet) // By id?
	return err
}

// TODO
func (db Database) EditSubnet(subnet *Subnet) error {
	err := db.C("subnets").UpdateId(subnet.ID, subnet)
	return err
}

func (db Database) GetReport(report *Report) (*Report, error) {
	err := db.C("reports").FindId(report.ID).One(report)
	return report, err
}

func (db Database) AddReport(report *Report) error {
	err := db.C("reports").Insert(report)
	return err
}

func (db Database) GetReportsForDevice(device *Device) ([]Report, error) {
	var reports []Report
	err := db.C("reports").Find(bson.M{"device_id": device.ID}).All(&reports)
	return reports, err
}

func OpenMongo(dsn string) (*mgo.Session, *mgo.Database, error) {
	var db *mgo.Database
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return nil, nil, err
	}

	db = session.DB(dsn)
	return session, db, nil
}
