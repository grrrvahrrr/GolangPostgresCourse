package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	ctx := context.Background()

	url := "postgres://bituser:bit@localhost:5433/bitmedb?sslmode=disable"

	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(err)
	}

	cfg.MaxConns = 2
	cfg.MinConns = 1

	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	duration := time.Duration(10 * time.Second)
	threads := 100
	fmt.Println("start attack")
	res := attack(ctx, duration, threads, dbpool)

	fmt.Println("duration:", res.Duration)
	fmt.Println("MaxConns:", cfg.MaxConns)
	fmt.Println("MinConns:", cfg.MinConns)
	fmt.Println("threads:", res.Threads)
	fmt.Println("queries:", res.QueriesPerformed)
	qps := res.QueriesPerformed / uint64(res.Duration.Seconds())
	fmt.Println("QPS:", qps)
}

type AttackResults struct {
	Duration         time.Duration
	Threads          int
	QueriesPerformed uint64
}

func attack(ctx context.Context, duration time.Duration, threads int, dbpool *pgxpool.Pool) AttackResults {
	var queries uint64

	attacker := func(stopAt time.Time) {
		for {
			_, err := search(ctx, dbpool, "3", 2)
			if err != nil {
				log.Printf("an error occurred: %v", err)
			}
			atomic.AddUint64(&queries, 1)
			if time.Now().After(stopAt) {
				return
			}
		}
	}

	var wg sync.WaitGroup
	wg.Add(threads)

	startAt := time.Now()
	stopAt := startAt.Add(duration)

	for i := 0; i < threads; i++ {
		go func() {
			attacker(stopAt)
			wg.Done()
		}()
	}
	wg.Wait()

	return AttackResults{
		Duration:         time.Since(startAt),
		Threads:          threads,
		QueriesPerformed: queries,
	}
}

func search(ctx context.Context, dbpool *pgxpool.Pool, param string, limit int) (string, error) {
	const sql = `select ip from bitme.urlusedata where ip_num_of_uses = $1 limit $2;`

	rows, err := dbpool.Query(ctx, sql, param, limit)
	if err != nil {
		return "", fmt.Errorf("failed to query data: %w", err)
	}
	defer rows.Close()

	var ipbase string

	for rows.Next() {
		var ip string
		err = rows.Scan(&ip)
		if err != nil {
			return "", fmt.Errorf("failed to scan row: %w", err)
		}
		ipbase += "ip" + "; "
		if rows.Err() != nil {
			return "", fmt.Errorf("failed to read response: %w", rows.Err())
		}
	}

	return ipbase, nil

}
