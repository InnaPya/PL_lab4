package main
import ("fmt"; "strconv")
type Token struct {
	data string
	recipient int
}
func grtn(i int, n int, t Token, c chan string) {
	fmt.Println("Thread", i, "is running")
    	if (i == t.recipient) {
		c<-"Thread " + strconv.Itoa(i) + " got " + t.data
		go grtn(i+1, n, t, c)
	} else if (i < n) {
		go grtn(i+1, n, t, c)
	}
}
func main() {
	c := make(chan string)
	N := 10
	dstn := 6
	t := Token{"a message", dstn}
	go grtn(1, N, t, c)
	fmt.Println(<-c)
}