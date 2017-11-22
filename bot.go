package irc

import (
	"fmt"
	"golang.org/x/net/proxy"
	"log"
	"net"
)

type BotInterface interface {
	Connect(conn net.Conn, err error)
	Send(string)
}

type Bot struct {
	proxy         string
	server        string
	Nick          string
	User          string
	Channel       string
	pass          string
	pread, pwrite chan string
	conn          net.Conn
}

func (bot *Bot) Send(command string) {
	fmt.Fprintf(bot.conn, "%s \r\n", command)
}

func NewBot(proxy string, server string, Nick string, User string, Channel string, pass string) *Bot {
	return &Bot{
		proxy:   proxy,
		server:  server,
		Nick:    Nick,
		User:    User,
		Channel: Channel,
		pass:    pass,
		conn:    nil,
	}
}

func (bot *Bot) Connect() (conn net.Conn, err error) {
	if len(bot.proxy) > 0 {
		dialer, err := proxy.SOCKS5("tcp", bot.proxy, nil, proxy.Direct)
		if err != nil {
			log.Fatal("unable to connect to proxy server ", err)
		}
		bot.conn, err = dialer.Dial("tcp", bot.server)
		if err != nil {
			log.Fatal("unable to connect to IRC server ", err)
		}
	} else {
		bot.conn, err = net.Dial("tcp", bot.server)
		if err != nil {
			log.Fatal("unable to connect to IRC server ", err)
		}
	}
	log.Printf("Connected to IRC server %s (%s) \n", bot.server, bot.conn.RemoteAddr())
	if len(bot.pass) > 0 {
		bot.Send(fmt.Sprintf("PASS %s \r\n", bot.pass))
	}
	bot.Send(fmt.Sprintf("USER %s 8 * :%s", bot.User, bot.User))
	bot.Send(fmt.Sprintf("NICK %s \r\n", bot.Nick))
	bot.Send(fmt.Sprintf("JOIN %s \r\n", bot.Channel))

	return bot.conn, err
}
