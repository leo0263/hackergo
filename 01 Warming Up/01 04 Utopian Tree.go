package main
import "fmt"

func main() {
    var jumSoal int
	var soal uint32
	fmt.Scan(&jumSoal);
	
    for i := 1; i <= jumSoal; i++ {
		var jawaban uint32
        
        fmt.Scan(&soal);
        jawaban = ^soal;
		
		fmt.Print(jawaban,"\n");
	}
}