package card

var List = map[string]string{
	"card": `{{define "card"}}
    <div class="card" id="{{.ID}}">
        <div class="card-body" id="{{.BodyID}}">
            <div class="card-index">
                <div class="card-top" id="{{.TopID}}" {{if eq .SubTitle ""}}style="height: 43px;"{{end}}>
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
                <div class="card-content" id="{{.ContentID}}">
                    {{.Content}}
                </div>
                <div class="card-footer" id="{{.FooterID}}">
                    {{.Footer}}
                </div>
            </div>
        </div>
    </div>
{{end}}`,
}
