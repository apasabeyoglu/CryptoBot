{{ template "layouts/_auth.tpl" . }}
{{ define "content" }}
        <br><br>
        <div class="row">
            <div class="col s12 m12 l6 offset-l3">
                <div class="card">
                    <div class="card-content">
                        <span class="card-title">Yeni Şifre Oluştur</span>
                        <div class="row">
                            <div class="input-field col s12">
                                <input id="password" type="password" class="validate">
                                <label for="password">Yeni Şifre</label>
                            </div>
                        </div>
                        <div class="row">
                            <div class="input-field col s12">
                                <input id="passwordConfirm" type="password" class="validate">
                                <label for="passwordConfirm">Yeni Şifre (Tekrar)</label>
                            </div>
                        </div>
                        <div class="row"> 
                            <input id="randomString" type="hidden" value="{{ .RandomString }}">
                            <button id="btnNewPassword" class="btn waves-effect waves-light teal ml10">Güncelle
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