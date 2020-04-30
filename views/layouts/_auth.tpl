<!DOCTYPE html>
  <html>
    <head> 
      <title>{{ .Title }} | CryptoBot</title>
      <link rel="icon" href="/static/favicon.ico" type="image/x-icon" />
      <link rel="shortcut icon" href="/static/favicon.ico" type="image/x-icon" />
      <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
      <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
      <link type="text/css" rel="stylesheet" href="/static/css/materialize.min.css"  media="screen,projection"/>
      <link rel="stylesheet" href="/static/css/iziToast.min.css">
      <link type="text/css" rel="stylesheet" href="/static/css/auth.css"/>
      <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
      {{ block "head" . }}{{ end }}
    </head>
    <body>
      <div class="container">
        {{ block "content" . }}{{ end }}
      </div>
      <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
      <script type="text/javascript" src="/static/js/materialize-auth.min.js"></script>
      <script type="text/javascript" src="/static/js/iziToast.min.js"></script> 
      <script type="text/javascript" src="/static/js/main.js"></script>
      {{ block "js" . }}{{ end }}
    </body>
  </html>  