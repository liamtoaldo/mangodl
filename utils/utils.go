package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	dl "mangodl/download"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	HELPFLAG            = "-h"
	HELPFLAGALT         = "--help"
	DOWNLOADFLAG        = "-D"
	DOWNLOADFLAGALT     = "--download"
	SEARCHFLAG          = "-S"
	SEARCHFLAGALT       = "--search"
	QUERYFLAG           = "-Q"
	QUERYFLAGALT        = "--query"
	DIRECTORYFLAG       = "-Dir"
	DIRECTORYFLAGALT    = "--directory"
	CHAPTERFLAG         = "-c"
	CHAPTERFLAGALT      = "--chapter"
	CHAPTERRANGEFLAG    = "-cr"
	CHAPTERRANGEFLAGALT = "--chapterrange"
	NOPLOTFLAG          = "-n"
	NOPLOTFLAGALT       = "--noplot"
)

//the variables based on the user's results
var (
	mangaName       string
	selectedMangaID string
	realMangaName   string
	currentState    byte //D for downloading, S for searching, Q for querying, H for help, E for error, F for directory/folder
	foundMangaIDs   []string
	foundMangaNames []string
	chosenDirectory string

	//optional
	chapterState  string //single or multiple or all
	singleChapter string
	chapterBegin  int
	chapterEnd    int
	plotState     string //no or yes
)

type DownloadedManga struct {
	title    string
	chapters string
}

func showHelp() {
	fmt.Println(`Usage: mangodl [FLAGS]...
Download manga using the terminal. The manga list is really big.

Arguments and flags:

	-h, --help			shows this message and exit

	Needed (one of them):
	-D, --download			downloads the manga specified after -D (e.g. mangodl -D jojo will search for 5 manga with that name and ask you which one to download)
	-S, --search			searches for the manga specified after this flag (e.g. mangodl -S "kanojo x kanojo" will search and display the manga found with that name)
	-Q, --query			show downloaded manga
	-Dir, --directory		sets the default directory to download manga (e.g. mangodl -Dir "~/Documents/manga/"), otherwise the default one would be "~/Downloaded Manga/" and the Desktop for Windows
	
	Optional:
	For -D:
	-c, --chapter			used to specify the chapter to download (if omitted it will download them all)
	-cr, --chapterrange		used to specify a range of chapters to download (e.g. mangodl -S -M "Martial Peak" -cr 1 99 will download chapters from 1 to 99 (included)
	
	For -S:
	-n, --noplot		do not print the plot of searched manga	
	`)
}

func isNextArg(index int) bool {
	if index >= len(os.Args) {
		return false
	}
	return true
}

//check and redirect the "states"
func checkArgs() {
	if len(os.Args) <= 1 {
		currentState = 'H'
		return
	}
	//detect if the user is running as sudo
	home, _ := os.UserHomeDir()
	if home == "/root" {
		fmt.Println("Avoid running mangodl as sudo")
		return
	}
	for i, s := range os.Args {
		//Help
		if s == HELPFLAG || s == HELPFLAGALT {
			currentState = 'H'
			return
		}
		//Download
		if s == DOWNLOADFLAG || s == DOWNLOADFLAGALT {
			if !isNextArg(i + 1) {
				currentState = 'E'
				fmt.Println("You should specify the manga to download")
				return
			}
			currentState = 'D'
			mangaName = os.Args[i+1]
			fmt.Println("Attempting to download the first manga result with that name, for better results, first use -S and search")
			continue
		}
		//Search for existing manga
		if s == SEARCHFLAG || s == SEARCHFLAGALT {
			if !isNextArg(i + 1) {
				currentState = 'E'
				fmt.Println("You should specify the manga to search for")
				return
			}
			currentState = 'S'
			mangaName = os.Args[i+1]
			fmt.Println("Searching the manga")
			continue
		}
		//Query
		if s == QUERYFLAG || s == QUERYFLAGALT {
			currentState = 'Q'
			break
		}
		//Directory selection
		if s == DIRECTORYFLAG || s == DIRECTORYFLAGALT {
			currentState = 'F'
			chosenDirectory = os.Args[i+1]
			break
		}

		//chapters
		if s == CHAPTERFLAG || s == CHAPTERFLAGALT {
			if !isNextArg(i + 1) {
				currentState = 'E'
				fmt.Println("Not enough chapter arguments")
				return
			}
			chapterState = "single"
			singleChapter = os.Args[i+1]
			break
		} else if s == CHAPTERRANGEFLAG || s == CHAPTERRANGEFLAGALT {

			if !isNextArg(i+1) && !isNextArg(i+2) {
				currentState = 'E'
				fmt.Println("Not enough chapter arguments")
				return
			}
			chapterState = "multiple"
			chapterBegin, _ = strconv.Atoi(os.Args[i+1])
			chapterEnd, _ = strconv.Atoi(os.Args[i+2])
			break
		} else {
			chapterState = "all"
		}

		//manga plot
		if s == NOPLOTFLAG || s == NOPLOTFLAGALT {
			plotState = "no"
		} else {
			plotState = "yes"
		}
	}
}

//looks for manga and displays them (5 for download and 10 for just searching)
func search(howMany int) {
	URL := fmt.Sprintf("https://ww.mangakakalot.tv/search/" + mangaName)
	currentState = 'D'
	res, err := http.Get(URL)
	if err != nil {
		log.Println("Unable to connect to website, error: ", err)
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	fmt.Println("Found: ")

	counter := 1 //the variable that counts the number of manga listed

	doc.Find("a").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		selectedMangaID, _ = selection.Attr("href")
		//5 is the default for the downloads
		if howMany == 5 {
			if strings.Contains(selectedMangaID, "/manga/") && (i == 60 || i == 64 || i == 68 || i == 72 || i == 76) {
				selectedMangaID = strings.Split(selectedMangaID, "/")[2]
				foundMangaIDs = append(foundMangaIDs, selectedMangaID)

				//search for the manga name, print the first 5 entries and let the user decide
				URL := fmt.Sprintf("https://ww.mangakakalot.tv/manga/" + selectedMangaID)
				currentState = 'D'
				res, err := http.Get(URL)
				if err != nil {
					log.Println("Unable to connect to website, error: ", err)
				}
				doc, _ := goquery.NewDocumentFromReader(res.Body)

				doc.Find("p").Each(func(j int, selection *goquery.Selection) {
					if strings.Contains(selection.Text(), "summary:") {
						realMangaName = strings.Trim(strings.Replace(selection.Text(), " summary:", "", -1), " ")
						foundMangaNames = append(foundMangaNames, realMangaName)
						fmt.Println("["+fmt.Sprint(counter)+"] -", realMangaName)
						counter++
						fmt.Println()
					}
				})

				return true
			}

		} else { //10 is the default for the search
			if strings.Contains(selectedMangaID, "/manga/") && (i == 60 || i == 64 || i == 68 || i == 72 || i == 76 || i == 80 || i == 84 || i == 88 || i == 92 || i == 96) {
				//fmt.Println(selectedMangaID)
				selectedMangaID = strings.Split(selectedMangaID, "/")[2]
				foundMangaIDs = append(foundMangaIDs, selectedMangaID)

				//search for the manga name, print the first 5 entries and let the user decide
				URL := fmt.Sprintf("https://ww.mangakakalot.tv/manga/" + selectedMangaID)
				currentState = 'D'
				res, err := http.Get(URL)
				if err != nil {
					log.Println("Unable to connect to website, error: ", err)
				}
				doc, _ := goquery.NewDocumentFromReader(res.Body)

				doc.Find("p").Each(func(j int, selection *goquery.Selection) {
					if strings.Contains(selection.Text(), "summary:") {
						realMangaName = strings.Trim(strings.Replace(selection.Text(), " summary:", "", -1), " ")
						foundMangaNames = append(foundMangaNames, realMangaName)
						fmt.Println("["+fmt.Sprint(counter)+"] -", realMangaName)
						counter++
					}
				})
				//plot
				if plotState == "yes" {
					doc.Find("#noidungm").Each(func(j int, selection *goquery.Selection) {
						if strings.Contains(selection.Text(), "summary:") {
							fmt.Println(strings.Trim(strings.Replace(strings.Replace(strings.Replace(strings.Trim(selection.Text(), " "), realMangaName, "", -1), "\n", "", -1), "  summary:  ", "", -1), " "))
							fmt.Println()
						}
					})
				}

				return true
			}
		}

		return true
	})

}

//the function to let the user choose the manga
func chooseManga() {
	fmt.Println("Which Manga do you want to download?")

	var inputChoice int
	fmt.Scan(&inputChoice)
	switch inputChoice {
	case 1:
		selectedMangaID = foundMangaIDs[0]
		realMangaName = foundMangaNames[0]
	case 2:
		selectedMangaID = foundMangaIDs[1]
		realMangaName = foundMangaNames[1]
	case 3:
		selectedMangaID = foundMangaIDs[2]
		realMangaName = foundMangaNames[2]
	case 4:
		selectedMangaID = foundMangaIDs[3]
		realMangaName = foundMangaNames[3]
	case 5:
		selectedMangaID = foundMangaIDs[4]
		realMangaName = foundMangaNames[4]
	}
}

//download the chosen manga
func download(chapter float32) bool {

	URL := fmt.Sprintf("https://ww.mangakakalot.tv/chapter/%s/chapter_%v", selectedMangaID, chapter)
	res, err := http.Get(URL)
	if err != nil {
		log.Println("Unable to connect to website, error: ", err)
	}
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	dir := ReadJSON() + realMangaName + "/Chapter " + fmt.Sprint(chapter)
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		log.Println(err)
	}
	selection := doc.Find("p")
	if strings.Contains(selection.Text(), "The server") {
		//if the chapter doesn't exist, delete the just created directory
		os.Remove(dir)
		return false
	}

	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		imageURL, _ := selection.Attr("data-src")
		fileName := fmt.Sprintf("%s/page%v.jpg", dir, i)

		err = dl.DownloadFile(imageURL, fileName)
		if err == nil {
			fmt.Println("downloading ::", fmt.Sprintf("%s/page%v.jpg", dir, i))
		}
	})
	return true
}

//the function used for --query (shows the downloaded manga)
func showDownloaded() {
	var downloaded []DownloadedManga
	dir := ReadJSON()
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println("You don't have any manga downloaded or they couldn't be recognised")
		log.Println(err)
	}

	for _, f := range files {
		if f.IsDir() {
			downloaded = append(downloaded, DownloadedManga{
				title:    f.Name(),
				chapters: "",
			})
			dir += f.Name()
			files, err = ioutil.ReadDir(dir)
			for _, f := range files {
				if f.IsDir() {
					downloaded[len(downloaded)-1].chapters += strings.Split(f.Name(), " ")[1] + " "
				}
			}
			dir = ReadJSON()
		}
	}
	for _, m := range downloaded {
		fmt.Println(m.title, "- Chapters:", m.chapters)
	}
}

//Execute is equivalent to a "main" since it does everything required to run and calls all other private functions
func Execute() {
	//put the default config file, which, for now, contains only the directory
	defaultJson()
	//Check if arguments are inputted correctly and change "states"
	checkArgs()

	if currentState == 'D' {
		go search(5) //redirect search in another goroutine
		chooseManga()
		if chapterState == "all" {
			i := 0
			for {
				i++
				if !download(float32(i)) {
					break
				}
			}
		} else if chapterState == "multiple" {
			for i := chapterBegin; i <= chapterEnd; i++ {
				if !download(float32(i)) {
					break
				}
			}
		} else if chapterState == "single" {
			tmp, _ := strconv.ParseFloat(singleChapter, 32)
			download(float32(tmp))
		}
	} else if currentState == 'Q' {
		showDownloaded()
	} else if currentState == 'F' {
		WriteJson(chosenDirectory)
	} else if currentState == 'S' {
		search(10)
	} else if currentState == 'H' {
		showHelp()
	} else if currentState == 'E' {
		return
	} else {
		fmt.Println("Unknown command, try mangodl --help for help")
	}

}
