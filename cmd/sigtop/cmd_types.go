package cmds

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/joelvaneenwyk/sigtop/pkg/getopt"
	"github.com/joelvaneenwyk/sigtop/pkg/safestorage"
	"github.com/joelvaneenwyk/sigtop/pkg/signal"
	"github.com/tbvdm/go-openbsd"
)

type cmdStatus int

const (
	CommandOK cmdStatus = iota
	CommandError
	CommandUsage
)

type cmdEntry struct {
	Name  string
	Alias string
	Usage string
	Execute  func([]string) cmdStatus
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

func encryptionKeyFromFile(keyfile getopt.Arg) (*safestorage.RawEncryptionKey, error) {
	if !keyfile.Set() {
		return nil, nil
	}

	system, file, found := strings.Cut(keyfile.String(), ":")
	if !found {
		system, file = file, system
	}

	f := os.Stdin
	if file != "-" {
		var err error
		if f, err = os.Open(file); err != nil {
			return nil, err
		}
		defer f.Close()
	}

	s := bufio.NewScanner(f)
	s.Scan()
	if s.Err() != nil {
		return nil, s.Err()
	}

	key := safestorage.RawEncryptionKey{
		Key: append([]byte{}, s.Bytes()...),
		OS:  system,
	}

	return &key, nil
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
