package nknovh_engine

import (
	"database/sql"
	"fmt"
	"time"
)

func chekBusyIp(o *NKNOVH, ip string) (bool, error) {

	getIPNodes, err := o.sql.db["main"].Query("SELECT * FROM nodes WHERE ip = ?", ip)
	if err != nil {
		return false, err
	}
	defer getIPNodes.Close()

	if getIPNodes.Next() {
		return false, nil
	}

	getIPWaitNodes, err := o.sql.db["main"].Query("SELECT * FROM wait_nodes WHERE ip = ?", ip)
	if err != nil {
		return false, err
	}
	defer getIPWaitNodes.Close()

	if getIPWaitNodes.Next() {
		return false, nil
	}

	return true, nil
}

func getAllNodes(o *NKNOVH) (*sql.Rows, *sql.Rows, error) {
	rows, err := o.sql.db["main"].Query("SELECT * FROM nodes")
	if err != nil {
		return nil, nil, err
	}

	rows_wait_nodes, err := o.sql.db["main"].Query("SELECT * FROM wait_nodes")
	if err != nil {
		return nil, nil, err
	}

	return rows, rows_wait_nodes, nil
}

func getBusyDirectoryNames(o *NKNOVH) ([]string, error) {
	data := make([]string, 0)
	activeNodesRows, err := o.sql.db["main"].Query("SELECT name FROM nodes")
	if err != nil {
		return data, err
	}
	defer activeNodesRows.Close()
	for activeNodesRows.Next() {
		var name string
		if err := activeNodesRows.Scan(&name); err != nil {
			return data, err
		}
		data = append(data, name)
	}

	waitNodesRows, err := o.sql.db["main"].Query("SELECT name FROM wait_nodes")
	if err != nil {
		return data, err
	}
	defer waitNodesRows.Close()
	for waitNodesRows.Next() {
		var name string
		if err := waitNodesRows.Scan(&name); err != nil {
			return data, err
		}
		data = append(data, name)
	}
	return data, nil

}

type ServerCreateRequest struct {
	Name     string
	Ip       string
	SSHKey   string
	User     string
	Password string
	UseProxy bool
	WaitTime int
}

func InsertNode(db *sql.DB, node ServerCreateRequest) (int64, error) {
	query := `INSERT INTO wait_nodes (name, ip, ssh_key, user, password, done, wait, use_proxy, created_at)
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?,?)`

	result, err := db.Exec(query, node.Name, node.Ip, node.SSHKey, node.User, node.Password, false, node.WaitTime, node.UseProxy, time.Now().UTC())
	if err != nil {
		return 0, fmt.Errorf("error inserting node: %v", err)
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting inserted ID: %v", err)
	}

	return insertedID, nil
}
