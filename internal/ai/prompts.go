package ai

const CommitMessagePrompt = "You are a git commit message generator. Analyze the following diff and generate a concise, conventional commit message. Format: <type>(<scope>): <description>\n\n" +
	"Types: feat, fix, docs, style, refactor, test, chore\n\n" +
	"Diff:\n%s\n\n" +
	"Generate 3 commit message options."

const PRDescriptionPrompt = "You are a pull request description generator. Based on the following commits and changes, generate a comprehensive PR description including:\n" +
	"- Summary\n" +
	"- Key changes (bullet points)\n" +
	"- Testing notes\n" +
	"- Breaking changes (if any)\n\n" +
	"Commits:\n%s\n\n" +
	"Changes:\n%s"

const CodeExplanationPrompt = "You are a code explainer. Explain the following code changes in simple terms. Focus on:\n" +
	"- What the code does\n" +
	"- Why it was changed\n" +
	"- Any potential impacts\n\n" +
	"Code:\n%s"
