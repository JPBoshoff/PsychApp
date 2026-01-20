package entries

import (
	"context"

	"github.com/JPBoshoff/PsychApp/services/api/internal/agent"
)

type PythonAnalyzer struct {
	client *agent.Client
}

func NewPythonAnalyzer(client *agent.Client) *PythonAnalyzer {
	return &PythonAnalyzer{client: client}
}

func (a *PythonAnalyzer) Analyze(ctx context.Context, text string, source string, metadata map[string]string) (map[string]any, error) {
	out, err := a.client.Analyze(ctx, agent.AnalyzeRequest{
		Text:     text,
		Source:   source,
		Metadata: metadata,
	})
	if err != nil {
		return nil, err
	}
	return out.Analysis, nil
}
