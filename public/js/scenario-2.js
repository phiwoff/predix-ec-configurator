$('#form-setup').submit(function (e) {
    // disabled inputs are not posted
    $('#form-setup :input').prop('disabled', false);
    if (!validateForm()) {
        e.preventDefault();
        return;
    } else {
        $('#form-setup')[0].submit();
    }
});

function validateForm() {
    var isValid = true;
    var elementsToBeValidated = ['uaaZoneID', 'uaaServiceURI', 'ecZoneID', 'ecServiceURI',
        'ecAdmToken', 'ecIDS', 'resourceHost', 'resourcePort']
    $.each(elementsToBeValidated, function (index, value) {
        var elem = '#' + value
        console.log(value);
        if ($(elem).val() === "") {
            console.log(value + "is blank");
            isValid = false;
        }
    });
    return isValid;
}

$(document).ready(function () {
    $('#predixOrgs').change(function () {
        $('#predixOrgSpaces').children('option:not(:first)').remove();
        $('#predixSpaceApps').children('option:not(:first)').remove();
        resetAll();
        var selectedOrgGuid = $('#predixOrgs').val();
        if (selectedOrgGuid != "") {
            doSpacesRequest(selectedOrgGuid);
        }
    });

    $('#predixOrgSpaces').change(function () {
        $('#predixSpaceApps').children('option:not(:first)').remove();
        resetAll();
        var selectedSpaceGuid = $('#predixOrgSpaces').val();
        if (selectedSpaceGuid != '') {
            doAppsRequest(selectedSpaceGuid);
        }
    });

    $('#predixSpaceApps').change(function () {
        resetAll();
        var selectedAppGuid = $('#predixSpaceApps').val();
        if (selectedAppGuid != '') {
            doAppEnvRequest(selectedAppGuid);
        }
    });

    $('#predixDataSource').change(function () {
        var selectedAppGuid = $('#predixSpaceApps').val();
        var selectedDataSource = $('#predixDataSource').val();
        if (selectedAppGuid != '' && selectedDataSource != '') {
            doServiceInfoRequest(selectedAppGuid, selectedDataSource);
        }
    });
});

function doSpacesRequest(orgGuid) {
    $.ajax({
        url: '/retrieveOrgSpaces',
        type: 'post',
        dataType: 'json',
        data: { ajax_post_data: orgGuid },
        success: function (data) {
            $('#predixOrgSpaces').children('option:not(:first)').remove();
            //$("#predixSpaces").find('option').remove().end();
            for (var i = 0; i < data.length; i++) {
                $('#predixOrgSpaces')
                    .append($('<option></option>')
                        .attr('value', data[i].guid)
                        .text(data[i].name));
            }
        },
    });
}

function doAppsRequest(spaceGuid) {
    $.ajax({
        url: '/retrieveSpaceApps',
        type: 'post',
        dataType: 'json',
        data: { ajax_post_data: spaceGuid },
        success: function (data) {
            $("#predixSpaceApps").children('option:not(:first)').remove();
            for (var i = 0; i < data.length; i++) {
                $('#predixSpaceApps')
                    .append($('<option></option>')
                        .attr('value', data[i].guid)
                        .text(data[i].name));
            }

        }
    });
}

function doAppEnvRequest(appGuid) {
    $.ajax({
        url: '/retrieveAppEnv',
        type: 'post',
        dataType: 'json',
        data: { ajax_post_data: appGuid },
        success: function (data) {
            validateBindings(data['VCAP_SERVICES'])
        }
    });
}

function doServiceInfoRequest(selectedAppGuid, selectedDataSource) {
    $.ajax({
        url: '/retrieveAppEnv',
        type: 'post',
        dataType: 'json',
        data: { ajax_post_data: selectedAppGuid },
        success: function (data) {
            var ss = data['VCAP_SERVICES'][selectedDataSource][0]
            var credentials = ss.credentials;
            if (isPostgreSQL(selectedDataSource)) {
                bindPostgreSQLServiceInfo(credentials);
            } else if (isRabbitMQ) {
                bindRabbitMQServiceInfo(credentials);
            }
        }
    });
}

function isPostgreSQL(datasource) {
    return (~datasource.indexOf('postgres')) ? true : false;
}

function isRabbitMQ(datasource) {
    return (~datasource.indexOf('rabbitmq')) ? true : false;
}

function bindPostgreSQLServiceInfo(data) {
    $('#resourceHost').val(data.hostname);
    $('#resourcePort').val(data.port);
}

function bindRabbitMQServiceInfo(data) {
    $('#resourceHost').val(data.protocols.amqp.host);
    $('#resourcePort').val(data.protocols.amqp.port);
}

function showSuccessMessage(message) {
    $('#successMessage').html(message);
    $('#successMessage').show();
    return;
}
function showWarningMessage(message) {
    $('#warningMessage').html(message);
    $('#warningMessage').show();
    return;
}

function showErrorMessage(message) {
    $('#errorMessage').html(message);
    $('#errorMessage').show();
    return;
}

function validateBindings(data) {
    if (!validateUAABinding(data)) return;
    if (!validateECBinding(data)) return;
    if (!validatePredixDataSourcesBinding(data)) return;
}

function validateUAABinding(data) {
    if (data.hasOwnProperty('predix-uaa')) {
        decodeUAA(data['predix-uaa']);
    } else {
        var errString = "<strong>ERROR!</strong> UAA with no binding to the selected App! Select a valid application.";
        showErrorMessage(errString);
        return false;
    }
    return true;
}

function decodeUAA(uaaData) {
    var uaaServiceName = uaaData[0].name;
    $('#uaaServiceName').html("[Service Name: <u>" + uaaServiceName + "</u>]");

    var uaaZoneID = uaaData[0].credentials.zone['http-header-value'];
    $('#uaaZoneID').val(uaaZoneID);

    var uaaServiceURI = uaaData[0].credentials.uri;
    $('#uaaServiceURI').val(uaaServiceURI);
}

function validateECBinding(data) {
    if (data.hasOwnProperty('enterprise-connect')) {
        decodeEC(data['enterprise-connect']);
    } else {
        var errString = "<strong>ERROR!</strong> Enterprise-Connect with no binding to the selected App! Select a valid application.";
        showErrorMessage(errString);
        return false;
    }
    return true;
}

function decodeEC(ecData) {
    var ecServiceName = ecData[0].name
    $('#ecServiceName').html("[Service Name: <u>" + ecServiceName + "</u>]");

    var ecZoneID = ecData[0].credentials.zone['http-header-value'];
    $('#ecZoneID').val(ecZoneID);

    var ecServiceURI = ecData[0].credentials['service-uri']
    $('#ecServiceURI').val(ecServiceURI);

    var ecAdmToken = ecData[0].credentials['ec-info']['adm_tkn'];
    $('#ecAdmToken').val(ecAdmToken);

    var ecIDS = ecData[0].credentials['ec-info']['ids'];
    $('#ecIDS').val(ecIDS)
}

function validatePredixDataSourcesBinding(data) {
    var isBinded = false;
    Object.keys(data).some(function (s) {
        if (isPostgreSQL(s) || isRabbitMQ(s)) {
            isBinded = true;
            var d = data[s][0];
            $('#predixDataSource')
                .append($('<option></option>')
                    .attr('value', d.label)
                    .text(d.name));
        }
    });
    if (!isBinded) {
        var errString = "<strong>No valid Predix Data Source (Postgres or RabbitMQ) with binding to the selected app. </strong>";
        showErrorMessage(errString);
        return false;
    }
    return true;
}

function resetAll() {
    // MESSAGES
    $('#successMessage').hide();
    $('#warningMessage').hide();
    $('#errorMessage').hide();
    // UAA
    $('#uaaServiceName').html('');
    $('#uaaZoneID').val('');
    $('#uaaServiceURI').val('');
    // EC
    $('#ecServiceName').html('');
    $('#ecZoneID').val('');
    $('#ecServiceURI').val('');
    $('#ecAdmToken').val('');
    $('#ecIDS').val('');
    // Predix DataSource
    $('#predixDataSource').children('option:not(:first)').remove();
    $('#resourceHost').val('');
    $('#resourcePort').val('');
}