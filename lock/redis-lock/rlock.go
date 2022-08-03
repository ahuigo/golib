package main
import (
  "fmt"
  "time"
  "context"
  "log"

  "github.com/bsm/redislock"
  "github.com/go-redis/redis/v8"
)

func main() {
	// Connect to redis.
	client := redis.NewClient(&redis.Options{
		Network:	"tcp",
		Addr:		"redis:6379",
	})
	defer client.Close()

	// Create a new lock client.
	locker := redislock.New(client)

	ctx := context.Background()

	// Try to obtain lock.
	lock, err := locker.Obtain(ctx, "my-key", 1000*time.Millisecond, nil)
	if err == redislock.ErrNotObtained {
		fmt.Println("Could not obtain lock!")
        return
	} else if err != nil {
        fmt.Println("obtain lock err:", err)
		log.Fatalln(err)
        return
	}

	// Don't forget to defer Release.
	defer lock.Release(ctx)
	fmt.Println("I have a lock!")

	// Sleep and check the remaining TTL.
	time.Sleep(50 * time.Millisecond)
	if ttl, err := lock.TTL(ctx); err != nil {
		log.Fatalln(err)
	} else if ttl > 0 {
		fmt.Println("Yay, I still have my lock!")
	}

	// Extend my lock.
	if err := lock.Refresh(ctx, 100*time.Millisecond, nil); err != nil {
		log.Fatalln(err)
	}

	// Sleep a little longer, then check.
	time.Sleep(100 * time.Millisecond)
	if ttl, err := lock.TTL(ctx); err != nil {
		log.Fatalln(err)
	} else if ttl == 0 {
		fmt.Println("Now, my lock has expired!")
	} else {
		fmt.Printf("Now, my lock ttl=%v!\n", ttl)
	}

}
