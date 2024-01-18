package macmap

import (
	"encoding/csv"
	"io"
)

type CReader struct {
	Reader     *csv.Reader
	skip       int
	limit      int
	isLimit    bool
	fieldNames []string
}

func NewReader(r io.Reader) *CReader {
	reader := &CReader{}
	reader.Reader = csv.NewReader(r)
	return reader
}

func (r *CReader) Init(originalCSV csv.Reader) {
	r.Reader.Comma = originalCSV.Comma
	r.Reader.LazyQuotes = originalCSV.LazyQuotes
	r.Reader.TrimLeadingSpace = originalCSV.TrimLeadingSpace
}

func (r *CReader) SetSkip(skip int) {
	r.skip = skip
}

func (r *CReader) SetLimit(limit int) {
	r.limit = limit
	r.isLimit = true
}
func (r *CReader) SetFieldNames(fieldNames []string) {
	r.fieldNames = fieldNames
}

func (r *CReader) GetFieldNames() (fieldNames []string, err error) {
	if len(r.fieldNames) == 0 {
		if fieldNames, err = r.Reader.Read(); err != nil {
			return nil, err
		}
	} else {
		return r.fieldNames, nil
	}

	emptyIndex := 0
	for i := len(fieldNames) - 1; i >= 0; i-- {
		if fieldNames[i] == "" {
			emptyIndex++
		} else {
			break
		}
	}
	fieldNames = fieldNames[:len(fieldNames)-emptyIndex]
	r.fieldNames = fieldNames
	return fieldNames, nil
}

func (r *CReader) Read1Line() (record []string, err error) {
	for {
		if len(r.fieldNames) == 0 {
			continue
		}
		if r.skip <= 0 {
			break
		}
		if _, err = r.Reader.Read(); err != nil {
			return nil, err
		}
		r.skip--
	}

	return r.Reader.Read()
}

func (r *CReader) ReadBlock() (records [][]string, err error) {
	for {
		if record, err := r.Reader.Read(); err != nil {
			break
		} else {
			if r.skip <= 0 {
				records = append(records, record)
				if r.isLimit && len(records) == r.limit {
					break
				}
				continue
			}
			r.skip--
		}
	}

	return
}

func (r *CReader) ReadAll() (records [][]string, err error) {
	var record []string
	for record, err = r.Reader.Read(); err == nil; record, err = r.Reader.Read() {
		records = append(records, record)
	}
	if err != nil && err != io.EOF {
		return nil, err
	}

	return records[1:], nil
}

func (r *CReader) ReadAll2Map(keyName string) (records map[string]interface{}, err error) {

	var keyNameInSlice = false
	var keyIndex = 0
	for key, value := range r.fieldNames {
		if value == keyName {
			keyNameInSlice = true
			keyIndex = key
		}
	}
	if !keyNameInSlice {
		return nil, err
	}
	records = make(map[string]interface{})
	for record, err := r.Reader.Read(); err == nil; record, err = r.Reader.Read() {
		records[record[keyIndex]] = record
	}
	if err != nil && err != io.EOF {
		return nil, err
	}

	return records, nil
}
