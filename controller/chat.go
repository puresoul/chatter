package controller


const chatpage = `
<html>
<body>
<div class="container" id="chat">
        <div class="left">
            <ul class="people">
                {{with .Usr}}
                {{range .}}
                <li onClick="setName(this)" class="person" id="{{.Name}}">
                    <span class="name">{{.Name}}</span>
                    <span class="time">{{.Time}}</span>
                </li>
                {{end}}
                {{end}}
            </ul>
        </div>
        <div class="right">
            <div class="top"><span><span class="name">NAME</span></span></div>
            <div class="chat">
                <div class="conversation-start">
                    <span>TIME</span>
                </div>
                {{with .Msg}}
                {{range .}}
                {{ if .Me }}
                <div class="bubble me">
                    ME
                </div>
                {{end}}
                {{ if .You }}
                <div class="bubble you">
                    YOU
                </div>
               {{end}}
               {{end}}
               {{end}}
            </div>
            <div class="write">
                <input type="text" />
                <a href="javascript:;" class="write-link send"></a>
            </div>
        </div>
</div>
</body>
</html>
`
