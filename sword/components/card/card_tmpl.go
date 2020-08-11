package card

var List = map[string]string{
	"card": `{{define "card"}}
    <div class="card">
        <div class="card-body">
            <div class="card-index">
                <div class="card-top" {{if eq .SubTitle ""}}style="height: 43px;"{{end}}>
                    <div class="card-meta">
                        <div class="card-title">
                            <span>{{.Title}}</span>
                            {{if ne .Action ""}}
                                <span class="card-title-action">
                                    {{.Action}}
                                </span>
                            {{end}}
                        </div>
                        {{if ne .SubTitle ""}}
                            <div class="card-subtitle"><span>{{.SubTitle}}</span></div>
                        {{end}}
                    </div>
                </div>
                <div class="card-content">
                    {{.Content}}
                </div>
                <div class="card-footer">
                    {{.Footer}}
                </div>
            </div>
        </div>
    </div>
{{end}}`,
}
