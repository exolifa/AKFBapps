package main

import (
	"AKFBapps/membersmod"
	"fmt"
)

func main() {

	fmt.Println("starting")
	membersmod.Display_member()
	fmt.Println(membersmod.Email_list())
	fmt.Println(membersmod.Members_list())

}
