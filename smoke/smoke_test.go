package smoke_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// TestExamplesBuildAndRun compiles each example binary and executes it
// with TRULAYER_DRY_RUN=true so the SDK does not attempt any HTTP calls.
// A non-zero exit from any example fails the test.
func TestExamplesBuildAndRun(t *testing.T) {
	examples := []string{
		"basic_trace",
		"openai_auto",
		"rag_pipeline",
		"agent",
		"feedback",
	}

	tmp := t.TempDir()
	repoRoot, err := filepath.Abs("..")
	if err != nil {
		t.Fatalf("abs: %v", err)
	}

	for _, name := range examples {
		name := name
		t.Run(name, func(t *testing.T) {
			binPath := filepath.Join(tmp, name)
			src := filepath.Join(repoRoot, "examples", name)

			build := exec.Command("go", "build", "-o", binPath, ".")
			build.Dir = src
			build.Env = append(os.Environ(), "GOFLAGS=-mod=mod")
			if out, err := build.CombinedOutput(); err != nil {
				t.Fatalf("build %s: %v\n%s", name, err, out)
			}

			run := exec.Command(binPath)
			run.Env = append(os.Environ(),
				"TRULAYER_DRY_RUN=true",
				"TRULAYER_API_KEY=tl_smoke",
			)
			if out, err := run.CombinedOutput(); err != nil {
				t.Fatalf("run %s: %v\n%s", name, err, out)
			}
		})
	}
}
