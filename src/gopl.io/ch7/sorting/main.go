package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Lenght time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return duration
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Lenght")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Lenght)
	}
	tw.Flush() // calculate column widths and print table
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (y byYear) Len() int           { return len(y) }
func (y byYear) Less(i, j int) bool { return y[i].Year < y[j].Year }
func (y byYear) Swap(i, j int)      { y[i], y[j] = y[j], y[i] }

// customSort combines a slice with a function
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	sort.Sort(byArtist(tracks))

	printTracks(tracks)

	fmt.Println()
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println()
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println()
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Lenght != y.Lenght {
			return x.Lenght < y.Lenght
		}
		return false
	}})
	printTracks(tracks)

	values := []int{3, 1, 4, 2}
	fmt.Println("\nSorted:", sort.IntsAreSorted(values))

	sort.Ints(values)
	fmt.Println(values)
	fmt.Println("Sorted:", sort.IntsAreSorted(values))

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)
	fmt.Println("Sorted:", sort.IntsAreSorted(values))

	fmt.Println()
	sort.Stable(byArtist(tracks))
	printTracks(tracks)

}
