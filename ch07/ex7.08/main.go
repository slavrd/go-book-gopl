package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

func main() {

	data := []*Data{
		{Col1: "c", Col2: "a", Col3: 5},
		{Col1: "a", Col2: "b", Col3: 2},
		{Col1: "a", Col2: "b", Col3: 1},
		{Col1: "b", Col2: "a", Col3: 4},
		{Col1: "c", Col2: "b", Col3: 2},
	}

	sdata := MultiTearSort{data: data}

	fmt.Println("Initial Data:")
	printData(data)

	updateSortOrder(&sdata, "Col3")
	sort.Sort(sdata)
	fmt.Println("Sorted data")
	fmt.Printf("Sort order: ")
	printSortOrder(&sdata)
	fmt.Println()
	printData(data)

	updateSortOrder(&sdata, "Col1")
	sort.Sort(sdata)
	fmt.Println("Sorted data")
	fmt.Printf("Sort order: ")
	printSortOrder(&sdata)
	fmt.Println()
	printData(data)

	updateSortOrder(&sdata, "Col2")
	sort.Sort(sdata)
	fmt.Println("Sorted data")
	fmt.Printf("Sort order: ")
	printSortOrder(&sdata)
	fmt.Println()
	printData(data)

}

type Data struct {
	Col1, Col2 string
	Col3       int
}

type DataColumns []string

func (d DataColumns) validate() bool {
	return true
}

type MultiTearSort struct {
	data         []*Data
	sFieldsOrder []string
}

func (d MultiTearSort) Len() int {
	return len(d.data)
}

func (d MultiTearSort) Swap(i, j int) {
	d.data[i], d.data[j] = d.data[j], d.data[i]
}

func (d MultiTearSort) Less(i, j int) bool {
	for k := len(d.sFieldsOrder) - 1; k >= 0; k-- {

		switch d.sFieldsOrder[k] {
		case "Col1":
			if d.data[i].Col1 < d.data[j].Col1 {
				return true
			} else if d.data[i].Col1 > d.data[j].Col1 {
				return false
			}
		case "Col2":
			if d.data[i].Col2 < d.data[j].Col2 {
				return true
			} else if d.data[i].Col2 > d.data[j].Col2 {
				return false
			}
		case "Col3":
			if d.data[i].Col3 < d.data[j].Col3 {
				return true
			} else if d.data[i].Col3 > d.data[j].Col3 {
				return false
			}
		}
	}
	return false
}

func updateSortOrder(t *MultiTearSort, c string) {
	if t.sFieldsOrder == nil {
		t.sFieldsOrder = make([]string, 0)
	}
	for i, f := range t.sFieldsOrder {
		if f == c {
			copy(t.sFieldsOrder[i:], t.sFieldsOrder[i+1:])
			t.sFieldsOrder[len(t.sFieldsOrder)-1] = c
			return
		}
	}
	t.sFieldsOrder = append(t.sFieldsOrder, c)
}

func printData(d []*Data) {

	format := "%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Col1", "Col2", "Col3")
	fmt.Fprintf(tw, format, "----", "----", "----")
	for _, i := range d {
		fmt.Fprintf(tw, format, i.Col1, i.Col2, i.Col3)
	}
	fmt.Fprintf(tw, "\n")
	tw.Flush()

}

func printSortOrder(ms *MultiTearSort) {
	for i := len(ms.sFieldsOrder) - 1; i >= 0; i-- {
		fmt.Printf("%s", ms.sFieldsOrder[i])
		if i != 0 {
			fmt.Printf(" -> ")
		}
	}
	fmt.Printf("\n")
}
