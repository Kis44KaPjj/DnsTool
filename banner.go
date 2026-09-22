package main

import "fmt"

const appbanner = `
 _____               _______               __ 
|     \.-----.-----.|_     _|.-----.-----.|  |
|  --  |     |__ --|  |   |  |  _  |  _  ||  |
|_____/|__|__|_____|  |___|  |_____|_____||__|									
`
const linebanner = `══════════════════════════════════════════════`

func showbanner() {
	fmt.Print(appbanner)
	fmt.Printf("                            [ Version: %s ]\n", version)
	fmt.Println(linebanner)
}
