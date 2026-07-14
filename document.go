package htmlbuilder

// Document represents a full HTML page: doctype, <html lang="...">,
// <head>, and <body>.
type Document struct {
	lang string
	head *Node
	body *Node
}

// New creates a new Document with sensible defaults: lang="en",
// empty <head> and <body>.
func New() *Document {
	return &Document{
		lang: "en",
		head: El("head"),
		body: El("body"),
	}
}

// StyleBlock appends a raw <style>...</style> tag to <head>. Use this
// for :root custom properties, @font-face, @import, or any CSS you
// want embedded directly rather than linked externally.
//
// Usage:
//
//	doc.StyleBlock(`
//	  :root {
//	    --primary: #1a1a1a;
//	    --font-body: "Inter", sans-serif;
//	  }
//	`)
func (d *Document) StyleBlock(css string) *Document {
	d.head.Child(El("style").Raw(css))
	return d
}

// Icon appends a <link rel="icon"> tag to <head>. Covers favicons,
// PNG icons, and Flaticon/CDN-hosted icon assets alike — just a URL
// and optional type/sizes.
//
// Usage:
//
//	doc.Icon("/favicon.ico", "")
//	doc.Icon("https://cdn-icons-png.flaticon.com/512/xxx/xxx.png", "image/png")
func (d *Document) Icon(href, mimeType string) *Document {
	l := El("link").Attr("rel", "icon").Attr("href", href)
	if mimeType != "" {
		l.Attr("type", mimeType)
	}
	d.head.Child(l)
	return d
}

// Script appends a raw <script>...</script> tag to <head> (or call
// on doc.Body() equivalent if you want it at the end of body instead —
// this version targets head).
func (d *Document) Script(js string) *Document {
	d.head.Child(El("script").Raw(js))
	return d
}

// Lang sets the <html lang="..."> attribute.
func (d *Document) Lang(l string) *Document {
	d.lang = l
	return d
}

// Title sets the <title> tag inside <head>. Repeated calls replace
// the previous title rather than appending a second one.
func (d *Document) Title(t string) *Document {
	for _, c := range d.head.children {
		if c.tag == "title" {
			c.text = t
			return d
		}
	}
	d.head.Child(El("title").Text(t))
	return d
}

// Meta appends a <meta> tag to <head> with the given attributes.
// Usage: doc.Meta(map[string]string{"charset": "utf-8"})
func (d *Document) Meta(attrs map[string]string) *Document {
	m := El("meta")
	for k, v := range attrs {
		m.Attr(k, v)
	}
	d.head.Child(m)
	return d
}

// Link appends a <link> tag to <head>, e.g. for stylesheets.
// Usage: doc.Link(map[string]string{"rel": "stylesheet", "href": "/style.css"})
func (d *Document) Link(attrs map[string]string) *Document {
	l := El("link")
	for k, v := range attrs {
		l.Attr(k, v)
	}
	d.head.Child(l)
	return d
}

// Head returns the <head> node for direct manipulation.
func (d *Document) Head() *Node {
	return d.head
}

// Body returns the <body> node so callers can .Child(...) into it.
func (d *Document) Body() *Node {
	return d.body
}
