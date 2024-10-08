<script setup>

import NavBar from '@/components/layout/shared/NavBar.vue'
import { useCounterStore } from '@/stores/counter.js'

const storeCounter = useCounterStore()

</script>

<template>
  <NavBar />
  <main>
  <div class="container">
    <div class="row">
      <div class="col-4">

        <p class="d-inline"><b> items in the cart:</b>
          <p class="d-inline" id="items-count">{{storeCounter.count}}</p>
        </p>

        <div class="btn-group me-2" role="group" aria-label="First group">
          <button type="button" @click="storeCounter.increment()" class="btn btn-primary">+</button>
          <button type="button" @click="storeCounter.decrement()" class="btn btn-primary">-</button>
        </div>

        <h3>double items:</h3>
        <h3>{{storeCounter.doubleCount}}</h3>


        <button type="button" @click="storeCounter.$reset()" class="btn btn-primary">Remove all cart items</button>

      </div>
    </div>
  </div>

  </main>

</template>

<style scoped>

</style>