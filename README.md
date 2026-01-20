Web:

1. cd web
2. npm run dev

API:

1. cd services/api
2. go run ./cmd/api
   or
   AGENT_DRIVER=python AGENT_URL="http://127.0.0.1:8091" REPO_DRIVER=postgres POSTGRES_DSN="postgres://$(whoami)@localhost:5432/psychapp?sslmode=disable" go run ./cmd/api

Agent:

1. cd services/agent
2. source .venv/bin/activate
3. uvicorn app.main:app --reload --port 8091

Test:

1. CREATE TWO ENTRIES
   curl -s -X POST http://localhost:8080/entries \
    -H "Content-Type: application/json" \
    -d '{"text":"Entry one for timeline","source":"open_journal"}' | cat

curl -s -X POST http://localhost:8080/entries \
 -H "Content-Type: application/json" \
 -d '{"text":"Entry two for timeline","source":"daily_checkin"}' | cat

2. SHOW ENTRIES
   curl -s "http://localhost:8080/entries?limit=10" | cat
