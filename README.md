## go-irc
Library to interact with IRC.

Forked from [gianarb/go-irc](https://github.com/gianarb/go-irc)


### PACKAGE DOCUMENTATION
```go
package irc
```
```go
type Bot struct {
    // contains filtered or unexported fields
}
```
Implement BotInterface 
```go
type BotInterface interface {
    Connect(conn net.Conn, err error)
    Send(string)
}
```
Bot struct contains : host, nickname, user, channel ...

```go
func NewBot(proxy string, server string, nick string, user string, channel string, pass string) *Bot
```
Returns new Bot


Bot Methods
* ``` func (bot *Bot) Connect() (conn net.Conn, err error)``` : Connect bot to irc server
* ``` func (bot *Bot) Send(command string) ``` : Send message to channel after connection

```
type Message struct {
    Draft string
}
```
Best practice for managing simple update of lib

### Use Case
```go
package main

import (
    "log"
    "fmt"
    "regexp"
    "bufio"
    "net/textproto"
    "github.com/gmolveau/go-irc"
)
 
func main(){
    myBot := NewBot(
        "",
        "irc.freenode.net:6667",
        "MyIRCBot",
        "MyIRCBot",
        "#channel",
        "",
    )
    conn, _ := myBot.Connect()
    defer conn.Close()
 
    verbose := true
    reader := bufio.NewReader(bot.conn)
    tp := textproto.NewReader(reader)
    for {
        line, err := tp.ReadLine()
        if err != nil {
            log.Fatal("unable to connect to IRC server ", err)
        }
 
        isPing, _ := regexp.MatchString("PING", line)
        if isPing  == true {
            bot.Send("PONG");
        }
        
        if verbose {
            fmt.Printf("%s\n", line)
        }
    }
}
```
