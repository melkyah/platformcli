package cmds

import (
	"os"

	"sigs.k8s.io/kind/pkg/cmd"
	"sigs.k8s.io/kind/pkg/cmd/kind"
	"sigs.k8s.io/kind/pkg/errors"
	"sigs.k8s.io/kind/pkg/exec"
	"sigs.k8s.io/kind/pkg/log"
)

func Kind() {
	if err := Run(cmd.NewLogger(), cmd.StandardIOStreams(), os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

// Run invokes the kind root command, returning the error.
// See: sigs.k8s.io/kind/pkg/cmd/kind
func Run(logger log.Logger, streams cmd.IOStreams, args []string) error {
	// actually run the command
	c := kind.NewCommand(logger, streams)
	c.SetArgs(args)
	if err := c.Execute(); err != nil {
		logError(logger, err)
		return err
	}
	return nil
}

// logError logs the error and the root stacktrace if there is one
func logError(logger log.Logger, err error) {
	colorEnabled := cmd.ColorEnabled(logger)
	if colorEnabled {
		logger.Errorf("\x1b[31mERROR\x1b[0m: %v", err)
	} else {
		logger.Errorf("ERROR: %v", err)
	}
	// Display Output if the error was from running a command ...
	if err := exec.RunErrorForError(err); err != nil {
		if colorEnabled {
			logger.Errorf("\x1b[31mCommand Output\x1b[0m: %s", err.Output)
		} else {
			logger.Errorf("\nCommand Output: %s", err.Output)
		}
	}
	// TODO: stacktrace should probably be guarded by a higher level ...?
	if logger.V(1).Enabled() {
		// Then display stack trace if any (there should be one...)
		if trace := errors.StackTrace(err); trace != nil {
			if colorEnabled {
				logger.Errorf("\x1b[31mStack Trace\x1b[0m: %+v", trace)
			} else {
				logger.Errorf("\nStack Trace: %+v", trace)
			}
		}
	}
}
