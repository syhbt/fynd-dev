// internal/banner/banner.go
package internal

import "fmt"

func DisplayBanner(silent bool) {
	if !silent {
		fmt.Println(`
        _______
        |  ____|
        | |__ _   _ _ __
        |  __| | | | '_ \
        | |  | |_| | | | |
        |_|   \__,_|_| |_| Bug Bounty Target Finder. v0.1
        `)
	}
}
