$(function () {
    $("#poke").click(poke);
    $("#say").click(say);
    
    var ws = new WebSocket("ws://localhost:8080/ws/example");
    ws.onopen = function () {
        add('system', 'Connected');
    }

    ws.onclose = function () {
        add('system', 'Disconnected');
    }

    ws.onmessage = function (message) {
        console.log("message", message);
        if (typeof message.data != 'undefined') {
            var data = $.parseJSON(message.data);
            if (data.topic && data.message) {
                add(data.topic, data.message);
            }
        }
    }

    ws.onerror = function () {
        console.error("websocket error!", arguments);
    }
});


function poke() {
    $.post('/api/poke', {'user': $('#user').val()});
}

function say() {
    $.post('/api/say', {'user': $('#user').val(),
                        'message': $('#message').val()});
}

function add(type, messagedata) {
    var message;
    if (messagedata != null && typeof messagedata == 'object') {
        message = $('<span>', {'class': 'from'}).html(messagedata.from);
        if (messagedata.message != null) {
            message.append($('<span>').html(' said '));
            message.append($('<span>', {'class': 'message'}).html(messagedata.message));
        }

    } else {
        message = messagedata;
    }
    $("#events").append($("<li>", {'class': type}).html(message));
}
