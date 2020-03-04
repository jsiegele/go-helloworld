package main

import (
	"errors"
	"fmt"
	"time"
)

// Post struct
type Post struct {
	body        string
	publishDate int64 // Unix timestamp
	next        *Post
}

// Feed struct
type Feed struct {
	length int // we'll use it later
	start  *Post
	end    *Post
}

// Append new Post to List
func (f *Feed) Append(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
		f.end = newPost
	} else {
		lastPost := f.end
		lastPost.next = newPost
		f.end = newPost
	}
	f.length++
}

// Remove Element from List
func (f *Feed) Remove(publishDate int64) {
	if f.length == 0 {
		panic(errors.New("Feed is empty"))
	}

	var previousPost *Post
	currentPost := f.start

	for currentPost.publishDate != publishDate {
		if currentPost.next == nil {
			panic(errors.New("no such post found"))
		}

		previousPost = currentPost
		currentPost = currentPost.next
	}
	previousPost.next = currentPost.next

	f.length--
}

// Insert Element
func (f *Feed) Insert(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
	} else {
		var previousPost *Post
		currentPost := f.start

		for currentPost.publishDate < newPost.publishDate {
			previousPost = currentPost
			currentPost = previousPost.next
		}

		previousPost.next = newPost
		newPost.next = currentPost
	}
	f.length++
}

// Inspect Feed
func (f *Feed) Inspect() {
	if f.length == 0 {
		fmt.Println("Feed is empty")
	}
	fmt.Println("========================")
	fmt.Println("Feed Length: ", f.length)

	currentIndex := 0
	currentPost := f.start

	for currentIndex < f.length {
		fmt.Printf("Item: %v - %v - %v - %v\n", currentIndex, currentPost.body, currentPost.publishDate, currentPost.next)
		currentPost = currentPost.next
		currentIndex++
	}
	fmt.Println("========================")
}

// Merge Feeds
func (f *Feed) Merge(u *Feed) {
	if f.length == 0 && u.length != 0 {
		f.start = u.start
		return
	} else if u.length == 0 && f.length != 0 {
		return
	}

	currentIndex := 0
	currentPost := u.start

	for currentIndex < u.length {
		f.Insert(currentPost)
		currentPost = currentPost.next
		currentIndex++
	}
}

// RecursiveMerge Merge Feeds
func (f *Post) RecursiveMerge(u *Post) *Post {
	if f == nil && u != nil {
		f = u
		return f
	} else if u == nil && f != nil {
		u = f
		return u
	}

	var result *Post

	if f.publishDate <= u.publishDate {
		result = f
		f.next = f.next.RecursiveMerge(u)
	} else {
		result = u
		u.next = f.RecursiveMerge(u.next)
	}
	return result
}

func main() {
	rightNow := time.Now().Unix()
	f := &Feed{}
	p1 := &Post{
		body:        "Lorem ipsum",
		publishDate: rightNow,
	}
	p2 := &Post{
		body:        "Dolor sit amet",
		publishDate: rightNow + 10,
	}
	p3 := &Post{
		body:        "consectetur adipiscing elit",
		publishDate: rightNow + 20,
	}
	p4 := &Post{
		body:        "sed do eiusmod tempor incididunt",
		publishDate: rightNow + 30,
	}
	f.Append(p1)
	f.Append(p2)
	f.Append(p3)
	f.Append(p4)

	f.Inspect()

	u := &Feed{}
	newPost := &Post{
		body:        "This is a new post",
		publishDate: rightNow + 15,
	}
	u.Insert(newPost)
	u.Inspect()

	// f.Merge(u)
	// f.Inspect()

	f.start.RecursiveMerge(u.start)
	// absolut nicht sauber :/
	f.length = f.length + u.length
	f.Inspect()
}
