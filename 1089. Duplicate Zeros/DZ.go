package main

func main()  {

	input:=[]int{1,0,2,3,0,4,5,0}
	duplicateZeros(input)
//test(input)
}

func duplicateZeros(arr []int) {

	for i := 0; i < len(arr); {
		if arr[i] == 0 {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i += 2
			continue
		}
		i++
	}

}