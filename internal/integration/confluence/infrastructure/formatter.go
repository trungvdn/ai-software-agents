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
	builder := markdown.MarkdownBuilder{}
	builder.H1("Vision")
	builder.Paragraph(r.buildVision(aggregate))
	builder.Divider()
	builder.H2("Goals")
	for _, goal := range r.buildGoals(aggregate) {
		builder.Bullet(goal)
	}
	builder.Divider()
	builder.H3("Epics")
	builder.Paragraph(r.buildEpics(aggregate))
	builder.Divider()
	builder.Bullet(r.buildStories(aggregate))

	return &domain.Page{
		Title:   r.buildTitle(aggregate),
		Content: builder.String(),
	}, nil
}

func (r *RequirementFormatter) buildTitle(aggregate *requirement.RequirementAggregate) string {
	return aggregate.Requirement.ProjectName
}

func (r *RequirementFormatter) buildVision(aggregate *requirement.RequirementAggregate) string {
	return aggregate.Requirement.Vision
}

func (r *RequirementFormatter) buildGoals(aggregate *requirement.RequirementAggregate) []string {
	goals := []string{}
	for _, goal := range aggregate.Requirement.Goals {
		goals = append(goals, goal.Description)
	}
	return goals
}

func (r *RequirementFormatter) buildEpics(aggregate *requirement.RequirementAggregate) string {
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
	epics := ""
	for _, epic := range aggregate.Epics {
		epics += "## " + epic.Name + "\n\n"
		epics += epic.Description + "\n\n"
		epics += "### Stories\n\n"
		for _, story := range epic.Stories {
			epics += "#### " + story.Title + "\n\n"
			epics += story.Title + "\n\n"
		}
	}

	return epics
}

func (r *RequirementFormatter) buildStories(aggregate *requirement.RequirementAggregate) string {
	stories := ""
	for _, story := range aggregate.Stories {
		stories += "- " + story.Title + ": " + story.Title + "\n"
	}
	return stories
}
