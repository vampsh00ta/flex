package main

func main() {
	cfg := LoadCondig()
	gitinit(cfg)
	add()
	commit("no cap and gown ait go to class")

	//fmt.Println(cfg.Authorization)
	//_, err := fetchSongs(cfg.Authorization, cfg.SpotifyAlbum)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
