			var patientDate = new Vue({
					delimiters: ['{', '}'],
					el:'#patientDate',
					data: {
						edit: {},
						editIndex: undefined,
						items: null,
						editIndex: null //Индекс редактируемой записи

					},
					mounted:function(){
						this.setDefaultItems();
					},
					  filters: {
						phone: function (value) {
						 // if (!value) return ''
							//  value = value.toString()
							//  return value.charAt(0).toUpperCase() + value.slice(1)
								 }
							 },
					methods: {
						setDefaultItems: function (){
							var urlParams = new URLSearchParams(window.location.search);
							console.log(urlParams.get('id')); // "edit"
							//var items;
							// this.$http.get('/edit?id='+document.forms["formx"]["id"].value+'').then(data =>{
								this.$http.get('/edit_get?id='+urlParams.get('id')+'').then(data =>{
								this.items = JSON.parse(data.body);
								console.log(data.body)
							});
						},

						inputValueFilterPhone: function() {
								$("#phone").mask("9(999) 999-9999");
							},
						editItem: function(item, $index, e){
							console.log($index)
							this.editIndex = $index; //Устанавливает индекс редактируемой записи
							this.edit = item;


						},

						successEdit: function (item, $index, e) {

							var urlParams = new URLSearchParams(window.location.search);
							$.ajax({
								url: '/editPatientPhoneEdit',
								type: 'post',
								data: {
									id: urlParams.get('id'),
									phone: document.getElementById("phone").value
								}
							}).done(function(val) {
								console.log(val);
							})
								//тут нужно добавить проверку, на ответ сервера.... очень важно!!!!!!!!!! не забивай!!!!
							this.items[this.editIndex] = this.edit
							this.items[this.editIndex].Phone = document.getElementById("phone").value;
							this.edit = {}
						},

						cancelEdit: function (item, $index, e) {
							this.edit = null;
						},



					}
				});
