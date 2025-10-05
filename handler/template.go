package handler

import (
	"html/template"
	"path/filepath"
	"strings"
)

func getFileIcon(name string, isDir bool) template.HTML {
	if isDir {
		// Folder
		return template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 16 16"><title>Folder-fill SVG Icon</title><path fill="#ffaa00" d="M9.828 3h3.982a2 2 0 0 1 1.992 2.181l-.637 7A2 2 0 0 1 13.174 14H2.825a2 2 0 0 1-1.991-1.819l-.637-7a2 2 0 0 1 .342-1.31L.5 3a2 2 0 0 1 2-2h3.672a2 2 0 0 1 1.414.586l.828.828A2 2 0 0 0 9.828 3m-8.322.12q.322-.119.684-.12h5.396l-.707-.707A1 1 0 0 0 6.172 2H2.5a1 1 0 0 0-1 .981z"/></svg>`)
	}

	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".html", ".htm", ".css", ".js", ".ts", ".py", ".rb", ".php", ".java", ".c", ".h", ".cpp", ".cc", ".cxx", ".hpp", ".cs", ".go", ".rs", ".kt", ".swift", ".m", ".mm", ".sh", ".bash", ".zsh", ".ps1", ".pl", ".pm", ".lua", ".r", ".jl", ".scala", ".clj", ".hs", ".erl", ".ex", ".exs", ".dart", ".vb", ".f90", ".f95", ".f03", ".sql", ".asm", ".s" :
		return template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 16 16"><title>File-earmark-code SVG Icon</title><g fill="#ffaa00"><path d="M14 4.5V14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h5.5zm-3 0A1.5 1.5 0 0 1 9.5 3V1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V4.5z"/><path d="M8.646 6.646a.5.5 0 0 1 .708 0l2 2a.5.5 0 0 1 0 .708l-2 2a.5.5 0 0 1-.708-.708L10.293 9L8.646 7.354a.5.5 0 0 1 0-.708m-1.292 0a.5.5 0 0 0-.708 0l-2 2a.5.5 0 0 0 0 .708l2 2a.5.5 0 0 0 .708-.708L5.707 9l1.647-1.646a.5.5 0 0 0 0-.708"/></g></svg>`)
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg":
		return template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 16 16"><title>Image SVG Icon</title><g fill="#ffaa00"><path d="M6.002 5.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0"/><path d="M2.002 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2zm12 1a1 1 0 0 1 1 1v6.5l-3.777-1.947a.5.5 0 0 0-.577.093l-3.71 3.71l-2.66-1.772a.5.5 0 0 0-.63.062L1.002 12V3a1 1 0 0 1 1-1z"/></g></svg>`)
	case ".mp3", ".wav", ".wma", ".aac", ".ogg", ".oga", ".flac", ".m4a", ".m4b", ".m4p", ".aiff", ".aif", ".aifc", ".opus", ".amr", ".au", ".snd", ".ra", ".mid", ".midi", ".rmi", ".ac3", ".dts", ".caf":
		return template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16 16"><title>File-earmark-music SVG Icon</title><g fill="#ffaa00"><path d="M11 6.64a1 1 0 0 0-1.243-.97l-1 .25A1 1 0 0 0 8 6.89v4.306A2.6 2.6 0 0 0 7 11c-.5 0-.974.134-1.338.377c-.36.24-.662.628-.662 1.123s.301.883.662 1.123c.364.243.839.377 1.338.377s.974-.134 1.338-.377c.36-.24.662-.628.662-1.123V8.89l2-.5z"/><path d="M14 14V4.5L9.5 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2M9.5 3A1.5 1.5 0 0 0 11 4.5h2V14a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h5.5z"/></g></svg>`)
	case ".mp4", ".m4v", ".mkv", ".webm", ".avi", ".mov", ".qt", ".wmv", ".flv", ".f4v", ".3gp", ".3g2", ".mpg", ".mpeg", ".mp2", ".mpe", ".mpv", ".ogv", ".rm", ".rmvb", ".vob", ".m2ts", ".mts", ".asf":
		return template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 16 16"><title>Film SVG Icon</title><path fill="#ffaa00" d="M0 1a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1zm4 0v6h8V1zm8 8H4v6h8zM1 1v2h2V1zm2 3H1v2h2zM1 7v2h2V7zm2 3H1v2h2zm-2 3v2h2v-2zM15 1h-2v2h2zm-2 3v2h2V4zm2 3h-2v2h2zm-2 3v2h2v-2zm2 3h-2v2h2z"/></svg>`)
	case ".pdf":
		return template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 16 16"><title>File-earmark-pdf SVG Icon</title><g fill="#ffaa00"><path d="M14 14V4.5L9.5 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2M9.5 3A1.5 1.5 0 0 0 11 4.5h2V14a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h5.5z"/><path d="M4.603 14.087a.8.8 0 0 1-.438-.42c-.195-.388-.13-.776.08-1.102c.198-.307.526-.568.897-.787a7.7 7.7 0 0 1 1.482-.645a20 20 0 0 0 1.062-2.227a7.3 7.3 0 0 1-.43-1.295c-.086-.4-.119-.796-.046-1.136c.075-.354.274-.672.65-.823c.192-.077.4-.12.602-.077a.7.7 0 0 1 .477.365c.088.164.12.356.127.538c.007.188-.012.396-.047.614c-.084.51-.27 1.134-.52 1.794a11 11 0 0 0 .98 1.686a5.8 5.8 0 0 1 1.334.05c.364.066.734.195.96.465c.12.144.193.32.2.518c.007.192-.047.382-.138.563a1.04 1.04 0 0 1-.354.416a.86.86 0 0 1-.51.138c-.331-.014-.654-.196-.933-.417a5.7 5.7 0 0 1-.911-.95a11.7 11.7 0 0 0-1.997.406a11.3 11.3 0 0 1-1.02 1.51c-.292.35-.609.656-.927.787a.8.8 0 0 1-.58.029m1.379-1.901q-.25.115-.459.238c-.328.194-.541.383-.647.547c-.094.145-.096.25-.04.361q.016.032.026.044l.035-.012c.137-.056.355-.235.635-.572a8 8 0 0 0 .45-.606m1.64-1.33a13 13 0 0 1 1.01-.193a12 12 0 0 1-.51-.858a21 21 0 0 1-.5 1.05zm2.446.45q.226.245.435.41c.24.19.407.253.498.256a.1.1 0 0 0 .07-.015a.3.3 0 0 0 .094-.125a.44.44 0 0 0 .059-.2a.1.1 0 0 0-.026-.063c-.052-.062-.2-.152-.518-.209a4 4 0 0 0-.612-.053zM8.078 7.8a7 7 0 0 0 .2-.828q.046-.282.038-.465a.6.6 0 0 0-.032-.198a.5.5 0 0 0-.145.04c-.087.035-.158.106-.196.283c-.04.192-.03.469.046.822q.036.167.09.346z"/></g></svg>`)
	case ".zip", ".tar", ".gz", ".rar", ".7z":
		return template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 16 16"><title>File-zip SVG Icon</title><g fill="#ffaa00"><path d="M6.5 7.5a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v.938l.4 1.599a1 1 0 0 1-.416 1.074l-.93.62a1 1 0 0 1-1.109 0l-.93-.62a1 1 0 0 1-.415-1.074l.4-1.599zm2 0h-1v.938a1 1 0 0 1-.03.243l-.4 1.598l.93.62l.93-.62l-.4-1.598a1 1 0 0 1-.03-.243z"/><path d="M2 2a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm5.5-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H9v1H8v1h1v1H8v1h1v1H7.5V5h-1V4h1V3h-1V2h1z"/></g></svg>`)
	default:
		// File default
		return template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" 
			fill="gray" viewBox="0 0 16 16" style="vertical-align: middle; margin-right: 6px;">
			<path d="M14 4.5V14a2 2 0 0 1-2 
			2H4a2 2 0 0 1-2-2V2a2 2 0 0 
			1 2-2h5.5L14 4.5z"/></svg>`)
	}
}

var dirTemplate = template.Must(template.New("dir").Parse(`
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>Index of {{.Path}}</title>
	<style>
		body { 
			font-family: Arial, sans-serif; 
			background: #fafafa;
			margin: 0;
			padding: 20px;
			font-size: 16px; /* sedikit lebih besar */
		}
		.container {
			width: 90%;              
			max-width: 1200px;       
			min-width: 300px;        
			margin: 0 auto;
			background: #fff;
			padding: 24px; /* lebih lega */
			border-radius: 8px;
			box-shadow: 0 2px 6px rgba(0,0,0,0.1);
		}
		.file-link {
			display: inline-flex;       /* bikin icon + text sejajar */
			align-items: center;        /* vertikal rata tengah */
			gap: 8px;                   /* jarak icon dengan teks */
		}

		.file-link img, 
		.file-link svg {
			width: 32px;                /* pastikan ukuran konsisten */
			height: 32px;
			flex-shrink: 0;             /* icon tidak mengecil */
			vertical-align: middle;     /* cadangan */
		}

		.filename {
			line-height: 1.2;           /* biar teks pas di tengah */
			font-size: 17px;
		}

		h1 { 
			font-size: 24px; /* lebih besar dari sebelumnya */
			margin-top: 0; 
		}
		table { 
			border-collapse: collapse; 
			width: 100%; 
			margin-top: 15px; 
			font-size: 15px; /* lebih jelas untuk tabel */
		}
		th, td { 
			padding: 10px 14px; 
			border-bottom: 1px solid #ddd; 
			text-align: left; 
		}
		a { 
			text-decoration: none; 
			color: #0073e6; 
		}
		a:hover { text-decoration: underline; }
		.footer {
			margin-top: 30px;
			font-size: 18px; /* sedikit lebih besar dari 12px */
			color: #666;
			text-align: center;
		}
		/* Responsif */
		@media (max-width: 600px) {
			body { font-size: 14px; padding: 10px; }
			h1 { font-size: 20px; }
			table { font-size: 13px; }
			th, td { padding: 8px 10px; }
		}
	</style>
</head>
<body>
	<div class="container">
		<h1>Index of {{.Path}}</h1>
		<table>
			<tr><th>Name</th><th>Size</th><th>Last Modified</th></tr>
			{{if .Parent}}
				<tr>
					<td><a href="{{.Parent}}">
				<!-- ikon folder naik -->
				<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="gray" viewBox="0 0 16 16" style="vertical-align: middle; margin-right: 6px;">
					<path d="M8.354 1.646a.5.5 0 0 1 0 .708L3.707 7H14.5a.5.5 0 0 1 0 1H3.707l4.647 4.646a.5.5 0 0 1-.708.708l-5.5-5.5a.5.5 0 0 1 0-.708l5.5-5.5a.5.5 0 0 1 .708 0z"/>
				</svg>
				Parent Directory
			</a></td>
					<td>-</td>
					<td>-</td>
				</tr>
			{{end}}
			{{range .Files}}
				<tr>
					<td><a href="{{.Name}}" class="file-link">
						{{.Icon}}
						<span class="filename">{{.Name}}</span>
					</a></td>
					<td>{{.Size}}</td>
					<td>{{.ModTime}}</td>
				</tr>
			{{end}}
		</table>
		<div class="footer">
			Powered by <strong>Auri HTTP Server</strong>
		</div>
	</div>
</body>
</html>

`))
