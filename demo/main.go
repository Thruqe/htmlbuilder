package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Thruqe/htmlbuilder"
)

var (
	mu      sync.Mutex
	counter int
)

func main() {
	http.HandleFunc("/", landing_page_handler)
	http.HandleFunc("/api/counter/increment", increment_handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(embeddedStatic())))
	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func landing_page_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, landing_page())
}

func increment_handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	current := counter
	mu.Unlock()

	w.Header().Set("Content-Type", "text/html")
	counterFragment(current).Render(w)
}

func counterFragment(n int) *htmlbuilder.Node {
	return htmlbuilder.El("span").
		Attr("id", "counter-value").
		Text(fmt.Sprintf("%d", n)).
		CSS(htmlbuilder.Style{
			FontWeight: "700",
			FontSize:   "1.4rem",
			Color:      "var(--accent)",
		})
}

func landing_page() string {
	doc := htmlbuilder.New().
		Title("GoHTML").
		MetaDefault().
		Meta(map[string]string{"name": "description", "content": "A fluent, type-safe HTML builder for Go. No templates, no string concatenation."}).
		Icon("https://cdn-icons-png.flaticon.com/512/186/186320.png", "image/png").
		Link(map[string]string{"rel": "preconnect", "href": "https://fonts.googleapis.com"}).
		Link(map[string]string{
			"rel":  "stylesheet",
			"href": "https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap",
		}).
		Link(map[string]string{
			"rel":  "stylesheet",
			"href": "https://cdn-uicons.flaticon.com/2.6.0/uicons-regular-rounded/css/uicons-regular-rounded.css",
		}).
		StyleBlock(`
		:root {
	    --bg: #fafafa;
		--surface: #ffffff;
		--fg: #111111;
		--muted: #666666;
		--border: #e5e5e5;
		--accent: #4f46e5;
		--shine: rgba(255,255,255,.45);
		--nav-bg: rgba(255,255,255,.72);
		}

		[data-theme="dark"] {
		--bg: #0a0a0a;
		--surface: #151515;
		--fg: #f5f5f5;
		--muted: #999999;
		--border: #2a2a2a;
		--accent: #9470ff;
		--shine: rgba(255,255,255,.25);
		--nav-bg: rgba(15,15,15,.72);
		}
		`)

	doc.Head().Child(
		htmlbuilder.El("script").Attr("src", "/static/hb.js").Attr("defer", "true"),
	)

	features := []struct {
		title string
		desc  string
	}{
		{"Node tree, not strings", "A real tree structure in Go — inspect, mutate, and reuse nodes before rendering, unlike raw string concatenation."},
		{"Escaping by default", "Text and attribute values are escaped automatically. Raw() is an explicit, opt-in escape hatch."},
		{"Typed CSS, every property", "Style is generated from MDN's CSS data — autocomplete on all ~500 CSS properties, no hand-maintained list."},
		{"Go talks to your HTML", "Server-rendered fragments, not a JSON API layer. One Go function renders the full page and its own live updates."},
	}

	doc.Body().CSS(htmlbuilder.Style{
		Margin:     "0",
		FontFamily: "'Inter', sans-serif",
		Background: "var(--bg)",
		Color:      "var(--fg)",
	})

	toggleBtn := htmlbuilder.El("button").
		Attr("id", "theme-toggle").
		Attr("aria-label", "Toggle theme").
		Child(htmlbuilder.El("i").Class("fi", "fi-rr-moon")).
		CSS(htmlbuilder.Style{
			Cursor:     "pointer",
			Border:     "none",
			Background: "none",
			FontSize:   "1.2rem",
			Color:      "var(--fg)",
		})
	themeToggle(toggleBtn, "#theme-toggle i")

	incrementBtn := htmlbuilder.El("button").
		Text("Increment").
		CSS(htmlbuilder.Style{
			Display:      "block",
			Margin:       "1rem auto 0",
			Background:   "none",
			Color:        "var(--fg)",
			Border:       "1px solid var(--border)",
			Padding:      "0.6rem 1.2rem",
			BorderRadius: "8px",
			Cursor:       "pointer",
		})
	onClick(incrementBtn, "/api/counter/increment", "counter-value")

	doc.Body().Child(
		htmlbuilder.El("nav").Child(
			htmlbuilder.Span("htmlbuilder").CSS(htmlbuilder.Style{
				FontWeight: "700",
				FontSize:   "1.1rem",
			}),
			toggleBtn,
		).CSS(htmlbuilder.Style{
			Display:        "flex",
			JustifyContent: "space-between",
			AlignItems:     "center",
			Padding:        "1.25rem 2rem",
			BorderBottom:   "1px solid var(--border)",
			Position:       "sticky",
			Top:            "0",
			ZIndex:         "1000",
			Background:     "var(--nav-bg)",
			BackdropFilter: "blur(16px)",
		}).SetStyle("-webkit-backdrop-filter", "blur(16px)").
			SetStyle("backdrop-filter", "blur(16px)"),

		htmlbuilder.El("section").Child(
			htmlbuilder.H1("Build HTML with Go — not templates.").CSS(htmlbuilder.Style{
				FontSize:   "2.5rem",
				LineHeight: "1.2",
				Margin:     "0 0 1rem",
			}),
			htmlbuilder.P("A fluent, type-checked HTML builder for Go. No string concatenation, no template files, no untyped CSS.").CSS(htmlbuilder.Style{
				Color:      "var(--muted)",
				FontSize:   "1.1rem",
				LineHeight: "1.6",
				Margin:     "0 0 2rem",
			}),
			htmlbuilder.A("Start Building").
				Attr("href", "https://github.com/Thruqe/htmlbuilder").
				CSS(htmlbuilder.Style{
					Display:        "inline-block",
					Background:     "var(--accent)",
					Color:          "var(--bg)",
					Padding:        "0.7rem 1.4rem",
					BorderRadius:   "8px",
					FontWeight:     "500",
					TextDecoration: "none",
				}).
				Hover(htmlbuilder.Style{
					Opacity: "0.85",
				}),
		).CSS(htmlbuilder.Style{
			MaxWidth:  "720px",
			Margin:    "0 auto",
			Padding:   "5rem 1.5rem 3rem",
			TextAlign: "center",
		}),

		htmlbuilder.El("section").Child(
			htmlbuilder.P("Server-rendered counter — click increments state in Go:").CSS(htmlbuilder.Style{
				Color:  "var(--muted)",
				Margin: "0 0 1rem",
			}),
			counterFragment(counter),
			incrementBtn,
		).CSS(htmlbuilder.Style{
			TextAlign: "center",
			Padding:   "2rem 1.5rem 4rem",
		}),

		htmlbuilder.El("section").Child(
			htmlbuilder.Each(features, func(f struct {
				title string
				desc  string
			}) *htmlbuilder.Node {
				return htmlbuilder.El("div").Child(
					htmlbuilder.H3(f.title).CSS(htmlbuilder.Style{
						FontSize: "1.05rem",
						Margin:   "0 0 0.5rem",
					}),
					htmlbuilder.P(f.desc).CSS(htmlbuilder.Style{
						Color:      "var(--muted)",
						FontSize:   "0.9rem",
						LineHeight: "1.5",
						Margin:     "0",
					}),
				).CSS(htmlbuilder.Style{
					Background:   "var(--surface)",
					Border:       "1px solid var(--border)",
					BorderRadius: "12px",
					Padding:      "1.5rem",

					Position:   "relative",
					Overflow:   "hidden",
					Transition: "transform .28s ease, box-shadow .28s ease",
				}).
					Hover(htmlbuilder.Style{
						Transform: "translateY(-8px)",
						BoxShadow: "0 24px 50px rgba(0,0,0,.14)",
					}).
					Pseudo("before", htmlbuilder.Style{
						Content:       `""`,
						Position:      "absolute",
						Top:           "-50%",
						Left:          "-150%",
						Width:         "45%",
						Height:        "220%",
						Background:    "linear-gradient(110deg, transparent 20%, transparent 35%, var(--shine) 50%, transparent 65%, transparent 80%)",
						Transform:     "rotate(25deg)",
						Filter:        "blur(5px)",
						Opacity:       "0.7",
						PointerEvents: "none",
						Transition:    "left .8s cubic-bezier(.22,.61,.36,1)",
					}).
					HoverPseudo("before", htmlbuilder.Style{
						Left: "170%",
					})

			})...,
		).CSS(htmlbuilder.Style{
			MaxWidth: "960px",
			Margin:   "0 auto",
			Padding:  "0 1.5rem 4rem",
			Display:  "grid",
			Gap:      "1.5rem",
		}),

		htmlbuilder.El("footer").Child(
			htmlbuilder.P("© 2026"),
			htmlbuilder.A("MIT Licensed").
				Attr("href", "https://github.com/Thruqe/htmlbuilder/blob/main/LICENSE").CSS(htmlbuilder.Style{Color: "var(--muted)"}),
		).CSS(htmlbuilder.Style{
			Padding:   "2rem",
			TextAlign: "center",
			Color:     "var(--muted)",
			BorderTop: "1px solid var(--border)",
		}),
	)

	return doc.String()
}
