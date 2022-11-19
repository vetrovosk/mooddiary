package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

const (
	shortForm = "2006-Jan-02"
)

type user struct {
	birth time.Time
	sex   string
	city  string
}

var gender = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("male"),
		tgbotapi.NewKeyboardButton("female"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("GetENV()")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	cahe := make(map[int64]user)
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		fmt.Println("map:", cahe)

		_, find := cahe[update.Message.From.ID]
		if !find {
			///регистрируйся падла давай
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ты теперь в Матрице")
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
			cahe[update.Message.From.ID] = user{}
			got := tgbotapi.NewMessage(update.Message.Chat.ID, "Дата рождения?(В форме: 2006-Jan-02)")
			if _, err := bot.Send(got); err != nil {
				panic(err)
			}
			continue
		}

		value := cahe[update.Message.From.ID]

		if time.Time.IsZero(value.birth) == true {
			value.birth, _ = time.Parse(shortForm, update.Message.Text)
			cahe[update.Message.From.ID] = value
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Твой пол?")
			msg.ReplyMarkup = gender
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
			continue
		}

		if value.sex == "" {
			value.sex = update.Message.Text
			cahe[update.Message.From.ID] = value

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Твой город?")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
			continue
		}

		if value.city == "" {
			value.city = update.Message.Text
			cahe[update.Message.From.ID] = value
		}
		break
	}
}

//потом сделать
// кажется тут нет смысла паниковать, надо залогировать ошибку и пойти на след цикл, ну кому-то не смогли ответить...
// очень жаль. по хорошему у нас должен быть механизм ретраев, если не вышло сейчас овтетить. но это все потом
