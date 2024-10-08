package nknovh_engine

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

func (o *configuration) configure() (*configuration, error) {
	//Static configuration
	f, err := os.Open("conf.json")
	if err != nil {
		return &configuration{}, err
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	conf := configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		return &configuration{}, err
	}
	//Container configuration
	if env := os.Getenv("ENVIRONMENT"); env == "dev" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("No .env files found. Using real environment")
		}

	}
	v := reflect.ValueOf(&conf).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		varName, _ := f.Tag.Lookup("env")
		if varName == "-" {
			continue
		}
		env, ok := os.LookupEnv(varName)
		if ok {
			v.Field(i).SetString(env)
		}

	}
	return &conf, nil
}
