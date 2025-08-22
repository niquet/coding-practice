package minimumswaps2

// MinimumSwaps loops over the unordered array
// If the integer at index $i$ is not at the correct position, swap integer
// at $i$ with integer at the correct position. Then increment swap counter
// by one.
// If integer at index $i$ is at the correct position, move pointer to the
// next bigger index.
func MinimumSwaps(arr []int32) int32 {
	var (
		cnt int32
		ptr int32
	)

	for ptr < int32(len(arr)) {
		if arr[ptr] != ptr+1 {
			tmp := arr[arr[ptr]-1]
			arr[arr[ptr]-1] = arr[ptr]
			arr[ptr] = tmp
			cnt++
			continue
		}
		ptr++
	}

	return cnt
}
