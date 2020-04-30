{{ template "layouts/_main.tpl" . }}
{{ define "content" }}
    <div class="row">
        <div class="col s12 m12 32">
            <div class="card-panel">
                <h4 class="header2">Hesaplarım</h4>
                <div class="row">
                    <table id="tableExchangeAccounts" class="responsive-table striped highlight">
                        <thead>
                            <tr>
                                <th>Borsa</th>
                                <th>İsim</th>
                                <th>Oluşturulma Zamanı</th>
                                <th>Durum</th>
                                <th>İşlemler</th>
                            </tr>
                        </thead>
                        <tbody id="tbodyExchangeAccounts">
                            
                        </tbody>
                    </table>
                    <div id="divNoExchangeAccounts" class="col s12">
                        <p>Borsa hesabı bulunamadı.</p>
                    </div>
                </div>
            </div>
        </div>
        {{ if .Exchanges }}
            <div class="col s12 m12 32">
                <div class="card-panel">
                    <h4 class="header2">Yeni Hesap</h4>
                    <div class="row">
                        <div class="col s12">
                            <div class="row">
                                <div class="input-field col s12">
                                    <select id="exchange">
                                        <option value="" disabled selected>Lütfen seçiminizi yapın</option>
                                        {{ range $key, $exchange := .Exchanges }}
                                            <option value="{{ $exchange.ID }}">{{ $exchange.Name }}</option>
                                        {{ end }}
                                    </select>
                                    <label>Borsa</label>
                                </div>
                            </div>
                            <div class="row">
                                <div class="input-field col s12">
                                    <input id="name" type="text">
                                    <label for="name">Hesap İsmi (Opsiyonel)</label>
                                </div>
                            </div>
                            <div class="row">
                                <div class="input-field col s12">
                                    <input id="apiKey" type="text">
                                    <label for="apiKey">API Key</label>
                                </div>
                            </div>
                            <div class="row">
                                <div class="input-field col s12">
                                    <input id="apiSecret" type="text">
                                    <label for="apiSecret">API Secret</label>
                                </div>
                            </div>
                            <div class="row">
                                <div class="input-field col s12">
                                    <button id="btnAddExchangeAccount" class="btn waves-effect waves-light right">Ekle
                                        <i class="material-icons right">send</i>
                                    </button>
                                </div>
                            </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        {{ end }}

        <div id="modalEditExchangeAccount" class="modal modal-fixed-footer">
            <div class="modal-content">
                <h4>Hesabı Düzenle</h4>
                <hr>
                <input id="editID" type="hidden">
                <div class="row">
                    <div class="input-field col s12">
                        <input id="editExchangeName" type="text" disabled></p>
                        <label for="editExchangeName">Borsa</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <input id="editName" type="text">
                        <label for="editName">Hesap İsmi (Opsiyonel)</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <input id="editKey" type="text">
                        <label for="editKey">API Key</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <input id="editSecret" type="text">
                        <label for="editSecret">API Secret</label>
                    </div>
                </div>
            </div>
            <div class="modal-footer">
                <a href="#!" class="modal-close waves-effect waves-red btn-flat">Vazgeç</a>
                <a id="editExchangeAccount" href="#!" class="waves-effect waves-green btn-flat">Kaydet</a>
            </div>
        </div>
    </div>
{{ end }}
{{ define "js" }}
    <script>
    $(document).ready(function() {
        refreshExchangeAccounts();
    });
    </script>
{{ end }}