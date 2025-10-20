package build

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
)

// ffiBuilder compiles a c-shared library for FFI use
type ffiBuilder struct {
	*BaseBuilder
}

func newFFIBuilder(options *Options) *ffiBuilder {
	return &ffiBuilder{BaseBuilder: NewBaseBuilder(options)}
}

// CompileProject builds a c-shared library into the configured output path
func (b *ffiBuilder) CompileProject(options *Options) error {
	// Reuse frontend step if needed
	// Generate runtime wrapper to keep parity, though not strictly required for FFI
	if err := generateRuntimeWrapper(options); err != nil {
		return err
	}

	// Ensure bin dir exists
	if err := os.MkdirAll(options.BinDirectory, 0o755); err != nil {
		return err
	}

	// Compute output path
	outputFile := b.OutputFilename(options)
	compiled := filepath.Join(options.BinDirectory, outputFile)
	options.CompiledBinary = compiled

	// Build args
	args := []string{"build", "-buildmode=c-shared"}

	// Tags
	tags := []string{"ffi", options.OutputType}
	if len(options.UserTags) > 0 {
		tags = append(tags, options.UserTags...)
	}
	args = append(args, "-tags", strings.Join(tags, ","))

	// Output
	args = append(args, "-o", compiled)

	// Verbose info
	pterm.Info.Println("Build command:", options.Compiler, commandPrettifier(args))

	cmd := exec.Command(options.Compiler, args...)
	cmd.Dir = b.projectData.Path
	cmd.Env = os.Environ()

	// Darwin: ensure minimum SDK like desktop builder
	// We intentionally keep environment minimal; consumers may override as needed

	cmd.Stderr = os.Stderr
	if options.Verbosity == VERBOSE {
		cmd.Stdout = os.Stdout
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffi build failed: %w", err)
	}

	pterm.Println("Done.")
	return nil
}
