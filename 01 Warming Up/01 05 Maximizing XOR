package main
import "fmt"

func main() {
    var jumSoal int
	var soal uint32
	fmt.Scan(&jumSoal);
	
    for i := 1; i <= jumSoal; i++ {
		var jawaban uint32
        
        fmt.Scan(&soal);
        jawaban = 1;
		for j := 1; j <= soal; j++ {
			if (j % 2 == 1) {
				jawaban *= 2;
			} else {
				jawaban += 1;
			}
		}
		
		fmt.Print(jawaban,"\n");
	}
}