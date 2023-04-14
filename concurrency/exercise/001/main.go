package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}

	m := &sync.RWMutex{}

	cacheChan := make(chan Book)
	dbChan := make(chan Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheChan)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryDatabase(id, m); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbChan)

		go func(cacheCh, dbCh <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("Source: Cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("Source: Database")
				fmt.Println(b)
			}
		}(cacheChan, dbChan)

		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()

			return b, true
		}
	}

	return Book{}, false
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished,
	)
}

var books = []Book{
	{
		ID:            1,
		Title:         "The Hitchhiker's Guide to the Galaxy",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},
	{
		ID:            2,
		Title:         "The Hobbit",
		Author:        "J.R.R Tolkein",
		YearPublished: 1937,
	},
	{
		ID:            3,
		Title:         "A Tale of Two Cities",
		Author:        "Charles Dickins",
		YearPublished: 1859,
	},
	{
		ID:            4,
		Title:         "Harry Potter and the Philosophers Stone",
		Author:        "J.K. Rowling",
		YearPublished: 1997,
	},
	{
		ID:            5,
		Title:         "Les Miserables",
		Author:        "Victor Hugo",
		YearPublished: 1862,
	},
	{
		ID:            6,
		Title:         "I, Robot",
		Author:        "Isaac Asamov",
		YearPublished: 1950,
	},
	{
		ID:            7,
		Title:         "The Gods Themselves",
		Author:        "Isaac Asamov",
		YearPublished: 1973,
	},
	{
		ID:            8,
		Title:         "The Moond is a Hash Mistress",
		Author:        "Robert A. Heinlein",
		YearPublished: 1966,
	},
	{
		ID:            9,
		Title:         "On Basilisk Station",
		Author:        "David Weber",
		YearPublished: 1993,
	},
	{
		ID:            10,
		Title:         "The Android's Dream",
		Author:        "John Scalzi",
		YearPublished: 2006,
	},
}
