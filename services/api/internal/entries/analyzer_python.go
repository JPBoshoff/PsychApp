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

func (a *PythonAnalyzer) Analyze(ctx context.Context, requestID string, entryID string, createdAt string, text string, source string, metadata map[string]string) (map[string]any, error) {
	out, err := a.client.Analyze(ctx, requestID, agent.AnalyzeRequest{
		Text:      text,
		Source:    source,
		Metadata:  metadata,
		EntryID:   entryID,
		CreatedAt: createdAt,
	})
	if err != nil {
		return nil, err
	}
	return out.Analysis, nil
}

