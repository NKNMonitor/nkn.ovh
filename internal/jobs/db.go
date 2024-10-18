package jobs

import (
	"database/sql"
	"time"
)

const (
	InitStep               = "init"
	IstallFirstNodeStep    = "execute-first-node-step"
	CopyFirstNodeFilesStep = "copy-first-node-files-step"
	StepFinished           = "finish"
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
	Step      string    `json:"step"`
	CreatedAt time.Time `json:"created_at"`
}

func (w *WaitNode) List(db *sql.DB) (nodes []WaitNode, err error) {

	q := "SELECT id, name, ip, ssh_key, user, password, done, wait, use_proxy,created_at, step FROM wait_nodes WHERE done = 0"

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var createdAt []byte
		var node WaitNode
		if err := rows.Scan(&node.Id, &node.Name, &node.Ip, &node.SSHKey, &node.User, &node.Password, &node.Done, &node.Wait, &node.UseProxy, &createdAt, &node.Step); err != nil {
			return nodes, err
		}
		node.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))

		nodes = append(nodes, node)
	}

	return

}

func (w *WaitNode) UpdateStep(db *sql.DB, step string, id int) (err error) {

	q := "UPDATE wait_nodes SET step= ? WHERE id=?"

	_, err = db.Exec(q, step, id)
	return
}

func (w *WaitNode) Finish(db *sql.DB, id int) (err error) {

	q := "UPDATE wait_nodes SET done= 1 WHERE id=?"

	_, err = db.Exec(q, id)
	return
}
