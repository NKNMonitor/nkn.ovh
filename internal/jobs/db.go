package jobs

import (
	"database/sql"
	"time"
)

type WaitNode struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Ip        string    `json:"ip"`
	SSHKey    string    `json:"ssh_key"`
	User      string    `json:"user"`
	Password  string    `json:"password"`
	Done      int       `json:"done"`
	Wait      int64     `json:"wait"`
	UseProxy  int       `json:"use_proxy"`
	CreatedAt time.Time `json:"created_at"`
}

func (w *WaitNode) List(db *sql.DB) (nodes []WaitNode, err error) {

	q := "SELECT id, name, ip, ssh_key, user, password, done, wait, use_proxy,created_at FROM wait_nodes WHERE done = 0"

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var node WaitNode
		if err := rows.Scan(&node.Id, &node.Name, &node.Ip, &node.SSHKey, &node.User, &node.Password, &node.Done, &node.Wait, &node.UseProxy, &node.CreatedAt); err != nil {
			return nodes, err
		}
		nodes = append(nodes, node)
	}

	return

}
