package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"time"
)

const s string = "constant"

type person struct {
	name string
	age  int
}

type rect struct {
	width, height float64
}

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func TestInterfaces() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}

func values() {
	fmt.Println("Go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	var cosa, cosa1 string = "a", "b"
	fmt.Println(cosa + " - " + cosa1)

}

func constants() {
	fmt.Println(s)
	const n = 500000000

	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

func looping() {
	i := 3

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func conditionals1() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func conditionals2() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")

}

func arrays() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("Set: ", a)
	fmt.Println("Get: ", a[4])

	fmt.Println("Len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func slices() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("Set: ", s)
	fmt.Println("Get: ", s[2])

	fmt.Println("Len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl3: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
	s = getSlice()
	fmt.Println("the s ", s)
}

func getSlice() []string {
	letters := []string{"a", "b", "c", "d"}
	return letters
}

func slices2() {
	c := make([]int, 2, 2)
	c[0] = 10
	c[1] = 9
	c = append(c, 8)
	//c[2] = 8
	fmt.Println(c)
	fmt.Println(cap(c))
	//append(c, )

	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	e[1] = 'm'

	fmt.Println(d)

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:]
	fmt.Println("1) capacity of s", cap(s))

	s = s[2:]
	fmt.Println("2) capacity of s", cap(s))

	s = make([]string, 5)
	fmt.Println("3) capacity of s", cap(s))

	s = s[2:4]
	fmt.Println("4) capacity of s", cap(s))

	s = s[:cap(s)]
	fmt.Println("5) capacity of s", cap(s))

	t := make([]string, len(s), (cap(s)+1)*2)
	for i := range s {
		t[i] = s[i]
	}
	s = t
	fmt.Println("5) capacity of s", cap(s))

	o := make([]string, 3)
	o[0] = "this"
	o[1] = " is a "
	o[2] = "test"
	copy(s, o)
	fmt.Println("o -> ", o)
	fmt.Println("s -> ", s)
	TestAppend()

}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	// if necessary, reallocate more space
	if n > cap(slice) {
		// allocate double wha'ts needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func TestAppend() {
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	fmt.Println(p)
}

func maps() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}

func RangeTest() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func TestFunc() {
	fmt.Println(plus(2, 3))
	fmt.Println(plusPlus(2, 3, 4))
}

func vals() (int, int) {
	return 3, 7
}

func TestVals() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

func sum1(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func TestSum1() {
	sum1(1, 2)
	sum1(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum1(nums...)
}

func IntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestIntSeq() {
	nextInt := IntSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := IntSeq()
	fmt.Println(newInts())

}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func TestPointers() {
	i := 1
	fmt.Println("Initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)
}

// we return a memory address to a pointer
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func TestPerson() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	fmt.Println(s.age)
}

func TestMethods() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func TestF1() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ": ", i)
	}
}

func TestGoRoutine() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func unBufferedChannels() {
	messages := make(chan string)
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

func channelBuffering() {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func worker0(done chan bool) {
	fmt.Print("working ...")
	time.Sleep(time.Second * 5)
	fmt.Println("done")
	done <- true
}

func TestWorker0() {
	done := make(chan bool, 1)
	go worker0(done)

	for {
		x := <-done
		if x == true {
			fmt.Println("x is true")
			break
		}
		fmt.Println("not done yet")
	}

}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

//https://gobyexample.com/channel-directions
func TestChannelDirections() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func TestSelect() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// https: //gobyexample.com/timeouts
func TestTimeouts() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

// https://gobyexample.com/non-blocking-channel-operations
func TestChannelOperations() {

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}

// https://gobyexample.com/closing-channels
func TestClosingChannel() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

// https://gobyexample.com/range-over-channels
func RangeOverChannels() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

// https://gobyexample.com/timers
func TestTimer() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

// https://gobyexample.com/tickers
func TestTickers() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func workerForPool(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "Started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "Finished job", j)
		results <- j * 2
	}
}

// https://gobyexample.com/worker-pools
func TestWorkerPools() {

	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go workerForPool(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Printf("Worker %d starting \n", id)

	time.Sleep(time.Second)
	fmt.Printf("worker %d done \n", id)

}

func TestWaitGroup() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	fmt.Println("we start to wait ")

	wg.Wait()
}
