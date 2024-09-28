// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func BaseLayout(title string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if !isPartial(ctx) {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html> <html lang=\"en\"><head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"description\" content=\"Adam&#39;s Blog. Programming, Vim, Photography, and more!\"><script src=\"https://unpkg.com/htmx.org@1.9.12\" integrity=\"sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2\" crossorigin=\"anonymous\" defer></script><link rel=\"stylesheet\" href=\"/static/css/syntax.min.css\" defer><link rel=\"stylesheet\" href=\"/static/css/main.css\" defer><title>Adam's Blog</title><link rel=\"icon\" href=\"data:,\"></head><body id=\"body\" class=\"flex flex-col justify-center bg-gray-900 lg:px-9 lg:pt-7\"><div class=\"flex flex-col justify-center items-center lg:p-6\"><div class=\"mb-3 w-full max-w-4xl bg-gray-800 shadow-xl sm:max-w-full md:max-w-4xl md:rounded-xl lg:max-w-6xl text-slate-300\"><div class=\"flex flex-row flex-nowrap pl-6 rounded-md shadow-md\"><div class=\"justify-start py-3 cursor-pointer shrink\" hx-get=\"/\" hx-push-url=\"true\" hx-on::after-request=\"document.getElementById(&#39;search-nav&#39;).value = &#39;&#39;;\" hx-trigger=\"click\" hx-target=\"#content\"><span class=\"font-bold text-sky-700\">::/</span> Adam's Things</div><div class=\"grow\">&nbsp;</div><input class=\"hidden flex-auto px-2 my-2 mr-4 font-normal bg-gray-700 rounded border border-gray-600 border-solid md:block outline-gray-600 max-w-56\" placeholder=\"Search...\" type=\"text\" name=\"search\" id=\"search-nav\" hx-post=\"/\" hx-trigger=\"input changed delay:200ms, search\" hx-target=\"#content\" hx-push-url=\"true\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = NavigationLinks().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div id=\"content\" class=\"p-6\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = scrollToTopButton().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body></html>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate