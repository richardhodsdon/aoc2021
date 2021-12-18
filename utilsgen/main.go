package utilsgen

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	out := make(map[K]V, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}
