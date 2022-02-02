# Собираем в гошке
FROM golang:1.17.6 as build

ENV BIN_FILE /opt/banner_rotator/banner_rotator-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ${CODE_DIR}

# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/banner_rotator/*

# На выходе тонкий образ
FROM alpine:3.9

LABEL ORGANIZATION="OTUS Online Education"
LABEL SERVICE="banner rotator"
LABEL MAINTAINERS="ALESTRAN"

ENV BIN_FILE "/opt/banner_rotator/banner_rotator-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/banner_rotator/config.json
ENV CONFIG_FILE2 ./configs/config.json
COPY ./configs/config.json ${CONFIG_FILE2}

CMD ${BIN_FILE}
