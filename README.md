# cclock
Time measure divisible by 100, not 60

## Installation

```
$ go get github.com/gmelodie/cclock
```

## Usage

Convert 1h 30min and 15s from weird (also called "normal") time to ctime
```
$ cclock 1 30 15
Normal Time:	234h	03min	03s
CTime:		234ch	05ct	09cs
```

Convert current time to ctime
```
$ cclock now
Normal Time:	11h	43min	06s
CTime:		11ch	71ct	84cs
```

Print a live clock that keeps refreshing
```
$ cclock clock
Normal Time:	11h	43min	06s
CTime:		11ch	71ct	84cs
```


## Inspiration

Why the heck is a minute 60 seconds? Why on earth is an hour 60 minutes? Wouldn't it make much more sense if a minute were 100 seconds, and an hour 100 minutes?

Introducing centhclock (cclock): a time measure that turns time counting into something that *actually* makes sense.

A centhclock is composed by three basic time measures: centhconds (`cs`), centhutes (`ct`) and centhours (`ch`), where

```
1ch = 100ct
1ct = 100cs
1ch = 1h
```

Which means that one hour should have `10^4` centhconds. Hence, all we have to do is convert the seconds in an hour to `10^4` centhconds:
```
1 hour = 60 * 60s = 3600s
```

So 1 hour, which should have `10^4` centhconds, has `3600` seconds:
```
10^4cs = 3600s
1cs = 3600/10^4s = 0.360s
1s = 10^4/3600cs = 2.778cs
```


Time measure | Centhconds | Centhutes | Centhours
---------------------|--------------------------|----------------|------------
1 Second | 2.777778     cs | 0.02777778   ct | 0.0002777778 ch
1 Minute | 166.66668    cs | 1.6666668    ct | 0.016666668  ch
1 Hour   | 10000.0008   cs | 100.000008   ct | 1.00000008   ch
1 Day    | 240000.01920 cs | 2400.0001920 ct | 24.000001920 ch


For usage purposes, we can also think about each centhclock measure in terms of seconds, minutes and hours

Time measure | Seconds | Minutes | Hours
---------------------|--------------------------|----------------|------------
1 Centhcond | 0.35999997 s | 0.00599999952 min | 0.00009999999 h
1 Centhute  | 35.999997  s | 0.599999952   min | 0.009999999   h
1 Centhour  | 3599.9997  s | 59.9999952    min | 0.9999999     h

















