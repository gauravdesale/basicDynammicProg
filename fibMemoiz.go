package main
import ("fmt")

func checkTable(a int) int {
	var lookUp [100]int
	var indexLookUp [100]int
	for i := 0; i < 100; i++ {
		if indexLookUp[i] == a {
			return lookUp[i]
		} 
		
	}
	return 0
}

//continue to add to this two more functions