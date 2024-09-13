<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>
<style>
.container {
  padding: 15px !important;
  background-color: #e9eff8;
}

.radius {
  border-radius: .25rem;
}

.my-card {
  padding: 20px;
  border-radius: .25rem;
  word-wrap: break-word;
  background-color: #fff;
  background-clip: border-box;
  border: 0 solid #f6f6f6;
  color: #666;
  box-shadow: 0 2px 4px rgb(15 34 58 / 12%);
}

.card-title {
  font-size: 16px;
  margin: 0 0 7px;
  font-weight: 600;
}
</style>
