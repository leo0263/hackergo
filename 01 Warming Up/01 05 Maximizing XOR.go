package main
import "fmt"

func main() {
    var start, stop int
	fmt.Scan(&start, &stop)
	
	var maxXOR int
	maxXOR = 0;
    for i := start; i <= stop; i++ {
		for j := i; j <= stop; j++ {
			var localXOR = i xor j
			if (localXOR > maxXOR) {
				maxXOR = localXOR;
			}
		}
	}
	
	fmt.Print(maxXOR);
}