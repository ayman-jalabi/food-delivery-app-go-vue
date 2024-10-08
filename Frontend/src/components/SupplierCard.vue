<script setup>
const props = defineProps({ name: String, picture: String })
</script>

<template>

  <div class="card">
    <div>{{ props.name }}</div>
    <div>{{ props.picture }}</div>
  </div>



  <a href="./products.html" class="list-group-item list-group-item-action m-2" aria-current="true">
    <div class="d-flex w-100 justify-content-between">
      <h5 class="mb-1">Restaurants</h5>
      <small>3 days ago</small>
    </div>
    <p class="mb-1">Some placeholder content in a paragraph.</p>
    <small>And some small print.</small>
  </a>
</template>

<style scoped>

</style>