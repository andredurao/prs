package git

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4/config"
	"io/ioutil"
	"log"
)

func ReadConfig() map[string]*config.RemoteConfig {
	data, err := ioutil.ReadFile("./.git/config")
	if err != nil {
		log.Fatal(err)
	}

	conf := config.NewConfig()
	err = conf.Unmarshal(data)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range conf.Remotes {
		fmt.Println("Key:", key, "Value:", value)
	}
	return conf.Remotes
}
