package main

import (
	"testing"

    "github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func generateMockDB() (*sqlx.DB, sqlmock.Sqlmock) {
	// err
	mockDB, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB,"sqlmock")
	return sqlxDB, mock
}

func generateIPRows() *sqlmock.Rows {
	columns := []string{"ip_id", "name", "ip"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(1, "ip1", "192.168.1.1")
	rows.AddRow(2, "ip2", "192.168.1.2")
	return rows
}

func generateDeviceRows() *sqlmock.Rows {
	columns := []string{"device_id", "name", "team", "owner", "location", "mac", "ip"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(0, "default", "none", "none", "none", "none", "0.0.0.0")
	rows.AddRow(1, "server", "team", "owner", "location", "AA:BB:CC:DD:EE:FF", "192.168.1.1")
	return rows
}

func TestSQLGetAllDevice(t *testing.T) {
	sqlxDB, mock := generateMockDB()
	defer sqlxDB.Close()

	rows := generateDeviceRows()

	mock.ExpectQuery("SELECT (.+) FROM devices").WillReturnRows(rows)

	devices := GetAllDevices(sqlxDB)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	if len(devices) != 2 {
		t.Error("Did not get the right number of devices from GetAllDevices")
	}
}
// not implemented yet
func TestSQLGetDevice(t *testing.T) {
	sqlxDB, mock := generateMockDB()
	defer sqlxDB.Close()

	rows := generateDeviceRows()
	mock.ExpectQuery("SELECT (.+) FROM devices WHERE device_id=?").WithArgs(1).WillReturnRows(rows)
	device := Device{}
	device.DeviceID = 1
	GetDevice(sqlxDB, device)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	rows = generateDeviceRows()
	mock.ExpectQuery("SELECT (.+) FROM devices WHERE name=?").WithArgs("server").WillReturnRows(rows)
	device.DeviceID = 0
	device.Name = "server"
	GetDevice(sqlxDB, device)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	rows = generateDeviceRows()
	device.Name = ""
	device.Owner = "owner"
	mock.ExpectQuery("SELECT (.+) FROM devices WHERE owner=?").WithArgs("owner").WillReturnRows(rows)
	GetDevice(sqlxDB, device)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestSQLGetAllSubnets(t *testing.T) {
	sqlxDB, mock := generateMockDB()
	defer sqlxDB.Close()
	columns := []string{"subnet_id", "name", "ip", "mask"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(0, "default", "0.0.0.0", "0")
	rows.AddRow(1, "team1", "192.168.1.0", "24")

	mock.ExpectQuery("SELECT (.+) FROM subnets").WillReturnRows(rows)

	subnets := GetAllSubnets(sqlxDB)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	if len(subnets) != 2 {
		t.Error("Did not get the right number of devices from GetAllSubnets")
	}
}
func TestSQLGetAllIPs(t *testing.T) {
	sqlxDB, mock := generateMockDB()
	defer sqlxDB.Close()

	rows := generateIPRows()

	mock.ExpectQuery("SELECT (.+) FROM ips").WillReturnRows(rows)

	ips := GetAllIPs(sqlxDB)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	if len(ips) != 2 {
		t.Error("Did not get the right number of devices from GetAllIPs")
	}
}
func TestSQLGetIP(t *testing.T) {
	sqlxDB, mock := generateMockDB()
	defer sqlxDB.Close()
	columns := []string{"ip_id", "name", "ip"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(1, "ip1", "192.168.1.1")
	rows.AddRow(2, "ip2", "192.168.1.2")
	mock.ExpectQuery("SELECT (.+) FROM ips").WillReturnRows(rows)

	ip := IP{}
	ip.IP = "192.168.1.2"
	ip = GetIP(sqlxDB, ip)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	if ip.Name == "ip2" && ip.IPID == 1 && ip.IP == "192.168.1.2" {
		t.Error("Did not get the right ip from GetIP")
	}
}
func TestSQLDeleteIP(t *testing.T) {
	sqlxDB, mock := generateMockDB()
	defer sqlxDB.Close()

	mock.ExpectExec("DELETE FROM ips WHERE ip_id=?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	ip := IP{}
	ip.IPID = 1
	DeleteIP(sqlxDB, ip)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestSQLAddIP(t *testing.T) {
	sqlxDB, mock := generateMockDB()
	defer sqlxDB.Close()
	columns := []string{"ip_id", "name", "ip"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(1, "ip1", "192.168.1.1")
	rows.AddRow(2, "ip2", "192.168.1.2")

	mock.ExpectExec("INSERT INTO ips \\(ip, name\\)").WithArgs("192.168.1.3", "ip3").WillReturnResult(sqlmock.NewResult(1,1))

	ip := IP{Name: "ip3", IP: "192.168.1.3"}
	AddIP(sqlxDB, ip)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
/*
func TestSQLEditIP(t *testing.T) {
	sqlxDB, mock := generateMockDB()
	defer sqlxDB.Close()
	columns := []string{"ip_id", "name", "ip"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(1, "ip1", "192.168.1.1")
	rows.AddRow(2, "ip2", "192.168.1.2")

	mock.ExpectExec("UPDATE ips").WithArgs("ip3", "192.168.1.3").WillReturnResult(sqlmock.NewResult(0, 1))

	ip := IP{IPID: 2, Name: "ip3", IP: "192.168.1.3"}
	EditIP(sqlxDB, ip)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
*/
