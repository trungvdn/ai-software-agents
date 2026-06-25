package markdown

import "strings"

type MarkdownBuilder struct {
	builder strings.Builder
}

func (b *MarkdownBuilder) H1(text string) *MarkdownBuilder {
	b.builder.WriteString("# " + text + "\n\n")
	return b
}

func (b *MarkdownBuilder) Bullet(text string) *MarkdownBuilder {
	b.builder.WriteString("- " + text + "\n")
	return b
}

func (b *MarkdownBuilder) H2(text string) *MarkdownBuilder {
	b.builder.WriteString("## " + text + "\n\n")
	return b
}

func (b *MarkdownBuilder) H3(text string) *MarkdownBuilder {
	b.builder.WriteString("### " + text + "\n\n")
	return b
}

func (b *MarkdownBuilder) H4(text string) *MarkdownBuilder {
	b.builder.WriteString("#### " + text + "\n\n")
	return b
}

func (b *MarkdownBuilder) Text(text string) *MarkdownBuilder {
	b.builder.WriteString(text)
	return b
}

func (b *MarkdownBuilder) Paragraph(text string) *MarkdownBuilder {
	b.builder.WriteString(text)
	b.builder.WriteString("\n\n")
	return b
}

func (b *MarkdownBuilder) Divider() {
	b.builder.WriteString("---\n\n")
}

func (b *MarkdownBuilder) Bold(text string) *MarkdownBuilder {
	b.builder.WriteString("**" + text + "**")
	return b
}

func (b *MarkdownBuilder) LabeledParagraph(bold string, text string) *MarkdownBuilder {
	b.builder.WriteString("**" + bold + "** " + text + "\n\n")
	return b
}

func (b *MarkdownBuilder) String() string {
	return b.builder.String()
}
