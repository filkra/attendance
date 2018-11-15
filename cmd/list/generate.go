package list

import (
	"fmt"
	"github.com/filkra/attendance/csv"
	"github.com/filkra/attendance/table"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var listGenerateCommand = &cobra.Command{
	Use:           "generate [appointment] [group_folder] [output_folder]",
	Short:         "Generates attendance lists using the specified groups",
	SilenceErrors: true,
	Args:          cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		appointment , err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		groups := csv.Load(fmt.Sprintf("%s/%d", args[1], appointment))
		for group, members := range groups {
			pdf := table.Generate(appointment, group, members)

			err := pdf.OutputFileAndClose(fmt.Sprintf("%s/%d/%s.pdf", args[2], appointment, group))
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}
