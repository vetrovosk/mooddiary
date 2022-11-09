// вот тут у нас работа связанныая с хранением данных
package storage

import (
	"go.uber.org/zap"
)

// SimpleMemStorage самое простое инмемори хранилище с которого мы начнем
type SimpleMemStorage struct {
	// храним настроение, где ключ это уникальный айди пользователя, а значение это массив чисел.
	// будто у нас от 0 до 100 дают оценку ну вот чисто как пример
	data map[string][]int

	logger *zap.SugaredLogger
}

func NewSimpleMemStorage(logger *zap.SugaredLogger) *SimpleMemStorage {
	return &SimpleMemStorage{
		data:   make(map[string][]int),
		logger: logger,
	}
}

func (s *SimpleMemStorage) Get(key string) ([]int, bool) {
	// хз что он подсвечивает, дурацекий вс код
	return s.data[key]
}

func (s *SimpleMemStorage) Set(key string, value []int) {
	s.data[key] = value
}

// MoreCleverStorage более умное хранилище, которое у нас при старте читает файлик и приодически сохраняет данные в файлик
type MoreCleverStorage struct {
	data     map[string][]int
	filename string

	logger *zap.SugaredLogger
}

func NewMoreCleverStorage(filename string, logger *zap.SugaredLogger) *MoreCleverStorage {

	mcs := MoreCleverStorage{
		filename: filename,
		logger:   logger,
	}
	mcs.data = mcs.initData()
	mcs.startDataSaveLoop()
	return &mcs
}

func (s *MoreCleverStorage) initData() map[string][]int {
	// читаем данные из файла s.file
	data := make(map[string][]int)
	// обработали их
	return data
}

func (s *MoreCleverStorage) startDataSaveLoop() {
	// тут мы сохраняем наши данные с каким то интервалом
}

func (s *MoreCleverStorage) Get(key string) ([]int, bool) {
	// хз что он подсвечивает, дурацекий вс код
	return s.data[key]
}

func (s *MoreCleverStorage) Set(key string, value []int) {
	s.data[key] = value
}
