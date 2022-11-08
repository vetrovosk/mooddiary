// а вот тут будет жить наш бот
package bot

// это интерфейс , который у нас должна реализовывать наша хранилка, которая живет в директории сторадж
type Storage interface {
	// тут понятно будут еще аргументы
	Get()
	//  и тут
	Set()
}

type Bot struct {
	storage Storage
}

func NewBot(token string, storageOur Storage) *Bot {
	// тут коннешн с тг устаналивается
	// какой-то инит делается
	// возвращается готовый интсан рабочего бота, которому в мейне надо только сделать run
	return &Bot{storage: storageOur}
}

func (b *Bot) Run() {
	// запускаем цикл работы бота
}
