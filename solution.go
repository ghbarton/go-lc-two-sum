package solution

func main() {
}

func solution(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, v := range nums {
		neededVal := target - v
		if _, ok := hashMap[neededVal]; ok {
			return []int{hashMap[neededVal], i}
		}
		hashMap[v] = i
	}
	return nil
}
