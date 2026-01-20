package entries

import "context"

func MockAnalyze(text string) map[string]any {
	return map[string]any{
		"quadrant_distribution": map[string]float64{
			"UL": 0.40,
			"UR": 0.20,
			"LL": 0.15,
			"LR": 0.25,
		},
		"themes": []string{"work_pressure", "self_criticism", "need_for_rest"},
		"mirror_reflection": map[string]any{
			"summary": "It sounds like today carried a lot of pressure and self-judgment. You’re noticing how that shows up internally, and you’re also aware your body is asking for rest. There may be a tension between what you feel you must do and what you can sustainably carry.",
			"clarifying_questions": []string{
				"What part of today felt most non-negotiable, and who decided that?",
				"When did you first notice the body signal that you needed rest?",
			},
		},
		"safety": map[string]any{
			"risk":               "none",
			"signals":            []string{},
			"recommended_action": "normal_flow",
		},
	}
}

type MockAnalyzer struct{}

func NewMockAnalyzer() *MockAnalyzer { return &MockAnalyzer{} }

func (a *MockAnalyzer) Analyze(ctx context.Context, text string, source string, metadata map[string]string) (map[string]any, error) {
	return MockAnalyze(text), nil
}
