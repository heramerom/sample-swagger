// +build sample_swagger

package sample_swagger

import (
	"html/template"
	"net/http"
)

var serverJson string

var htmlTemp = `
<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Swagger UI</title>
    <link rel="stylesheet" type="text/css"
          href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.18.2/swagger-ui.css">
    <link rel="icon" type="image/png" href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAEPElEQVR4Ab1XA8xlWQx+a9u2bcRex17btj1G8Nu2bdu2bfOi02/yzjN/NWnuYdtbH42joKrqqbKy+aSsSD/yN0KS16t4PAKU5I0a/kZhD2dwVrNToKjyhUz4e0WR+ohUcgRwlgX5ke9ebI++nT+W3mYiU4Iwr9HwdAOVtHpSYuXPFF70IZDHP1Fxqwfv1RPOCMBdpvEuaDnJXLmY/yBOEFvfXKa8puP0b8Td9Jmnxib+E3EX5TYe4ztLOqFZiATQdJT51XyhQVyu6gqhX4KvAXGn8Jegq6myMwg0hFmamPa1dv9cMGfHopD8t0FsWxic/yaxwxoIoV5sy+ZxgrlHxisgsCPonvaioRAJFn0CDgd1AUPy38LFHcWgvDdI0Idjmqr+QvbYCWKo6go2uvi1zzlU0RFAq+tzNDRVa5fR4FQNrW7Mw/70te+5RnsVnQG66DAKUcS58HZTh4su/YIA43NtlN1wyK4A2fUHaWy2hQCxZV8b7f0cdJUuOpC0dLZnu/RgMY/Dx5QgogCwP/ohh9X9X9R9BKjtiTDby2k4ovUFue+kLyB1ipBDDJteaBlMIcDvoTc5LAC0CGgdTDPb+zv8Dl1oMu+nNVAFJshiOGAuQBoBfgu53mEBfgy8HEyodSjd0j58SUTEz9BABCYlrR4WD/eNl0FS+srnbIcF+NL7DIQdDUxWWdwvanUTGojSoKphgnxuenBf9APICdQ7Xoo5fe51mpEgYAQUc+zhDMZdowVgYNF3Eip+IADzrtOgnGKComJ4KLNuH3F40vLaNB2KfRxr1D6cSSvrs/RryHVgTJPzXTS10MPjM9nu1/LZGeoYycFZOhj7KC2tToIGZdUfMKIdVvi+MMGYNQH40n5cpqW1KSb2mFaAbOQD+AMEAHOaXuyFAPQrCwDhOkfycJYOxDxCi6sToIHQtC4Aq9iqCfbHPKw1QQnm9AWr92sjE5wJNDLBFzoT5MMEEN6uCSIwKbbmhBPCCc9y2gn7JystO2GLcEIpZnfCMMBmGCJdCwEQhpJBIrpzDxLR7QaJSHpam4rlPizkNh41u1DdFUqAfc6k4sh7baTiw8IBB8DbpBgtoZMxuhBT+iW2UGAQGXaZI3xHZ5sJEFf2jUkxupLWNhdFEvrZqPNVVGUKG5WdgWblGKV1bWOBhqbr7AmANMtnF6E5+sb3PKO98g5/AoCXWcesGDQkaKN2uiEJzHvNtCGx2JIlaOOT3NNf2jHmrmnPg6ZwvCTwstqUonEUQkAT2//z1w36QbnFbnuO1lkIAXXBJ34WjukEovtBK8c0DJir1zr8MNF2rySiI7fxCP1tIU9YaDgQanxn0eBhspkEmlt5mr1r8jRDJKCeI5+jqAAxxhoiwORppsww8/e39VhFuOChycIMkIOAJIM439bj1IJGTkcPhzaKvzGoZCinQIxRWJDbGZ/FWUfpngCleTNdmkrhIgAAAABJRU5ErkJggg==
" sizes="32x32"/>
    <link rel="icon" type="image/png" href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAABhElEQVR4AZVTA0x1cRz9jPHD7DCnKc1htms2spv5smvINS/bjXF4xpRtzHvvf0//82zc7fLg/vjJ/1AUES8UW6NQxKG8P/F0PjcScxMDhco3SVIJYbVqj6cxvl2I9rlM8ByTz9rjKRCTZo3kBhPP3jyb0TSVipKBT0HPRondPJsgTeZ9TOQHFcX14/9JDHvWjf9zmTS6c2Zorj+vaxpwcncQIDy528eGptEZSQqosdeEeTNnFzFUJLVjf3H7YnG/a44nHVGwwiySC3h8O0bl8O8Ag/KhH3h8P3G9y8IWgFpG8MRK8+PkbjHMF2tgOn3LeeiYy0LHfDZ6l3LRNJ0G0/mK5JSQi7bZDFDrY0DQfL4qyTTIp5gnzWgA49kypnZLfQxCpPArRArHgSmwEJrjqdiLeOQooruNbA0Btoot8zc4vt3HprbJ2cZk2Fxt9AySCXVRDtL1s9Hxd99RFrM0YSShxMQoVrxG2d+kkemwJiwSK82TzxwcYtzKAHF068wtFIn+/A9tMmqI7MxzGAAAAABJRU5ErkJggg==
" sizes="16x16"/>
    <style>
        html {
            box-sizing: border-box;
            overflow: -moz-scrollbars-vertical;
            overflow-y: scroll;
        }

        *,
        *:before,
        *:after {
            box-sizing: inherit;
        }

        body {
            margin: 0;
            background: #fafafa;
        }
    </style>
</head>

<body>
<div id="swagger-ui"></div>
<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.18.2/swagger-ui-bundle.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.18.2/swagger-ui-standalone-preset.js"></script>
<script>

    var spec = {{.Spec}};

	spec = JSON.parse(spec);

    window.onload = function () {
        const ui = SwaggerUIBundle({
            spec: spec,
            dom_id: '#swagger-ui',
            deepLinking: true,
            presets: [
                SwaggerUIBundle.presets.apis,
                SwaggerUIStandalonePreset.slice(1) // here
            ],
            layout: "StandaloneLayout"
        });
        window.ui = ui;
    }
</script>
</body>
</html>
`

func ServerHTTP(w http.ResponseWriter, r *http.Request) {
	if serverJson == "" {
		serverJson = parse()
		if serverJson == "" {
			serverJson = "{}"
		}
	}
	t, err := template.New("swagger").Parse(htmlTemp)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	err = t.Execute(w, map[string]string{"Spec": serverJson})
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
