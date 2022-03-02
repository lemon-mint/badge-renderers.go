FROM golang:latest as build

ADD . /build
WORKDIR /build

RUN CGO_ENABLED=0 go build -ldflags '-s -w' -o /build/app.exe

FROM scratch

COPY --from=build /build/app.exe /app.exe

ENTRYPOINT ["/app.exe"]
