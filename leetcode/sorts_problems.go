package leetcode

//func FindAllRecipesAlgorithmKana_2115(recipes []string, ingredients [][]string, supplies []string) []string {
//	countIngredients := make(map[string]int, 0)
//	recipesIngredients := make(map[string][]string, 0)
//
//	result := make([]string, 0)
//
//	for recipe_index, ingredient := range ingredients {
//		countIngredients[recipes[recipe_index]] = len(ingredient)
//		recipesIngredients[recipes[recipe_index]] = ingredient
//	}
//
//	for _, ingridient := range supplies {
//		//TODO: если я буду добавлять элементы в supplies - сломается ли цикл при переаллокации
//	}
//
//}

// Counting sort https://leetcode.com/problems/sort-an-array/solutions/5532275/explanations-no-one-will-give-you-3-detailed-approaches-extremely-simple-and-effective
func SortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	count := make([]int, max-min+1)
	for _, num := range nums {
		count[num-min]++
	}

	index := 0
	for i := 0; i < len(count); i++ {
		for count[i] > 0 {
			nums[index] = i + min
			index++
			count[i]--
		}
	}
	return nums
}
