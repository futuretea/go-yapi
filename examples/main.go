package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/futuretea/go-yapi"
	"golang.org/x/crypto/ssh/terminal"
)

func getAPI() (string, string) {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("API URL: ")
	apiURL, _ := r.ReadString('\n')
	apiURL = strings.TrimSpace(apiURL)

	fmt.Print("API Token: ")
	byteAPIToken, _ := terminal.ReadPassword(int(syscall.Stdin))
	apiToken := strings.TrimSpace(string(byteAPIToken))
	fmt.Println()

	return apiURL, apiToken
}

func main() {
	apiURL, apiToken := getAPI()
	yapiClient, err := yapi.NewClient(nil, apiURL, apiToken)
	if err != nil {
		panic(err)
	}
	project, _, _ := yapiClient.Project.Get()
	catMenus, _, _ := yapiClient.CatMenu.Get(project.Data.ID)
	for _, catmenu := range catMenus.Data {
		interfaceListParam := new(yapi.InterfaceListParam)
		interfaceListParam.CatID = catmenu.ID
		interfaceListParam.Page = 1
		interfaceListParam.Limit = 1000
		interfaces, _, _ := yapiClient.Interface.GetList(interfaceListParam)
		for _, i := range interfaces.Data.List {
			result, _, _ := yapiClient.Interface.Get(i.ID)
			fmt.Printf("project_id=%d, catmenu id=%d, interface id=%d, interface title=%s\n",
				project.Data.ID,
				catmenu.ID,
				i.ID,
				result.Data.Title,
			)
		}
	}
}
