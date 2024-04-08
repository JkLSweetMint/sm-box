package size

import (
	"testing"
)

func TestInt64(t *testing.T) {
	var (
		values = []FileSize{
			"4KB",
			"4MB",
			"4GB",
			"4TB",
			"1.5KB",
			"1.5MB",
			"1.5GB",
			"1.5TB",
		}
		result = []int64{
			4096,
			4194304,
			4294967296,
			4398046511104,
			1536,
			1572864,
			1610612736,
			1649267441664,
		}
	)

	for i, v := range values {
		var res = result[i]

		if v.Int64() != res {
			t.Fatalf("Invalid value %d after converting file size %s. ", res, v)
		}
	}
}
