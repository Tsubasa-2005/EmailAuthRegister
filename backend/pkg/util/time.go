package util

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// NowJST 日本に夏時間が導入されたり、USTとの時差が変わらない限りこの実装で問題ない。(そのような可能性は十分低いのでこの実装で問題ない)
func NowJST() time.Time {
	return time.Now().In(time.FixedZone("JST", 9*60*60))
}

func Timestamptz(time time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:             time,
		InfinityModifier: pgtype.Finite,
		Valid:            true,
	}
}
