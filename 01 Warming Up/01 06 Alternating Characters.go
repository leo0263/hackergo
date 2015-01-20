// https://www.hackerrank.com/challenges/alternating-characters
// by.0263

package main
import "fmt"

func main() {
    var jumSoal int
	var inputan string
    fmt.Scan(&jumSoal)

    for soal := 1; soal <= jumSoal; soal++ {
		fmt.Scan(&inputan);
		
		var currentChar byte;
		var deletions int;
		for i := 0; i < len(inputan); i++ {
			if (i == 0) {
                deletions = 0
				currentChar = inputan[i]
			} else {
				if (currentChar == inputan[i])	{
					deletions++
				} else {
    				currentChar = inputan[i]
				}
			}
		}

		fmt.Println(deletions)
	}
	
}