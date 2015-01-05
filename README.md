Go
==

This is a server, written in [Go](https://golang.org/), for playing the game [Go](http://en.wikipedia.org/wiki/Go_%28game%29).

I built this so I could program AIs for playing the game. The eventual goal is to stand this up publically and let people play against AIs.

Fair warning: I don't actually know how to play the game nor do I know the language. Pull requests more than welcome.

Eventual Features:
==
1. Two humans, same computer
2. Two humans, different computers
3. A human against his or her own AI
4. A human against other AIs, without having access to their source
5. AIs against each other, neither having access to the other's source
6. All interested parties should be able to access all game logs


To Run:
==
```
cd source
go run http.go
```
You should be able to open up a browser pointed at localhost:8000 to play


Or, if you use [gin](https://github.com/codegangsta/gin):

```
gin --port 4000 --appPort 8000 -i go run http.go
```

In which case you should hit localhost:4000
