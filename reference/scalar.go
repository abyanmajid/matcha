package reference

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func safeJSONConfiguration(options *Options) string {
	jsonData, _ := json.Marshal(options)
	escapedJSON := strings.ReplaceAll(string(jsonData), `"`, `&quot;`)
	return escapedJSON
}

func specContentHandler(specContent interface{}) string {
	switch spec := specContent.(type) {
	case func() map[string]interface{}:
		result := spec()
		jsonData, _ := json.Marshal(result)
		return string(jsonData)
	case map[string]interface{}:
		jsonData, _ := json.Marshal(spec)
		return string(jsonData)
	case string:
		return spec
	default:
		return ""
	}
}

func ScalarHTML(optionsInput *Options, r *http.Request) (string, error) {
	opts := defaultOptions(*optionsInput)

	if opts.SpecContent == nil && opts.Source != "" {
		spec, err := fetchContentFromURL(opts.Source, r)
		if err != nil {
			return "", err
		}

		opts.SpecContent = spec
	}

	jsonCfg := safeJSONConfiguration(opts)
	html := specContentHandler(opts.SpecContent)

	var title string

	if opts.PageTitle != "" {
		title = opts.PageTitle
	} else {
		title = "Matcha API Reference"
	}

	customThemeCss := CustomThemeCSS

	if opts.Theme != "" {
		customThemeCss = ""
	}

	return fmt.Sprintf(`
    <!DOCTYPE html>
    <html>
      <head>
        <title>%s</title>
        <meta charset="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <style>%s</style>
      </head>
      <body>
        <script id="api-reference" type="application/json" data-configuration="%s">%s</script>
        <script src="%s"></script>
      </body>
    </html>
  `, title, customThemeCss, jsonCfg, html, opts.CDN), nil
}
