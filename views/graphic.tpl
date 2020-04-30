{{ template "layouts/_main.tpl" . }}
{{ define "content" }}
<div class="row">
    <div class="col s12 m12 32">
        <div class="card-panel">
            <h4 class="header2">Grafik</h4>
            <div class="row">
                <div id="divExchangeAccounts">
                    <div class="input-field col s12">
                        <select id="exchangeAccounts">
                            <option value="" disabled selected>Lütfen borsa hesabınızı seçin</option>
                        </select>
                        <label>Borsa</label>
                    </div>
                </div>
                <div id="divNoChartData" class="col s12">
                    <p>Yeterli veri bulunamadı.</p>
                </div>
                <div id="balanceChart" style="width:100%;"></div>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "js" }}
    <script>
    $(document).ready(function() {
        getExchangeAccounts();
    });
    </script>
    <script src="/static/lib/amcharts/amcharts.js" type="text/javascript"></script>
    <script src="/static/lib/amcharts/serial.js" type="text/javascript"></script>
    <script src="/static/lib/amcharts/themes/light.js" type="text/javascript"></script>
    <script src="/static/lib/amcharts/lang/tr.js" type="text/javascript"></script>
{{ end }}