package m

import (
	"testing"
	"time"

	"github.com/gorhill/cronexpr"
)

func TestA(t *testing.T) {
    //expr := cronexpr.MustParse("*/3 */4 * * *")
    expr := cronexpr.MustParse("15 17 * * *")
    now := time.Now()
    nextTime := expr.Next(now)
    t.Log(nextTime)
    t.Log(nextTime.Sub(now))
}
