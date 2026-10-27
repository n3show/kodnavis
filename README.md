# Kodnavis

Web platform for learning programming.

## Quick start

    cp .env.example .env
    docker compose up --build

- Web: http://localhost:5173
- API health: http://localhost:8080/health
- Postgres: localhost:5432

## Useful commands

    docker compose logs -f api   # follow API logs
    docker compose down          # stop
    docker compose down -v       # stop and delete DB data

