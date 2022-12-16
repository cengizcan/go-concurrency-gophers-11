package steps

import (
	"context"
	"testing"
)

func TestStatS7(t *testing.T) {
	imageCnt := []int{10, 20, 25, 5}
	ctx := context.TODO()
	stat := S7(ctx, 3, imageCnt)
	t.Logf("%v", stat)
	if stat.successCnt != 60 {
		t.Errorf("Want 60 got %d", stat.successCnt)
	}

}
