// Встраивание: CachingFeed со встроенным *Feed и кэшем документов
package main

import (
	"fmt"
	"log"
	"time"
)

type Document struct {
	Key   string
	Title string
}

type Feed struct{}

func (f *Feed) Count() int {
	return 42
}

func (f *Feed) Fetch(key string) (Document, error) {
	time.Sleep(time.Second)

	doc := Document{
		Key:   key,
		Title: "Title for " + key,
	}
	return doc, nil
}

type FetchCounter interface {
	Fetch(key string) (Document, error)
	Count() int
}

func process(fc FetchCounter) {
	fmt.Printf("There are %d documents\n", fc.Count())

	keys := []string{"a", "a", "a", "b", "b", "b"}

	for _, key := range keys {
		doc, err := fc.Fetch(key)
		if err != nil {
			log.Printf("Could not fetch %s : %v", key, err)
			return
		}

		fmt.Printf("%s : %v\n", key, doc)
	}
}

type CachingFeed struct {
	*Feed
	docs map[string]Document
}

func NewCachingFeed(f *Feed) *CachingFeed {
	return &CachingFeed{
		Feed: f,
		docs: make(map[string]Document),
	}
}

func (cf *CachingFeed) Fetch(key string) (Document, error) {
	if doc, ok := cf.docs[key]; ok {
		return doc, nil
	}

	doc, err := cf.Feed.Fetch(key)
	if err != nil {
		return Document{}, err
	}

	cf.docs[key] = doc
	return doc, nil
}

func main() {
	fmt.Println("Using Feed directly")
	start := time.Now()
	process(&Feed{})
	fmt.Println("Время:", time.Since(start).Round(time.Second))

	fmt.Println("Using CachingFeed")
	start = time.Now()
	process(NewCachingFeed(&Feed{}))
	fmt.Println("Время:", time.Since(start).Round(time.Second))
}
