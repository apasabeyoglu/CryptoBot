{{ template "layouts/_auth.tpl" . }}
{{ define "content" }}
    <br><br>
    <div class="row">
        <div class="col s12 m6 offset-m3">
            <div class="card">
                <div class="card-content">
                 <span class="card-title">Yeni Kullanıcı Kaydı</span>
                    <div class="row">
                        <div class="input-field col s12">
                            <input id="email" type="email" class="validate">
                            <label for="email">Email</label>
                        </div>
                    </div>
                    <div class="row">
                        <div class="input-field col s12">
                            <input id="username" type="text" class="validate">
                            <label for="username">Kullanıcı Adı</label>
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
                            <input id="passwordConfirm" type="password" class="validate">
                            <label for="passwordConfirm">Şifre (Tekrar)</label>
                        </div>
                    </div>
                    <div class="row">
                        <button id="btnRegister" class="btn waves-effect waves-light teal ml10">Kayıt Ol
                            <i class="material-icons right">add_box</i>
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