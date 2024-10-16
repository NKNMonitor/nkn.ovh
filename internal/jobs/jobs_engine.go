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
	db       *sql.DB
	ctx      context.Context
	interval time.Duration
	log      *zap.Logger
}

func NewJobsEngine(dbConnectionString string, ctx context.Context, interval time.Duration, logger *zap.Logger) (*JobsEngine, error) {
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

	return &JobsEngine{db: db, ctx: ctx, interval: interval, log: logger}, nil

}

func (j *JobsEngine) Run() {
	go func() {
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
				if wait.After(now) {

				}
			}

		}
	}()

}
