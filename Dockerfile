########################################
# Backend build stage
########################################
FROM golang:alpine AS backend-builder

WORKDIR /src

COPY . ./
RUN go mod download
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/kinklist ./cmd

########################################
# Frontend build stage
########################################
FROM node:alpine AS frontend-builder

WORKDIR /app

COPY frontend/package*.json ./
RUN npm ci

COPY frontend ./
RUN npm run build

########################################
# Final runtime image
########################################
FROM alpine:3.20

WORKDIR /app

RUN addgroup -S app && adduser -S app -G app

COPY --from=backend-builder /out/kinklist ./kinklist
COPY --from=frontend-builder /app/dist ./frontend/dist
COPY --from=frontend-builder /app/public ./frontend/public/

RUN mkdir -p /data && chown -R app:app /app /data
ENV KL_DB_PATH=/data/kinklist.db
USER app

EXPOSE 8080

CMD ["./kinklist"]
