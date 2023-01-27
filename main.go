package main

import (
	"context"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/navikt/nada-pg-test/pkg/database"
	"github.com/sirupsen/logrus"
)

const (
	numInserts = 100
)

func main() {
	ctx := context.Background()
	connString := getConnString()

	log := logrus.New()
	repo, err := database.New(connString, log.WithField("subsystem", "database"))
	if err != nil {
		panic(err)
	}

	data := map[string]any{"test": map[string]any{"key": "data"}}

	log.Infof("Running %v db inserts", numInserts)
	times := []time.Duration{}
	for i := 0; i < numInserts; i++ {
		elapsed, err := repo.InsertData(ctx, data)
		if err != nil {
			panic(err)
		}
		times = append(times, elapsed)
		time.Sleep(20 * time.Millisecond)
	}

	min, max, avg := minMaxAvg(times)

	log.Info("Minimum time: ", min)
	log.Info("Maximum time: ", max)
	log.Info("Average time (ms): ", avg/1000000)

	for {
		time.Sleep(5 * time.Second)
	}
}

func minMaxAvg(times []time.Duration) (time.Duration, time.Duration, float64) {
	min := times[0]
	max := times[0]
	sum := 0
	for _, number := range times {
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
		sum += int(number)
	}
	return min, max, float64(sum) / float64(len(times))
}

func getConnString() string {
	if os.Getenv("NAIS_DATABASE_NADA_DB_TESTA_NADA_DB_TEST_URL") == "" {
		return "postgres://postgres:postgres@localhost:5432/nada?sslmode=disable"
	}

	return os.Getenv("NAIS_DATABASE_NADA_DB_TESTA_NADA_DB_TEST_URL") + "?sslmode=disable"
}
