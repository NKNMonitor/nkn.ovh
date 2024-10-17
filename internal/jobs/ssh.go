package jobs

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/proxy"
)

type SSHClient struct {
	sshClient *ssh.Client
}

func (c *SSHClient) Close() {
	c.sshClient.Close()
}

func (c *SSHClient) ExecuteCommand(command string) error {
	session, err := c.sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var stdoutBuf, stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf

	if err := session.Run(command); err != nil {
		return fmt.Errorf("failed to run command '%s': %v, stderr: %s", command, err, stderrBuf.String())
	}
	return nil
}

func NewProxySSHClient(proxyAddress, proxyLogin, proxyPassword string, sshAddress, sshLogin, sshPassword, keySsh string) (*SSHClient, error) {
	dialer, err := proxy.SOCKS5("tcp", proxyAddress, &proxy.Auth{User: proxyLogin, Password: proxyPassword}, proxy.Direct)
	if err != nil {
		return nil, fmt.Errorf("Create Proxy Error: %s", err)
	}
	var sshConfig *ssh.ClientConfig
	if keySsh == "" {
		sshConfig = &ssh.ClientConfig{
			User: sshLogin,
			Auth: []ssh.AuthMethod{
				ssh.Password(sshPassword),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	} else {
		byteSshKey := stringToBytes(keySsh)
		signer, err := ssh.ParsePrivateKey(byteSshKey)
		if err != nil {
			return nil, err
		}
		sshConfig = &ssh.ClientConfig{
			User: sshLogin,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	}
	conn, err := dialer.Dial("tcp", sshAddress)
	if err != nil {
		return nil, fmt.Errorf("TCP Dialler Error: %s", err)
	}

	c, chans, reqs, err := ssh.NewClientConn(conn, sshAddress, sshConfig)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("ServerConnection Error: %s", err)
	}
	client := ssh.NewClient(c, chans, reqs)

	return &SSHClient{
		sshClient: client,
	}, nil

}

func NewCommonSSHClient(sshAddress, sshLogin, sshPassword, keySsh string) (*SSHClient, error) {
	var sshConfig *ssh.ClientConfig

	if keySsh == "" {
		sshConfig = &ssh.ClientConfig{
			User: sshLogin,
			Auth: []ssh.AuthMethod{
				ssh.Password(sshPassword),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	} else {
		byteSshKey := []byte(keySsh)
		signer, err := ssh.ParsePrivateKey(byteSshKey)
		if err != nil {
			return nil, fmt.Errorf("Ошибка при парсинге SSH ключа: %s", err)
		}
		sshConfig = &ssh.ClientConfig{
			User: sshLogin,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	}

	conn, err := ssh.Dial("tcp", sshAddress, sshConfig)
	if err != nil {
		return nil, fmt.Errorf("Ошибка TCP соединения: %s", err)
	}

	return &SSHClient{
		sshClient: conn,
	}, nil
}

func stringToBytes(str string) []byte {
	return []byte(str)
}
