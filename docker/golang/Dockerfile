FROM golang:1.20

WORKDIR /app

RUN go install github.com/volatiletech/sqlboiler@latest
RUN go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

# この辺はプロジェクトのディレクトリ構成に合わせてください
COPY ./src . 
