package confluence

import (
	"context"

	"github.com/trungvdn/ai-software-agents/internal/intergration/confluence/domain"
)

type ConfluenceClient interface {
	CreatePage(
		ctx context.Context,
		page domain.Page,
	) error
}
