FROM golang:1.22.4 as builder

RUN mkdir /home/app

WORKDIR /home/app

RUN mkdir ./bin
RUN mkdir ./sbin
RUN mkdir ./sbin/services
RUN mkdir ./sbin/tools
RUN mkdir ./src

COPY src/go.mod src/go.sum ./src/
RUN cd ./src; go mod download && go mod verify

COPY src ./src

WORKDIR /home/app/src

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build" go build -o /home/app/bin/box ./cmd/app

FROM ubuntu:24.10

RUN mkdir /home/app

WORKDIR /home/app

COPY --from=builder /home/app/bin/box ./bin/box
COPY --from=builder /home/app/sbin/tools ./sbin/tools

COPY etc .
COPY system .
COPY var .

ENTRYPOINT ["./bin/box"]