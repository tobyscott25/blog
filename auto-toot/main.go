package main

import "os"

func main() {
	println("THE ROYAL TOOT!")
	// Send a toot to Mastodon
	println(os.Getenv("MASTODON_ACCESS_TOKEN"))

}
