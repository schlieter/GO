const socket = io();

new Vue({
    el: '#chat-app',
    data: {
        mensaje: '',
        mensajes: []
    },
    created() {
        const vm = this;
        socket.on('chat mensaje', function (msg) {
            vm.mensajes.push({
                text:msg,
                date : new Date().toLocaleDateString()
            })
        })
    },
    methods: {
        enviarMensaje() {
            socket.emit('chat mensaje', this.mensaje);
            this.mensaje = '';
        }
    }
});
