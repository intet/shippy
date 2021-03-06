# Мы используем официальное изображение golang, которое содержит все
# правильные инструменты сборки и библиотеки. Обратите внимание на `as builder`,
# это дает этому контейнеру имя, на которое мы можем ссылаться позже.
FROM golang:alpine as builder

RUN apk --no-cache add git

# Установит рабочий каталог на наш текущий сервис в gopath
WORKDIR /app/user-service

# Скопирует текущий код в рабочий каталог
COPY . .

RUN go mod download

# Создаст двоичный файл с флагами, который позволит
# нам запустить этот двоичный файл в Alpine.
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service -a -installsuffix cgo main.go repository.go handler.go database.go token_service.go

# Здесь мы используем второй оператор FROM,
# это говорит Docker начать новый процесс сборки с этим же образом.
FROM alpine:latest

# Пакет, связанный с безопасностью - хорошо бы его иметь
RUN apk --no-cache add ca-certificates

# Как и прежде, создайте каталог для нашего приложения.
RUN mkdir /app
WORKDIR /app

# Здесь вместо того, чтобы копировать двоичный файл с нашего хоста,
# мы извлекаем двоичный файл из контейнера с именем `builder`
# Это позволяет заглянуть в наш предыдущий образ,
# и найти двоичный файл, который мы создали ранее,
# и поместить его в этот контейнер. Удивительно!
COPY --from=builder /app/user-service/user-service .

# Запустит бинарный файл как обычно! На этот раз с бинарной сборкой в
# отдельном контейнере со всеми правильными зависимостями и
# run time библиотеками.
CMD ["./user-service"]

