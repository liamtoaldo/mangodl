package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	dl "mangodl-git/download"
	outl "mangodl-git/output"
	"math"
	"net/http"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
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
	OUTPUTFLAG          = "-o"
	OUTPUTFLAGALT       = "--output"
	SPECIALFLAG         = "-s"
	SPECIALFLAGALT      = "--special"
	SELECTFIRSTFLAG     = "-f"
	SELECTFIRSTFLAGALT  = "--first"
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
	alreadyChecked  = false

	//optional
	chapterState  string //single or multiple or all
	singleChapter string
	chapterBegin  string
	chapterEnd    string
	plotState     string  //no or yes
	output        = "img" //img, pdf or cbz, default is image
	special       = false
	selectFirst   = false
)

type DownloadedManga struct {
	title    string
	chapters []float64
}

func showHelp() {
	fmt.Println(`Usage: mangodl [FLAGS]...
Download manga using the terminal. The manga list is really big.

Arguments and flags:

	-h, --help			shows this message and exit

	Needed (one of them):
	-D, --download			downloads the manga specified after -D (e.g. mangodl -D jojo will search for 10 manga with that name and ask you which one to download)
	-S, --search			searches for the manga specified after this flag (e.g. mangodl -S "kanojo x kanojo" will search and display the manga found with that name)
	-Q, --query			show downloaded manga
	-Dir, --directory		sets the default directory to download manga (e.g. mangodl -Dir "$HOME/Documents/manga/"), otherwise the default one would be "$HOME/Downloaded Manga/" and the Desktop for Windows
	
	Optional:
	For -D:
	-c, --chapter			used to specify the chapter to download (if omitted it will download them all)
	-cr, --chapterrange		used to specify a range of chapters to download (e.g. mangodl -D "Martial Peak" -cr 1 99 will download chapters from 1 to 99 (included)
	-o, --output			used to specify the file output of the pages (img, pdf or cbz), e.g. mangodl -D "Tokyo Revengers" -o pdf will create a pdf for every chapter. By default, it's images.
							Remember that this flag and any other flags must be used before the chapter selection flag, otherwise they wouldn't be detected
	-s, --special			used to download "special" chapters too, the ones with floating point values (13.1, 14.7, 99.3, etc). Makes the downloads slower, so use this only if needed
	-f, --first				used to skip the selection phase and select the first manga found. (e.g. mangodl -D "Chainsaw" -f)
	
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
			fmt.Println("Attempting to download a manga with that name, for better results, first use -S and search")
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
		//Skip selection (first)
		if (s == SELECTFIRSTFLAG || s == SELECTFIRSTFLAGALT) && currentState == 'D' {
			selectFirst = true
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
			chapterBegin = os.Args[i+1]
			chapterEnd = os.Args[i+2]
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
		//output
		if s == OUTPUTFLAG || s == OUTPUTFLAGALT {
			if !isNextArg(i + 1) {
				currentState = 'E'
				fmt.Println("You should specify the output (img or pdf)")
				return
			}
			output = os.Args[i+1]
		}
		//special chapters
		if s == SPECIALFLAG || s == SPECIALFLAGALT {
			special = true
		}
	}
}

//looks for manga and displays them (10 for download and 10 for just searching)
func search(howMany int) {
	URL := fmt.Sprintf("https://ww.mangakakalot.tv/search/" + mangaName)
	currentState = 'D'
	res, err := http.Get(URL)
	if err != nil {
		log.Println("Unable to connect to website, error: ", err)
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	fmt.Println("Found: ")

	counter := 1   //the variable that counts the number of manga listed in the search.
	previousI := 0 //the previous index, it's used to avoid doubles
	doc.Find("a").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		selectedMangaID, _ = selection.Attr("href")
		//this condition checks if the current analyzed manga has the tag title (which means it's not a searched manga) and if the i is greater than the previous i + 1 (to avoid duplicates) but, because of this, the first option gets ignored, so I just added another condition
		if strings.Contains(selectedMangaID, "/manga/") && selection.AttrOr("title", "y") == "y" && counter <= howMany && ((counter > 1 && i > previousI+1) || counter == 1) {
			selectedMangaID = strings.Split(selectedMangaID, "/")[2]
			foundMangaIDs = append(foundMangaIDs, selectedMangaID)

			//search for the manga name, print the first 10 entries and let the user decide
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
					//If the skip flag is selected, skip other manga. I placed this code just after the first manga selection.
					if selectFirst {
						return
					}
					fmt.Println("["+fmt.Sprint(counter)+"] -", realMangaName)
					counter++
				}
			})
			//If the skip flag is selected, skip other manga and just download the first possible.
			if selectFirst {
				return false
			}
			//added wait option because the host has become slow, and requests to it need to be slow as well
			time.Sleep(300 * time.Millisecond)
			//plot
			if plotState == "yes" {
				doc.Find("#noidungm").Each(func(j int, selection *goquery.Selection) {
					if strings.Contains(selection.Text(), "summary:") {
						fmt.Println(strings.Trim(strings.Replace(strings.Replace(strings.Replace(strings.Trim(selection.Text(), " "), realMangaName, "", -1), "\n", "", -1), "  summary:  ", "", -1), " "))
						fmt.Println()
					}
				})
			}
			previousI = i
			return true
		}

		return true

	})
}

//the function to let the user choose the manga
func chooseManga() {
	//If the select first flag is true, skip the manga selection and take the first manga.
	if selectFirst {
		selectedMangaID = foundMangaIDs[0]
		realMangaName = foundMangaNames[0]
	} else {
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
		case 6:
			selectedMangaID = foundMangaIDs[5]
			realMangaName = foundMangaNames[5]
		case 7:
			selectedMangaID = foundMangaIDs[6]
			realMangaName = foundMangaNames[6]
		case 8:
			selectedMangaID = foundMangaIDs[7]
			realMangaName = foundMangaNames[7]
		case 9:
			selectedMangaID = foundMangaIDs[8]
			realMangaName = foundMangaNames[8]
		case 10:
			selectedMangaID = foundMangaIDs[9]
			realMangaName = foundMangaNames[9]
		}
	}
	if output != "img" {
		fmt.Println("Downloading images to be converted...")
	}
}

//download the chosen manga
func download(chapter float32) bool {
	var chapterNumber string
	if chapter-float32(int(chapter)) > 0 {
		chapterNumber = strconv.FormatFloat(float64(chapter), 'f', 1, 64)
	} else {
		chapterNumber = strconv.FormatFloat(float64(chapter), 'f', -1, 64)
	}
	URL := fmt.Sprintf("https://ww.mangakakalot.tv/chapter/%s/chapter-%v", selectedMangaID, fmt.Sprint(chapterNumber))
	res, err := http.Get(URL)
	if err != nil {
		log.Println("Unable to connect to website, error: ", err)
	}
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	if strings.Contains(realMangaName, ":") && runtime.GOOS == "windows" {
		realMangaName = strings.ReplaceAll(realMangaName, ":", " -")
	}
	dir := ReadJSON() + realMangaName + "/Chapter " + fmt.Sprint(chapterNumber)
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		log.Println(err)
	}
	selection := doc.Find("span")
	if strings.Contains(selection.Text(), "Error") {
		//wait for half a second, otherwise the checking would be too fast and would skip some chapters
		time.Sleep(300 * time.Millisecond)
		//if the chapter doesn't exist, delete the just created directory
		err := os.Remove(dir)
		if err != nil {
			log.Println(err)
		}
		return false
	}

	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		imageURL, _ := selection.Attr("data-src")
		fileName := fmt.Sprintf("%s/page%03d.jpg", dir, i)

		err = dl.DownloadFile(imageURL, fileName)
		if err == nil && output == "img" {
			fmt.Println("Downloading ::", fileName)
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
				chapters: make([]float64, 0),
			})
			dir += f.Name()
			files, err = ioutil.ReadDir(dir)
			for _, f := range files {
				if f.IsDir() {
					tmp, err := strconv.ParseFloat(strings.Split(f.Name(), " ")[1], 32)
					if err != nil {
						log.Println(err)
					}
					downloaded[len(downloaded)-1].chapters = append(downloaded[len(downloaded)-1].chapters, math.Ceil(tmp*100)/100) //rounding the number to the first decimal digit because some chapters have decimals
				}
			}
			dir = ReadJSON()
		}
	}
	for _, m := range downloaded {
		//sort the chapters, so that they get displayed correctly ( 1 2 3 ) instead of ( 1 10 11 2 23 3 4 ...)
		sort.Float64s(m.chapters)
		fmt.Println(m.title, "- Chapters:", m.chapters)
	}
}

//this function checks if the output was redirected to pdf or cbz, and if it was, it takes care of it
func prepareOutput(i float64) {
	if output != "img" {
		var chapterNumber string
		if i-float64(int(i)) > 0 {
			chapterNumber = strconv.FormatFloat(i, 'f', 1, 64)
		} else {
			chapterNumber = strconv.FormatFloat(i, 'f', -1, 64)
		}
		dir := ReadJSON() + realMangaName + "/Chapter " + chapterNumber
		pageNumber := outl.GetNumberOfPages(dir)
		var pages []string
		for j := 1; j <= pageNumber; j++ {
			pages = append(pages, fmt.Sprintf("%s/page%03d.jpg", dir, j))
		}
		if output == "pdf" {
			fmt.Println("Converting the downloaded images to PDF in", fmt.Sprintf("%s/Chapter%v.pdf", dir, chapterNumber))
			outl.ConvertToPDF(pages, fmt.Sprintf("%s/Chapter%v.pdf", dir, chapterNumber))
			//move the chapter one directory up.
			err := os.Rename(fmt.Sprintf("%s/Chapter%v.pdf", dir, chapterNumber), fmt.Sprintf(strings.Replace(dir, fmt.Sprintf("Chapter %v", chapterNumber), "", 1))+fmt.Sprintf("Chapter%v.pdf", chapterNumber))
			if err != nil {
				log.Println(err)
			}
		} else if output == "cbz" {
			fmt.Println("Converting the downloaded images to CBZ in", fmt.Sprintf("%s/Chapter%v.cbz", dir, chapterNumber))
			outl.ConvertToCBZ(pages, fmt.Sprintf("%s/Chapter%v.cbz", dir, chapterNumber))
			//move the chapter one directory up.
			err := os.Rename(fmt.Sprintf("%s/Chapter%v.cbz", dir, chapterNumber), fmt.Sprintf(strings.Replace(dir, fmt.Sprintf("Chapter %v", chapterNumber), "", 1))+fmt.Sprintf("Chapter%v.cbz", chapterNumber))
			if err != nil {
				log.Println(err)
			}
		}
		//delete the folder and images, (issue #21)
		fmt.Println("Removing previously downloaded images...")
		os.RemoveAll(dir)
		fmt.Println("Done")
	}
}

//Execute is equivalent to a "main" since it does everything required to run and calls all other private functions
func Execute() {
	//put the default config file, which, for now, contains only the directory
	defaultJson()
	//Check if arguments are inputted correctly and change "states"
	checkArgs()

	if currentState == 'D' {

		plotState = "no"
		search(10) //redirect search in another goroutine
		chooseManga()
		if chapterState == "all" {
			if special {
				i := 0
				for {
					i++
					tmpDownloaded := download(float32(i))
					if tmpDownloaded {
						prepareOutput(float64(i))
					}
					alreadyChecked = true
					for j := 0.1; j <= 0.9; j += 0.1 {
						if download(float32(float64(i) + j)) {
							alreadyChecked = false
							prepareOutput(float64(i) + j)
						}
					}
					if !tmpDownloaded && alreadyChecked {
						break
					}
				}
			} else {
				i := 0
				for {
					i++
					tmpDownloaded := download(float32(i))
					if tmpDownloaded {
						prepareOutput(float64(i))
					}
					if !tmpDownloaded && alreadyChecked {
						break
					} else if !tmpDownloaded && !alreadyChecked {
						fmt.Println("Checking for weird naming conventions...")
						for j := 0.1; j <= 0.9; j += 0.1 {
							if download(float32(float64(i) + j)) {
								prepareOutput(float64(i) + j)
							}
						}
						alreadyChecked = true
					}
				}
			}
		} else if chapterState == "multiple" {
			if special {
				tmp, _ := strconv.ParseFloat(chapterEnd, 32)
				for i, _ := strconv.ParseFloat(chapterBegin, 32); i <= tmp; i++ {
					tmpDownloaded := download(float32(i))
					if tmpDownloaded {
						prepareOutput(i)
					}
					alreadyChecked = true
					for j := i - float64(int(i)); j <= 0.9; j += 0.1 {
						//if the chapter is a whole number, it gets redownloaded, so add +0.1 to avoid that.
						if j == 0.0 {
							j = 0.1
						}
						//out of the loop if the chapter being downloaded is greater than the last chapter.
						if i+j >= tmp {
							break
						}
						if download(float32(i + j)) {
							alreadyChecked = false
							prepareOutput(i + j)
						}
					}
					if !tmpDownloaded && alreadyChecked {
						break
					}
					i = float64(int(i))
				}
				//checks if the chapterEnd (tmp) is a whole number, if it isn't, then it downloads the last chapter (which is decimal)
				if tmp != float64(int64(tmp)) {
					if download(float32(tmp)) {
						prepareOutput(tmp)
					}
				}
			} else {
				tmp, _ := strconv.ParseFloat(chapterEnd, 32)
				for i, _ := strconv.ParseFloat(chapterBegin, 32); i <= tmp; i++ {
					tmpDownloaded := download(float32(i))
					if tmpDownloaded {
						prepareOutput(i)
					}
					if !tmpDownloaded && alreadyChecked {
						break
					} else if !tmpDownloaded && !alreadyChecked {
						fmt.Println("Checking for weird naming conventions...")
						for j := 0.1; j <= 0.9; j += 0.1 {
							if download(float32(i + j)) {
								prepareOutput(i)
							}
						}
						alreadyChecked = true
					}
					i = float64(int(i))
				}
				//checks if the chapterEnd (tmp) is a whole number, if it isn't, then it downloads the last chapter (which is decimal)
				if tmp != float64(int64(tmp)) {
					if download(float32(tmp)) {
						prepareOutput(tmp)
					}
				}
			}
		} else if chapterState == "single" {
			tmp, _ := strconv.ParseFloat(singleChapter, 32)
			download(float32(tmp))
			prepareOutput(tmp)
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
