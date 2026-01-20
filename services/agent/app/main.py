from __future__ import annotations

from fastapi import FastAPI, HTTPException, Request

from .analysis_mock import make_entry_id, make_mock_analysis
from .models import AnalyzeRequest, AnalyzeResponse


app = FastAPI(title="PsychApp Agent Service", version="0.1.0")


@app.get("/health")
def health() -> dict:
    return {"status": "ok"}


@app.post("/analyze", response_model=AnalyzeResponse)
def analyze(req: AnalyzeRequest, request: Request) -> AnalyzeResponse:
    request_id = request.headers.get("x-request-id")
    if not req.text.strip():
        raise HTTPException(status_code=400, detail="text is required")

    entry_id = req.entry_id or make_entry_id()
    created_at = req.created_at or AnalyzeResponse.now_iso()
    analysis = make_mock_analysis(req.text)

    if request_id:
        print(f"[agent] x-request-id={request_id}")

    return AnalyzeResponse(
        entry_id=entry_id,
        created_at=created_at,
        analysis=analysis,
        request_id=request_id,
        mock_notice="mock analysis - python agent scaffold (dev only)",
    )
