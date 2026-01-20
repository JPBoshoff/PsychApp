from __future__ import annotations

from fastapi import FastAPI, HTTPException

from .analysis_mock import make_entry_id, make_mock_analysis
from .models import AnalyzeRequest, AnalyzeResponse


app = FastAPI(title="PsychApp Agent Service", version="0.1.0")


@app.get("/health")
def health() -> dict:
    return {"status": "ok"}


@app.post("/analyze", response_model=AnalyzeResponse)
def analyze(req: AnalyzeRequest) -> AnalyzeResponse:
    if not req.text.strip():
        raise HTTPException(status_code=400, detail="text is required")

    entry_id = make_entry_id()
    analysis = make_mock_analysis(req.text)

    return AnalyzeResponse(
        entry_id=entry_id,
        created_at=AnalyzeResponse.now_iso(),
        analysis=analysis,
    )
