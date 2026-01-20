package entries

import "context"

type Analyzer interface {
	Analyze(
		ctx context.Context,
		requestID string,
		entryID string,
		createdAt string,
		text string,
		source string,
		metadata map[string]string,
	) (analysis map[string]any, err error)
}

