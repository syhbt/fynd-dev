// internal/banner/banner.go
package internal

import "fmt"

func DisplayBanner(silent bool) {
	if !silent {
		fmt.Println(`

	    ______                __
	   / ____/_  ______  ____/ /
	  / /_  / / / / __ \/ __  /
	 / __/ / /_/ / / / / /_/ /
	/_/    \__, /_/ /_/\__,_/
	      /____/	Fynd v0.1

   	         github.com/syhbt

        `)
	}
}
