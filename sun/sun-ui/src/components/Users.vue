<template>
  <div>
    <h2>{{ msg }}</h2>

    <div class="row row-cards-pf">
      <div class="col-xs-12 col-sm-6 col-md-3">
        <div class="card-pf card-pf-accented card-pf-aggregate-status">
          <h2 class="card-pf-title">
            <span class="fa fa-shield"></span>
            <span class="card-pf-aggregate-status-count">0</span> Ipsum
          </h2>
          <div class="card-pf-body">
            <p class="card-pf-aggregate-status-notifications">
              <span class="card-pf-aggregate-status-notification">
                <a href="#" class="add" data-toggle="tooltip" data-placement="top" title="Add Ipsum">
                  <span class="pficon pficon-add-circle-o"></span>
                </a>
              </span>
            </p>
          </div>
        </div>

      </div>
      <div class="col-xs-12 col-sm-6 col-md-3">
        <div class="card-pf card-pf-accented card-pf-aggregate-status">
          <h2 class="card-pf-title">
            <a href="#">
              <span class="fa fa-shield"></span>
              <span class="card-pf-aggregate-status-count">20</span> Amet</a>
          </h2>
          <div class="card-pf-body">
            <p class="card-pf-aggregate-status-notifications">
              <span class="card-pf-aggregate-status-notification">
                <a href="#">
                  <span class="pficon pficon-error-circle-o"></span>4</a>
              </span>
              <span class="card-pf-aggregate-status-notification">
                <a href="#">
                  <span class="pficon pficon-warning-triangle-o"></span>1</a>
              </span>
            </p>
          </div>
        </div>

      </div>
      <div class="col-xs-12 col-sm-6 col-md-3">
        <div class="card-pf card-pf-accented card-pf-aggregate-status">
          <h2 class="card-pf-title">
            <a href="#">
              <span class="fa fa-shield"></span>
              <span class="card-pf-aggregate-status-count">9</span> Adipiscing</a>
          </h2>
          <div class="card-pf-body">
            <p class="card-pf-aggregate-status-notifications">
              <span class="card-pf-aggregate-status-notification">
                <span class="pficon pficon-ok"></span>
              </span>
            </p>
          </div>
        </div>

      </div>
      <div class="col-xs-12 col-sm-6 col-md-3">
        <div class="card-pf card-pf-accented card-pf-aggregate-status">
          <h2 class="card-pf-title">
            <a href="#">
              <span class="fa fa-shield"></span>
              <span class="card-pf-aggregate-status-count">12</span> Lorem</a>
          </h2>
          <div class="card-pf-body">
            <p class="card-pf-aggregate-status-notifications">
              <a href="#">
                <span class="card-pf-aggregate-status-notification">
                  <span class="pficon pficon-error-circle-o"></span>1</span>
              </a>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Users',
    data() {
      return {
        msg: 'Welcome to Users View'
      }
    }
  }
</script>

<style scoped>

</style>