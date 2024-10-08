<script setup>

</script>

<template>
  <nav class="navbar navbar-expand-lg bg-body-tertiary">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">Ayman food</a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link
              class="nav-link"
              :to="{name: 'home'}"
              active-class="active-link"
            >Home</router-link>
          </li>
          <li class="nav-item">
            <router-link
              class="nav-link"
              :to="{name: 'login'}"
              active-class="active-link"
            >Login</router-link>
          </li>
          <li class="nav-item">
            <router-link
              class="nav-link"
              :to="{name: 'register'}"
              active-class="active-link"
            >Register</router-link>
          </li>
          <li class="nav-item">
            <router-link
              class="nav-link"
              :to="{name: 'cart'}"
              active-class="active-link"
            >Cart</router-link>
          </li>
        </ul>

        <form class="d-flex" role="search">
          <input class="form-control me-2" type="search" placeholder="Search" aria-label="Search" />
          <button class="btn btn-outline-success" type="submit">Search</button>
        </form>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.navbar{
  background-color: papayawhip !important;
  border-radius: 15px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  margin-bottom: 15px;
}
.navbar-brand{
  color: #f39c12;
  font-size: 1.3rem;
}
.breadcrumb {
  background-color: transparent;
}
.breadcrumb-item + .breadcrumb-item::before {
  content: '>';
}
.nav-link {
  text-decoration: none; /* Remove default link underline */
  color: #666; /* Default link color */
}

.active-link {
  color: orange; /* Active link color */
  font-weight: bold; /* Optional: Make active link bold */
}
</style>