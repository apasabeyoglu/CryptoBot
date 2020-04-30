{{ template "layouts/_auth.tpl" . }}
{{ define "content" }}
        <br><br>
        <div class="row">
            <div class="col s12 m12 l6 offset-l3">
                <div class="card">
                    <div class="card-content">
                        <span class="card-title">Yeni Şifre Talebi</span>
                        <div class="row">
                            <div class="input-field col s12">
                                <input id="email" type="email" class="validate">
                                <label for="email">Email</label>
                            </div>
                        </div>
                        <div class="row">
                            <button id="btnForgotPassword" class="btn waves-effect waves-light teal ml10">Gönder
                                <i class="material-icons right">send</i>
                            </button>
                            <br><br>
                            <hr>
                            <p class="center-align">
                                <a href="/kullanici/giris">Giriş Yap</a>
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
{{ end }}
{{ define "js" }}
    <script type="text/javascript" src="/static/js/jquery.cookie.min.js"></script>
{{ end }}