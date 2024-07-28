package cmds

import (
	"path/filepath"

	"github.com/joelvaneenwyk/sigtop/pkg/signal"
	"github.com/tbvdm/go-openbsd"
)

type cmdStatus int

const (
	CommandOk cmdStatus = iota
	CommandError
	CommandUsage
)

type cmdEntry struct {
	Name    string
	Alias   string
	Usage   string
	Execute func([]string) cmdStatus
}

var cmdEntries = []cmdEntry{
	cmdCheckDatabaseEntry,
	cmdExportAvatarsEntry,
	cmdExportAttachmentsEntry,
	cmdExportDatabaseEntry,
	cmdExportMessagesEntry,
	cmdQueryDatabaseEntry,
}

func Command(name string) *cmdEntry {
	for _, cmd := range cmdEntries {
		if name == cmd.Name || name == cmd.Alias {
			return &cmd
		}
	}
	return nil
}

func unveilSignalDir(dir string) error {
	if err := openbsd.Unveil(dir, "r"); err != nil {
		return err
	}

	// SQLite/SQLCipher needs to create the WAL and shared-memory files if
	// they don't exist already. See https://www.sqlite.org/tempfiles.html.

	walFile := filepath.Join(dir, signal.DatabaseFile+"-wal")
	shmFile := filepath.Join(dir, signal.DatabaseFile+"-shm")

	if err := openbsd.Unveil(walFile, "rwc"); err != nil {
		return err
	}

	if err := openbsd.Unveil(shmFile, "rwc"); err != nil {
		return err
	}

	return nil
}

func recipientFilename(rpt *signal.Recipient, ext string) string {
	return sanitiseFilename(rpt.DetailedDisplayName() + ext)
}
