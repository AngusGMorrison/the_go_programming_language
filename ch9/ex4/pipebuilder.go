// Construct a pipeline that connects an arbitrary number of goroutines with channels. What is the
// maximum number of pipeline stages you can create without running out of memory? How long does
// it take a value to transit the entire pipeline?
package pipebuilder

func build(receive, end chan struct{}, count int) {
	if count > 0 {
		count--
		send := make(chan struct{})
		go build(send, end, count)
		send <- <-receive
	} else {
		end <- struct{}{} // send ready signal to original caller
		end <- <-receive
	}
}
