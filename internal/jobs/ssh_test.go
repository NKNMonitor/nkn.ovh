package jobs

import "testing"

func TestCommonSSH(t *testing.T) {

	client, err := NewCommonSSHClient("193.17.183.140:22", "root", "UAvuWijNF", "")
	if err != nil {
		t.Error(err)
		return
	}
	if err := client.ExecuteCommand("mkdir hello"); err != nil {
		t.Error(err)
	}

}

func TestSSLSSH(t *testing.T) {

	client, err := NewProxySSHClient("oproxy.site:11435", "nAfypY", "SaUSEx7uk5TA", "193.17.183.140:22", "root", "UAvuWijNF", "")
	if err != nil {
		t.Error(err)
		return
	}
	if err := client.ExecuteCommand("mkdir helloproxy"); err != nil {
		t.Error(err)
	}

}

func TestCopyFile(t *testing.T) {
	client, err := NewCommonSSHClient("193.17.183.140:22", "root", "UAvuWijNF", "")
	if err != nil {
		t.Error(err)
		return
	}
	if err := client.CopyFile("./db.go", "api.md", "0655"); err != nil {
		t.Error(err)
	}
}
