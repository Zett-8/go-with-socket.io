// init form

document.getElementById('root').innerHTML = `
    <form id="fm" action="">
    <input id="m" autocomplete="off" /><button>Send</button>
    </form>
    <ul id="messages"></ul>
`


let MessageArea = document.getElementById('messages')
let InputForm = document.getElementById('m')
const roomName = document.getElementsByTagName('title')[0].innerText

document.body.addEventListener('onload', () => InputForm.focus())



const socket = io.connect()

if (!socket.connect) {
    console.log('failed')
}




socket.emit('room_in', roomName)

socket.on('reply', function(msg){
    MessageArea.innerHTML += '<li>' + msg + '</li>'
})

document.getElementById('fm').addEventListener('submit', e => {
    e.preventDefault()
    const msg = InputForm.value
    const data = {
        room: roomName,
        message: msg
    }

    socket.emit('notice', JSON.stringify(data))

    InputForm.value = ''
})