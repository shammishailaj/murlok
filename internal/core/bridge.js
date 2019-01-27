const murlok = {
    url: '{{.LocalServerURL}}',
    rpcURL: '{{.LocalServerURL}}/murlok',
    onEvent: function (name = '', arg = {}) {
        console.log('murlok event not handled => ' + name + ': ' + JSON.stringify(arg));
    },
};