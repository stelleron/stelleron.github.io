package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/pelletier/go-toml/v2"
)

// Post type enum
type PostType int

const (
	SpecialPost PostType = iota
	BlogPost
	ProjectPost
	HomepagePost
	BlogHomePost
)

// Structs
type SidebarData struct {
	Github   string
	Linkedin string
}

type SiteConfig struct {
	Name        string
	Username    string
	Pronouns    string
	Description string
	ProfilePic  string
	Sidebar     SidebarData
}

type MarkdownFile struct {
	FileName string
	FileText string
}

type ProjectFolder struct {
	SourceDir      string
	DestinationDir string
	ProjectType    PostType
	MarkdownFiles  []MarkdownFile
}

type Frontmatter struct {
	FileName    string
	Title       string
	Date        string
	Description string
	SortDate    time.Time
}

type ProjectFrontmatter struct {
	FileName     string
	Title        string
	Date         string
	Description  string
	Link         string
	ProjectImage string
	SortDate     time.Time
}

type FrontmatterList []Frontmatter
type ProjectFrontmatterList []ProjectFrontmatter

func (f FrontmatterList) Len() int           { return len(f) }
func (f FrontmatterList) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f FrontmatterList) Less(i, j int) bool { return f[i].SortDate.After(f[j].SortDate) }

func (f ProjectFrontmatterList) Len() int           { return len(f) }
func (f ProjectFrontmatterList) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ProjectFrontmatterList) Less(i, j int) bool { return f[i].SortDate.After(f[j].SortDate) }

// Global variables
var blog_name string
var pfp_path string
var username string
var pronouns string
var description string
var links_sidebar string
var blogposts_data FrontmatterList
var projects_data ProjectFrontmatterList

// Constants
const AnimDelayIncrement = 0.15
const SiteBasePath = "site/"
const SiteBlogPath = "site/blog/"
const BlogPathForLinks = "/site/blog/"
const SiteProjectPath = "site/projects/"
const ProjectPathForLinks = "/site/projects/"
const GitHubSidebar = "<a href=\"%s\"><img src=\"/images/base/github-mark.svg\" class=\"icon\" width=\"32\" height=\"32\"></a>"
const LinkedInSidebar = "<a href=\"%s\"><img src=\"/images/base/linkedin-mark.svg\" class=\"icon\" width=\"32\" height=\"32\"></a>"
const HtmlTemplate = `<!DOCTYPE html>
<head>
    <title> %s </title>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css"/>
	<link rel="stylesheet" href="/style/style.css">
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.3.1/styles/default.min.css">
	<link rel="preconnect" href="https://fonts.googleapis.com">
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
	<link href="https://fonts.googleapis.com/css2?family=PT+Serif&display=swap" rel="stylesheet">
	<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.3.1/highlight.min.js"></script>
</head>
<body>
	<div class="sidebar">
		<div class="sidebar-cont">
		%s
		</div>
	</div>
	<div class="content">
		%s
	</div>
	<script>
		hljs.highlightAll();
	</script>
</body>
`
const HtmlBlogpostTemplate = `
<div class="index-post" style="animation-delay: %fs">
<a href="%s" class="index-post-title">%s</a>
<div class="index-post-date">%s</div>
<p class="index-post-desc">%s</p>
%s
</div>
`

const HtmlProjectTemplate = `
<div class="index-post" style="animation-delay: %fs">
<a href="%s" class="index-post-title">%s</a>
<div class="index-post-date">%s</div>
%s
<p class="index-post-desc">%s</p>
%s
</div>
`

func assemble_sidebar() string {
	sidebar := ""
	// Profile pic
	sidebar += fmt.Sprintf("<img class=\"profile-pic\" src=\" %s \">", pfp_path)
	// Poster's name
	sidebar += fmt.Sprintf("<h1 class=\"name\"> %s </h1>", blog_name)
	// Poster's username and pronouns
	sidebar += "<div>"
	sidebar += fmt.Sprintf("<span class=\"username\"> %s </span>", username)
	sidebar += "<span class=\"username\"> · </span>"
	sidebar += fmt.Sprintf("<span class=\"pronouns\"> %s </span>", pronouns)
	sidebar += "</div>"
	// Poster's description
	sidebar += fmt.Sprintf("<p class=\"description\"> %s </p>", description)
	// Links to other parts of the page
	sidebar += "<p class=\"sidebar-link-cont\"><a href=\"/\" class=\"sidebar-links\">Home</a></p>"
	sidebar += "<p class=\"sidebar-link-cont\"><a href=\"/site/about.html\" class=\"sidebar-links\">About</a></p>"
	sidebar += "<p class=\"sidebar-link-cont\"><a href=\"/site/resume.html\" class=\"sidebar-links\">Resume</a></p>"
	sidebar += "<p class=\"sidebar-link-cont\"><a href=\"/site/blog.html\" class=\"sidebar-links\">Blog</a></p>"
	// Add links
	sidebar += links_sidebar
	// Add UMA Ring
	sidebar += "<script id=\"umaring_js\" src=\"https://umaring.mkr.cx/ring.js?id=stelleron\"></script><div id=\"umaring\"></div>"
	return sidebar
}

func assemble_webpage(page_title string, content string) string {
	return fmt.Sprintf(HtmlTemplate, page_title, assemble_sidebar(), content)
}

func process_md_file(md_file MarkdownFile) (string, Frontmatter) {
	// First find the frontmatter data
	md_data := strings.Replace(md_file.FileText, "\n", "", -1)
	frontmatter_data := regexp.MustCompile(`===(.*)===`).FindStringSubmatch(md_data)[1]

	// Then parse it
	title_loc := strings.Index(frontmatter_data, "title")
	date_loc := strings.Index(frontmatter_data, "date")
	description_loc := strings.Index(frontmatter_data, "description")

	frontmatter_obj := Frontmatter{
		FileName:    md_file.FileName,
		Title:       frontmatter_data[title_loc+len("title:")+1 : date_loc],
		Date:        frontmatter_data[date_loc+len("date:")+1 : description_loc],
		Description: frontmatter_data[description_loc+len("description:"):],
	}

	var err error
	frontmatter_obj.SortDate, err = time.Parse("01-02-2006", frontmatter_obj.Date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}

	end_ptr := strings.Index(strings.Replace(md_file.FileText, "===", "xxx", 1), "===")

	return md_file.FileText[end_ptr+len("==="):], frontmatter_obj
}

func process_project_md_file(md_file MarkdownFile) (string, ProjectFrontmatter) {
	// First find the frontmatter data
	md_data := strings.Replace(md_file.FileText, "\n", "", -1)
	frontmatter_data := regexp.MustCompile(`===(.*)===`).FindStringSubmatch(md_data)[1]

	// Then parse it
	title_loc := strings.Index(frontmatter_data, "title")
	date_loc := strings.Index(frontmatter_data, "date")
	description_loc := strings.Index(frontmatter_data, "description")
	link_loc := strings.Index(frontmatter_data, "link")
	project_img_loc := strings.Index(frontmatter_data, "project-image")

	frontmatter_obj := ProjectFrontmatter{
		FileName:     md_file.FileName,
		Title:        frontmatter_data[title_loc+len("title:")+1 : date_loc],
		Date:         frontmatter_data[date_loc+len("date:")+1 : description_loc],
		Description:  frontmatter_data[description_loc+len("description:")+1 : link_loc],
		Link:         frontmatter_data[link_loc+len("link:")+1 : project_img_loc],
		ProjectImage: frontmatter_data[project_img_loc+len("project-image:")+1:],
	}

	if frontmatter_obj.ProjectImage != "None" {
		frontmatter_obj.ProjectImage = fmt.Sprintf("<img src=\"images/%s\" width=\"450\" height=\"300\">", frontmatter_obj.ProjectImage)
	} else {
		frontmatter_obj.ProjectImage = ""
	}

	var err error
	frontmatter_obj.SortDate, err = time.Parse("01-02-2006", frontmatter_obj.Date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}

	end_ptr := strings.Index(strings.Replace(md_file.FileText, "===", "xxx", 1), "===")

	return md_file.FileText[end_ptr+len("==="):], frontmatter_obj
}

func generate_sidebar(cfg SiteConfig) {
	// Look for a GitHub and LinkedIn link
	if cfg.Sidebar.Github != "" {
		github_logo := GitHubSidebar
		links_sidebar += fmt.Sprintf(github_logo, cfg.Sidebar.Github)
		links_sidebar += "\n"
	}
	if cfg.Sidebar.Linkedin != "" {
		linkedin_logo := LinkedInSidebar
		links_sidebar += fmt.Sprintf(linkedin_logo, cfg.Sidebar.Linkedin)
	}

	links_sidebar = fmt.Sprintf("<div class=\"footer\">%s</div>", links_sidebar)
}

func load_blog_pages(source_dir string, dest_dir string, post_type PostType) ProjectFolder {
	var project_folder ProjectFolder

	dir, err := os.Open(source_dir)
	if err != nil {
		fmt.Println("Error opening directory:", err)
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(0)
	if err != nil {
		fmt.Println("Error reading directory:", err)
	}

	for _, fileInfo := range fileInfos {
		// Check if it's a regular file
		if fileInfo.Mode().IsRegular() {
			filePath := filepath.Join(source_dir, fileInfo.Name())

			// Read and print the file content
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
			}

			project_folder.MarkdownFiles = append(project_folder.MarkdownFiles, MarkdownFile{strings.Replace(fileInfo.Name(), "md", "html", -1), string(fileData)})
		}
	}

	project_folder.SourceDir = source_dir
	project_folder.DestinationDir = dest_dir
	project_folder.ProjectType = post_type

	return project_folder
}

func md_to_html(md_file MarkdownFile, p_type PostType) string {
	md_str, md_frontmatter := process_md_file(md_file)

	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md_str))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	if p_type == BlogPost {
		blogposts_data = append(blogposts_data, md_frontmatter)
	}
	return assemble_webpage(blog_name+" · "+md_frontmatter.Title, string(markdown.Render(doc, renderer)))
}

func publish_folder(p_folder ProjectFolder) {
	for _, fileData := range p_folder.MarkdownFiles {
		html_file := md_to_html(fileData, p_folder.ProjectType)
		err := os.WriteFile(filepath.Join(p_folder.DestinationDir, fileData.FileName), []byte(html_file), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
			return
		}
	}

}

func generate_blog_homepage() {
	sort.Sort(blogposts_data)
	html_data := "<h1>My Blog</h1>"
	anim_delay := 0.0
	for x, blogpost := range blogposts_data {
		var blogpost_html string
		if x < len(blogposts_data)-1 {
			blogpost_html = fmt.Sprintf(HtmlBlogpostTemplate, anim_delay, path.Join(BlogPathForLinks, blogpost.FileName), blogpost.Title, blogpost.Date, blogpost.Description, "<hr class=\"index-post-hr\">")
		} else {
			blogpost_html = fmt.Sprintf(HtmlBlogpostTemplate, anim_delay, path.Join(BlogPathForLinks, blogpost.FileName), blogpost.Title, blogpost.Date, blogpost.Description, "")
		}
		html_data += blogpost_html
		anim_delay += AnimDelayIncrement
	}
	err := os.WriteFile("site/blog.html", []byte(assemble_webpage(blog_name+" · Blog", html_data)), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
		return
	}
}

func generate_projects_homepage(p_folder ProjectFolder) {
	for _, fileData := range p_folder.MarkdownFiles {
		_, md_frontmatter := process_project_md_file(fileData)
		projects_data = append(projects_data, md_frontmatter)
	}

	sort.Sort(projects_data)
	html_data := "<h1>My Projects</h1>"
	anim_delay := 0.0
	for x, project := range projects_data {
		var blogpost_html string
		if x < len(projects_data)-1 {
			blogpost_html = fmt.Sprintf(HtmlProjectTemplate, anim_delay, project.Link, project.Title, project.Date, project.ProjectImage, project.Description, "<hr class=\"index-post-hr\">")
		} else {
			blogpost_html = fmt.Sprintf(HtmlProjectTemplate, anim_delay, project.Link, project.Title, project.Date, project.ProjectImage, project.Description, "")
		}
		html_data += blogpost_html
		anim_delay += AnimDelayIncrement
	}
	err := os.WriteFile("index.html", []byte(assemble_webpage(blog_name+" · Homepage", html_data)), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
		return
	}
}

func main() {
	fmt.Println("Building site...")

	// Get config.toml
	config_txt, err := os.ReadFile("config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Load the config.toml
	var cfg SiteConfig
	err = toml.Unmarshal([]byte(config_txt), &cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing config: %v\n", err)
		os.Exit(1)
	}

	// Get some config data
	blog_name = cfg.Name
	username = cfg.Username
	pronouns = cfg.Pronouns
	description = cfg.Description
	pfp_path = cfg.ProfilePic

	// Generate a sidebar of links
	generate_sidebar(cfg)

	// Then iterate through and convert the posts to HTML
	// == First convert the about and resume pages
	{
		special_pages := load_blog_pages("posts/special/", SiteBasePath, SpecialPost)
		publish_folder(special_pages)
	}
	{
		blogpost_pages := load_blog_pages("posts/blog/", SiteBlogPath, BlogPost)
		publish_folder(blogpost_pages)
		generate_blog_homepage()
	}
	{
		projects_pages := load_blog_pages("posts/projects/", SiteProjectPath, ProjectPost)
		generate_projects_homepage(projects_pages)
	}
}
