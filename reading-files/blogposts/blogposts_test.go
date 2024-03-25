package blogposts_test

import (
	"fmt"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/nachiket87/learn-go-with-tests/reading-files/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world2.md": {Data: []byte("Title: Post 2")},
		"hello-world.md":  {Data: []byte("title: Post 1")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	got := posts[0]
	fmt.Println(got)
	want := blogposts.Post{Title: "Post 1"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)

	}
}
