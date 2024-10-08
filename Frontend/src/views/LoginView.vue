<script setup>
import NavBar from '@/components/layout/shared/NavBar.vue'
</script>

<template>

  <NavBar />

  <div class="container-lg">

    <main>
      <form class="container mt-5">
        <div class="row m-4">
          <div class="col-6">
            <label for="email">Enter your email: </label>
            <input type="email" class="form-control" placeholder="email" aria-label="email">
          </div>


        </div>
        <div class="row m-4">
          <div class="col-6">
            <label for="password">Enter your password: </label>
            <input type="password" id="inputPassword5" class="form-control" placeholder="Password"
                   aria-describedby="passwordHelpBlock">
          </div>
        </div>

<!--login button-->
        <button class="btn btn-primary ms-5 my-3" type="submit">Login</button>


      </form>
      </main>
  </div>
</template>

<style scoped>

</style>