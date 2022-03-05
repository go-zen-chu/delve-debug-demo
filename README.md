# go-delve-debug-sample

Go sample repo for demonstrating TDD and delve debugging.

## What you can learn from this repo

1. About VSCode Go extension and delve
1. delve tutorial (Breakpoint, Hit Count, goroutine, Local vars)
1. Go race detector

## Explanation

### Understanding problem

1. Explain about main.go and count.go
1. Run `go run main.go`
1. In count.go, open Command Palette and run `Go: generate test for file` & `Go: generate test for this function`
1. Implement count_test.go and run `go test -v ./count`
1. Open Command Palette and run `Go: Toggle Test Coverage`
1. Discuss why the test failed.

### Answer

1. Create break point on count with Expression `90 < count && count < 100` (break) and `990 < count && count < 1000` (doesn't always break)
1. Set break point in count_debug.go and run Parallel1000countDebug from count_test.go
1. Fix count.go and run test with success.

## Appendix

Try, `go run --race main.go`

```shell
==================
WARNING: DATA RACE
Read at 0x00c00013a018 by goroutine 8:
  github.com/go-zen-chu/delve-debug-demo/count.Parallel1000count.func1()
      /Users/amasuda/dev/delve-sample/count/count.go:17 +0x8d

Previous write at 0x00c00013a018 by goroutine 7:
  github.com/go-zen-chu/delve-debug-demo/count.Parallel1000count.func1()
      /Users/amasuda/dev/delve-sample/count/count.go:17 +0xa4

Goroutine 8 (running) created at:
  github.com/go-zen-chu/delve-debug-demo/count.Parallel1000count()
      /Users/amasuda/dev/delve-sample/count/count.go:15 +0x84
  main.main()
      /Users/amasuda/dev/delve-sample/main.go:10 +0x29

Goroutine 7 (finished) created at:
  github.com/go-zen-chu/delve-debug-demo/count.Parallel1000count()
      /Users/amasuda/dev/delve-sample/count/count.go:15 +0x84
  main.main()
      /Users/amasuda/dev/delve-sample/main.go:10 +0x29
==================
1000 にならんわけないやろww
-> 998
なんでや!!
Found 1 data race(s)
exit status 66
```
