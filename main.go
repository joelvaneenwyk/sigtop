// Copyright (c) 2021, 2023 Tim van der Molen <tim@kariliq.nl>
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

package main

import (
	"log"
	"os"

	cmds "github.com/joelvaneenwyk/sigtop/cmd/sigtop"
	"github.com/tbvdm/go-cli"
)

func main() {
	cli.SetLog()

	if len(os.Args) < 2 {
		cli.ExitUsage("command", "[argument ...]")
	}

	cmd := cmds.Command(os.Args[1])
	if cmd == nil {
		log.Fatalln("invalid command:", os.Args[1])
	}

	switch cmd.Execute(os.Args[2:]) {
	case cmds.CommandError:
		os.Exit(1)
	case cmds.CommandUsage:
		cli.ExitUsage(cmd.Name, cmd.Usage)
	}
}
