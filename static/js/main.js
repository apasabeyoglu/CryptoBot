$(document).ready(function () {
    $('.modal').modal();
    
    $('#btnLogin').click(function() {
        email = $('#email').val();
        password = $('#password').val();
    
        if(!validateEmail(email)) {
            showErrorToast("Lütfen geçerli bir email adresi girin.");
        } else if(password.length < 4 || password.length > 20) {
            showErrorToast("Lütfen geçerli bir şifre girin.");
        } else {
            $.ajax({
                url: "/kullanici/giris",
                type: "POST",
                async: true,
                cache: false,
                dataType: "json",
                headers: {
                    'X-Xsrftoken': getToken(),
                },
                data: {email: email, password: password, captcha: grecaptcha.getResponse() },
                beforeSend: function() {
                    $('#btnLogin').attr('disabled', true);
                },
                success: function(res) {
                    if(res.Result == 0)
                    {
                        showSuccessToast(res.Message);
                        setTimeout(function() { window.location = "/"; }, 2000);
                    } 
                    else 
                    {
                        showErrorToast(res.Message);
                        grecaptcha.reset();
                    }

                },
                error: function(res) {
                    showErrorToast("Beklenmeyen bir hata oluştu.");
                },
                complete: function() {
                    $('#btnLogin').attr('disabled', false);
                }
            });
        }
    });

    $('#btnRegister').click(function() {
        email = $('#email').val();
        username = $('#username').val();
        password = $('#password').val();
        passwordConfirm = $('#passwordConfirm').val();

        if(!validateEmail(email)) {
            showErrorToast("Lütfen geçerli bir email adresi girin.");
        } else if(username.length < 2 || username.length > 50) {
            showErrorToast("Kullanıcı adı 2-50 karakter aralığında olmalıdır.");
        } else if(password.length < 6 || password.length > 20) {
            showErrorToast("Şifre 6-20 karakter aralığında olmalıdır.");
        } else if(password != passwordConfirm) {
            showErrorToast("Girilen şifreler aynı değil.");
        } else {
            $.ajax({
                url: "/kullanici/kayit",
                type: "POST",
                async: true,
                cache: false,
                dataType: "json",
                headers: {
                    'X-Xsrftoken': getToken(),
                },
                data: {email: email, username: username, password: password},
                beforeSend: function() {
                    $('#btnRegister').attr('disabled', true);
                },
                success: function(res) {
                    if(res.Result == 0)
                    {
                        showSuccessToast(res.Message);
                        setTimeout(function() { window.location = "/"; }, 5000);
                    } 
                    else 
                    {
                        showErrorToast(res.Message);
                    }

                },
                error: function(res) {
                    showErrorToast("Beklenmeyen bir hata oluştu.");
                },
                complete: function() {
                    $('#btnRegister').attr('disabled', false);
                }
            });
        }  
    });

    $('#btnForgotPassword').click(function() {
        email = $('#email').val();

        if(!validateEmail(email)) {
            showErrorToast("Lütfen geçerli bir email adresi girin.");
        } else {
            $.ajax({
                url: "/kullanici/sifremi-unuttum",
                type: "POST",
                async: true,
                cache: false,
                dataType: "json",
                headers: {
                    'X-Xsrftoken': getToken(),
                },
                data: {email: email},
                beforeSend: function() {
                    $('#btnForgotPassword').attr('disabled', true);
                },
                success: function(res) {
                    if(res.Result == 0)
                    {
                        showSuccessToast(res.Message);
                        setTimeout(function() { window.location = "/"; }, 5000);
                    } 
                    else 
                    {
                        showErrorToast(res.Message);
                    }

                },
                error: function(res) {
                    showErrorToast("Beklenmeyen bir hata oluştu.");
                },
                complete: function() {
                    $('#btnForgotPassword').attr('disabled', false);
                }
            });
        }
    });

    $('#btnNewPassword').click(function() {
        password = $('#password').val();
        passwordConfirm = $('#passwordConfirm').val();
        randomString = $('#randomString').val();

        if(password.length < 6 || password.length > 20) {
            showErrorToast("Şifre 6-20 karakter aralığında olmalıdır.");
        } else if(password != passwordConfirm) {
            showErrorToast("Girilen şifreler aynı değil.");
        } else if(randomString.length != 100) {
            showErrorToast("Beklenmeyen bir hata oluştu.");
        } else {
            $.ajax({
                url: "/kullanici/yeni-sifre",
                type: "POST",
                async: true,
                cache: false,
                dataType: "json",
                headers: {
                    'X-Xsrftoken': getToken(),
                },
                data: {password: password, randomString: randomString},
                beforeSend: function() {
                    $('#btnNewPassword').attr('disabled', true);
                },
                success: function(res) {
                    if(res.Result == 0)
                    {
                        showSuccessToast(res.Message);
                        setTimeout(function() { window.location = "/"; }, 5000);
                    } 
                    else 
                    {
                        showErrorToast(res.Message);
                        $('#btnNewPassword').attr('disabled', false);
                    }
                },
                error: function(res) {
                    showErrorToast("Beklenmeyen bir hata oluştu.");
                    $('#btnNewPassword').attr('disabled', false);
                }
            });
        }

    });

    $('#btnAddExchangeAccount').click(function() {
        exchangeID =  $('#exchange').val();
        name =  $('#name').val();
        apiKey = $('#apiKey').val();
        apiSecret = $('#apiSecret').val();

        if(!exchangeID) {
            showErrorToast("Lütfen hesap eklemek isteğiniz borsayı seçin.");
        } else if(name && name.length > 20) {
            showErrorToast("İsim alanı 20 karakterden uzun olamaz.");
        }  else if(apiKey.length < 32 || apiKey.length > 64) {
            showErrorToast("Lütfen geçerli bir api key değeri girin.");
        } else if(apiSecret.length < 32 || apiSecret.length > 64) {
            showErrorToast("Lütfen geçerli bir api secret değeri girin.");
        } else {
            $.ajax({
                url: "/borsa-hesabi/ekle",
                type: "POST",
                async: true,
                cache: false,
                dataType: "json",
                headers: {
                    'X-Xsrftoken': getToken(),
                },
                data: {exchangeID: exchangeID, name: name, apiKey: apiKey, apiSecret: apiSecret},
                beforeSend: function() {
                    $('#btnAddExchangeAccount').attr('disabled', true);
                },
                success: function(res) {
                    if(res.Result == 0)
                    {
                        $('#name').val('');
                        $('#apiKey').val('');
                        $('#apiSecret').val('');
                        showSuccessToast(res.Message);
                        refreshExchangeAccounts();
                    }  
                    else 
                    {
                        showErrorToast(res.Message);
                        $('#btnNewPassword').attr('disabled', false);
                    }
                },
                error: function(res) {
                    showErrorToast("Beklenmeyen bir hata oluştu.");
                    $('#btnAddExchangeAccount').attr('disabled', false);
                },
                complete: function() {
                    $('#btnAddExchangeAccount').attr('disabled', false);
                }
            });
        }
    });

    $('#editExchangeAccount').click(function() {
        exchangeAccountID = $('#editID').val();
        name = $('#editName').val();
        apiKey = $('#editKey').val();
        apiSecret = $('#editSecret').val();

        if(!exchangeAccountID) {
            showErrorToast("Hesap bulunamadı.");
            refreshExchangeAccounts();
        } else if(name && name.length > 20) {
            showErrorToast("İsim alanı 20 karakterden uzun olamaz.");
        }  else if(apiKey.length < 32 || apiKey.length > 64) {
            showErrorToast("Lütfen geçerli bir api key değeri girin.");
        } else if(apiSecret.length < 32 || apiSecret.length > 64) {
            showErrorToast("Lütfen geçerli bir api secret değeri girin.");
        } else {
            $.ajax({
                url: "/borsa-hesabi/duzenle",
                type: "POST",
                async: true,
                cache: false,
                dataType: "json",
                headers: {
                    'X-Xsrftoken': getToken(),
                },
                data: {exchangeAccountID: exchangeAccountID, name: name, apiKey: apiKey, apiSecret: apiSecret},
                beforeSend: function() {
                    $('#editExchangeAccount').attr('disabled', true);
                },
                success: function(res) {
                    if(res.Result == 0)
                    {
                        $('#editID').val('');
                        $('#editName').val('');
                        $('#editKey').val('');
                        $('#editSecret').val('');
                        showSuccessToast(res.Message);
                        refreshExchangeAccounts();
                        $('.modal.open').modal('close');
                    }  
                },
                error: function(res) {
                    showErrorToast("Beklenmeyen bir hata oluştu.");
                    $('#editExchangeAccount').attr('disabled', false);
                },
                complete: function() {
                    $('#editExchangeAccount').attr('disabled', false);
                }
            });
        }
    });

    $('body').on('click', '.exchangeAccountSwitch', function() {
        if(this.checked) {
            isActive = 1;
            message = "aktif";
        } else {
            isActive = 0;
            message = "pasif";
        }

        if (!confirm('Hesabı ' + message + ' hale getirmek istediğinizden emin misiniz?')) {
            this.checked = !isActive;
            return false;
        }

        exchangeAccountID = $(this).data("id");
        
        if(Number.isInteger(exchangeAccountID)) {
            $.ajax({
                url: "/borsa-hesabi/aktiflestir",
                type: "POST",
                async: true,
                cache: false,
                dataType: "json",
                headers: {
                    'X-Xsrftoken': getToken(),
                },
                data: {exchangeAccountID: exchangeAccountID, isActive: isActive},
                success: function(res) {
                    if(res.Result == 0)
                    {
                        showSuccessToast(res.Message);
                    } 
                    else 
                    {
                        showErrorToast(res.Message);
                    }
                },
                error: function(res) {
                    showErrorToast("Beklenmeyen bir hata oluştu.");
                }
            });
        } else {
            showErrorToast("Beklenmeyen bir hata oluştu.");
        }
    });

    $('body').on('click', '.btnEditExchangeAccount', function() {
        exchangeAccountID = $(this).data("id");
        if(Number.isInteger(exchangeAccountID)) {
            $.ajax({
                url: "/borsa-hesabi/duzenle",
                type: "GET",
                async: true,
                cache: false,
                dataType: "json",
                headers: {
                    'X-Xsrftoken': getToken(),
                },
                data: {exchangeAccountID: exchangeAccountID},
                success: function(res) {
                    if(res) 
                    {
                        id = res["ID"];
                        exchangeName = res["ExchangeName"];
                        name = res["Name"];
                        key = res["Key"];
                        secret = res["Secret"];

                        $('#editID').val(id);
                        $('#editExchangeName').val(exchangeName);
                        $('#editName').val(name);
                        $('#editKey').val(key);
                        $('#editSecret').val(secret);
                        Materialize.updateTextFields();
                    } 
                    else 
                    {
                        showErrorToast("Hesap bulunamadı.");
                        refreshExchangeAccounts();
                    }
                },
                error: function(res) {
                    showErrorToast("Beklenmeyen bir hata oluştu.");
                }
            });
        } else {
            showErrorToast("Beklenmeyen bir hata oluştu.");
        }
    });

    $('body').on('click', '.btnDeleteExchangeAccount', function() {
        if (!confirm('Hesabı silmek istediğinizden emin misiniz?')) return false;
        
        exchangeAccountID = $(this).data("id");
        
        if(Number.isInteger(exchangeAccountID)) {
            $.ajax({
                url: "/borsa-hesabi/sil",
                type: "POST",
                async: true,
                cache: false,
                dataType: "json",
                headers: {
                    'X-Xsrftoken': getToken(),
                },
                data: {exchangeAccountID: exchangeAccountID},
                success: function(res) {
                    if(res.Result == 0)
                    {
                        showSuccessToast(res.Message);
                        refreshExchangeAccounts();
                    } 
                    else 
                    {
                        showErrorToast(res.Message);
                    }
                },
                error: function(res) {
                    showErrorToast("Beklenmeyen bir hata oluştu.");
                }
            });
        } else {
            showErrorToast("Beklenmeyen bir hata oluştu.");
        }
    });

    $('#exchangeAccounts').on('change', function() {
        exchangeAccountID = parseInt($(this).val());
        
        if(Number.isInteger(exchangeAccountID)) {
            loadChart(exchangeAccountID);
        } else {
            showErrorToast("Beklenmeyen bir hata oluştu.");
        }
    });
});

function refreshExchangeAccounts() {
    $.ajax({
        url: "/borsa-hesabi/listele",
        type: "GET",
        async: true,
        cache: false,
        dataType: "json",
        headers: {
            'X-Xsrftoken': getToken(),
        },
        success: function(res) {
            if(res.length > 0) {
                var accountsHtml = '';
                for (var i = 0; i < res.length; i++) {
                    id = res[i]["ID"];
                    exchangeName = res[i]["ExchangeName"];
                    name = res[i]["Name"];
                    key = res[i]["Key"];
                    createdAt = res[i]["CreatedAt"];
                    isActive = res[i]["IsActive"];

                    if(name.length > 0) {
                        nameString = name;
                    } else {
                        nameString = key.substring(0, 12) + "...";
                    }

                    if(isActive) { 
                        switchCheckedString = "checked"; 
                    } else {
                        switchCheckedString = ""; 
                    } 
                    
                    accountsHtml += '<tr>';
                    accountsHtml += '<td>' + exchangeName + '</td>';
                    accountsHtml += '<td>' + nameString + '</td>';
                    accountsHtml += '<td>' + createdAt + '</td>';
                    accountsHtml += '<td><div class="switch"><label>Pasif<input data-id="' + id +'" class="exchangeAccountSwitch" ' + switchCheckedString + ' type="checkbox"><span class="lever"></span>Aktif</label></div></td>';
                    accountsHtml += '<td><a data-id="' + id +'" class="btnEditExchangeAccount btn-floating waves-effect waves-light cyan modal-trigger" href="#modalEditExchangeAccount"><i class="material-icons">mode_edit</i></a><a data-id="' + id +'" class="btnDeleteExchangeAccount btn-floating waves-effect waves-light red accent-2"><i class="material-icons">clear</i></a></td>';
                    accountsHtml += '</tr>';
                }

                $('#tbodyExchangeAccounts').html(accountsHtml);
                $('#tableExchangeAccounts').fadeIn();
                $('#divNoExchangeAccounts').fadeOut();
            } else {
                $('#tableExchangeAccounts').fadeOut();
                $('#divNoExchangeAccounts').fadeIn();
            }
        },
        error: function(res) {
            showErrorToast("Beklenmeyen bir hata oluştu.");
        }
    });
}

function getExchangeAccounts() {
    $.ajax({
        url: "/borsa-hesabi/listele",
        type: "GET",
        async: true,
        cache: false,
        dataType: "json",
        headers: {
            'X-Xsrftoken': getToken(),
        },
        success: function(res) {
            var accountsHtml = '';
            if(res.length > 0) {
                for (var i = 0; i < res.length; i++) {
                    id = res[i]["ID"];
                    exchangeName = res[i]["ExchangeName"];
                    name = res[i]["Name"];
                    key = res[i]["Key"];
                    isActive = res[i]["IsActive"];

                    if(name.length > 0) {
                        nameString = name;
                    } else {
                        nameString = key.substring(0, 12) + "...";
                    }

                    if(isActive) {
                        isActiveText = "Aktif";
                    } else {
                        isActiveText = "Pasif";
                    }

                    accountsHtml += '<option value="' + id  + '">' + exchangeName + ' - ' + nameString + ' (' + isActiveText + ')</option>';
                }

                $("#exchangeAccounts").append(accountsHtml);
                $("#exchangeAccounts").material_select();
                $('#divExchangeAccounts').fadeIn();
            } else {
                $("#divExchangeAccounts").fadeOut();
                $('#divNoExchangeAccounts').fadeIn();
            }
        },
        error: function(res) {
            showErrorToast("Beklenmeyen bir hata oluştu.");
        }
    });
}

function getBalances() {
    $.ajax({
        url: "/bakiye/listele",
        type: "GET",
        async: true,
        cache: false,
        dataType: "json",
        headers: {
            'X-Xsrftoken': getToken(),
        },
        success: function(res) {
            if(res.length > 0) {
                var balancesHtml = '';
                var btcPricesHtml = '';
                var btcUsdtBinance;
                for (var i = 0; i < res.length; i++) {

                    exchange = res[i].Exchange;

                    if(exchange === "Binance") {
                        btcUsdtBinance = res[i].BtcUsdt;
                    }

                    name = res[i].Name;
                    logoImagePath = res[i].LogoImagePath;
                    totalBtc = res[i].TotalBtc;
                    totalUsdt = res[i].TotalUsdt;

                    balancesHtml += '<div class="card-panel">';
                    balancesHtml +=     '<div>';
                    balancesHtml +=         '<span style="line-height: 50px; font-size: 16px; font-weight: bold;">' + name +'</span>';
                    balancesHtml +=         '<span class="right">';
                    balancesHtml +=             '<img width="100" src="' + logoImagePath + '">';
                    balancesHtml +=         '</span>';
                    balancesHtml +=     '</div>';
                    balancesHtml +=     '<div class="row">';
                    balancesHtml +=         '<table class="striped">';
                    balancesHtml +=             '<thead>';
                    balancesHtml +=                 '<tr>';
                    balancesHtml +=                     '<th>Kripto Para</th>';
                    balancesHtml +=                     '<th>Boşta</th>';
                    balancesHtml +=                     '<th>Toplam</th>';
                    balancesHtml +=                     '<th>Birim Fiyatı</th>';
                    balancesHtml +=                     '<th>BTC Karşılığı</th>';
                    balancesHtml +=                     '<th>USDT Karşılığı</th>';
                    balancesHtml +=                 '</tr>';
                    balancesHtml +=             '</thead>';
                    balancesHtml +=             '<tbody>';

                    balances =  res[i].Balances;
 
                    for (var j = 0; j < balances.length; j++) {
                        currency = balances[j].Currency;
                        tradeUrl = balances[j].TradeURL;
                        free = balances[j].Free;
                        total = balances[j].Total;
                        btcEquivalent = balances[j].BtcEquivalent;
                        usdtEquivalent = balances[j].UsdtEquivalent;
                        unitPrice = balances[j].UnitPrice;

                        balancesHtml +=             '<tr>';
                        if(tradeUrl.length > 0) {
                            balancesHtml +=             '<td><a target="_blank" href="'+ tradeUrl +'">' + currency + '</a></td>';
                        } else {
                            balancesHtml +=             '<td>' + currency + '</td>';
                        }
                        balancesHtml +=                 '<td>' + free.toFixed(8) + ' ' + currency + '</td>';
                        balancesHtml +=                 '<td>' + total.toFixed(8) + ' ' + currency + '</td>';
                        balancesHtml +=                 '<td>' + unitPrice.toFixed(8) +' BTC</td>';
                        balancesHtml +=                 '<td>' + btcEquivalent.toFixed(8) +' BTC</td>';
                        balancesHtml +=                 '<td>' + usdtEquivalent.toFixed(2) +' USDT</td>';
                        balancesHtml +=             '</tr>';
                    }

                    balancesHtml +=                 '<tr class="tr-total">';
                    balancesHtml +=                     '<td colspan=4><b>Toplam</b></td>';
                    balancesHtml +=                     '<td><b>' +  totalBtc.toFixed(8) + ' BTC</b></td>';
                    balancesHtml +=                    '<td><b>' +  totalUsdt.toFixed(2) + ' USDT</b></td>';
                    balancesHtml +=                 '</tr>';
                    balancesHtml +=             '</tbody>';
                    balancesHtml +=         '</table>';
                    balancesHtml +=     '</div>';
                    balancesHtml += '</div>';

                }

                btcPricesHtml += '<div class="card gradient-45deg-amber-amber gradient-shadow min-height-100 white-text">';
                btcPricesHtml +=     '<div class="padding-1">';
                btcPricesHtml +=         '<h5>Binance</h5><p>1 BTC = ' + btcUsdtBinance.toFixed(2) + ' USDT</p>';
                btcPricesHtml +=     '</div>';
                btcPricesHtml += '</div>';

                $('#divBalances').html(btcPricesHtml + balancesHtml);
                $('#divBalances').fadeIn();
                $('#divNoActiveExchangeAccounts').fadeOut();
            } else {
                $('#divBalances').fadeOut();
                $('#divNoActiveExchangeAccounts').fadeIn();
            }
        },
        error: function(res) {
            showErrorToast("Beklenmeyen bir hata oluştu.");
        }
    });
}

function loadChart(exchangeAccountID) {
    $.ajax({
        url: "/bakiye/grafik/goruntule",
        type: "GET",
        async: true,
        cache: false,
        dataType: "json",
        headers: {
            'X-Xsrftoken': getToken(),
        },
        data: {exchangeAccountID: exchangeAccountID},
        success: function(res) {
            if(res.length > 2) {
                var chartData = res;

                var chart = AmCharts.makeChart("balanceChart", {
                    "type": "serial",
                    "theme": "light",
                    "language": "tr",
                    "mouseWheelZoomEnabled": true,
                    "legend": {
                        "useGraphSettings": true
                    },
                    "dataProvider": chartData,
                    "synchronizeGrid":true,
                    "dataDateFormat": "YYYY-MM-DD HH:NN:SS",
                    "valueAxes": [{
                        "id":"v1",
                        "axisColor": "#FF6600",
                        "axisThickness": 2,
                        "axisAlpha": 1,
                        "position": "left"
                    }, {
                        "id":"v2",
                        "axisColor": "#FCD202",
                        "axisThickness": 2,
                        "axisAlpha": 1,
                        "position": "right"
                    }],
                    "graphs": [{
                        "valueAxis": "v1",
                        "lineColor": "#FF6600",
                        "bullet": "round",
                        "bulletBorderThickness": 1,
                        "hideBulletsCount": 100,
                        "title": "BTC",
                        "valueField": "Btc",
                        "fillAlphas": 0
                    }, {
                        "valueAxis": "v2",
                        "lineColor": "#FCD202",
                        "bullet": "round",
                        "bulletBorderThickness": 1,
                        "hideBulletsCount": 100,
                        "title": "USDT",
                        "valueField": "Usdt",
                        "fillAlphas": 0
                    }],
                    "chartScrollbar": {},
                    "chartCursor": {
                        "cursorPosition": "mouse"
                    },
                    "categoryField": "Date",
                    "categoryAxis": {
                        "minPeriod": "mm",
                        "parseDates": true,
                        "axisColor": "#DADADA",
                        "minorGridEnabled": true
                    },
                    "export": {
                        "enabled": true,
                        "position": "bottom-right"
                        }
                });

                chart.addListener("rendered", zoomChart);
                zoomChart();
                
                function zoomChart() {
                    chart.zoomToIndexes(chartData.length - 96, chartData.length - 1);
                }
                $('#divNoChartData').fadeOut();
                $("#balanceChart").height(400);
            } else {
                $('#divNoChartData').fadeIn();
            }
        },
        error: function(res) {
            showErrorToast("Beklenmeyen bir hata oluştu.");
        }
    });
}

function getOrders() {
    $.ajax({
        url: "/emir/listele",
        type: "GET",
        async: true,
        cache: false,
        dataType: "json",
        headers: {
            'X-Xsrftoken': getToken(),
        },
        success: function(res) {
            if(res && res.length > 0) {
                var ordersHtml = '';
                for (var i = 0; i < res.length; i++) {
                    clientOrderId = res[i].clientOrderId;
                    executedQty = res[i].executedQty;
                    icebergQty = res[i].icebergQty;
                    orderId = res[i].orderId;
                    origQty = res[i].origQty;
                    price = res[i].price;
                    side = res[i].side;
                    status = res[i].status;
                    stopPrice = res[i].stopPrice;
                    symbol = res[i].symbol;
                    type = res[i].type;

                    if(side == "SELL") {
                        sideText = "Satış";
                    } else if(side == "BUY") {
                        sideText = "Alış";
                    } else {
                        sideText = side;
                    }

                    ordersHtml += '<tr>';
                    ordersHtml +=     '<td>' + symbol + '</td>';
                    ordersHtml +=     '<td>' + sideText + '</td>';
                    ordersHtml +=     '<td>' + capitalizeFirstLetter(type) + '</td>';
                    ordersHtml +=     '<td>' + origQty + '</td>';
                    ordersHtml +=     '<td>' + price + '</td>';
                    ordersHtml += '</tr>';
                }
            
                $('#tbodyOrders').html(ordersHtml);
                $('#tableOrders').fadeIn();
                $('#divNoOrders').fadeOut();
            } else {
                $('#tableOrders').fadeOut();
                $('#divNoOrders').fadeIn();
            }
        },
        error: function(res) {
            showErrorToast("Beklenmeyen bir hata oluştu.");
        }
    });
}

function validateEmail(email) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}

function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1).toLowerCase();
}

function showErrorToast(message) {
    iziToast.error({
        title: 'Hata',
        position: 'topRight',
        message: message
    });
}

function showSuccessToast(message) {
    iziToast.success({
        title: 'Başarılı',
        position: 'center',
        message: message
    });
}

function getToken() {
    var xsrf = $.cookie('_xsrf');
    if (xsrf !== null) {
        xsrf = xsrf.split('|')[0];
        return base64Decode(xsrf);
    }
    return "";
}

function base64Decode(data) {
    var b64 = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=';
    var o1, o2, o3, h1, h2, h3, h4, bits, i = 0,
        ac = 0,
        dec = '',
        tmp_arr = [];
    if (!data) {
        return data;
    }
    data += '';
    do { 
        h1 = b64.indexOf(data.charAt(i++));
        h2 = b64.indexOf(data.charAt(i++));
        h3 = b64.indexOf(data.charAt(i++));
        h4 = b64.indexOf(data.charAt(i++));
        bits = h1 << 18 | h2 << 12 | h3 << 6 | h4;
        o1 = bits >> 16 & 0xff;
        o2 = bits >> 8 & 0xff;
        o3 = bits & 0xff;
        if (h3 == 64) {
            tmp_arr[ac++] = String.fromCharCode(o1);
        } else if (h4 == 64) {
            tmp_arr[ac++] = String.fromCharCode(o1, o2);
        } else {
            tmp_arr[ac++] = String.fromCharCode(o1, o2, o3);
        }
    } while (i < data.length);
    dec = tmp_arr.join('');
    return dec.replace(/\0+$/, ''); 
}