package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Engine struct {
	Template *template.Template
}

var engineList = map[string]Engine{}

func Init(name, location string) error {
	if _, ok := engineList[name]; ok {
		return fmt.Errorf("Engine name is already initialized.")
	}

	engine := Engine{}
	var funcMap = template.FuncMap{
		"Currency":     engine.currencyFormat,
		"CallTemplate": engine.callTemplate,
	}

	_, err := os.Stat(location)
	if err != nil {
		return err
	}

	if os.IsNotExist(err) {
		return fmt.Errorf("Directory not found.")
	}

	templateFiles := location + "*.html"
	templates := template.Must(template.New("").Funcs(funcMap).ParseGlob(templateFiles))
	engine.Template = templates

	engineList[name] = engine

	return nil
}

func Get(name string) (Engine, error) {
	if engine, ok := engineList[name]; ok {
		return engine, nil
	}
	return Engine{}, fmt.Errorf("Engine %s haven't initialized.", name)
}

func (e Engine) RenderWithLayout(w http.ResponseWriter, renderName string, context map[string]interface{}) {
	context["include"] = renderName
	err := e.Template.ExecuteTemplate(w, "layout", context)
	if err != nil {
		println(err.Error())
	}
}

func (e Engine) RenderPlain(w http.ResponseWriter, renderName string, context map[string]interface{}) {
	e.Template.ExecuteTemplate(w, renderName, context)
}

func (e Engine) currencyFormat(n string) string {
	na := strings.Split(n, ".")
	r := []rune(na[0])
	b := []rune{}
	for i, j := len(r)-1, 3; i >= 0; i-- {
		if j == 0 && r[i] != '-' {
			b = append([]rune{'.'}, b...)
			j = 3
		}
		j--
		b = append([]rune{r[i]}, b...)
	}
	return string(b)
}

func (e Engine) callTemplate(name string, data interface{}) (template.HTML, error) {
	buf := bytes.NewBuffer([]byte{})
	err := e.Template.ExecuteTemplate(buf, name, data)
	ret := template.HTML(buf.String())

	return ret, err
}
