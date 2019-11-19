const socket = io.connect();

if (!socket.connect) {
    console.log('failed')
}

socket.on('reply', function(msg){
    document.getElementById('messages').innerHTML += '<li>' + msg + '</li>';
});

$('form').submit(function(){
    socket.emit('notice', $('#m').val());
    $('#m').val('');
    return false;
});

$('#bt').click(function (){
    console.log('button click')
    socket.emit('bye', null);
    return false;
});



///// prac
document.getElementById('num').innerText = 'change!'
