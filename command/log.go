package command

import (
	"fmt"

	. "github.com/heroku/force/error"
	. "github.com/heroku/force/lib"
)

var cmdLog = &Command{
	Run:   getLog,
	Usage: "log",
	Short: "Fetch debug logs",
	Long: `
Fetch debug logs

Examples:

  force log [list]

  force log <id>

  force log delete <id>
`,
}

func init() {
}

func getLog(cmd *Command, args []string) {
	force, _ := ActiveForce()
	if len(args) == 0 || args[0] == "list" {
		records, err := force.QueryLogs()
		if err != nil {
			ErrorAndExit(err.Error())
		}
		force.DisplayAllForceRecords(records)
	} else if args[0] == "delete" {
		if len(args) != 2 {
			ErrorAndExit("You need to provide the id of a debug log to delete.")
		}
		err := force.DeleteToolingRecord("ApexLog", args[1])
		if err != nil {
			ErrorAndExit(err.Error())
		}
		fmt.Println("Debug log deleted")
	} else {
		logId := args[0]
		log, err := force.RetrieveLog(logId)
		if err != nil {
			ErrorAndExit(err.Error())
		}
		fmt.Println(log)
	}
}
