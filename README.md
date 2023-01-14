# ghost-to-md

Convert your Ghost exported JSON file into separate Markdown files for each post.

## Installation
1. Clone this repository to your local machine

```
git@github.com:yesjinu/ghost-to-md.git
```

2. Install Go. You can download at [HERE](https://golang.org/doc/install)

3. Run the main.go file, passing in the path to your Ghost exported JSON file as an argument

```go
go run main.go YOUR_JSON_FILE.json
```

## Usage
The tool will parse your JSON file and create a new Markdown file for each post in the same directory as the JSON file. The file name will be the post's title with a ".md" extension.

If you need specific kinds of frontmatter, you can modify the `mdFormat` and `toContentString` in `main.go` according to your needs. You can also create an issue and I'll try to add it.

## Contributions
Feel free to submit pull requests or issues if you have any improvements or bug fixes.

## License
This project is licensed under the MIT License