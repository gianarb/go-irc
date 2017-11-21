## go-irc
Library to interact with IRC.

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
    "bufio"
    "fmt"
    irc "github.com/gmolveau/go-irc"
    parser "gopkg.in/sorcix/irc.v2"
    "log"
    "net/textproto"
    "time"
)

func main() {
    // don't forget to `go get "gopkg.in/sorcix/irc.v2"`
    // and to run Tor
    bot := irc.NewBot(
        "127.0.0.1:9050",
        "freenodeok2gncmy.onion:6667",
        "MySuperBot",
        "MySuperBot",
        "#go-nuts",
        "",
    )
    conn, _ := bot.Connect()
    defer conn.Close()
    verbose := false
    reader := bufio.NewReader(conn)
    tp := textproto.NewReader(reader)
    for {
        line, err := tp.ReadLine()
        if err != nil {
            log.Fatal("unable to connect to IRC server ", err)
        }
        message := parser.ParseMessage(line)
        if verbose {
            fmt.Printf("%v \n", message)
        }

        if message.Command == "PING" {
            bot.Send(fmt.Sprint("PONG %d", time.Now().UnixNano()))
        }

        if message.Command == "PRIVMSG" {
            if message.Params[0] == bot.Nick {
                // this private message is of the bot
                msg := message.Params[1]
                // fmt.Println(msg)
                // Do Something with this msg
            }
        }
    }
}
```
