package io_bound

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/profile"
	"strconv"
	"testing"
	"time"
	"vu/benchmark/queue/helpers"
	"vu/benchmark/queue/internal"
	"vu/benchmark/queue/tasks"
)

func BenchmarkTestIOBound(b *testing.B) {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	poolSize := 128
	capacity := 1_000

	ip, _ := helpers.GetOutboundIP()

	payload, err := json.Marshal(tasks.SlowAPITaskInput{Addr: fmt.Sprintf("%s:%d", ip, 8082)})
	if err != nil {
		b.Fatalf("marshal hash input: %v", err)
	}

	queue := internal.NewQueue(capacity, poolSize, true)
	pending := make(chan (<-chan internal.Output), capacity)
	for i := 0; i < 5; i++ {
		go func() {
			for ch := range pending {
				out := <-ch

				if out.Err != nil {
					b.Logf("task failed: %v", out.Err)
				} else {
					//b.Logf("task ok: %s", out.Res)
				}
			}
		}()
	}

	i := 0
	b.ResetTimer()
	for i < 100 {
		task := tasks.Task{
			Id:    strconv.Itoa(i),
			Type:  tasks.SlowAPITaskType,
			Input: payload,
		}

		ch, err := queue.Put(&task)
		if err != nil {
			//b.Logf("Error when put task %d: %v", i, err)
			time.Sleep(2 * time.Second)
			continue
		}

		//b.Logf("put task %d", i)
		pending <- ch
		i++
	}

	queue.Shutdown()
	close(pending)
	b.StopTimer()
	time.Sleep(1 * time.Second)
}

func TestIOBound(t *testing.T) {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	poolSize := 8
	capacity := 1_000

	ip, _ := helpers.GetOutboundIP()

	payload, err := json.Marshal(tasks.SlowAPITaskInput{Addr: fmt.Sprintf("%s:%d", ip, 8082)})
	if err != nil {
		t.Fatalf("marshal hash input: %v", err)
	}

	queue := internal.NewQueue(capacity, poolSize, true)
	pending := make(chan (<-chan internal.Output), capacity)
	for i := 0; i < 5; i++ {
		go func() {
			for ch := range pending {
				out := <-ch

				if out.Err != nil {
					t.Logf("task failed: %v", out.Err)
				} else {
					//t.Logf("task ok: %s", out.Res)
				}
			}
		}()
	}

	i := 0
	for i < 100 {
		task := tasks.Task{
			Id:    strconv.Itoa(i),
			Type:  tasks.SlowAPITaskType,
			Input: payload,
		}

		ch, err := queue.Put(&task)
		if err != nil {
			//t.Logf("Error when put task %d: %v", i, err)
			time.Sleep(2 * time.Second)
			continue
		}

		//t.Logf("put task %d", i)
		pending <- ch
		i++
	}

	queue.Shutdown()
	close(pending)
	time.Sleep(1 * time.Second)
}
