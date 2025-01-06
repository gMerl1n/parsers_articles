package parserhb

// Semaphore — структура для управления количеством параллельных горутин
type Semaphore struct {
	semaChannels chan struct{}
}

// NewSemaphore — создает новый семафор с заданной максимальной емкостью
func NewSemaphore(max int) *Semaphore {
	return &Semaphore{
		semaChannels: make(chan struct{}, max),
	}
}

// Acquire — резервирует место в семафоре
func (s *Semaphore) Acquire() {
	s.semaChannels <- struct{}{}
}

// Release — освобождает место в семафоре
func (s *Semaphore) Release() {
	<-s.semaChannels
}
