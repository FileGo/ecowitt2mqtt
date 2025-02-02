FROM golang:alpine AS build-env
WORKDIR /app
ADD . /app/
RUN go mod download
RUN go build -o /ecowitt2mqtt

FROM gcr.io/distroless/base
COPY --from=build-env /ecowitt2mqtt /
ENV GIN_MODE=release
EXPOSE 55904
CMD ["/ecowitt2mqtt"]