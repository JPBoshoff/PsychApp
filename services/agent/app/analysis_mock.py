from __future__ import annotations

from datetime import datetime, timezone
from typing import Any, Dict


def make_entry_id() -> str:
    now = datetime.now(timezone.utc)
    # Similar feel to your Go IDs
    return "entry_" + now.strftime("%Y%m%d_%H%M%S.%f")[:-3]


def make_mock_analysis(text: str) -> Dict[str, Any]:
    # For now, deterministic mock. Later this becomes your real AQAL agent pipeline.
    return {
        "quadrant_distribution": {"UL": 0.40, "UR": 0.20, "LL": 0.15, "LR": 0.25},
        "themes": ["work_pressure", "self_criticism", "need_for_rest"],
        "mirror_reflection": {
            "summary": (
                "It sounds like today carried a lot of pressure and self-judgment. "
                "You’re noticing how that shows up internally, and you’re also aware "
                "your body is asking for rest. There may be a tension between what "
                "you feel you must do and what you can sustainably carry."
            ),
            "clarifying_questions": [
                "What part of today felt most non-negotiable, and who decided that?",
                "When did you first notice the body signal that you needed rest?",
            ],
        },
        "safety": {"risk": "none", "signals": [], "recommended_action": "normal_flow"},
    }
