{{ template "layouts/_auth.tpl" . }}
{{ define "head" }}
    <script src='https://www.google.com/recaptcha/api.js'></script>
{{ end }}
{{ define "content" }}
        <br><br>
        <div class="row">
            <div class="col s12 m12 l6 offset-l3">
                <div class="card">
                    <div class="card-content">
                        {{ if .flash.success }}
                            <div class="card-panel teal accent-3">{{ .flash.success }}</div>
                        {{ end }}
                        <span class="card-title">Kullanıcı Girişi</span>
                        <div class="row">
                            <div class="input-field col s12">
                                <input id="email" type="email" class="validate">
                                <label for="email">Email</label>
                            </div>
                        </div>
                        <div class="row">
                            <div class="input-field col s12">
                                <input id="password" type="password" class="validate">
                                <label for="password">Şifre</label>
                            </div>
                        </div>
                        <div class="row">
                            <div class="input-field col s12">
                                <div class="g-recaptcha" data-sitekey="{{ .SiteKey }}"></div>
                            </div>
                        </div>
                        <div class="row">
                            <button id="btnLogin" class="btn waves-effect waves-light teal ml10">Giriş Yap
                                <i class="material-icons right">vpn_key</i>
                            </button>
                            <br><br>
                            <hr>
                            <a class="left" href="/kullanici/sifremi-unuttum">Şifremi unuttum</a>
                            <a class="right" href="/kullanici/kayit">Henüz üye değilim</a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
{{ end }}
{{ define "js" }}
    <script type="text/javascript" src="/static/js/jquery.cookie.min.js"></script>
{{ end }}