package helpers

import (
	"bufio"
	"encoding/csv"
	"io"
)

func ReadCsvFile(rs io.ReadSeeker) ([][]string, error) {
	// Skip first row (line)
	firstRow, err := bufio.NewReader(rs).ReadSlice('\n')

	if err != nil {
		return nil, err
	}
	_, err = rs.Seek(int64(len(firstRow)), io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Read remaining rows
	r := csv.NewReader(rs)
	r.FieldsPerRecord = -1

	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}
