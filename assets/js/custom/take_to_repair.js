Vue.component('crumb-up', {
    delimiters: ['{', '}'],
    template: '#crumb-up',
    props: {
        selected_status: String
    },
    data: function() {
        return {

        };
    }
})

Vue.component('selected-status-patient', {
    delimiters: ['{', '}'],
    template: '#selected-status-patient',
    data: function() {
        return {
            selected_status: 'Пациенты в очереди'
        };
    },
    watch: {
        selected_status: function() {
            demo.setDefaultItems(this.selected_status)
        }
    }
})

Vue.component('modal-nar', {
    delimiters: ['{', '}'],
    template: '#modal-template-nar',
    data: function() {
        return {
            edit_vrach_ortoped: '',
            edit_vrach_technic: '',
            edit_date_open_nar: '',
            edit_date_start_production: '',
            edit_date_close_nar: '',
            edit_sum: '',
            edit_number_nar: '',
            add_nar: false,
            flag: false,
            lastIdTable: '',
            selectedVrachOrotped: '',
            selectedVrachTechnic: '',
            buttonAddRow: true,
            number_nar: '',
            date_open_nar: '',
            date_start_production: '',
            date_close_nar: '',
            sum: '',
            addItemRow: false,
            addItemRowSelected: false,
            edit: {},
            addNarRow: {},
            editIndex: undefined,
            items: null,
            // selected: 'ВТ',
            items1: null,
            itemsFioVrach: null,
            itemsFioVrachTechnic: null,
            editIndex: null,
            currentId: '', //Индекс редактируемой записи
            full_name: demo.full_name,
            date_birth: demo.date_birth,
            sum_all: 0.0,
            selected_fio_reg: navBar.nameReg
        };
    },
    mounted: function() {
        this.setDefaultItems();
    },
    watch: {
        sum: function() {
            this.sum = this.sum.toString().replace(/\$|\./g, '');
            if (isNaN(this.sum))
                this.sum = "0";
            sign = (this.sum == (this.sum = Math.abs(this.sum)));
            this.sum = Math.floor(this.sum * 100 + 0.50000000001);
            cents = this.sum % 100;
            this.sum = Math.floor(this.sum / 100).toString();
            if (cents < 10)
                cents = "0" + cents;
            for (var i = 0; i < Math.floor((this.sum.length - (1 + i)) / 3); i++)
                this.sum = this.sum.substring(0, this.sum.length - (4 * i + 2)) + '.' +
                this.sum.substring(this.sum.length - (4 * i + 2));
        },
        edit_sum: function() {
            this.edit_sum = this.edit_sum.toString().replace(/\$|\./g, '');
            if (isNaN(this.edit_sum))
                this.edit_sum = "0";
            sign = (this.edit_sum == (this.edit_sum = Math.abs(this.edit_sum)));
            this.edit_sum = Math.floor(this.edit_sum * 100 + 0.50000000001);
            cents = this.edit_sum % 100;
            this.edit_sum = Math.floor(this.edit_sum / 100).toString();
            if (cents < 10)
                cents = "0" + cents;
            for (var i = 0; i < Math.floor((this.edit_sum.length - (1 + i)) / 3); i++)
                this.edit_sum = this.edit_sum.substring(0, this.edit_sum.length - (4 * i + 2)) + '.' +
                this.edit_sum.substring(this.edit_sum.length - (4 * i + 2));
            console.log("!!!")
        }
    },
    methods: {
        setDefaultItems: function() {
            console.log(demo.currentId)
            this.$http.get('/edit_listNar_get?id=' + demo.currentId + '').then(data => {
                this.items = JSON.parse(data.body);
            var json = JSON.parse(data.body);
             this.sum_all = 0.0;
            for(i = 0; i <= json.length; i++){
                this.sum_all += parseFloat(json[i].sum)
            }
            });
            this.$http.get('/listlgotcat').then(data => {
                this.items1 = JSON.parse(data.body);;
            });
            this.$http.get('/edit_fiovrach_get').then(data => {
                this.itemsFioVrach = JSON.parse(data.body);
            });
            this.$http.get('/edit_fiovrach_technic_get').then(data => {
                this.itemsFioVrachTechnic = JSON.parse(data.body);
            });

        },

        noNumber: function(evt) {
            var regex = new RegExp("-?(0|([1-9]\d*))(\.\d+)?");
            var key = String.fromCharCode(!evt.charCode ? evt.which : evt.charCode);
            if (!regex.test(key)) {
                event.preventDefault();
                return false;
            }
        },
        inputValueFilterSum: function() {
            // $('#sum1').priceFormat({
            //     prefix: '',
            //     allowNegative: true
            // });
        },
        inputValueFilterSumSave: function() {
            // this.sum = $('#sum1').priceToFloat()
        },
        btn_add_nar: function() {
            this.add_nar = true;
        },
        close_nar: function() {
            this.add_nar = false;
            this.lastIdTable = "";
            this.number_nar = "";
            this.date_open_nar = "";
            this.selectedVrachOrotped = "";
            this.date_start_production = "";
            this.selectedVrachTechnic = "";
            this.date_close_nar = "";
            this.sum = "";
        },
        successAddNar: function() {
            this.currentId = demo.currentId;
            var me = this;
            this.sum = String(this.sum).replace(/(\d+)\.(?=\d+\.\d+)/g, '$1');
            console.log(this.sum)
            if (parseFloat(this.sum) > 999999.99) {
                alertify.alert("Сумма не может быть больше 999999.99 !", function() {
                    alertify.message('OK');
                });
                return;
            }
            if (me.number_nar == '' || me.date_open_nar == '' || me.selectedVrachOrotped == '' || me.sum == '' || me.sum == '0.0' || me.sum == '0') {
                alertify.alert("Заполнены не все поля.", function() {
                    alertify.message('OK');
                });
                return;
            }

            $.ajax({
                url: '/edit_add_nar',
                type: 'post',
                data: {
                    id_patient: me.currentId,
                    vrach_ortoped: me.selectedVrachOrotped,
                    vrach_technic: me.selectedVrachTechnic,
                    number_nar: me.number_nar,
                    date_open_nar: me.date_open_nar,
                    date_start_production: me.date_start_production,
                    date_close_nar: me.date_close_nar,
                    sum: me.sum,
                    name_reg: navBar.nameReg
                }
            }).done(function(data, status) {
                console.log(status)
                if (status == 'success') {
                    me.setDefaultItems();
                    me.addItemRow = false;
                    me.number_nar = "";
                    me.date_open_nar = "";
                    me.selectedVrachOrotped = "";
                    me.date_start_production = "";
                    me.selectedVrachTechnic = "";
                    me.date_close_nar = "";
                    me.sum = "";
                    me.edit_sum = "";
                    me.buttonAddRow = true;
                    me.flag = false
                    alertify.set({ delay: 10000 });
                    alertify.success('Наряд добавлен');
                } else {
                    alertify.error('Произошла ошибка на сервере');
                }
            })
        },
        cancelEdit: function(item, $index, e) {
            this.edit = {};
            this.setDefaultItems();
            this.addItemRow = false;
            this.number_nar = "";
            this.date_open_nar = "";
            this.selectedVrachOrotped = "";
            this.date_start_production = "";
            this.selectedVrachTechnic = "";
            this.date_close_nar = "";
            this.sum = "";
            this.edit_sum = "";
            this.buttonAddRow = true;
            this.flag = false
        },
        cancelAddItemRowSelected: function() {
            this.addItemRow = false;
            this.lastIdTable = "";
            this.number_nar = "";
            this.date_open_nar = "";
            this.selectedVrachOrotped = "";
            this.date_start_production = "";
            this.selectedVrachTechnic = "";
            this.date_close_nar = "";
            this.sum = "";
            this.buttonAddRow = true;
        },
        editItem: function(item, $index) {
            this.edit_vrach_ortoped = item.vrach_ortoped;
            this.edit_vrach_technic = item.vrach_technic;
            this.editIndex = $index; //Устанавливает индекс редактируемой записи
            this.edit = item;
            this.buttonAddRow = false;
        },
        addItem: function() {
            this.addItemRow = true;
            this.lastIdTable = this.items[this.items.length - 1] + 1;
            this.buttonAddRow = false;
        },
        successEdit: function(item, $index, e) {
            this.currentId = demo.currentId;
            var me = this;
            this.edit_sum = String(this.edit_sum).replace(/(\d+)\.(?=\d+\.\d+)/g, '$1');
            if (parseFloat(this.edit_sum) > 999999.99) {
                alertify.alert("Сумма не может быть больше 999999.99 !", function() {
                    alertify.message('OK');
                });
                return;
            }


            var date_open_nar;
            var date_start_production;
            var date_close_nar;
            var vrach_ortoped;
            var vrach_technic;
            var sum;

            if (item.date_open_nar == '') {
                date_open_nar = this.edit_date_open_nar;
            } else {
                date_open_nar = item.date_open_nar;
            }

            if (item.date_start_production == '') {
                date_start_production = this.edit_date_start_production;
            } else {
                date_start_production = item.date_start_production;
            }

            if (item.date_close_nar == '') {
                date_close_nar = this.edit_date_close_nar;
            } else {
                date_close_nar = item.date_close_nar;
            }

            if (item.sum == '') {
                sum = this.edit_sum;
            } else {
                sum = item.sum;
            }

            if (this.selectedVrachOrotped == '') {
                vrach_ortoped = this.edit_vrach_ortoped;
            } else {
                vrach_ortoped = this.selectedVrachOrotped;
            }

            if (this.selectedVrachTechnic == '') {
                vrach_technic = this.edit_vrach_technic;
            } else {
                vrach_technic = this.selectedVrachTechnic;
            }

            $.ajax({
                url: '/edit_list_nar_edit_save',
                type: 'post',
                data: {
                    id: item.id,
                    id_patient: me.currentId,
                    date_open_nar: date_open_nar,
                    vrach_ortoped: vrach_ortoped,
                    date_start_production: date_start_production,
                    vrach_technic: vrach_technic,
                    date_close_nar: date_close_nar,
                    sum: sum,
                    name_reg: navBar.nameReg
                }
            }).done(function(data, status) {
                if (status == "success") {
                    me.items[this.editIndex] = this.edit
                    me.edit = {}
                    me.setDefaultItems();
                    me.edit_date_open_nar = '';
                    me.edit_date_start_production = '';
                    me.edit_date_close_nar = '';
                    me.edit_sum = '';
                    me.selectedVrachOrotped = '';
                    me.selectedVrachTechnic = '';
                    me.number_nar = '';
                    me.edit_vrach_otroped = '';
                    me.edit_vrach_technic = '';
                    me.buttonAddRow = true;
                    alertify.set({ delay: 10000 });
                    alertify.success("Наряд обновлен");
                } else {
                    alertify.set({ delay: 10000 });
                    alertify.success("Ошибка на сервере")
                }
            })
        }
    }
})

Vue.component('modal-call', {
    delimiters: ['{', '}'],
    template: '#modal-template-call',
    data: function() {
        return {
            items: null,
            itemsComments: null,
            date: '',
            comment: '',
            patient_refuse: false,
            flag_comments: false
        };
    },
    mounted: function() {
        this.setDefaultItems();
        this.setDefaultItemsComments();
    },
    methods: {
        setDefaultItems: function() {
            this.$http.get('/taket_to_repair_get_info_patient?id=' + demo.currentId + '').then(data => {
                this.items = JSON.parse(data.body);

            });
            this.$http.get('/taket_to_repair_get_info_patient_comments?id=' + demo.currentId + '').then(data => {
                this.itemsComments = JSON.parse(data.body);
                console.log(data.body)

        });
        },
        setDefaultItemsComments: function() {
            this.$http.get('/taket_to_repair_get_info_patient_comments?id=' + demo.currentId + '').then(data => {
                this.itemsComments = JSON.parse(data.body);
            console.log(data.body)

        })
            ;
        },
        add_comment: function (id) {
            if(this.comment == ''){
                alertify.alert("Комментарий не может быть пустым!.", function() {
                    alertify.message('OK');
                });
                return;
            }else {
                var me = this;
                $.ajax({
                    url: '/take_to_repair_comment_put',
                    type: 'post',
                    data: {
                        id: demo.currentId,
                        comment: this.comment,
                        name_reg: navBar.nameReg
                    }
                }).done(function (data, status) {
                    if (status == 'success') {
                        alertify.set({delay: 10000});
                        alertify.success("Комментарий добавлен");
                        me.setDefaultItemsComments();
                        me.comment = '';
                    } else {
                        alertify.set({delay: 50000});
                        alertify.success("Произошла ошибка");
                    }
                })
            }
        },
        add_call: function() {
            if(this.date == '' && this.patient_refuse == false){
                alertify.alert("Не выбрано действие.", function() {
                    alertify.message('OK');
                });
                return;
            }
            if(this.date != '' || this.patient_refuse == true) {
                var me = this;
                $.ajax({
                    url: '/take_to_repair_invitation_put',
                    type: 'post',
                    data: {
                        id: demo.currentId,
                        date_invitation: this.date,
                        comment: this.comment,
                        patient_refuse: this.patient_refuse
                    }
                }).done(function (data, status) {
                    console.log("status = " + status);
                    if (status == 'success') {
                        alertify.set({delay: 10000});
                        if (me.patient_refuse == false) {
                            alertify.success("Пациент приглашен на примем");
                        } else {
                            alertify.success("Пациент отменил посещение");
                            me.patient_refuse = false;
                        }
                        demo.showModalCall = false;
                        demo.setDefaultItems(demo.selected_status);
                    } else {
                        alertify.set({delay: 50000});
                        alertify.success("Произошла ошибка");
                        demo.showModalCall = false;
                    }

                    // console.log(val);

                })
                //тут нужно добавить проверку, на ответ сервера.... очень важно!!!!!!!!!! не забивай!!!!
            }
        }
    }
})



Vue.component('modal-profile', {
    delimiters: ['{', '}'],
    template: '#modal-template-profile',
    data: function() {
        return {
            items: null,
            itemsComments: null,
            date: '',
            comment: '',
            patient_refuse: false
        };
    },
    mounted: function() {
        this.setDefaultItems();
        this.setDefaultItemsComments();
    },
    methods: {
        setDefaultItems: function() {
            console.log(demo.currentId)
            this.$http.get('/taket_to_repair_get_info_patient?id=' + demo.currentId + '').then(data => {
                this.items = JSON.parse(data.body);

            });
        },
        setDefaultItemsComments: function() {
            this.$http.get('/taket_to_repair_get_info_patient_comments?id=' + demo.currentId + '').then(data => {
                this.itemsComments = JSON.parse(data.body);
            console.log(data.body)

        })
            ;
        },
        add_comment: function (id) {
            if(this.comment == ''){
                alertify.alert("Комментарий не может быть пустым!.", function() {
                    alertify.message('OK');
                });
                return;
            }else {
                var me = this;
                $.ajax({
                    url: '/take_to_repair_comment_put',
                    type: 'post',
                    data: {
                        id: demo.currentId,
                        comment: this.comment,
                        name_reg: navBar.nameReg
                    }
                }).done(function (data, status) {
                    if (status == 'success') {
                        alertify.set({delay: 10000});
                        alertify.success("Комментарий добавлен");
                        me.setDefaultItemsComments();
                        me.comment = '';
                    } else {
                        alertify.set({delay: 50000});
                        alertify.success("Произошла ошибка");
                    }
                })
            }
        },
        add_call: function() {
            if(this.date == '' && this.patient_refuse == false){
                alertify.alert("Не выбрано действие.", function() {
                    alertify.message('OK');
                });
                return;
            }
            if(this.date != '' || this.patient_refuse == true) {
                var me = this;
                $.ajax({
                    url: '/take_to_repair_invitation_put',
                    type: 'post',
                    data: {
                        id: demo.currentId,
                        date_invitation: this.date,
                        comment: this.comment,
                        patient_refuse: this.patient_refuse
                    }
                }).done(function (data, status) {
                    console.log("status = " + status);
                    if (status == 'success') {
                        alertify.set({delay: 10000});
                        if (me.patient_refuse == false) {
                            alertify.success("Пациент приглашен на примем");
                        } else {
                            alertify.success("Пациент отменил посещение");
                            me.patient_refuse = false;
                        }
                        demo.showModalCall = false;
                        demo.setDefaultItems(demo.selected_status);
                    } else {
                        alertify.set({delay: 50000});
                        alertify.success("Произошла ошибка");
                        demo.showModalCall = false;
                    }

                    // console.log(val);

                })
                //тут нужно добавить проверку, на ответ сервера.... очень важно!!!!!!!!!! не забивай!!!!
            }
        }
    }
});


Vue.component('modal', {
    delimiters: ['{', '}'],
    template: '#modal-template',
    data: function() {
        return {
            fam: '',
            name: '',
            lastname: '',
            date_birth: '',
            number_phone: '',
            home_adres: '',
            number_ud: '',
            number_pasport: '',
            selected_lgot_cat: '',
            items_lgot_cat: null,
            selected_fio_reg: navBar.nameReg,
            items_fio_reg: null,
            note: ''
        };
    },
    mounted: function() {
        this.setDefaultItems();
    },
    methods: {
        setDefaultItems: function() {
            //var items;
            this.$http.get('/catalog_lgot_cat_get').then(data => {
                this.items_lgot_cat = JSON.parse(data.body);
            });
            // this.$http.get('/catalog_med_reg_get').then(data => {
            //     this.items_fio_reg = JSON.parse(data.body);
            // });
        },
        inputValueFilterPhone: function() {
            var me = this;
            jQuery(function($) {
                $("#phone").mask("9 (999) 999-9999", { completed: function() { me.number_phone = this.val(); } });
            });
        },
        inputValueFilterPasport: function () {
            var me = this;
            jQuery(function($) {
                $("#number_pasport").mask('99-99 999999', { completed: function() { me.number_pasport = this.val(); } });
            });
        },
        inputValueFilterAdres: function() {
            var me = this;
            $("#homeAdres").suggestions({
                token: "bd6f46fc9a074c3dfc16d8a44f5822f03dff0bc0",
                type: "ADDRESS",
                count: 5,
                /* Вызывается, когда пользователь выбирает одну из подсказок */
                onSelect: function(suggestion) {
                    console.log(suggestion);
                    me.home_adres = suggestion.value;
                }
            });
        },
        inputValueFilter: function(id) {
            var input = document.getElementById(id);
            input['oninput' in input ? 'oninput' : 'onpropertychange'] = function() {
                var str = this.value,
                    reg = /[^-а-яё\d\s]/g,
                    regUp = /(^|\s)(\S)/g;
                str = str.toLowerCase().replace(reg, '');
                str = str.replace(regUp, function g(a, b, c) {
                    return b + c.toUpperCase()
                });
                this.value = str;
            };
        },
        add_patient: function() {
            // Insert AJAX call here...
            var me = this;
            $.ajax({
                url: '/take_to_repair_patient_put',
                type: 'post',
                data: {
                    fam: this.fam,
                    name: this.name,
                    lastname: this.lastname,
                    date_birth: this.date_birth,
                    number_phone: this.number_phone,
                    home_adres: this.home_adres,
                    number_ud: this.number_ud,
                    number_pasport: this.number_pasport,
                    lgot_cat: this.selected_lgot_cat,
                    fio_reg: this.selected_fio_reg
                }
            }).done(function(val) {
                alertify.set({ delay: 10000 });
                alertify.success("Пациент успешно добавлен");
                me.fam = '';
                me.name = '';
                me.lastname = '';
                me.date_birth = '';
                me.number_phone = '';
                me.home_adres = '';
                me.number_ud = '';
                me.number_pasport = '';
                me.selected_lgot_cat = '';
                // me.selected_fio_reg = '';
                me.note = '';

                demo.setDefaultItems(demo.selected_status);
            })
            if(me.note !== '') {
                $.ajax({
                    url: '/take_to_repair_comment_put',
                    type: 'post',
                    data: {
                        id: demo.currentId,
                        comment: this.note,
                        name_reg: navBar.nameReg
                    }
                }).done(function (data, status) {
                    if (status == 'success') {
                        alertify.set({delay: 10000});
                        alertify.success("Комментарий добавлен");
                        me.note = '';
                    } else {
                        alertify.set({delay: 50000});
                        alertify.success("Произошла ошибка");
                    }
                })
            }

        }
    }
});


// register the grid component
Vue.component('demo-grid', {
    delimiters: ['{', '}'],
    template: '#grid-template',
    props: {
        data: Array,
        columns: Array,
        filterKey: String,
        selected_status: String
    },
    data: function() {
        var sortOrders = {}
        this.columns.forEach(function(key) {
            sortOrders[key] = 1
        })
        return {
            sortKey: '',
            sortOrders: sortOrders,
            translate: {Id:'Очередь', FullName:'ФИО', Datebirth:'Дата рождения', Lgotcat:'Льгот. кат.' }
        }

    },
    computed: {
        filteredData: function() {
            var sortKey = this.sortKey
            var filterKey = this.filterKey && this.filterKey.toLowerCase()
            var order = this.sortOrders[sortKey] || 1
            var data = this.data
            if (filterKey) {
                data = data.filter(function(row) {
                    return Object.keys(row).some(function(key) {
                        return String(row[key]).toLowerCase().indexOf(filterKey) > -1
                    })
                })
            }
            if (sortKey) {
                data = data.slice().sort(function(a, b) {
                    a = a[sortKey]
                    b = b[sortKey]
                    return (a === b ? 0 : a > b ? 1 : -1) * order
                })
            }
            return data
        }
    },
    filters: {
        capitalize: function(str) {
            print: String
            return str.charAt(0).toUpperCase() + str.slice(1)
        }
    },
    methods: {
        sortBy: function(key) {
            this.sortKey = key
            this.sortOrders[key] = this.sortOrders[key] * -1
        },
        add_nar_patient: function(item) {

            demo.currentId = item.Id;
            demo.full_name = item.FullName;
            demo.date_birth = item.Datebirth;
            demo.showModalnar = true;
        },
        patient_complite: function (item) {
            alertify.confirm("Отметить пациента:" + item.FullName + " как пролеченного?", function(e) {
                if (e) {
                    var me = this;
                    $.ajax({
                        url: '/take_to_repair_patient_complite',
                        type: 'post',
                        data: {
                            id: item.Id
                        }
                    }).done(function (val) {
                        alertify.set({delay: 10000});
                        alertify.success("Пациент пролечен");


                        demo.setDefaultItems(demo.selected_status);
                    })
                }
            })

        },
        move_patient_treatment: function(item) {
            var me = this;
            alertify.confirm("Переместить пациента:" + item.FullName + " в категорию \"Пациенты на лечении\" ?", function(e) {
                if (e) {
                    // user clicked "ok"
                    $.ajax({
                        url: '/take_to_repair_move_patient_treatment',
                        type: 'post',
                        data: {
                            id: item.Id
                        }
                    }).done(function(data, status) {
                        if (status == "success") {
                            demo.setDefaultItems(me.selected_status)
                            alertify.set({ delay: 10000 });
                            alertify.success("Пациент: " + item.FullName + " перемещен \nв \" Пациенты на лечении \"");
                        } else {
                            alertify.set({ delay: 10000 });
                            alertify.success("Ошибка на сервере")
                        }
                    })
                } else {
                    console.log("");
                }
            })

        },
        call_patient: function(item) {
            demo.currentId = item;
            demo.showModalCall = true;
        },
		profile_patient: function(item) {
            demo.currentId = item;
            demo.showModalProfile = true;
        },
		
        print: function(id) {
            $("#mySelector").printThis({
                debug: false, // show the iframe for debugging
                importCSS: true, // import page CSS
                importStyle: false, // import style tags
                printContainer: true, // grab outer container as well as the contents of the selector
                loadCSS: "path/to/my.css", // path to additional css file - use an array [] for multiple
                pageTitle: "", // add title to print page
                removeInline: false, // remove all inline styles from print elements
                printDelay: 333, // variable print delay; depending on complexity a higher value may be necessary
                header: null, // prefix to html
                footer: null, // postfix to html
                base: false, // preserve the BASE tag, or accept a string for the URL
                formValues: true, // preserve input/form values
                canvas: false, // copy canvas elements (experimental)
                doctypeString: "...", // enter a different doctype for older markup
                removeScripts: false // remove script tags from print content
            });
        }
    }
})

// bootstrap the demo
var demo = new Vue({
    el: '#demo',
    delimiters: ['{', '}'],
    data: {
        searchQuery: '',
        gridColumns: ['Id', 'FullName', 'Datebirth', 'Lgotcat'],
        gridData: null,
        showModal: false,
        showModalnar: false,
        showModalCall: false,
		showModalProfile: false,
        currentId: '',
        selected_status: '',
        full_name: '',
        date_birth: ''
    },
    mounted: function() {
        this.setDefaultItems('Пациенты в очереди');
    },
    methods: {
        setDefaultItems: function(selected_status) {
            if (selected_status == 'Пациенты в очереди') {
                this.$http.get('/take_to_repair_queue_get').then(data => {
                    this.gridData = JSON.parse(data.body);

                });
                this.selected_status = selected_status;
            }
            if (selected_status == 'Приглашенные пациенты') {
                this.$http.get('/take_to_repair_invitation_get').then(data => {
                    this.gridData = JSON.parse(data.body);
                });
                this.selected_status = selected_status;
            }
            if (selected_status == 'Пациенты на лечении') {
                this.$http.get('/take_to_repair_treatment_get').then(data => {
                    this.gridData = JSON.parse(data.body);
                });
                this.selected_status = selected_status;
            }

        }
    }
})