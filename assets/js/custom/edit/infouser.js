var infoUser = new Vue({
					delimiters: ['{', '}'],
					el:'#infoUser',
					data: {
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
						editIndex: null //Индекс редактируемой записи
					},

					mounted:function(){
						this.setDefaultItems();
					},
					watch:{
						edit: function(){
							$(".born").maskMoney();
						}
					},
					methods: {
						setDefaultItems: function (){
							//var items;
							var urlParams = new URLSearchParams(window.location.search);
							this.$http.get('/edit_listNar_get?id='+urlParams.get('id')+'').then(data =>{
								this.items = JSON.parse(data.body);
								 if (data.body = '[]'){
									 this.flag = true;
									 console.log("!!!!!!!!!!!!!"+this.flag);
								 }
							});
							this.$http.get('/listlgotcat').then(data =>{
								this.items1 = JSON.parse(data.body);;
							});
							this.$http.get('/edit_fiovrach_get').then(data =>{
								this.itemsFioVrach = JSON.parse(data.body);
							});
							this.$http.get('/edit_fiovrach_technic_get').then(data =>{
								this.itemsFioVrachTechnic = JSON.parse(data.body);
							});
						},
						editItem: function(item, $index, e){
							this.editIndex = $index; //Устанавливает индекс редактируемой записи
							this.edit = item;
						},
						addItem: function(){
                            this.addItemRow = true;
                            this.lastIdTable = this.items[this.items.length - 1] +1;
							this.buttonAddRow = false;
							// this.addItemRow = true;
							// this.addItemRowSelected = true;
							// this.addNarRow = true;

						},

						successEdit: function (item, $index, e) {
                            var urlParams = new URLSearchParams(window.location.search);
                            console.log(item.id.String)
							$.ajax({
								url: '/edit_list_nar_edit_save',
								type: 'post',
								data: {
									id: item.id.String,
                                    id_patient: urlParams.get('id'),
									date_open_nar: item.date_open_nar.String,
									vrach_otroped: item.vrach_ortoped.String,
									date_start_production: item.date_start_production.String,
									vrach_technic: item.vrach_technic.String,
									date_close_nar: item.date_close_nar.String,
									sum: item.sum.String
								}
							}).done(function(val) {
								// console.log(val);
							})
								//тут нужно добавить проверку, на ответ сервера.... очень важно!!!!!!!!!! не забивай!!!!
							this.items[this.editIndex] = this.edit
							this.edit = {}
						},
						successAddNar: function(){
							var urlParams = new URLSearchParams(window.location.search);
							me = this;
							$.ajax({
								url: '/edit_add_nar',
								type: 'post',
								data: {
									id_patient: urlParams.get('id'),
									vrach_ortoped:	this.selectedVrachOrotped,
									vrach_technic: this.selectedVrachTechnic,
									number_nar: this.number_nar,
									date_open_nar: this.date_open_nar,
									date_start_production: this.date_start_production,
									date_close_nar: this.date_close_nar,
									sum: this.sum,
								}
							}).done(function(data, status) {
								 console.log("status = "+ status);
								 me.setDefaultItems();
							})


                // this.lastIdTable = Number(this.items[this.items.length - 1].id.String) + 1;
                 this.addItemRow = false;
                // this.items.push({"id":{"String": ""+this.lastIdTable+"","Valid":true},
								// "number_nar":{"String":""+this.number_nar+"","Valid":false},
								// "date_open_nar":{"String":""+this.date_open_nar+"","Valid":false},
								// "vrach_ortoped":{"String":""+this.selectedVrachOrotped+"","Valid":false},
								// "date_start_production":{"String":""+this.date_start_production+"","Valid":false},
								// "vrach_technic":{"String":""+this.selectedVrachTechnic+"","Valid":false},
								// "date_close_nar":{"String":""+this.date_close_nar+"","Valid":false},
								// "sum":{"String":""+this.sum+"","Valid":false}});

                            this.number_nar = "";
                            this.date_open_nar = "";
                            this.selectedVrachOrotped = "";
                            this.date_start_production = "";
                            this.selectedVrachTechnic = "";
                            this.date_close_nar = "";
                            this.sum = "";
                            this.buttonAddRow = true;
														this.flag = false;
                        },
							deleteEdit: function (item, $index, e) {
							$.ajax({
								url: '/listNarEditSave',
								type: 'post',
								data: {
									id: item.Id,
									patient_id: document.forms["formx"]["id"].value,
									CloseOpen: item.CloseOpen,
									Sum: item.Sum
								}
							}).done(function(val) {
								console.log(val);
							})
								//тут нужно добавить проверку, на ответ сервера.... очень важно!!!!!!!!!! не забивай!!!!
							this.items[this.editIndex] = this.edit
							this.edit = {}
						},

						cancelEdit: function (item, $index, e) {
							this.edit = {};
						},
						cancelAddItemRowSelected: function (){
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



					}
				});
