package jobs

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"go.uber.org/zap"
)

type JobsEngine struct {
	db            *sql.DB
	ctx           context.Context
	interval      time.Duration
	log           *zap.Logger
	proxyAddress  string
	proxyLogin    string
	proxyPassword string
	nodesPatch    string
}

const (
	CREATE_FIRST_NODE_CMD = "wget -O install.sh 'https://nknx.org/api/v1/fast-deploy/install/72ecb02d-e6ac-4e31-a107-7c045fff4f8a/linux-amd64/My-Node-1'; bash install.sh"
	CREATE_THIRD_NODE_CMD = "wget https://download.npool.io/npool.sh && sudo chmod +x npool.sh && sudo ./npool.sh 4qivfmoYZJq9zJIM"
)

func NewJobsEngine(dbConnectionString string, ctx context.Context, interval time.Duration, logger *zap.Logger, proxyAddress, proxyLogin, proxyPass, nodesPath string) (*JobsEngine, error) {
	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(1)
	db.SetConnMaxIdleTime(2 * time.Second)

	_, err = db.Exec("SELECT version()")
	if err != nil {
		switch err.(type) {
		case *net.OpError:
			unwrapped := errors.Unwrap(err)
			switch unwrapped.(type) {
			case *net.DNSError:
				return nil, fmt.Errorf("DNS ERROR. RDS HOST DOES NOT EXISTS")

			case *os.SyscallError:
				return nil, fmt.Errorf("RDS CONNECTION TIMEOUT. MAYBE ACCESS DENIED AT THE NETWORK LEVEL")
			default:
				return nil, fmt.Errorf("RDS CONNECTION OPERATION ERROR %e", err)
			}
		default:
			return nil, fmt.Errorf("RDS CONNECTION ERROR %e", err)

		}
	}

	return &JobsEngine{db: db, ctx: ctx, interval: interval, log: logger, proxyAddress: proxyAddress, proxyLogin: proxyLogin, proxyPassword: proxyPass, nodesPatch: nodesPath}, nil

}

func (j *JobsEngine) Run() {
	go func() {
		for {
			select {
			case <-j.ctx.Done():
				j.log.Info("JOB DOWN")
			default:
				wn := WaitNode{}
				list, err := wn.List(j.db)
				if err != nil {
					j.log.Error("Failed to list nodes", zap.Error(err))
				}
				for _, i := range list {
					now := time.Now().UTC()
					wait := i.CreatedAt.Add(time.Duration(i.Wait) * time.Hour)
					if now.After(wait) {
						//Create SSH Client
						var client *SSHClient
						var err error
						if i.UseProxy == 1 {
							client, err = NewProxySSHClient(j.proxyAddress, j.proxyLogin, j.proxyPassword, fmt.Sprintf("%s:22", i.Ip), i.User, i.Password, i.SSHKey)
							if err != nil {
								j.log.Error("Failed to create SSH client", zap.Error(err))
								continue
							}
						} else {
							client, err = NewCommonSSHClient(fmt.Sprintf("%s:22", i.Ip), i.User, i.Password, i.SSHKey)
							if err != nil {
								j.log.Error("Failed to create SSH client", zap.Error(err))
								continue
							}
						}
					STEP_LOOP:
						switch i.Step {
						case InitStep:
							var cmd, nextStep string

							if i.Name == "" {
								cmd = CREATE_THIRD_NODE_CMD
								nextStep = StepFinished
							} else {
								cmd = CREATE_FIRST_NODE_CMD
								nextStep = IstallFirstNodeStep
							}
							if err := client.ExecuteCommand(cmd); err != nil {
								j.log.Error("Failed to execute init step comand", zap.Error(err))
								continue
							}
							i.Step = nextStep
							if err := i.UpdateStep(j.db, nextStep, int(i.Id)); err != nil {
								j.log.Error("Failed to update step", zap.Error(err))
							}
							goto STEP_LOOP
						case IstallFirstNodeStep:
							if err := client.CopyFile(fmt.Sprintf("%s/%s/wallet.json", j.nodesPatch, i.Name), "/home/nknx/nkn-commercial/services/nkn-node/wallet.json", "0655"); err != nil {
								j.log.Error("Failed to copy wallet.json", zap.Error(err))
								continue
							}
							if err := client.CopyFile(fmt.Sprintf("%s/%s/wallet.pswd", j.nodesPatch, i.Name), "/home/nknx/nkn-commercial/services/nkn-node/wallet.pswd", "0655"); err != nil {
								j.log.Error("Failed to copy wallet.json", zap.Error(err))
								continue
							}
							i.Step = StepFinished
							if err := i.UpdateStep(j.db, StepFinished, int(i.Id)); err != nil {
								j.log.Error("Failed to update step", zap.Error(err))
							}
							goto STEP_LOOP
						case StepFinished:
							if err := i.Finish(j.db, int(i.Id)); err != nil {
								j.log.Error("Failed to update step", zap.Error(err))
							}

						}

					}
				}

			}
			time.Sleep(20 * time.Second)
		}
	}()

}
