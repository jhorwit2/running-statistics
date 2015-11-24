# go-running-statistics

Running Statistics implementation in Go which is a direct port of [this implementation](http://www.johndcook.com/blog/skewness_kurtosis/).

## Install

`go get github.com/jhorwit2/running-statistic`

## Usage

```
r := runningstats.New()
r.Add(1)
r.Add(2)
r.Add(3)
mean := r.Mean()
stddev := r.Stddev()
```
