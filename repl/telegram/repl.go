package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/toolchain/default"
	"log"
	"os"
	"time"
)

var bot *tgbotapi.BotAPI

const banner = `
Bem-vindo ao golisp, um dialeto de lisp feito em Go!

Como introdução tente realizar operações matemáticas, como por exemplo (+ 2 2)

Outros comandos incluem (env-dump) para despejar o objeto de env que está sendo usado na vm atualmente, que consequentemente expoe todos os elementos visíveis no escopo.

Funções de cálculo, como +, -, * e / só aceitam dois argumentos. Para realizar com mais elementos utilize a função reduce. Ex: (reduce + '(2 3 4)).

Digite /help para reexibir esta mensagem
`

func main() {
	timeout := time.Second * 5
	ctx := context.Background()
	var err error
	if len(os.Args) < 2 {
		panic("Missing telegram bot api key")
	}
	tc := tdefault.NewDefaultToolchain(nil)
	bot, err = tgbotapi.NewBotAPI(os.Args[1])
	if err != nil {
		panic(err)
	}
	// bot.Debug = true
	log.Printf("Authorized: @%s", bot.Self.UserName)
	updchan, err := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))
	if err != nil {
		panic(err)
	}
	for update := range updchan {
		if update.Message == nil {
			continue
		}
		log.Printf("MSG from: @%s: %s", update.Message.From.UserName, update.Message.Text)
		if update.Message.Command() == "help" {
			reply(&update, banner)
			continue
		}
		if update.Message.Command() == "start" {
			reply(&update, banner)
			continue
		}
		stmt := ""
		resp := func(v data.LispValue) string {
			return v.Repr()
		}
		if update.Message.Command() == "spew" {
			resp = func(v data.LispValue) string {
				return spew.Sdump(v)
			}
			stmt = update.Message.CommandArguments()
		} else {
			stmt = update.Message.Text
		}
		if update.Message.Command() == "parse" {
			stmt = update.Message.CommandArguments()
		}
		ctx, cancel := context.WithTimeout(ctx, timeout)
		ast, err := tc.ParseString(ctx, stmt)
		if err != nil {
			reply(&update, fmt.Sprintf("🤔 %s", err.Error()))
			continue
		}
		if update.Message.Command() == "parse" {
			reply(&update, spew.Sdump(ast))
			continue
		}
		res, err := tc.Eval(ctx, ast)
		if err != nil {
			reply(&update, fmt.Sprintf("🤯 %s", err.Error()))
			continue
		}
		cancel() // Apenas chamar no final. Evitar context leak
		reply(&update, fmt.Sprintf("👍 %s", resp(res)))
	}
}

func reply(u *tgbotapi.Update, txt string) {
	m := tgbotapi.NewMessage(int64(u.Message.From.ID), txt)
	m.ReplyToMessageID = u.Message.MessageID
	_, err := bot.Send(m)
	if err != nil {
		reply(u, fmt.Sprintf("📡 %s", err.Error()))
	}
}
