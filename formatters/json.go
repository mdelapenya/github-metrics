package formatters

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mdelapenya/github-metrics/types"
)

var data = []*types.LabelResponse{}

type JsonFormatter struct {
	FileBaseFormatter
}

// Format reads the output file, gets the array of labels and appends the current one, writing it to disk
// this could not be very efficient, but let's start from here
func (cf JsonFormatter) Format(lr *types.LabelResponse) {
	data = append(data, lr)

	content, _ := json.Marshal(data)
	_ = ioutil.WriteFile(cf.file, content, 0644)
}

func (cf JsonFormatter) read() ([]types.LabelResponse, error) {
	file, _ := ioutil.ReadFile(cf.file)

	data := []types.LabelResponse{}
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
