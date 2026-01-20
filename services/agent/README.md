# PsychApp Agent Service (FastAPI)

## Run

python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt

uvicorn app.main:app --reload --port 8091

## Test

curl http://localhost:8091/health

curl -X POST http://localhost:8091/analyze \
 -H "Content-Type: application/json" \
 -d '{"text":"Today I felt overwhelmed at work.","source":"open_journal"}'
