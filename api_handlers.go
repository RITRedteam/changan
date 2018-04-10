package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/koalatea/changan/pkg/forms"
	"github.com/koalatea/changan/pkg/models"
	"gopkg.in/mgo.v2/bson"
)

func (app *App) apiViewDevices(w http.ResponseWriter, r *http.Request) {
	// TODO better error handling in api
	// TODO better json objects
	devices, err := app.Mongo.GetAllDevices()
	if err != nil {
		app.APIServerError(w, err)
		return
	}

	//renderTemplate(w, r, "templates/devices.html",
	data := &APIData{Devices: devices}

	app.ReturnAPI(w, r, data)
}

func (app *App) apiAddDevices(w http.ResponseWriter, r *http.Request) {
	newDevice := &forms.NewDevice{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newDevice)
	if err != nil {
		if err.Error() == "EOF" {
			app.APIServerError(w, errors.New("incorrect json"))
		} else {
			app.APIServerError(w, err)
			return
		}
	}

	if !newDevice.Valid() {
		// TODO gotta do this guy
		http.Error(w, "test", http.StatusInternalServerError)
		return
	}

	// convert form.Interfaces and form.IPs to models
	var interfaces []models.Interface
	for _, newDeviceInterface := range newDevice.Interfaces {
		var newInterface models.Interface
		newInterface.MAC = newDeviceInterface.MAC
		newInterface.Name = newDeviceInterface.Name
		var ips []models.IP
		for _, ip := range newDeviceInterface.IPs {
			id2 := bson.NewObjectId() // TODO when subnets are implemented make this auto figure out
			newIP := models.IP{
				IP:       ip.IP,
				SubnetID: id2,
			}
			ips = append(ips, newIP)
		}
		newInterface.IPs = ips
		interfaces = append(interfaces, newInterface)
	}

	id := bson.NewObjectId()
	device := &models.Device{ //should I make this a pointer?
		ID:         id,
		Name:       newDevice.Name,
		Team:       newDevice.Team,
		Owner:      newDevice.Owner,
		Location:   newDevice.Location,
		Interfaces: interfaces,
	}

	err = app.Mongo.AddDevice(device)
	if err != nil {
		app.APIServerError(w, err)
		return
	}

	data := &APIData{Result: "device added"}
	app.ReturnAPI(w, r, data)
}

func (app *App) apiDeleteDevices(w http.ResponseWriter, r *http.Request) {
	device := &models.Device{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(device)
	if err != nil {
		if err.Error() == "EOF" {
			app.APIServerError(w, errors.New("incorect json"))
		} else {
			app.APIServerError(w, err)
		}
	}

	if device.ID != bson.ObjectId("") {
		app.Mongo.DeleteDevice(device)
	} else if device.Name != "" {
		device2, err := app.Mongo.GetDeviceByName(device) // can try deleting by name too later TODO
		if err != nil {
			app.APIServerError(w, err)
			return
		}
		if device2.ID != bson.ObjectId("") { //TODO FIX POINTER THINGS
			app.Mongo.DeleteDevice(device2)
		}
		// do not know if I actually want this TODO
	} else {
		app.APIServerError(w, errors.New("json must have an device_id or device_name"))
	}
	data := &APIData{
		Result: "device deleted",
	}
	app.ReturnAPI(w, r, data)
}

func (app *App) apiEditDevices(w http.ResponseWriter, r *http.Request) {
	device := &models.Device{}
	newDevice := &models.Device{} // TODO better naming
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(device)
	if err != nil {
		if err.Error() == "EOF" {
			app.APIServerError(w, errors.New("incorect json"))
		} else {
			app.APIServerError(w, err)
		}
	}
	app.Logger.Debugf("Device sent to API Edit Device %+v", device)

	if device.ID == bson.ObjectId("") {
		newDevice, err = app.Mongo.GetDeviceByName(device)
		if err != nil {
			app.APIServerError(w, err)
		} else if newDevice.ID == bson.ObjectId("") {
			app.APIServerError(w, errors.New("unsure what happened in apiEditDevice"))
		}
		app.Logger.Debugf("Device found from name in API EditDevice: %+v", newDevice)
		device.ID = newDevice.ID
	}
	app.Logger.Debugf("Final Device object to edit: %+v", device)
	err = app.Mongo.EditDevice(device)
	if err != nil {
		app.APIServerError(w, err)
	}

	data := &APIData{
		Result: "Device edited",
	}
	app.ReturnAPI(w, r, data)
}

func (app *App) apiViewSubnets(w http.ResponseWriter, r *http.Request) {
	subnets, err := app.Mongo.GetAllSubnets()
	if err != nil {
		app.APIServerError(w, err)
		return
	}

	//renderTemplate(w, r, "templates/devices.html",
	data := &APIData{Subnets: subnets}

	app.ReturnAPI(w, r, data)
}

func (app *App) apiAddSubnets(w http.ResponseWriter, r *http.Request) {
	newSubnet := &forms.NewSubnet{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newSubnet)
	if err != nil {
		if err.Error() == "EOF" {
			app.APIServerError(w, errors.New("incorrect json"))
		} else {
			app.APIServerError(w, err)
			return
		}
	}

	if !newSubnet.Valid() {
		// TODO gotta do this guy
		http.Error(w, "test", http.StatusInternalServerError)
		return
	}

	id := bson.NewObjectId()
	subnet := &models.Subnet{
		ID:   id,
		Name: newSubnet.Name,
		IP:   newSubnet.IP,
		Mask: newSubnet.Mask,
	}

	err = app.Mongo.AddSubnet(subnet)
	if err != nil {
		app.APIServerError(w, err)
		return
	}

	data := &APIData{Result: "subnet added"}
	app.ReturnAPI(w, r, data)
}

func (app *App) apiDeleteSubnet(w http.ResponseWriter, r *http.Request) {
	subnet := &models.Subnet{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(subnet)
	if err != nil {
		if err.Error() == "EOF" {
			app.APIServerError(w, errors.New("incorect json"))
		} else {
			app.APIServerError(w, err)
		}
	}

	if subnet.ID != bson.ObjectId("") {
		err = app.Mongo.DeleteSubnet(subnet)
		if err != nil {
			app.APIServerError(w, err)
			return
		}
	} else if subnet.Name != "" {
		deleteSubnet, err := app.Mongo.GetSubnetByName(subnet) // can try deleting by name too later TODO
		if err != nil {
			app.APIServerError(w, err)
			return
		}
		if deleteSubnet.ID != bson.ObjectId("") { //TODO FIX POINTER THINGS
			app.Mongo.DeleteSubnet(deleteSubnet)
			if err != nil {
				app.APIServerError(w, err)
				return
			}
		}
		// do not know if I actually want this TODO
	} else {
		app.APIServerError(w, errors.New("json must have a subnet_id or subnet_name"))
	}
	data := &APIData{
		Result: "subnet deleted",
	}
	app.ReturnAPI(w, r, data)
}

func (app *App) apiEditSubnet(w http.ResponseWriter, r *http.Request) {
	subnet := &models.Subnet{}
	newSubnet := &models.Subnet{} // TODO better naming
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(subnet)
	if err != nil {
		if err.Error() == "EOF" {
			app.APIServerError(w, errors.New("incorect json"))
		} else {
			app.APIServerError(w, err)
		}
	}
	app.Logger.Debugf("Subnet sent to API Edit Subnet %+v", subnet)

	if subnet.ID == bson.ObjectId("") {
		newSubnet, err = app.Mongo.GetSubnetByName(subnet)
		if err != nil {
			app.APIServerError(w, err)
		} else if newSubnet.ID == bson.ObjectId("") {
			app.APIServerError(w, errors.New("unsure what happened in apiEditSubnet"))
		}
		app.Logger.Debugf("Subnet found from name in API EditSubnet: %+v", newSubnet)
		subnet.ID = newSubnet.ID
	}
	app.Logger.Debugf("Final Subnet object to edit: %+v", subnet)
	err = app.Mongo.EditSubnet(subnet)
	if err != nil {
		app.APIServerError(w, err)
	}

	data := &APIData{
		Result: "Subnet edited",
	}
	app.ReturnAPI(w, r, data)
}
