FROM golang:1.22-alpine AS build
ENV TZ=Asia/Tokyo

WORKDIR /opt/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /bin/server ./cmd/server


FROM gcr.io/distroless/base:nonroot AS runner
ENV TZ=Asia/Tokyo
ENV GOENV=production
# should be set by the environment
ENV PORT=1323 

COPY --from=build /bin/server /bin/server

EXPOSE ${PORT}
USER nonroot
ENTRYPOINT ["/bin/server"]
