package infrastructure

import (
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/domain"
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/infrastructure/markdown"
	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/requirement"
)

type Formatter interface {
	Format(
		aggregate *requirement.RequirementAggregate,
	) (*domain.Page, error)
}

type RequirementFormatter struct {
}

func NewRequirementFormatter() *RequirementFormatter {
	return &RequirementFormatter{}
}

func (r *RequirementFormatter) Format(
	aggregate *requirement.RequirementAggregate,
) (*domain.Page, error) {
	// Implement the logic to format the requirement aggregate into a Confluence page
	// You can use the aggregate parameter to access the requirement, epics, and stories
	// For example, you might create a new domain.Page with the appropriate title and content based on the aggregate
	/*
		# Intelligent Investment System

		---

		## Vision

		Build an AI-powered investment platform...

		---

		## Goals

		- Track portfolio
		- Analyze stocks
		- Provide recommendations

		---

		# Epic

		## Portfolio Management

		Manage investment portfolio.

		### Stories

		#### Add Stock

		As a retail investor

		I want to add a stock

		So that I can track my investments

		---

		#### Remove Stock

		...
	*/
	builder := &markdown.MarkdownBuilder{}
	r.buildVision(builder, aggregate)
	builder.Divider()
	r.buildGoals(builder, aggregate)
	builder.Divider()
	r.buildEpics(builder, aggregate)
	builder.Divider()

	return &domain.Page{
		Title:   r.buildTitle(aggregate),
		Content: builder.String(),
	}, nil
}

func (r *RequirementFormatter) buildTitle(aggregate *requirement.RequirementAggregate) string {
	return aggregate.Requirement.ProjectName
}

func (r *RequirementFormatter) buildVision(builder *markdown.MarkdownBuilder, aggregate *requirement.RequirementAggregate) {
	builder.H1("Vision")
	builder.Paragraph(aggregate.Requirement.Vision)
}

func (r *RequirementFormatter) buildGoals(
	builder *markdown.MarkdownBuilder,
	aggregate *requirement.RequirementAggregate,
) {
	builder.H2("Goals")

	for _, goal := range aggregate.Requirement.Goals {
		builder.Bullet(goal.Description)
	}
}

func (r *RequirementFormatter) buildEpics(builder *markdown.MarkdownBuilder, aggregate *requirement.RequirementAggregate) {
	/*
		## Portfolio Management

			Manage investment portfolio.

			### Stories

			#### Add Stock

			As a retail investor

			I want to add a stock

			So that I can track my investments

			---

		#### Remove Stock
	*/
	builder.H3("Epics")
	for _, epic := range aggregate.Epics {
		builder.H2(epic.Name).String()
		builder.Paragraph(epic.Description)
		builder.H3("Stories")
		for _, story := range epic.Stories {
			builder.H4(story.Title)
			builder.Text("**As a** " + story.AsA + "\n\n")
			builder.Text("**I want** " + story.IWant + "\n\n")
			builder.Text("**So that** " + story.SoThat + "\n\n")
		}
	}
}
