source .env
goose -dir="data/pgmigration" postgres "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DB" $1