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
		Title("GoHTML — Vercel Style").
		MetaDefault().
		Link(map[string]string{
			"rel":  "stylesheet",
			"href": "https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800&display=swap",
		}).
		Link(map[string]string{
			"rel":  "stylesheet",
			"href": "https://cdn-uicons.flaticon.com/2.6.0/uicons-regular-rounded/css/uicons-regular-rounded.css",
		}).
		StyleBlock(`
			* {
				box-sizing: border-box;
			}

			/* CSS Color Tokens for Vercel Theme Switching */
			:root {
				--bg-page: #ffffff;
				--bg-nav: rgba(255, 255, 255, 0.8);
				--bg-card: #ffffff;
				--bg-card-hover: #fafafa;
				--bg-badge: rgba(0, 0, 0, 0.05);
				--border-color: #e5e5e5;
				--border-hover: #000000;
				--text-main: #000000;
				--text-muted: #666666;
				--btn-bg: #000000;
				--btn-text: #ffffff;
				--btn-border: #000000;
				--btn-hover-bg: #ffffff;
				--btn-hover-text: #000000;
				--grid-line: rgba(0, 0, 0, 0.06);
				--gradient-text: linear-gradient(180deg, #000000 0%, #666666 100%);
			}

			[data-theme="dark"] {
				--bg-page: #000000;
				--bg-nav: rgba(0, 0, 0, 0.8);
				--bg-card: #0a0a0a;
				--bg-card-hover: #111111;
				--bg-badge: rgba(255, 255, 255, 0.08);
				--border-color: #222222;
				--border-hover: #444444;
				--text-main: #ffffff;
				--text-muted: #888888;
				--btn-bg: #ffffff;
				--btn-text: #000000;
				--btn-border: #ffffff;
				--btn-hover-bg: #000000;
				--btn-hover-text: #ffffff;
				--grid-line: rgba(255, 255, 255, 0.05);
				--gradient-text: linear-gradient(180deg, #FFFFFF 0%, #A1A1AA 100%);
			}

			html {
				scroll-behavior: smooth;
				background-color: var(--bg-page);
				color: var(--text-main);
				transition: background-color 0.2s ease, color 0.2s ease;
			}

			@media (min-width: 769px) {
				html {
					scroll-snap-type: y mandatory;
				}
				section, footer {
					scroll-snap-align: start;
					scroll-snap-stop: always;
				}
			}

			/* Background Grid */
			.bg-grid {
				background-image: linear-gradient(to right, var(--grid-line) 1px, transparent 1px),
				                  linear-gradient(to bottom, var(--grid-line) 1px, transparent 1px);
				background-size: 64px 64px;
				mask-image: radial-gradient(ellipse 60% 50% at 50% 0%, #000 70%, transparent 100%);
				-webkit-mask-image: radial-gradient(ellipse 60% 50% at 50% 0%, #000 70%, transparent 100%);
			}

			.hero-title {
				background: var(--gradient-text);
				-webkit-background-clip: text;
				-webkit-text-fill-color: transparent;
			}

			.btn-primary {
				background: var(--btn-bg);
				color: var(--btn-text);
				border: 1px solid var(--btn-border);
				padding: 0.75rem 1.75rem;
				border-radius: 6px;
				font-weight: 600;
				font-size: 0.9rem;
				text-decoration: none;
				transition: all 0.15s ease;
			}
			.btn-primary:hover {
				background: var(--btn-hover-bg);
				color: var(--btn-hover-text);
			}

			.btn-secondary {
				background: transparent;
				color: var(--text-muted);
				border: 1px solid var(--border-color);
				padding: 0.75rem 1.75rem;
				border-radius: 6px;
				font-weight: 500;
				font-size: 0.9rem;
				text-decoration: none;
				transition: all 0.15s ease;
			}
			.btn-secondary:hover {
				color: var(--text-main);
				border-color: var(--border-hover);
			}

			.feature-card {
				background: var(--bg-card);
				border: 1px solid var(--border-color);
				border-radius: 8px;
				padding: 2rem;
				transition: all 0.2s ease;
			}
			.feature-card:hover {
				border-color: var(--border-hover);
				background: var(--bg-card-hover);
			}

			.mobile-menu {
				display: flex;
				flex-direction: column;
				align-items: flex-start;
				justify-content: center;
				gap: 1.25rem;
				position: fixed;
				top: 64px;
				left: 0;
				right: 0;
				background: var(--bg-nav);
				border-bottom: 1px solid var(--border-color);
				padding: 0 2rem;
				z-index: 999;
				backdrop-filter: blur(20px);
				-webkit-backdrop-filter: blur(20px);
				max-height: 0;
				opacity: 0;
				overflow: hidden;
				transition: max-height 0.3s cubic-bezier(0.16, 1, 0.3, 1),
				            opacity 0.25s ease,
				            padding 0.3s ease;
			}

			.mobile-menu.active {
				max-height: 240px;
				opacity: 1;
				padding: 2rem;
			}

			@media (max-width: 768px) {
				.desktop-links {
					display: none !important;
				}
				.hamburger-btn {
					display: flex !important;
				}
				h1 {
					font-size: 2.5rem !important;
					letter-spacing: -0.04em !important;
				}
				h2 {
					font-size: 1.75rem !important;
					margin-bottom: 2rem !important;
				}
				.features-grid {
					grid-template-columns: 1fr !important;
					gap: 1rem !important;
				}
				section {
					padding: 6rem 1.5rem 3rem !important;
					height: auto !important;
					min-height: 100vh !important;
				}
			}
		`).
		Script(`
			function toggleTheme() {
				const currentTheme = document.documentElement.getAttribute('data-theme');
				const targetTheme = currentTheme === 'dark' ? 'light' : 'dark';
				document.documentElement.setAttribute('data-theme', targetTheme);
				document.getElementById('theme-icon').className = targetTheme === 'dark' ? 'fi fi-rr-sun' : 'fi fi-rr-moon';
			}
		`)

	doc.Body().
		Attr("data-theme", "light").
		CSS(htmlbuilder.Style{
			Margin:     "0",
			Padding:    "0",
			FontFamily: "'Inter', -apple-system, BlinkMacSystemFont, sans-serif",
			Background: "var(--bg-page)",
			Color:      "var(--text-main)",
		})

	// Theme Toggle Button
	themeToggleBtn := htmlbuilder.El("button").
		Attr("aria-label", "Toggle Theme").
		Attr("onclick", "toggleTheme()").
		Child(
			htmlbuilder.El("i").Attr("id", "theme-icon").Class("fi", "fi-rr-moon").CSS(htmlbuilder.Style{
				FontSize: "1rem",
				Color:    "var(--text-main)",
			}),
		).
		CSS(htmlbuilder.Style{
			Background:     "transparent",
			Border:         "1px solid var(--border-color)",
			BorderRadius:   "6px",
			Cursor:         "pointer",
			Padding:        "6px 10px",
			Display:        "flex",
			AlignItems:     "center",
			JustifyContent: "center",
		})

	hamburgerBtn := htmlbuilder.El("button").
		Class("hamburger-btn").
		Attr("aria-label", "Toggle Menu").
		Attr("onclick", "document.getElementById('mobile-dropdown').classList.toggle('active')").
		Child(
			htmlbuilder.El("i").Class("fi", "fi-rr-menu-burger").CSS(htmlbuilder.Style{
				FontSize: "1.1rem",
				Color:    "var(--text-main)",
			}),
		).
		CSS(htmlbuilder.Style{
			Display:      "none",
			Background:   "transparent",
			Border:       "1px solid var(--border-color)",
			BorderRadius: "6px",
			Cursor:       "pointer",
			Padding:      "6px 10px",
		})

	// 1. Navigation Header
	navBar := htmlbuilder.El("nav").Child(
		htmlbuilder.El("div").Child(
			htmlbuilder.Span("GoHTML").CSS(htmlbuilder.Style{
				FontWeight:    "700",
				FontSize:      "1.05rem",
				Color:         "var(--text-main)",
				LetterSpacing: "-0.02em",
			}),
		),
		htmlbuilder.El("div").CSS(htmlbuilder.Style{
			Display:    "flex",
			AlignItems: "center",
			Gap:        "1.5rem",
		}).Child(
			htmlbuilder.El("div").Class("desktop-links").Child(
				htmlbuilder.A("Overview").Attr("href", "#home").CSS(navLinkStyle()),
				htmlbuilder.A("Features").Attr("href", "#features").CSS(navLinkStyle()),
				htmlbuilder.A("GitHub").Attr("href", "https://github.com/Thruqe/htmlbuilder").CSS(navLinkStyle()),
			).CSS(htmlbuilder.Style{
				Display: "flex",
				Gap:     "2rem",
			}),
			themeToggleBtn,
			hamburgerBtn,
		),
	).CSS(htmlbuilder.Style{
		Display:        "flex",
		JustifyContent: "space-between",
		AlignItems:     "center",
		Padding:        "0 2rem",
		Height:         "64px",
		MaxWidth:       "1200px",
		Margin:         "0 auto",
		Width:          "100%",
		Position:       "fixed",
		Top:            "0",
		Left:           "0",
		Right:          "0",
		ZIndex:         "1000",
		Background:     "var(--bg-nav)",
		BorderBottom:   "1px solid var(--border-color)",
	}).SetStyle("backdrop-filter", "blur(12px)").
		SetStyle("-webkit-backdrop-filter", "blur(12px)")

	// Mobile Menu
	mobileDropdown := htmlbuilder.El("div").
		Class("mobile-menu").
		Attr("id", "mobile-dropdown").
		Child(
			htmlbuilder.A("Overview").
				Attr("href", "#home").
				Attr("onclick", "document.getElementById('mobile-dropdown').classList.remove('active')").
				CSS(mobileNavLinkStyle()),
			htmlbuilder.A("Features").
				Attr("href", "#features").
				Attr("onclick", "document.getElementById('mobile-dropdown').classList.remove('active')").
				CSS(mobileNavLinkStyle()),
			htmlbuilder.A("GitHub").
				Attr("href", "https://github.com/Thruqe/htmlbuilder").
				CSS(mobileNavLinkStyle()),
		)

	// 2. Hero Section
	heroSection := htmlbuilder.El("section").
		Attr("id", "home").
		Child(
			htmlbuilder.El("div").Class("bg-grid").CSS(htmlbuilder.Style{
				Position: "absolute",
				Top:      "0",
				Left:     "0",
				Right:    "0",
				Bottom:   "0",
				ZIndex:   "-1",
			}),
			htmlbuilder.El("div").Child(
				htmlbuilder.El("div").Child(
					htmlbuilder.Span("GoHTML v1.0").CSS(htmlbuilder.Style{
						Background:   "var(--bg-badge)",
						Border:       "1px solid var(--border-color)",
						BorderRadius: "100px",
						Padding:      "4px 12px",
						FontSize:     "0.8rem",
						FontWeight:   "500",
						Color:        "var(--text-muted)",
					}),
				).CSS(htmlbuilder.Style{MarginBottom: "1.5rem"}),

				htmlbuilder.H1("Pure Go. Zero Templates.").
					Class("hero-title").
					CSS(htmlbuilder.Style{
						FontSize:      "4.5rem",
						FontWeight:    "800",
						LineHeight:    "1.05",
						LetterSpacing: "-0.04em",
						Margin:        "0 0 1.5rem",
					}),

				htmlbuilder.P("Develop, test, and deploy type-safe web components natively in Go. Engineered for maximum runtime speed without HTML parser overhead.").CSS(htmlbuilder.Style{
					Color:         "var(--text-muted)",
					FontSize:      "1.25rem",
					LineHeight:    "1.6",
					Margin:        "0 auto 2.5rem",
					MaxWidth:      "620px",
					FontWeight:    "400",
					LetterSpacing: "-0.01em",
				}),

				htmlbuilder.El("div").Child(
					htmlbuilder.A("Get Started").Attr("href", "#features").Class("btn-primary"),
					htmlbuilder.A("Documentation").Attr("href", "https://github.com/Thruqe/htmlbuilder").Class("btn-secondary"),
				).CSS(htmlbuilder.Style{
					Display:        "flex",
					Gap:            "1rem",
					JustifyContent: "center",
					AlignItems:     "center",
				}),
			).CSS(htmlbuilder.Style{
				TextAlign: "center",
			}),
		).CSS(htmlbuilder.Style{
		Display:        "flex",
		AlignItems:     "center",
		JustifyContent: "center",
		Height:         "100vh",
		MinHeight:      "100vh",
		Position:       "relative",
		Padding:        "0 1.5rem",
	})

	// 3. Features Section
	featuresData := []struct {
		title string
		desc  string
	}{
		{"Declarative Tree API", "Construct UI DOM elements directly using Go functions, eliminating string parsing overhead during response rendering."},
		{"Type-Safe Styling", "Maintain structured inline design properties using Go structs, completely preventing invalid CSS attribute emission."},
		{"Zero Runtime Overhead", "Compiles directly into binary execution flow. Eliminates Node.js toolchains, bundlers, and dynamic template resolution."},
	}

	featuresSection := htmlbuilder.El("section").
		Attr("id", "features").
		Child(
			htmlbuilder.El("div").Child(
				htmlbuilder.H2("Engineered for Go performance").CSS(htmlbuilder.Style{
					FontSize:      "2.5rem",
					FontWeight:    "700",
					TextAlign:     "center",
					LetterSpacing: "-0.03em",
					Margin:        "0 0 3rem",
					Color:         "var(--text-main)",
				}),
				htmlbuilder.El("div").Class("features-grid").Child(
					htmlbuilder.Each(featuresData, func(f struct {
						title string
						desc  string
					}) *htmlbuilder.Node {
						return htmlbuilder.El("div").Class("feature-card").Child(
							htmlbuilder.H3(f.title).CSS(htmlbuilder.Style{
								FontSize:      "1.1rem",
								FontWeight:    "600",
								Margin:        "0 0 0.75rem",
								Color:         "var(--text-main)",
								LetterSpacing: "-0.01em",
							}),
							htmlbuilder.P(f.desc).CSS(htmlbuilder.Style{
								Color:      "var(--text-muted)",
								FontSize:   "0.9rem",
								LineHeight: "1.6",
								Margin:     "0",
							}),
						)
					})...,
				).CSS(htmlbuilder.Style{
					Display:             "grid",
					GridTemplateColumns: "repeat(auto-fit, minmax(300px, 1fr))",
					Gap:                 "1.5rem",
					Width:               "100%",
				}),
			).CSS(htmlbuilder.Style{
				MaxWidth:  "1100px",
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
		Padding:        "0 2rem",
	})

	// 4. Footer
	footer := htmlbuilder.El("footer").Child(
		htmlbuilder.El("div").Child(
			htmlbuilder.P("© 2026 GoHTML. All rights reserved.").CSS(htmlbuilder.Style{
				Color:      "var(--text-muted)",
				FontSize:   "0.85rem",
				Margin:     "0",
				FontWeight: "400",
			}),
		),
	).CSS(htmlbuilder.Style{
		Display:        "flex",
		AlignItems:     "center",
		JustifyContent: "center",
		Height:         "20vh",
		BorderTop:      "1px solid var(--border-color)",
	})

	doc.Body().Child(navBar, mobileDropdown, heroSection, featuresSection, footer)

	return doc.String()
}

func navLinkStyle() htmlbuilder.Style {
	return htmlbuilder.Style{
		Color:          "var(--text-muted)",
		TextDecoration: "none",
		FontWeight:     "500",
		FontSize:       "0.9rem",
		Transition:     "color 0.15s ease",
	}
}

func mobileNavLinkStyle() htmlbuilder.Style {
	return htmlbuilder.Style{
		Color:          "var(--text-muted)",
		TextDecoration: "none",
		FontWeight:     "500",
		FontSize:       "1rem",
		Transition:     "color 0.15s ease",
	}
}
