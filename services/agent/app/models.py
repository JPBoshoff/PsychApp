from __future__ import annotations

from datetime import datetime, timezone
from typing import Any, Dict, List, Optional

from pydantic import BaseModel, Field


class AnalyzeRequest(BaseModel):
    text: str = Field(..., min_length=1)
    source: Optional[str] = None
    metadata: Optional[Dict[str, str]] = None

    # Optional - provided by upstream (Go API) for traceability
    entry_id: Optional[str] = None
    created_at: Optional[str] = None


class MirrorReflection(BaseModel):
    summary: str
    clarifying_questions: List[str] = Field(default_factory=list)


class Safety(BaseModel):
    risk: str = "none"  # none|low|medium|high later
    signals: List[str] = Field(default_factory=list)
    recommended_action: str = "normal_flow"


class AnalyzeResponse(BaseModel):
    entry_id: str
    created_at: str
    analysis: Dict[str, Any]
    mock_notice: str = "mock analysis - python agent scaffold (dev only)"
    request_id: str | None = None


    @staticmethod
    def now_iso() -> str:
        return datetime.now(timezone.utc).isoformat().replace("+00:00", "Z")
