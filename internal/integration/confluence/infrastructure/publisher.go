package infrastructure

import (
	"context"

	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/domain"
	"github.com/trungvdn/ai-software-agents/internal/requirement/application/publish_requirement"
)

type ConfluencePublisher struct {
	formatter        Formatter
	confluenceClient domain.ConfluenceClient
}

func NewConfluencePublisher(
	formatter Formatter,
	confluenceClient domain.ConfluenceClient,
) *ConfluencePublisher {
	return &ConfluencePublisher{
		formatter:        formatter,
		confluenceClient: confluenceClient,
	}
}
func (c *ConfluencePublisher) Publish(
	ctx context.Context,
	request publish_requirement.PublishRequirementRequest,
) error {
	// Implement the logic to publish the content to Confluence
	// For example, you can use the Confluence API to create a new page or update an existing one
	// You may need to authenticate with the Confluence API and handle any errors that may occur during the publishing process
	page, err := c.formatter.Format(request.RequirementAggregate)
	if err != nil {
		return err
	}

	_, err = c.confluenceClient.CreatePage(ctx, *page)
	if err != nil {
		return err
	}

	return nil
}
