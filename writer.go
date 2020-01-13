package main

import (
	"encoding/csv"
	"io"
	"time"
)

func writeCsv(logs []entry, file io.Writer) error {
	// todo sort by time

	csvwriter := csv.NewWriter(file)
	csvwriter.Write([]string{
		"time",
		"result",
		"commit",
	})
	for _, row := range logs {
		err := csvwriter.Write([]string{
			time.Unix(int64(row.timestamp), 0).Format(time.RFC3339),
			row.data,
			row.commit,
		})
		if err != nil {
			return err
		}
	}

	csvwriter.Flush()

	return nil
}
