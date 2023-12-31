[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/ou2pDf_w)
### Задача 1: Средняя-Простая
**Заголовок**: Создание HTTP-сервера с базовой маршрутизацией

**Описание**:
Создайте простой HTTP-сервер на Go, который будет отвечать на HTTP-запросы на разные URL-адреса (маршруты). Сервер должен иметь следующую функциональность:

- Ответ на запрос к корневому пути ("/") с приветственным сообщением.
- Ответ на запрос к пути "/time" текущим временем сервера.
- Обработка 404 ошибки для несуществующих маршрутов.

**Требования**:
- Используйте стандартный пакет `net/http` для создания сервера.
- Выводите сообщения в консоль при получении запросов.

**Требования к клиенту**:
- Создайте HTTP-клиент на Go, который будет отправлять запросы к указанным выше маршрутам.
- Клиент должен выводить ответ сервера в консоль.

---

### Задача 2: Средняя-Сложная
**Заголовок**: Разработка HTTP-сервера с поддержкой JSON и обработкой запросов

**Описание**:
Разработайте HTTP-сервер на Go, который будет обрабатывать POST-запросы с JSON-данными. Сервер должен принимать JSON с информацией о пользователе (имя и возраст), сохранять эту информацию в памяти и предоставлять ее по GET-запросу. Требуемая функциональность:

- POST-запрос на "/user" принимает JSON вида `{"name": "Alice", "age": 30}` и сохраняет информацию о пользователе.
- GET-запрос на "/user" возвращает список всех сохраненных пользователей в формате JSON.
- Использование структур Go для представления пользователя и обработки JSON.
- Обработка ошибок, включая неверный формат JSON и неверные данные.

**Требования**:
- Используйте стандартные пакеты `net/http` и `encoding/json`.
- Обеспечьте корректную обработку ошибок и соответствующие HTTP-статусы ответов.
- Для хранения данных о пользователях используйте внутреннюю структуру данных (например, слайс или мапу).

**Требования к клиенту**:
- Реализуйте клиент на Go, который отправляет POST-запросы для создания новых пользователей и GET-запросы для получения списка пользователей.
- Клиент должен корректно обрабатывать и выводить ответы от сервера.

