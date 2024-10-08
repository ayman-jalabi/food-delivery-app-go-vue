<script setup>
import { RouterLink, RouterView } from 'vue-router'
import './assets/base.css'
import './assets/main.css'
</script>

<template>
  <div class="main-container">
    <RouterLink />
    <RouterView />
  </div>

</template>

<style scoped>
.main-container {
  background-color: #472b00;
  min-height: 100vh;
  width: 80%;        /* Set the desired percentage width */
  margin: 0 auto;    /* Centers the container horizontally */
  padding: 0 20px;
  border-radius: 15px;
}
</style>
