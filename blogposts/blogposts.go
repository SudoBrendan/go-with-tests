package blogposts

import (
	"bufio"
	"bytes"
	"io"
	"io/fs"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Post struct {
	Title       string   `yaml:"Title"`
	Description string   `yaml:"Description"`
	Tags        []string `yaml:"Tags"`
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := parsePostFromFile(fileSystem, f)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func parsePostFromFile(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}
	post, err := parsePostData(postData)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

func parsePostData(postData []byte) (Post, error) {
	var post Post
	post, err := parseHeader(postData, post)
	if err != nil {
		return Post{}, err
	}
	post, err = parseBody(postData, post)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

func parseHeader(postData []byte, post Post) (Post, error) {
	err := yaml.UnmarshalStrict(postData, &post)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

func parseBody(postData []byte, post Post) (Post, error) {
	scanner := bufio.NewScanner(bytes.NewReader(postData))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			break
		}
	}
	var sb strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		sb.WriteString(line)
		sb.WriteString("\n")
	}
	post.Body = strings.TrimSuffix(sb.String(), "\n")
	return post, nil
}
