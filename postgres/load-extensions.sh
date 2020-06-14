psql -v ON_ERROR_STOP=1 --dbname "$POSTGRES_DB" --username "$POSTGRES_USER" <<EOF
create extension "uuid-ossp";
select * FROM pg_extension;
EOF