# HTML Builder

## About

HTML Builder is a lightweight Go package designed to construct fully featured HTML5 markup, CSS styling, and vanilla JavaScript components—all using Go syntax. Write your entire front-end structure directly in Go and generate clean, performant, and standard-compliant web layouts without leaving your Go environment. [View an Example](https://htmlbuilder-1j4g.onrender.com/)

## Features

- Build nested structures using Go syntax, catching tag and syntax errors at compile time.
- Define component styles directly alongside your markup, keeping everything in Go.
- Easily inject script blocks and event handlers to make your components interactive.
- Lightweight and self-contained, built purely on the Go standard library.
- Instantly compile your Go structures into clean, production-ready HTML5 strings.

## Usage

The [godoc](https://pkg.go.dev/github.com/Thruqe/htmlbuilder) includes docs for all methods and event types.

<details>
  <summary>Click here to view a simple code example!</summary>

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Thruqe/htmlbuilder"
)

var PORT = ":8080"

func main() {
	http.HandleFunc("/", handler)
	log.Println("listening on http://localhost:8080")
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, simple_landing_page())
}

func simple_landing_page() string {
	doc := htmlbuilder.New().
		Title("Simple Landing Page").
	    MetaDefault().
		Link(map[string]string{
			"rel":  "stylesheet",
			"href": "https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap",
		}).
		Link(map[string]string{
			"rel":  "stylesheet",
			"href": "https://cdn-uicons.flaticon.com/2.6.0/uicons-regular-rounded/css/uicons-regular-rounded.css",
		}).
		// Cleaned up mobile responsive layout math
		StyleBlock(`
			html {
				scroll-behavior: smooth;
			}

			/* Only enable snap-scrolling on large screens where content comfortably fits */
			@media (min-width: 769px) {
				html {
					scroll-snap-type: y mandatory;
				}
				section, footer {
					scroll-snap-align: start;
					scroll-snap-stop: always;
				}
			}

			/* Smooth-slide & Fade-in/out Dropdown Animation */
			.mobile-menu {
				display: flex;
				flex-direction: column;
				align-items: center;
				justify-content: center;
				gap: 1.25rem;
				position: fixed;
				top: 65px;
				left: 0;
				right: 0;
				background: rgba(248, 250, 252, 0.96);
				border-bottom: 1px solid rgba(226, 232, 240, 0);
				padding: 0;
				z-index: 999;
				backdrop-filter: blur(16px);
				-webkit-backdrop-filter: blur(16px);
				max-height: 0;
				opacity: 0;
				overflow: hidden;
				transition: max-height 0.35s cubic-bezier(0.4, 0, 0.2, 1),
				            opacity 0.25s ease,
				            padding 0.35s ease,
				            border-color 0.35s ease;
			}

			.mobile-menu.active {
				max-height: 220px;
				opacity: 1;
				padding: 2rem 1.5rem;
				border-color: rgba(226, 232, 240, 1);
			}

			/* Mobile Tweaks */
			@media (max-width: 768px) {
				.desktop-links {
					display: none !important;
				}
				.hamburger-btn {
					display: flex !important;
				}
				h1 {
					font-size: 2.25rem !important; /* Smaller headings so they don't break wraps */
				}
				h2 {
					font-size: 1.75rem !important;
					margin-bottom: 1.5rem !important;
				}
				.features-grid {
					grid-template-columns: 1fr !important;
					gap: 1rem !important;
				}
				section {
					padding: 6rem 1rem 3rem !important; /* Extra breathing padding at top for the fixed header */
					height: auto !important;           /* Reset strict height to avoid cutoffs */
					min-height: 100vh !important;       /* Fallback layout target */
				}
			}
		`)

	// Global Reset & Light Theme Styles
	doc.Body().CSS(htmlbuilder.Style{
		Margin:     "0",
		Padding:    "0",
		FontFamily: "'Plus Jakarta Sans', sans-serif",
		Background: "#f8fafc",
		Color:      "#0f172a",
	})

	// Hamburger Menu Button
	hamburgerBtn := htmlbuilder.El("button").
		Class("hamburger-btn").
		Attr("aria-label", "Toggle Menu").
		Attr("onclick", "document.getElementById('mobile-dropdown').classList.toggle('active')").
		Child(
			htmlbuilder.El("i").Class("fi", "fi-rr-menu-burger").CSS(htmlbuilder.Style{
				FontSize: "1.25rem",
				Color:    "#0f172a",
			}),
		).
		CSS(htmlbuilder.Style{
			Display:    "none",
			Background: "none",
			Border:     "none",
			Cursor:     "pointer",
			Padding:    "4px",
		})

	// 1. Navigation / Header Component
	navBar := htmlbuilder.El("nav").Child(
		htmlbuilder.Span("GoHTML").CSS(htmlbuilder.Style{
			FontWeight: "800",
			FontSize:   "1.4rem",
			Color:      "#087ea4",
		}),
		htmlbuilder.El("div").Class("desktop-links").Child(
			htmlbuilder.A("Home").Attr("href", "#home").CSS(navLinkStyle()),
			htmlbuilder.A("Features").Attr("href", "#features").CSS(navLinkStyle()),
			htmlbuilder.A("GitHub").Attr("href", "https://github.com/Thruqe/htmlbuilder").CSS(navLinkStyle()),
		).CSS(htmlbuilder.Style{
			Display: "flex",
			Gap:     "1.5rem",
		}),
		hamburgerBtn,
	).CSS(htmlbuilder.Style{
		Display:        "flex",
		JustifyContent: "space-between",
		AlignItems:     "center",
		Padding:        "1.25rem 2rem",
		MaxWidth:       "1200px",
		Margin:         "0 auto",
		Width:          "100%",
		BoxSizing:      "border-box",
		Position:       "fixed",
		Top:            "0",
		Left:           "0",
		Right:          "0",
		ZIndex:         "1000",
		Background:     "rgba(248, 250, 252, 0.8)",
		BorderBottom:   "1px solid rgba(226, 232, 240, 0.8)",
	}).SetStyle("backdrop-filter", "blur(12px)").
		SetStyle("-webkit-backdrop-filter", "blur(12px)")

		// Mobile Dropdown Menu Component
	mobileDropdown := htmlbuilder.El("div").
		Class("mobile-menu").
		Attr("id", "mobile-dropdown").
		Child(
			htmlbuilder.A("Home").
				Attr("href", "#home").
				Attr("onclick", "document.getElementById('mobile-dropdown').classList.remove('active')").
				CSS(mobileNavLinkStyle()). // Sets size, decoration, etc.
				// Set the base color using custom style properties (so they don't block hover)
				SetStyle("color", "inherit").
				Hover(htmlbuilder.Style{Color: "#087ea4"}),
			htmlbuilder.A("Features").
				Attr("href", "#features").
				Attr("onclick", "document.getElementById('mobile-dropdown').classList.remove('active')").
				CSS(mobileNavLinkStyle()).
				SetStyle("color", "inherit").
				Hover(htmlbuilder.Style{Color: "#087ea4"}),
			htmlbuilder.A("GitHub").
				Attr("href", "https://github.com/Thruqe/htmlbuilder").
				CSS(mobileNavLinkStyle()).
				SetStyle("color", "inherit").
				Hover(htmlbuilder.Style{Color: "#087ea4"}),
		)

	// 2. Hero Section
	heroSection := htmlbuilder.El("section").
		Attr("id", "home").
		Child(
			htmlbuilder.El("div").Child(
				htmlbuilder.H1("Pure Go. Zero Templates.").
					CSS(htmlbuilder.Style{
						FontSize:   "3.5rem",
						FontWeight: "800",
						LineHeight: "1.2",
						Margin:     "0 0 1.5rem",
						Background: "linear-gradient(to right, #087ea4, #149eca)",
					}).
					SetStyle("-webkit-background-clip", "text").
					SetStyle("-webkit-text-fill-color", "transparent"),
				htmlbuilder.P("Build beautifully structured, type-safe web pages straight from your Go backend without touching a single raw template file.").CSS(htmlbuilder.Style{
					Color:      "#475569",
					FontSize:   "1.2rem",
					LineHeight: "1.6",
					Margin:     "0 auto 2.5rem",
					MaxWidth:   "620px",
					FontWeight: "500",
				}),
				htmlbuilder.A("Get Started").
					Attr("href", "#features").
					CSS(htmlbuilder.Style{
						Background:     "#087ea4",
						Color:          "#ffffff",
						Padding:        "0.8rem 2rem",
						BorderRadius:   "10px",
						FontWeight:     "600",
						TextDecoration: "none",
						Transition:     "all 0.2s ease",
						BoxShadow:      "0 4px 12px rgba(8, 126, 164, 0.15)",
					}).
					Hover(htmlbuilder.Style{
						Transform: "translateY(-2px)",
						BoxShadow: "0 6px 20px rgba(8, 126, 164, 0.25)",
					}),
			),
		).CSS(htmlbuilder.Style{
		Display:        "flex",
		AlignItems:     "center",
		JustifyContent: "center",
		Height:         "100vh",
		MinHeight:      "100vh",
		BoxSizing:      "border-box",
		Padding:        "0 1.5rem",
	})

	// 3. Features Section
	featuresData := []struct {
		title string
		desc  string
	}{
		{"Declarative API", "Define UI structures cleanly with nested Go functions instead of parsing HTML fragments at runtime."},
		{"Type-Safe CSS", "Construct inline designs leveraging typed style structures, reducing property typos and layout errors."},
		{"Zero Dependencies", "Compiles down instantly to your Go binary, avoiding bloated external toolchains or node modules."},
	}

	featuresSection := htmlbuilder.El("section").
		Attr("id", "features").
		Child(
			htmlbuilder.El("div").Child(
				htmlbuilder.H2("Everything you need, built natively").CSS(htmlbuilder.Style{
					FontSize:   "2.25rem",
					FontWeight: "800",
					TextAlign:  "center",
					Margin:     "0 0 3rem",
					Color:      "#0f172a",
				}),
				htmlbuilder.El("div").Class("features-grid").Child(
					htmlbuilder.Each(featuresData, func(f struct {
						title string
						desc  string
					}) *htmlbuilder.Node {
						return htmlbuilder.El("div").Child(
							htmlbuilder.H3(f.title).CSS(htmlbuilder.Style{
								FontSize:   "1.2rem",
								FontWeight: "700",
								Margin:     "0 0 0.75rem",
								Color:      "#0f172a",
							}),
							htmlbuilder.P(f.desc).CSS(htmlbuilder.Style{
								Color:      "#475569",
								FontSize:   "0.95rem",
								LineHeight: "1.5",
								Margin:     "0",
							}),
						).CSS(htmlbuilder.Style{
							Background:   "#ffffff",
							Border:       "1px solid #e2e8f0",
							BorderRadius: "12px",
							Padding:      "2rem",
							Transition:   "all 0.2s ease",
						}).Hover(htmlbuilder.Style{
							BorderColor: "#087ea4",
							BoxShadow:   "0 10px 25px rgba(0,0,0,0.03)",
						})
					})...,
				).CSS(htmlbuilder.Style{
					Display:             "grid",
					GridTemplateColumns: "repeat(auto-fit, minmax(280px, 1fr))",
					Gap:                 "2rem",
					Width:               "100%",
				}),
			).CSS(htmlbuilder.Style{
				MaxWidth:  "1000px",
				Margin:    "0 auto",
				Width:     "100%",
				BoxSizing: "border-box",
			}),
		).CSS(htmlbuilder.Style{
		Display:        "flex",
		AlignItems:     "center",
		JustifyContent: "center",
		Height:         "100vh",
		MinHeight:      "100vh",
		BoxSizing:      "border-box",
		Padding:        "0 2rem",
	})

	// 4. Footer
	footer := htmlbuilder.El("footer").Child(
		htmlbuilder.El("div").Child(
			htmlbuilder.P("© 2026 Built with love").CSS(htmlbuilder.Style{
				Color:      "#94a3b8",
				FontSize:   "0.9rem",
				Margin:     "0",
				FontWeight: "500",
			}),
		),
	).CSS(htmlbuilder.Style{
		Display:        "flex",
		AlignItems:     "center",
		JustifyContent: "center",
		Height:         "25vh",
		BorderTop:      "1px solid #e2e8f0",
		Background:     "#f1f5f9",
	})

	// Build Page
	doc.Body().Child(navBar, mobileDropdown, heroSection, featuresSection, footer)

	return doc.String()
}

func navLinkStyle() htmlbuilder.Style {
	return htmlbuilder.Style{
		Color:          "#475569",
		TextDecoration: "none",
		FontWeight:     "600",
		Transition:     "color 0.15s ease",
	}
}

// Mobile Dropdown Nav Link Style Helper (Lighter, smaller, normal weight)
func mobileNavLinkStyle() htmlbuilder.Style {
	return htmlbuilder.Style{
		// Color is omitted here so SetStyle("color", "inherit") and .Hover() can work without specificity blockages!
		TextDecoration: "none",
		FontWeight:     "400",  // Normal weight, not bold
		FontSize:       "1rem", // Smaller font size
		Transition:     "color 0.15s ease",
	}
}
```

</details>

## Contributing

I love pull requests! Feel free to submit a PR to help make this project better.

## License

This project is [MIT LICENSED](./LICENSE)
