package git

import (
	"gopkg.in/src-d/go-git.v4/config"
	"io/ioutil"
	"log"
	"regexp"
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

	return conf.Remotes
}

func RemoteURL() string {
	remotes := ReadConfig()
	return remotes["origin"].URLs[0]
}

func RepositoryPath() string {
	re := regexp.MustCompile("git@github.com:(.*).git")
	return re.FindStringSubmatch(RemoteURL())[1]
}
