// +build js,wasm

package main

import (
	"github.com/shurcooL/github_flavored_markdown"
	"syscall/js"
)

var document = js.Global().Get("document")

func getElementById(id string) js.Value {
	return document.Call("getElementById", id)
}

func renderEditor(parent js.Value) js.Value {
	editorMarkup := `
		<div id="editor" style="display: flex; flex-flow: row wrap;">
			<textarea id="markdown" style="width: 50%; height: 400px">This *is* a **test**.</textarea>
			<div id="preview" style="width: 50%;"></div>
			<button id="render">Render Markdown</button>
		</div>
	`
	parent.Call("insertAdjacentHTML", "beforeend", editorMarkup)
	return getElementById("editor")
}

func renderMarkdown(markdown js.Value, preview js.Value) {
	md := markdown.Get("value").String()
	html := github_flavored_markdown.Markdown([]byte(md))
	preview.Set("innerHTML", string(html))
}

func main() {
	done := make(chan struct{})

	renderEditor(document.Get("body"))
	markdown := getElementById("markdown")
	preview := getElementById("preview")
	renderButton := getElementById("render")
	renderButton.Set("onclick", js.NewCallback(func([]js.Value) {
		renderMarkdown(markdown, preview)
	}))
	renderMarkdown(markdown, preview)

	<-done
}
