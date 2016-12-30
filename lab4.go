package main
import ("fmt"; "strconv")
type Token struct {
	data string
	recipient int
}
func grtn(i int, t Token, c chan string) {
    	if (i != t.recipient) {
		go grtn(i+1, t, c)
	} else {
		c<-strconv.Itoa(i) + " " + t.data
	}
}
func main() {
	c := make(chan string)
	N := 10
	t := Token{"message", N}
	go grtn(1, t, c)
	fmt.Println(<-c)
}