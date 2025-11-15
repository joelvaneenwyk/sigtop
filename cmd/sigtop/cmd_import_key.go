// Copyright (c) 2025 Tim van der Molen <tim@kariliq.nl>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package cmds

import (
	"io"
	"log"
	"os"

	"github.com/joelvaneenwyk/sigtop/pkg/getopt"
	"github.com/joelvaneenwyk/sigtop/pkg/signal"
	"github.com/tbvdm/go-openbsd"
)

var cmdImportKeyEntry = cmdEntry{
	Name:    "import-key",
	Alias:   "",
	Usage:   "[-B] [-d signal-directory] [file]",
	Execute: cmdImportKey,
}

func cmdImportKey(args []string) cmdStatus {
	getopt.ParseArgs("Bd:", args)
	var dArg getopt.Arg
	Bflag := false
	for getopt.Next() {
		switch opt := getopt.Option(); opt {
		case 'B':
			Bflag = true
		case 'd':
			dArg = getopt.OptionArg()
		}
	}

	if err := getopt.Err(); err != nil {
		log.Fatal(err)
	}

	args = getopt.Args()
	if len(args) > 1 {
		return CommandUsage
	}

	var key []byte
	if len(args) == 0 || args[0] == "-" {
		var err error
		if key, err = io.ReadAll(os.Stdin); err != nil {
			log.Fatalf("cannot read from standard input: %v", err)
		}
	} else {
		var err error
		if key, err = os.ReadFile(args[0]); err != nil {
			log.Fatal(err)
		}
	}

	signalDir, err := signalDirFromArgument(dArg, Bflag)
	if err != nil {
		log.Fatal(err)
	}

	if err := openbsd.Unveil(signalDir, "r"); err != nil {
		log.Fatal(err)
	}

	if err := openbsd.Pledge("stdio rpath"); err != nil {
		log.Fatal(err)
	}

	if err := signal.ImportEncryptionKey(Bflag, signalDir, key); err != nil {
		log.Fatal(err)
	}

	return CommandOK
}
