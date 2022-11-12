package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// вот енв лучше вынести отдельно, потому что очевидно, что если у нас нев пусто(==""), то кто-то что-то не заполнил и надо просто заканчивать работу.
	bot, err := tgbotapi.NewBotAPI(os.Getenv("os.GetENV()"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	// такое выносится в константы или конфиг
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			// кажется тут нет смысла паниковать, надо залогировать ошибку и пойти на след цикл, ну кому-то не смогли ответить...
			// очень жаль. по хорошему у нас должен быть механизм ретраев, если не вышло сейчас овтетить. но это все потом
			panic(err)
		}
	}
}


//
func main2() {
	// подгатавливаем данные
	// считываем конфиги
	// читаем енв
	// настраиваем уровень логирование
	// и потом начинаем наше приложение что-то типа
	data = {}

	// потом инициализируем ботаcon
	bot, err  := NewBot(data)
	if err != nil {
		log.Panicf("can't init bot, error: %s", err)
	}
	bot.Run()

	// то есть мейн такой голенький
	// It's common to have a small main function that imports and invokes the code from the /internal and /pkg directories and nothing else.


	// сам бот будет жить где-то в папке internal
	// https://github.com/golang-standards/project-layout

}