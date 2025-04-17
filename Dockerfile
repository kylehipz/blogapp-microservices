FROM golang:1.24-alpine AS build

ARG DIR

WORKDIR /app

COPY libs libs
COPY ${DIR} ${DIR}

RUN go work init ./libs ./${DIR}
RUN go mod download
RUN go build -o bin/main ./${DIR}

FROM gcr.io/distroless/static-debian12
ARG DIR

COPY --from=build /app/bin/main /app/main

ENTRYPOINT ["/app/main"]

