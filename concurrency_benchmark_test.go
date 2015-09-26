package gospeed

import "testing"

func BenchmarkConcurrentMakeChannelBoolUnbuffered(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(chan bool)
	}
}

func BenchmarkConcurrentMakeChannelBool1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(chan bool, 1)
	}
}

func BenchmarkConcurrentMakeChannelBool10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(chan bool, 10)
	}
}

func BenchmarkConcurrentMakeChannelStringUnbuffered(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(chan string)
	}
}

func BenchmarkConcurrentMakeChannelString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(chan string, 1)
	}
}

func BenchmarkConcurrentMakeChannelString10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(chan string, 10)
	}
}

func BenchmarkConcurrentGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {}()
	}
}

func BenchmarkConcurrentMakeSyncChannelAndClose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		close(make(chan bool))
	}
}

func BenchmarkConcurrentMakeASyncChannelAndClose1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		close(make(chan bool, 1))
	}
}

func BenchmarkConcurrentMakeASyncChannelAndClose10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		close(make(chan bool, 10))
	}
}

func BenchmarkConcurrentMakeASyncChannelAndClose100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		close(make(chan bool, 100))
	}
}

func BenchmarkConcurrentMakeASyncChannelAndClose1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		close(make(chan bool, 1000))
	}
}

func BenchmarkConcurrentSyncBoolChannelWrite1(b *testing.B) {
	b.StopTimer()
	c := make(chan bool)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c <- true
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncBoolChannelWrite10(b *testing.B) {
	b.StopTimer()
	c := make(chan bool)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 10; j > 0; j-- {
			c <- true
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncBoolChannelWrite100(b *testing.B) {
	b.StopTimer()
	c := make(chan bool)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 100; j > 0; j-- {
			c <- true
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncBoolChannelWrite1000(b *testing.B) {
	b.StopTimer()
	c := make(chan bool)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 1000; j > 0; j-- {
			c <- true
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncBoolChannelWrite1(b *testing.B) {
	b.StopTimer()
	c := make(chan bool, 1)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c <- true
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncBoolChannelWrite10(b *testing.B) {
	b.StopTimer()
	c := make(chan bool, 10)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 10; j > 0; j-- {
			c <- true
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncBoolChannelWrite100(b *testing.B) {
	b.StopTimer()
	c := make(chan bool, 100)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 100; j > 0; j-- {
			c <- true
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncBoolChannelWrite1000(b *testing.B) {
	b.StopTimer()
	c := make(chan bool, 1000)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 1000; j > 0; j-- {
			c <- true
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncStringChannelWrite0x1(b *testing.B) {
	b.StopTimer()
	c := make(chan string)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c <- ""
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncStringChannelWrite0x10(b *testing.B) {
	b.StopTimer()
	c := make(chan string)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 10; j > 0; j-- {
			c <- ""
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncStringChannelWrite0x100(b *testing.B) {
	b.StopTimer()
	c := make(chan string)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 100; j > 0; j-- {
			c <- ""
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncStringChannelWrite0x1000(b *testing.B) {
	b.StopTimer()
	c := make(chan string)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 1000; j > 0; j-- {
			c <- ""
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncStringChannelWrite0x1(b *testing.B) {
	b.StopTimer()
	c := make(chan string, 1)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c <- ""
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncStringChannelWrite0x10(b *testing.B) {
	b.StopTimer()
	c := make(chan string, 10)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 10; j > 0; j-- {
			c <- ""
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncStringChannelWrite0x100(b *testing.B) {
	b.StopTimer()
	c := make(chan string, 100)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 100; j > 0; j-- {
			c <- ""
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncStringChannelWrite0x1000(b *testing.B) {
	b.StopTimer()
	c := make(chan string, 1000)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 1000; j > 0; j-- {
			c <- ""
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncStringChannelWrite1x1(b *testing.B) {
	b.StopTimer()
	c := make(chan string)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c <- "h"
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncStringChannelWrite1x10(b *testing.B) {
	b.StopTimer()
	c := make(chan string)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 10; j > 0; j-- {
			c <- "h"
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncStringChannelWrite1x100(b *testing.B) {
	b.StopTimer()
	c := make(chan string)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 100; j > 0; j-- {
			c <- "h"
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncStringChannelWrite1x1000(b *testing.B) {
	b.StopTimer()
	c := make(chan string)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 1000; j > 0; j-- {
			c <- "h"
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncStringChannelWrite1x1(b *testing.B) {
	b.StopTimer()
	c := make(chan string, 1)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c <- "h"
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncStringChannelWrite1x10(b *testing.B) {
	b.StopTimer()
	c := make(chan string, 10)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 10; j > 0; j-- {
			c <- "h"
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncStringChannelWrite1x100(b *testing.B) {
	b.StopTimer()
	c := make(chan string, 100)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 100; j > 0; j-- {
			c <- "h"
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncStringChannelWrite1x1000(b *testing.B) {
	b.StopTimer()
	c := make(chan string, 1000)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 1000; j > 0; j-- {
			c <- "h"
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncStringChannelWrite10x1(b *testing.B) {
	b.StopTimer()
	c := make(chan string)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c <- "hellohello"
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncStringChannelWrite10x10(b *testing.B) {
	b.StopTimer()
	c := make(chan string)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 10; j > 0; j-- {
			c <- "hellohello"
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncStringChannelWrite10x100(b *testing.B) {
	b.StopTimer()
	c := make(chan string)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 100; j > 0; j-- {
			c <- "hellohello"
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentSyncStringChannelWrite10x1000(b *testing.B) {
	b.StopTimer()
	c := make(chan string)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 1000; j > 0; j-- {
			c <- "hellohello"
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncStringChannelWrite10x1(b *testing.B) {
	b.StopTimer()
	c := make(chan string, 1)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c <- "hellohello"
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncStringChannelWrite10x10(b *testing.B) {
	b.StopTimer()
	c := make(chan string, 10)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 10; j > 0; j-- {
			c <- "hellohello"
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncStringChannelWrite10x100(b *testing.B) {
	b.StopTimer()
	c := make(chan string, 100)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 100; j > 0; j-- {
			c <- "hellohello"
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}

func BenchmarkConcurrentASyncStringChannelWrite10x1000(b *testing.B) {
	b.StopTimer()
	c := make(chan string, 1000)
	go func() {
		for _ = range c {
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 1000; j > 0; j-- {
			c <- "hellohello"
		}
	}
	b.StopTimer()
	close(c)
	b.StartTimer()
}
