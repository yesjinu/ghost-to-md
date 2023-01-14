package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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

type GhostExportedFileFormat struct { // TODO: sync with Ghost export format
	DB []MyDB `json:"db"`
}

func extractFormat() string {
	return `Id: %s
Uuid: %s
Title: %s
Slug: %s
MobileDoc: %s
Html: %s
CommentID: %s
Plaintext: %s
FeatureImage: %s
Featured: %d
Status: %s
Locale: %s
Visibility: %s
CreatedAt: %s
UpdatedAt: %s
PublishedAt: %s
`
}

func RemoveSpecialCharacters(str string) string {
	reg, _ := regexp.Compile("[^a-zA-Zㄱ-ㅎ가-힣0-9]+")
	newStr := reg.ReplaceAllString(str, "-")
	return newStr
}

func main() {
	file, _ := ioutil.ReadFile("config.json") // TODO: get from cli args
	var data GhostExportedFileFormat
	json.Unmarshal(file, &data)

	dirName := "posts"
	os.Mkdir(dirName, os.ModePerm)

	var posts = data.DB[0].Data.Posts
	for _, post := range posts {
		fileName := fmt.Sprintf("%s/%s.md", dirName, RemoveSpecialCharacters(post.Title))
		content := fmt.Sprintf(extractFormat(), post.ID, post.Uuid, post.Title, post.Slug, post.MobileDoc, post.Html, post.CommentID, post.Plaintext, post.FeatureImage, post.Featured, post.Status, post.Locale, post.Visibility, post.CreatedAt, post.UpdatedAt, post.PublishedAt)
		ioutil.WriteFile(fileName, []byte(content), 0644)
	}
}
