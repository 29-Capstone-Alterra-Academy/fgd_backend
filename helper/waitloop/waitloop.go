package waitloop

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func WaitPersistence(dialector gorm.Dialector, duration time.Duration) (*gorm.DB, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutExceeded := time.After(duration)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", duration)

		case <-ticker.C:
			db, err := gorm.Open(dialector, &gorm.Config{})
			if err == nil {
				return db, nil
			}
		}
	}
}
