package src

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetFiles returns files under root that match "Includes" and don't match "Excludes."
func GetFiles(root string, cfg *Config) ([]string, error) {
	var files []string

	includes := normalizePatterns(cfg.Includes)
	excludes := normalizePatterns(cfg.Excludes)

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, werr error) error {
		if werr != nil {
			log.Printf("Skipping %q: %v", path, werr)
			return nil
		}
		if d.IsDir() {
			return nil
		}
		cleaned := filepath.ToSlash(filepath.Clean(path))
		if matchesPattern(cleaned, includes) && !matchesPattern(cleaned, excludes) {
			files = append(files, cleaned)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk error at %q: %w", root, err)
	}

	return files, nil
}

func normalizePatterns(patterns []string) []string {
	out := make([]string, len(patterns))
	for i, p := range patterns {
		out[i] = filepath.ToSlash(filepath.Clean(p))
	}
	return out
}

func matchesPattern(path string, patterns []string) bool {
	for _, p := range patterns {
		if path == p || matchesParentPattern(path, p) || matchesGlobPattern(path, p) {
			return true
		}
	}
	return false
}

func matchesParentPattern(path, pattern string) bool {
	rel, err := filepath.Rel(pattern, path)
	return err == nil && !strings.HasPrefix(rel, "..")
}

func matchesGlobPattern(path, pattern string) bool {
	matched, err := filepath.Match(pattern, filepath.Base(path))
	return err == nil && matched
}
