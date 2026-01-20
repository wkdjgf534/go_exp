package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Color string

const (
	Green  Color = "green"
	Orange Color = "orange"
	Yellow Color = "yellow"
)

var (
	validColors = map[Color]struct{}{
		Green:  {},
		Orange: {},
		Yellow: {},
	}
)

func (c Color) String() string {
	return string(c)
}

func (c Color) MarshalJSON() ([]byte, error) {
	_, ok := validColors[c]
	if !ok {
		return nil, fmt.Errorf("invalid color: %s", c)
	}

	return []byte(`"` + c + `"`), nil
}

func (c *Color) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}

	_, ok := validColors[Color(s)]
	if !ok {
		return fmt.Errorf("invalid color: %s", s)
	}
	*c = Color(s)
	return nil
}

func main() {
	c := Green
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println("error marshalling color: ", err)
		os.Exit(1)
	}

	fmt.Println(string(b))

	var unmarshalledColor Color
	err = json.Unmarshal([]byte(`"orange"`), &unmarshalledColor)
	if err != nil {
		fmt.Println("error unmarshalling color type: ", err)
		os.Exit(1)
	}

	fmt.Println(unmarshalledColor)
}
