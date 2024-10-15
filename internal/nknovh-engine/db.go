package nknovh_engine

import (
	"database/sql"
	"fmt"
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

func getAllNodes(o *NKNOVH) (*sql.Rows, *sql.Rows) {
	rows, err := o.sql.db["main"].Query("SELECT * FROM nodes")
	if err != nil {
		panic(err.Error())
	}

	rows_wait_nodes, err := o.sql.db["main"].Query("SELECT * FROM wait_nodes")
	if err != nil {
		panic(err.Error())
	}

	return rows, rows_wait_nodes
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
	query := `INSERT INTO wait_nodes (name, ip, ssh_key, user, password, done, wait, use_proxy)
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(query, node.Name, node.Ip, node.SSHKey, node.User, node.Password, false, node.WaitTime, node.UseProxy)
	if err != nil {
		return 0, fmt.Errorf("error inserting node: %v", err)
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting inserted ID: %v", err)
	}

	return insertedID, nil
}
