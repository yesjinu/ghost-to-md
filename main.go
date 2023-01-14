package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

type PostType struct {
	ID           string `json:"id"`
	Uuid         string `json:"uuid"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	MobileDoc    string `json:"mobiledoc"`
	Html         string `json:"html"`
	CommentID    string `json:"comment_id"`
	Plaintext    string `json:"plaintext"`
	FeatureImage string `json:"feature_image"`
	Featured     int    `json:"featured"`
	Status       string `json:"status"`
	Locale       string `json:"locale"`
	Visibility   string `json:"visibility"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	PublishedAt  string `json:"published_at"`
}

type DataType struct {
	Posts []PostType `json:"posts"`

	/*
	 * NOTE: data below is not used for now
	 */

	// PostsMeta []PostMetaType `json:"posts_meta"`
	// Users []UserType `json:"users"`
	// PostsAuthors []PostAuthorType `json:"posts_authors"`
	// Roles []RoleType `json:"roles"`
	// RolesUsers []RoleUserType `json:"roles_users"`
}

type MyDB struct {
	Data DataType `json:"data"`
}

type GhostExportedFileFormat struct {
	DB []MyDB `json:"db"`
}

func extractFormat() string {
	return `---
title: %s
slug: %s
createdAt: %s
updatedAt: %s
publishedAt: %s
FeatureImage: %s
---
%s
`
}

func html2md(html string) string {
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}

	return markdown
}

func RemoveSpecialCharacters(str string) string {
	reg, _ := regexp.Compile("[^a-zA-Zㄱ-ㅎ가-힣0-9]+")
	newStr := reg.ReplaceAllString(str, "-")
	return newStr
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide a JSON file to convert")
		os.Exit(1)
	}

	file, _ := ioutil.ReadFile(os.Args[1])
	var data GhostExportedFileFormat
	json.Unmarshal(file, &data)

	dirName := "posts"
	os.Mkdir(dirName, os.ModePerm)

	var posts = data.DB[0].Data.Posts
	for _, post := range posts {
		fileName := fmt.Sprintf("%s/%s.md", dirName, RemoveSpecialCharacters(post.Title))
		content := fmt.Sprintf(extractFormat(),
			post.Title,
			post.Slug,
			post.CreatedAt,
			post.UpdatedAt,
			post.PublishedAt,
			post.FeatureImage,
			html2md(post.Html),
		)
		ioutil.WriteFile(fileName, []byte(content), 0644)
	}
}
