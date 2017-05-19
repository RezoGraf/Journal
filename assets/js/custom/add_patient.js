    <!-- 		<transition name="modal"> -->
    <div class="modal-mask" transition="modal">
      <div class="modal-wrapper">
        <div class="modal-container-700px">

          <div class="modal-header">
            <slot name="header">


            </slot>
          </div>
          <div class="modal-body">
            <slot name="body">
              <div class="row">

              <div class="col-xs-6 col-sm-4 col-md-4">
                <input type="text" v-model.trim="thepost" id="fam" class="form-control" name="fam" placeholder="" value="">

              </div>
              <div class="col-xs-6 col-sm-4 col-md-4">
                <input type="text" v-model.trim="thepost" id="name" class="form-control" name="name" placeholder="" value="">
              </div>
              <div class="col-xs-6 col-sm-4 col-md-4">
                <input type="text" v-model.trim="thepost" id="lastname" class="form-control" name="lastname" placeholder="" value="">
              </div>
              </div>
              <div class="row">
              <div class="col-xs-9 col-sm-6 col-md-6">
                <input type="date" v-model="thepost" id="date_birth"  class="form-control" name="date_birth" placeholder=" " value="">
              </div>
              <div class="col-xs-9 col-sm-6 col-md-6">
                <input type="text" v-model.trim.number="thepost" id="numberPhone" name="phone" class="form-control" placeholder=" " value="">
              </div>
              </div>
              <div class="row">
              <div class="col-xs-18 col-sm-12 col-md-12">
                <link href="https://cdn.jsdelivr.net/jquery.suggestions/17.2/css/suggestions.css" type="text/css" rel="stylesheet" />
                <input type="text" v-model="thepost" id="homeAdres" class="form-control" name="homeadres" placeholder=" ">
                <span style='color:red' id='homeadresf'></span>
                <!-- <input type="text" v-model="thepost" id="thepost" class="form-control" name="thepost" placeholder=" " value=""> -->
              </div>
              </div>
              <div class="row">
              <div class="col-xs-6 col-sm-4 col-md-4">
                <input type="text" v-model.trim="thepost" id="thepost" class="form-control" name="thepost" placeholder=" " value="">
              </div>
              <div class="col-xs-6 col-sm-4 col-md-4">
              <select  name="select_name" v-model="selected" class="form-control">
                <option disabled value=""> </option>
                <option v-for="(item, $index) in items" v-bind:selected="item.value" size="5">
                {item.ThePost}
                </option>
              </select>
              </div>
              <div class="col-xs-6 col-sm-4 col-md-4">
              <select  name="select_name" v-model="selected" class="form-control">
                <option disabled value=""> </option>
                <option v-for="(item, $index) in items" v-bind:selected="item.value" size="5">
                {item.ThePost}
                </option>
              </select>
              </div>
              </div>
              <div class="row">
              <div class="col-xs-18 col-sm-12 col-md-12">
                <textarea type="text"  v-model="thepost" id="thepost" class="form-control" name="thepost" placeholder="" value=""></textarea>
              </div>
              </div>
            </slot>
          </div>

          <div class="modal-footer">
            <slot name="footer">
              <button class="btn btn-default" @click="$emit('close') ">

              </button>
              <button class="btn btn-primary" @click="savePost()">

              </button>
            </slot>
          </div>
        </div>
      </div>
    </div>
    <!-- 		</transition> -->
