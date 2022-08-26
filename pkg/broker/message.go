package broker

import (
	"encoding/json"
)

type Message struct {
	Value []byte
}

func (m Message) Unmarshal(deliveryBody []byte, i interface{}) error {
	err := json.Unmarshal(deliveryBody, &m)
	if err != nil {
		return err
	}

	return json.Unmarshal(m.Value, i)
}
