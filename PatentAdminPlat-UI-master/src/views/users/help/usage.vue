<template>
  <div class="container">
    <div class="my-card">
      <div v-for="i in 9" :key="i">
        <img
          :src="`http://www.sriptb.com/static/docs/usage/${i}.jpg`"
          style="height: 800px;width: 100%"
        >
        <hr>
      </div>

    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      list: []
    }
  }
}
</script>
