<!DOCTYPE html>
<html lang="en">
  <head>
    <title>{{ .Title }} | CryptoBot</title>
    <link rel="icon" href="/static/favicon.ico" type="image/x-icon" />
    <link rel="shortcut icon" href="/static/favicon.ico" type="image/x-icon" />
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="msapplication-tap-highlight" content="no">
    <link href="/static/css/materialize.css" type="text/css" rel="stylesheet">
    <link href="/static/css/style.css" type="text/css" rel="stylesheet">
    <link href="/static/css/perfect-scrollbar.css" type="text/css" rel="stylesheet">
    <link rel="stylesheet" href="/static/css/iziToast.min.css">
    <link href="/static/css/custom.css" type="text/css" rel="stylesheet">
    {{ block "head" . }}{{ end }}
  </head>
  <body>
    <header id="header" class="page-topbar">
      <div class="navbar-fixed">
        <nav class="navbar-color gradient-45deg-light-blue-cyan">
          <div class="nav-wrapper">
            <ul class="left"> 
              <li>
                <h1 class="logo-wrapper">
                  <a href="/" class="brand-logo darken-1">
                    <img src="/static/images/logo.png" alt="materialize logo">
                    <span class="logo-text hide-on-med-and-down">CryptoBot</span>
                  </a>
                </h1>
              </li>
            </ul>
            <ul class="right hide-on-med-and-down">
              <li>
                <a href="javascript:void(0);" class="waves-effect waves-block waves-light toggle-fullscreen" title="Tam ekran">
                  <i class="material-icons">settings_overscan</i>
                </a>
              </li>
            </ul>
          </div>
        </nav>
      </div>
    </header>

    <div id="main">
      <div class="wrapper">
        <aside id="left-sidebar-nav">
          <ul id="slide-out" class="side-nav fixed leftside-navigation">
            <li class="no-padding">
              <ul class="collapsible" data-collapsible="accordion">
                <li class="bold">
                  <a href="/bakiye" class="waves-effect waves-cyan">
                      <i class="material-icons">account_balance_wallet</i>
                      <span class="nav-text">Bakiyelerim</span>
                    </a>
                </li>
                <li class="bold">
                  <a href="/emir" class="waves-effect waves-cyan">
                      <i class="material-icons">compare_arrows</i>
                      <span class="nav-text">Emirlerim</span>
                    </a>
                </li>
                <li class="bold">
                  <a href="/bakiye/grafik" class="waves-effect waves-cyan">
                      <i class="material-icons">show_chart</i>
                      <span class="nav-text">Grafik</span>
                    </a>
                </li>
                <li class="bold">
                  <a href="/borsa-hesabi" class="waves-effect waves-cyan">
                      <i class="material-icons">account_balance</i>
                      <span class="nav-text">Hesaplarım</span>
                    </a>
                </li>
                <li class="bold">
                    <form id="logoutForm" method="post" action="/kullanici/cikis">
                        {{ .XsrfData }}
                        <a id="btnLogout" onclick="$('#logoutForm').submit();" class="waves-effect waves-cyan">
                            <i class="material-icons">keyboard_tab</i>
                            <span class="nav-text">Çıkış</span>
                        </a>
                    </form>
                </li>
              </ul>
            </li>
          </ul>
          <a href="#" data-activates="slide-out" class="sidebar-collapse btn-floating btn-medium waves-effect waves-light hide-on-large-only">
            <i class="material-icons">menu</i>
          </a>
        </aside>

        <section id="content">
          <div class="container">
			      {{ block "content" . }}{{ end }}
          </div>
        </section>
      </div>
    </div>

    <footer class="page-footer gradient-45deg-light-blue-cyan">
        <div class="footer-copyright">
          <div class="container">
            <span>
              <script type="text/javascript">
                document.write(new Date().getFullYear());
              </script> - CryptoBot</span>
          </div>
        </div>
    </footer>

    <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
    <script type="text/javascript" src="/static/js/materialize.min.js"></script>
    <script type="text/javascript" src="/static/js/perfect-scrollbar.min.js"></script>
    <script type="text/javascript" src="/static/js/iziToast.min.js"></script> 
    <script type="text/javascript" src="/static/js/jquery.cookie.min.js"></script>
    <script type="text/javascript" src="/static/js/plugins.js"></script>
    <script type="text/javascript" src="/static/js/main.js"></script>
    {{ block "js" . }}{{ end }}
  </body>
</html>