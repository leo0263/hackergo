// https://www.hackerrank.com/challenges/the-love-letter-mystery
// by.0263

package main
import "fmt"


func main() {
    var jumSoal, changes, diff int
	var inputan string
    fmt.Scan(&jumSoal)

    for soal := 1; soal <= jumSoal; soal++ {
		fmt.Scan(&inputan)
		
		changes = 0
		for i := 0; i < len(inputan)/2; i++ {
			diff = int(inputan[i]) - int(inputan[len(inputan)-i-1])
			if (diff < 0) {
				changes -= diff
			} else {
				changes += diff
			}
		}

		fmt.Println(changes)
	}
	
}