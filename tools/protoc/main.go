package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type build struct {
	Inputs []string `json:"inputs"`
}

func buildGo(wd string, input string) {

	protocCmd := exec.Command("protoc", fmt.Sprintf("--proto_path=%s/protos", wd), fmt.Sprintf("--go_out=%s/go/pkg", wd), "--go_opt=paths=source_relative", fmt.Sprintf("%s/protos/%s", wd, input))
	protocCmd.Stderr = os.Stderr

	output, err := protocCmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(output)

}

func gocompile(wd string, input string) {

	protocCmd := exec.Command("protoc", fmt.Sprintf("--proto_path=%s/protos", wd), fmt.Sprintf("--go_out=%s/go/pkg", wd), "--go_opt=paths=source_relative", fmt.Sprintf("%s/protos/%s", wd, input))
	protocCmd.Stderr = os.Stderr

	protocCmd.Output()

}

func javacompile(wd string, input string) {
	protocCmd := exec.Command("protoc", fmt.Sprintf("--proto_path=%s/protos", wd), fmt.Sprintf("--java_out=%s/java/src/main/java", wd), fmt.Sprintf("%s/protos/%s", wd, input))
	protocCmd.Stderr = os.Stderr

	protocCmd.Output()
}

func main() {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to get working directory. Reason: %s", err.Error())
	}

	buildFilePath := fmt.Sprintf("%s/build/build.json", wd)
	builddata, err := ioutil.ReadFile(buildFilePath)
	if err != nil {
		log.Fatalf("Unable to read build json. Reason: %s", err.Error())
	}

	var build build
	err = json.Unmarshal(builddata, &build)
	if err != nil {
		log.Fatalf("Unable to unmarshal json. Reason: %s", err)
	}

	for _, input := range build.Inputs {
		buildGo(wd, input)
		javacompile(wd, input)
	}

}
