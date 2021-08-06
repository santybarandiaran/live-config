package parser

import "github.com/mitchellh/mapstructure"

func Map(json map[string]interface{}, out *interface{}) error {
	//TODO check if it's possible to handle Time attributes
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &out,
	}

	d, err := mapstructure.NewDecoder(config)

	if err != nil {
		return err
	}

	err = d.Decode(json)

	if err != nil {
		return err
	}

	return nil
}


