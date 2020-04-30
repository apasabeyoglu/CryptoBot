{{ template "layouts/_main.tpl" . }}
{{ define "content" }}
<div class="row">
    <div class="col s12 m12 32">
        <div class="card-panel">
            <h4 class="header2">Emirlerim</h4>
            <div class="row">
                <table id="tableOrders" class="responsive-table striped highlight">
                    <thead>
                        <tr>
                            <th>Market</th>
                            <th>İşlem</th>
                            <th>Tür</th>
                            <th>Miktar</th>
                            <th>Fiyat</th>
                        </tr>
                    </thead>
                    <tbody id="tbodyOrders">
                        
                    </tbody>
                </table>
                <div id="divNoOrders" class="col s12">
                    <p>Aktif emir bulunamadı.</p>
                </div>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "js" }}
    <script>
    $(document).ready(function() {
        getOrders();
        setInterval(function(){ 
            getOrders();
        }, 10000);
    });
    </script>
{{ end }}