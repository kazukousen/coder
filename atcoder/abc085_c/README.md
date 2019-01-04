https://atcoder.jp/contests/abc085/tasks/abc085_c

Nのなかでaがいくつ、bがいくつ・・・みたいなのは
探索のときに

```go
for a := 0; a < N; a++ {
    for b := 0; b < (N - a); b++ {
        // do something ...
    }
}
```

1秒間にできるforループは10^8回
