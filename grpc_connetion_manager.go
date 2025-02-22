package s_utils

import (
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"sync"
)

// client Хранит в себе конект, клиента и кол-во активных запросов
type client[T any] struct {
	conn   *grpc.ClientConn
	client T

	activeRequests int
}

// ClientManager отвечает за работу с несколькими конектами
type ClientManager[T any] struct {
	current *client[T]
	old     map[string]*client[T]

	f      func(cc grpc.ClientConnInterface) T
	target string
	opts   []grpc.DialOption
	mu     sync.RWMutex
}

// GetClient Возвращает активный конект и функцию для отключения текущего запроса
func (m *ClientManager[T]) GetClient() (T, func()) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	//Достаем текущий конект
	mc := m.current

	//Добавляем ему активный запрос
	mc.activeRequests++

	//Создаем функцию для очистки
	cleanup := func() {
		m.mu.Lock()
		defer m.mu.Unlock()

		//Вычитаем активный запрос
		mc.activeRequests--
		//Удаляем неактивные конекты
		m.deleteInactiveConnections()
	}

	return mc.client, cleanup
}

// NewConnection Сохраняет настройки для последующих обновлений подключений и вызывает RefreshConnection
func (m *ClientManager[T]) NewConnection(f func(cc grpc.ClientConnInterface) T, target string, opts ...grpc.DialOption) error {
	m.f = f
	m.target = target
	m.opts = opts

	return m.RefreshConnection()
}

// RefreshConnection Текущее подключение переносит в мапу старых подключений, создает новое подключение и делает его активным
func (m *ClientManager[T]) RefreshConnection() error {
	// Создаём новое gRPC-соединение
	conn, err := grpc.Dial(m.target, m.opts...)
	if err != nil {
		return err
	}

	//Создаем нового клиента
	newManaged := &client[T]{
		conn:   conn,
		client: m.f(conn),
	}

	// Лочим менеджер на время переключения
	m.mu.Lock()
	defer m.mu.Unlock()

	// Текущее «уходит» в old, а на его место ставим новое
	if m.current != nil {
		if m.old == nil {
			m.old = make(map[string]*client[T])
		}
		m.old[uuid.New().String()] = m.current
	}
	m.current = newManaged

	//Удаляем неактиные подключения
	m.deleteInactiveConnections()

	return nil
}

// deleteInactiveConnections Удаляет из мапы неактивные подключения
func (m *ClientManager[T]) deleteInactiveConnections() {
	for k, v := range m.old {
		if v != nil && v.activeRequests == 0 {
			v.conn.Close()
			delete(m.old, k)
		}
	}
}
