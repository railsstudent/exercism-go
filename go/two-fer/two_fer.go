package twofer

import "fmt"

func ShareWith(name string) string {
	who := "you"
	if name != "" {
		who = name
	}
	return fmt.Sprintf("One for %s, one for me.", who)
}
