package main

import "mangodl/utils"

func main() {
	//start the program using the internal "utils" package
	utils.Execute()
}

//TODO add option to convert chapters into pdf
//TODO fix chapters sorting in -Q ( 1 2 3 instead of 1 10 11 2 3)
//TODO add double check to see if chapters were skipped (e.g. chapter 37 exists, chapter 38 doesn't but the rest of the series does)
//TODO change host to https://www.mangahere.cc/
//TODO fix macOS installation process
