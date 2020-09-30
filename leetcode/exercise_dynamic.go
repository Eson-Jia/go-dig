package dance

func ExerciseMemoizedCutRod(price []int, len int) int {
	revenue := make([]int, len+1)
	revenue[0] = 0
	exerciseMemoizedCutRod(price, revenue, len)
	return 0
}

func exerciseMemoizedCutRod(price []int, revenue []int, len int) int {
	if revenue[len] != 0 {
		return revenue[len]
	}
	return 0
}

func ExerciseButtonUpCutRod(price []int, length int) int {
	record := make([]int, length+1)
	record[0] = 0
	for i := 0; i < length; i++ {

	}
	return 0
}
