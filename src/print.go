package src

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// PrintMarkdown creates a Markdown file at outputPath and writes a
// file-structure overview and file contents for the given paths.
func PrintMarkdown(outputPath string, paths []string) error {
	mdFile, err := createMarkdownFile(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create markdown file: %w", err)
	}
	defer mdFile.Close()

	if err := writeHeader(mdFile, "Project"); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	if err := printFileStructure(mdFile, paths); err != nil {
		return fmt.Errorf("failed to print file structure: %w", err)
	}

	if err := printFileContent(mdFile, paths); err != nil {
		return fmt.Errorf("failed to print file content: %w", err)
	}

	return nil
}

func createMarkdownFile(outputPath string) (*os.File, error) {
	mdFile, err := os.Create(outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create markdown file: %w", err)
	}
	return mdFile, nil
}

func writeHeader(w io.Writer, title string) error {
	_, err := fmt.Fprintf(w, "# %s\n\n", title)
	return err
}

func printFileStructure(w io.Writer, paths []string) error {
	if _, err := fmt.Fprintf(w, "## File Structure\n\n"); err != nil {
		return fmt.Errorf("failed to write heading: %w", err)
	}

	tree := buildFileTree(paths)
	return printTree(w, tree, 0)
}

func buildFileTree(paths []string) map[string]interface{} {
	tree := make(map[string]interface{})

	for _, path := range paths {
		parts := strings.Split(path, "/")
		current := tree

		for i, part := range parts {
			// If this is the last part, it's a file; store as nil.
			if i == len(parts)-1 {
				current[part] = nil
				continue
			}
			// Otherwise, it's a directory node.
			if _, exists := current[part]; !exists {
				current[part] = make(map[string]interface{})
			}
			current = current[part].(map[string]interface{})
		}
	}
	return tree
}

func printTree(w io.Writer, tree map[string]interface{}, indentLevel int) error {
	keys := make([]string, 0, len(tree))
	for key := range tree {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		isDir := tree[key] != nil
		displayName := key
		if isDir {
			displayName += "/"
		}

		indent := strings.Repeat("  ", indentLevel)
		if _, err := fmt.Fprintf(w, "%s- %s\n", indent, displayName); err != nil {
			return fmt.Errorf("failed to write tree node: %w", err)
		}

		if isDir {
			subTree := tree[key].(map[string]interface{})
			if err := printTree(w, subTree, indentLevel+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func printFileContent(w io.Writer, paths []string) error {
	if _, err := fmt.Fprintf(w, "## File Content\n\n"); err != nil {
		return fmt.Errorf("failed to write heading: %w", err)
	}

	for _, path := range paths {
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Skipping file %s due to read error: %v\n", path, err)
			continue
		}

		lang := detectLanguage(path)

		if _, err := fmt.Fprintf(w, "### %s\n\n", path); err != nil {
			return fmt.Errorf("failed to write file title: %w", err)
		}
		if _, err := fmt.Fprintf(w, "```%s\n%s\n```\n\n", lang, content); err != nil {
			return fmt.Errorf("failed to write code block: %w", err)
		}
	}
	return nil
}

func detectLanguage(filePath string) string {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".go":
		return "go"
	case ".js":
		return "javascript"
	case ".ts":
		return "typescript"
	case ".py":
		return "python"
	case ".java":
		return "java"
	case ".c":
		return "c"
	case ".cpp":
		return "cpp"
	case ".html":
		return "html"
	case ".css":
		return "css"
	case ".sh":
		return "bash"
	case ".json":
		return "json"
	case ".yaml", ".yml":
		return "yaml"
	case ".xml":
		return "xml"
	case ".md":
		return "markdown"
	default:
		return ""
	}
}
