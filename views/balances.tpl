{{ template "layouts/_main.tpl" . }}
{{ define "content" }}
<div class="row">
    <div class="col s12 m12 32">
        <div class="card-panel">
            <h4 class="header2">Bakiyelerim</h4>
            <div class="row">
                <div id="divNoActiveExchangeAccounts" class="col s12">
                    <p>Aktif borsa hesabı bulunamadı. Eklemek için <a href="/borsa-hesabi">tıklayın.</a></p>
                </div>
                <div id="divBalances"></div>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "js" }}
    <script>
    $(document).ready(function() {
        getBalances();
        setInterval(function(){ 
            getBalances();
        }, 5000);
    });
    </script>
{{ end }}