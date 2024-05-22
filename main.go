package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/navikt/nada-pg-test/pkg/database"
)

func main() {
	ctx := context.Background()
	connString := getConnString()

	numInserts, err := strconv.Atoi(os.Getenv("NUM_INSERTS"))
	if numInserts == 0 || err != nil {
		numInserts = 100
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	repo, err := database.New(connString)
	if err != nil {
		panic(err)
	}

	data := map[string]any{"test": map[string]any{"key": "data"}}

	for {
		log.Info(fmt.Sprintf("Running %v inserts", numInserts))
		times, err := runInserts(ctx, repo, data, numInserts)
		if err != nil {
			log.Error(err.Error())
		}

		min, max, avg := minMaxAvg(times)

		log.Info(fmt.Sprintf("Minimum time: %v", min))
		log.Info(fmt.Sprintf("Maximum time: %v", max))
		log.Info(fmt.Sprintf("Average time (ms): %v", avg/1000000))

		time.Sleep(5 * time.Second)
	}
}

func runInserts(ctx context.Context, repo *database.Repo, data map[string]any, numInserts int) ([]time.Duration, error) {
	times := []time.Duration{}
	for i := 0; i < numInserts; i++ {
		elapsed, err := repo.InsertData(ctx, data)
		if err != nil {
			return times, err
		}
		times = append(times, elapsed)
		time.Sleep(20 * time.Millisecond)
	}
	return times, nil
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
	if os.Getenv("DB_URL") == "" {
		return "postgres://postgres:postgres@localhost:5432/nada?sslmode=disable"
	}

	return os.Getenv("DB_URL") + "?sslmode=disable"
}
