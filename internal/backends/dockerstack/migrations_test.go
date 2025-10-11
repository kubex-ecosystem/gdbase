package dockerstack

import (
	"strings"
	"testing"

	"github.com/kubex-ecosystem/gdbase/internal/bootstrap"
)

func TestParseSQLHandlesDollarQuotedBlocks(t *testing.T) {
	content, err := bootstrap.MigrationFiles.ReadFile("embedded/002_hardening.sql")
	if err != nil {
		t.Fatalf("failed to read embedded migration file: %v", err)
	}

	mgr := &MigrationManager{}
	stmts := mgr.parseSQL(string(content))

	if len(stmts) == 0 {
		t.Fatalf("expected statements but got none")
	}

	// Ensure that at least one statement contains the DO $$ ... $$ block
	found := false
	for _, s := range stmts {
		if strings.Contains(s.SQL, "DO $$") {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("expected a dollar-quoted DO $$ block to be parsed as a single statement; got %d statements", len(stmts))
	}
}

// Test helper to print parsed statements for debugging; run explicitly when needed.
func TestParseSQL_PrintStatements(t *testing.T) {
	content, err := bootstrap.MigrationFiles.ReadFile("embedded/002_hardening.sql")
	if err != nil {
		t.Fatalf("failed to read embedded migration file: %v", err)
	}

	mgr := &MigrationManager{}
	stmts := mgr.parseSQL(string(content))

	t.Logf("parsed %d statements:\n", len(stmts))
	for i, s := range stmts {
		// Print first 200 chars for brevity
		out := s.SQL
		if len(out) > 200 {
			out = out[:200] + "..."
		}
		t.Logf("[%02d] line %d: %s\n", i+1, s.Line, out)
	}
}
