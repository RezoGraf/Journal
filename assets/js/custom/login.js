/**
 * Created by farex on 09.05.2017.
 */
Vue.component('modal-call', {
    delimiters: ['{', '}'],
    template: '#modal-template-call',
    data: function() {
        return {
            items: null,
            pwd: '',
            showModalCall: true,
            selectedReg: ''
        };
    },
    mounted: function() {
         this.setDefaultItems();
    },
    methods: {
        setDefaultItems: function () {
            this.$http.get('/catalog_med_reg_get').then(data => {
                this.items = JSON.parse(data.body);
                console.log(data.body)
        });
        },
        entrance: function () {
            if (this.selectedReg == '' || this.password == ''){
                alertify.alert("Заполнены не все поля.", function() {
                    alertify.message('OK');
                });
                return;
            }
            var me = this
            console.log(this.selectedReg, this.pwd)
            $.ajax({
                url: '/login_auth',
                type: 'post',
                data: {
                    id: me.selectedReg,
                    pwd: me.pwd
                }
            }).done(function(data, status) {
                if (status == "success") {
                    window.location.href = "take_to_repair";
                } else {
                    alertify.error('Произошла ошибка на сервере');
                }
            })
        }
        }

})

var demo = new Vue({
    el: '#demo',
    delimiters: ['{', '}'],
    data: {
        showModalCall: true
    }
})