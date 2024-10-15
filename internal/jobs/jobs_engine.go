package jobs

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

type JobsEngine struct {
	db       *sql.DB
	ctx      context.Context
	interval time.Duration
}

func NewJobsEngine(dbConnectionString string, ctx context.Context, interval time.Duration) (*JobsEngine, error) {
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

	return &JobsEngine{db: db, ctx: ctx, interval: interval}, nil

}
