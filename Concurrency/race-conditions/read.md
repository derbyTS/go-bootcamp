## Run `go run -race <file_name>` to check race condition

### Sample:

```sh
go run -race main.go
==================
WARNING: DATA RACE
Read at 0x00c000098048 by goroutine 13:
  main.(*counter).increment()
      /Users/user/Development/Gogogo/go-bootcamp/Concurrency/race-conditions/main.go:45 +0xb0
  main.main.func1()
      /Users/user/Development/Gogogo/go-bootcamp/Concurrency/race-conditions/main.go:70 +0xa8

Previous write at 0x00c000098048 by goroutine 14:
  main.(*counter).increment()
      /Users/user/Development/Gogogo/go-bootcamp/Concurrency/race-conditio
  main.main.func1()
      /Users/user/Development/Gogogo/go-bootcamp/Concurrency/race-conditio

Goroutine 13 (running) created at:
  main.main()
      /Users/user/Development/Gogogo/go-bootcamp/Concurrency/race-conditio

Goroutine 14 (running) created at:
  main.main()
      /Users/user/Development/Gogogo/go-bootcamp/Concurrency/race-conditio
ns/main.go:67 +0x70
==================
==================
WARNING: DATA RACE
Write at 0x00c000098048 by goroutine 7:
  main.(*counter).increment()
      /Users/user/Development/Gogogo/go-bootcamp/Concurrency/race-conditio
ns/main.go:45 +0xc4
  main.main.func1()
      /Users/user/Development/Gogogo/go-bootcamp/Concurrency/race-conditio
ns/main.go:70 +0xa8

Previous write at 0x00c000098048 by goroutine 13:
  main.(*counter).increment()
      /Users/user/Development/Gogogo/go-bootcamp/Concurrency/race-conditio
ns/main.go:45 +0xc4
  main.main.func1()
      /Users/user/Development/Gogogo/go-bootcamp/Concurrency/race-conditio
ns/main.go:70 +0xa8

Goroutine 7 (running) created at:
  main.main()
      /Users/user/Development/Gogogo/go-bootcamp/Concurrency/race-conditio
ns/main.go:67 +0x70

Goroutine 13 (running) created at:
  main.main()
      /Users/user/Development/Gogogo/go-bootcamp/Concurrency/race-conditio
ns/main.go:67 +0x70
==================
Final counter value: 7054
Found 2 data race(s)
exit status 66
```
