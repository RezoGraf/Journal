/**
 * Created by localadmin on 12.05.17.
 */

Vue.component('print-page', {
    delimiters: ['{', '}'],
    template: '#print-page',
    props: {
        print: false
    }
})

Vue.component('cat-sum', {
    delimiters: ['{', '}'],
    template: '#cat-sum',
    props: {
        cats: null,
        date_start: '',
        date_end: '',
        selected_status: ''
    },
    methods: {
        print: function () {
            demo.print = true;
            setTimeout(function() {$("#print").printThis({
                // debug: false, // show the iframe for debugging
                // importCSS: true, // import page CSS
                // importStyle: false, // import style tags
                // printContainer: true, // grab outer container as well as the contents of the selector
                // loadCSS: "path/to/my.css", // path to additional css file - use an array [] for multiple
                //pageTitle: "", // add title to print page
                 removeInline: false, // remove all inline styles from print elements
                // printDelay: 333, // variable print delay; depending on complexity a higher value may be necessary
                 header: null, // prefix to html
                 footer: null // postfix to html
                // base: false, // preserve the BASE tag, or accept a string for the URL
                // formValues: true, // preserve input/form values
                // canvas: false, // copy canvas elements (experimental)
                // doctypeString: "...", // enter a different doctype for older markup
                // removeScripts: false // remove script tags from print content
            })}, 1000)
            // demo.print = false;
        }
    }
})

Vue.component('search-btn', {
    delimiters: ['{', '}'],
    template: '#search-btn',
    props: {
        selected_status: String
    },
    methods: {
        searchQuery: function () {
            if(demo.date_start == '' || demo.date_end == '' || demo.selected_status == ''){
                alertify.alert('Не выбран "Период отчета или", или "Вид отчета"', function() {
                    alertify.message('OK');
                });
                return;
            }
            demo.gridData = null;
            var me = this;
            switch(demo.selected_status){
                case 'По вызванным пациентам':
                    me.$http.get('/reports_invitation_get?date_start='+demo.date_start+'&date_end='+demo.date_end+'').then(data => {
                        demo.gridData = JSON.parse(data.body);
                    var data = JSON.parse(data.body);
                    var cats = {};
                    data.forEach(function(v) {
                        if(v.Lgotcat in cats) {
                            cats[v.Lgotcat]++;
                        } else {
                            cats[v.Lgotcat] = 1;
                        }
                    });
                    demo.cats = cats;
            });
            break;
        }
        }
    }
})


Vue.component('crumb-up', {
    delimiters: ['{', '}'],
    template: '#crumb-up',
    props: {
        selected_status: String
    }
})

Vue.component('selected-status-patient', {
    delimiters: ['{', '}'],
    template: '#selected-status-patient',
    data: function() {
        return {
            selected_status: ''
        };
    },
    watch: {
        selected_status: function () {
            demo.selected_status = this.selected_status;
        }
    }
})


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
            translate: {Id:'Очередь',DateRecord:'Дата записи на прием', FullName:'ФИО', DateInvitation:'Дата приглашения на прием', Lgotcat:'Льгот. кат.' }
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
    el: '#demo1',
    delimiters: ['{', '}'],
    data: {
        searchQuery: '',
        gridColumns: ['Id', 'DateRecord','FullName','DateInvitation', 'Lgotcat'],
        gridData: null,
        showModal: false,
        showModalnar: false,
        showModalCall: false,
        currentId: '',
        selected_status: '',
        full_name: '',
        date_birth: '',
        cats: null,
        date_start: '',
        date_end: '',
        show: false,
        print: false
    }
})