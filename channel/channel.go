package channel

import "fmt"

func ChannelBasic() {
	var newChan = make(chan string)

	var sayHello = func(val string) {
		var resp = fmt.Sprintf("Hi %s", val)
		newChan <- resp
	}

	go sayHello("Foo")
	go sayHello("Baar")
	fmt.Println(<-newChan)
	fmt.Println(<-newChan)
}

func ChannelAsATypeAndParam() {
	var msg = make(chan string)
	nameUsers := []string{"John", "Wick", "Alex", "Ferguson"}
	for _, name := range nameUsers {
		go func(who string) {
			msg <- fmt.Sprintf("Hello %s", who)
		}(name)

		printData(msg)
	}
}

func printData(msg chan string) {
	fmt.Println(<-msg)
}
