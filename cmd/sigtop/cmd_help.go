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

package cmds

import (
	"fmt"
)

var cmdHelpEntry = cmdEntry{
	Name:    "help",
	Alias:   "h",
	Usage:   "",
	Execute: cmdHelp,
}

func cmdHelp(args []string) cmdStatus {
	if len(args) != 0 {
		return CommandUsage
	}

	fmt.Println(`sigtop - Export messages, attachments, and other data from Signal Desktop

USAGE:
    sigtop <command> [options] [arguments]

COMMANDS:
    check              Check the integrity of the Signal Desktop database
    att                Export attachments
    avt                Export avatars
    db                 Export and decrypt the Signal Desktop database
    key                Export the encryption key for the database
    msg                Export messages
    query              Run an SQL query on the database

OPTIONS:
    -B                 Use Signal Desktop Beta encryption key and default directory
    -D                 Export database key instead of encryption key
    -d <dir>           Specify the Signal Desktop data directory
    -c <conv>          Specify conversation(s) to export data from (can be used multiple times)
    -f <format>        Specify message export format (json, text, text-short)
    -s <interval>      Specify time interval for export (YYYY-MM-DD,YYYY-MM-DD)
    -i                 Perform an incremental export (skip previously exported data)
    -m                 Set modification time to the received time (for attachments)
    -M                 Set modification time to the sent time (for attachments)
    -o <outfile>       Specify output file for SQL query results

EXAMPLES:
    sigtop msg -f json messages        # Export all messages in JSON format
    sigtop att -c alice -c bob         # Export attachments from Alice and Bob's conversations
    sigtop att -s 2021-02,             # Export attachments sent from February 2021 onwards
    sigtop db -B signal-beta.db        # Export the database from Signal Desktop Beta
    sigtop query "SELECT * FROM messages;" -o output.txt  # Run a database query and save results

For more details, visit: https://github.com/tbvdm/sigtop`)

	return CommandOK
}
