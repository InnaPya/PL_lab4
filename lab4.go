package main
import ("fmt"; "strconv")
type Token struct {
	data string
	recipient int
}
type Chain struct {
	i int
	ch chan Token
	next *Chain
}
func grtn(chain Chain, c chan string) {
	fmt.Println("Tread ", chain.i, " running")
	t := <-chain.ch
	if chain.i == t.recipient {
		c <-"Thread " + strconv.Itoa(chain.i) + " got " + t.data
	} else if chain.next != nil {
		chain.next.ch <- t
	} else {
		c <-"Wrong addresse"
	}
}
func main() {
	N := 10
	chain := make([]Chain, N)
	c := make(chan string)
	for i := N - 1; i >= 0; i-- {
		channel := make(chan Token)
		if i == N - 1 {
			chain[i] = Chain{i, channel, nil}
		} else {
			chain[i] = Chain{i, channel, &chain[i+1]}
		}	
		go grtn(chain[i], c)
	}
	chain[0].ch <- Token{"a message", 5}
	fmt.Println(<-c)
}