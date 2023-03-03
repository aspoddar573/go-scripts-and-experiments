package understandingContext

import (
	"context"
	"fmt"
)

func ValuesInContext() {
	mp := make(map[string]int)
	mp["hey"] = 1
	mp["there"] = 2

	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	ctx = context.WithValue(ctx, "key1", "value1")
	ctx = context.WithValue(ctx, "key2", mp)

	fmt.Println(ctx.Value("key"))
	fmt.Println(ctx.Value("key1"))
	fmt.Println(ctx.Value("key2"))
}
