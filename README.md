# ghost-to-md

Are you looking for a way to convert your Ghost exported JSON file into separate Markdown files for each post? Look no further! This tool will help you do just that.

## Installation
1. First things first, let's get a copy of this tool on your local machine by cloning this repository:

```bash
git clone git@github.com:yesjinu/ghost-to-md.git
```

2. Make sure you have Go installed on your machine. If not, head over to [this link](https://golang.org/doc/install) to get it set up: 

3. Now that you have Go, running this tool is as easy as pie. All you need to do is run the main.go file, passing in the path to your Ghost exported JSON file as an argument. Like this:

```bash
go run main.go YOUR_JSON_FILE.json
```

## Usage
Once you've run the code, it will parse your JSON file and create a new Markdown file for each post in the `posts/` directory. The file name will be the post's title with a ".md" extension, so you'll be able to find them easily.

For example, when you export your content from Ghost blog, your JSON file should look like [this](https://github.com/yesjinu/ghost-to-md/blob/master/sample/ghost_exported_file.json). Ghost exports the file with flatten JSON format.

```
{"db":[{"meta":{"exported_on":1673670577626,"version":"4.47.4"},"data":{"posts":[{"id":"63c22e7bb534ee0210985d41","uuid":"32785e92-196e-4e2e-8159-189195e92fcc","title":"Your Title","slug":"some-slug","mobiledoc":"{\"version\":\"0.3.1\",\"atoms\":[]...
```

After you run this code, Boom! The tool will create files like this:

```
---
title: Your Title
slug: some-slug
createdAt: 2023-01-14 04:24:27
updatedAt: 2023-01-14 04:41:30
publishedAt: 2023-01-14 04:28:17
FeatureImage: some_image_url
---
Your content

```


If you need specific kinds of frontmatter, you can modify the `mdFormat` and `toContentString` in `main.go` according to your needs. You can also create an issue and I'll try to add it.

## Contributions
I'm always looking for ways to improve this tool, so if you have any ideas or found any bugs, please don't hesitate to submit a pull request or open an issue.

## License
This project is licensed under the MIT License, so feel free to use it for whatever you need.
