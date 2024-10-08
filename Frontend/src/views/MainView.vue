<script setup>
import NavBar from '@/components/layout/shared/NavBar.vue'
import { onMounted, onBeforeUnmount, computed, ref } from 'vue';
import { useSupplierStore } from '@/stores/supplierStore.js'; // Adjust the path
import { storeToRefs } from 'pinia';
import SupplierCard from '@/components/layout/SupplierCard.vue';

// Use the Pinia store
const supplierStore = useSupplierStore();
const { suppliers, loading, error, page, pageSize } = storeToRefs(supplierStore);

onMounted(() => {
  supplierStore.fetchSuppliers(); // Load the first page
});

// Methods for pagination
const nextPage = () => {
  supplierStore.nextPage();
};

const previousPage = () => {
  supplierStore.previousPage();
};

const handleClick = () => {
  supplierStore.setSelectedSupplier(props.supplier); // Set the selected supplier in the store
  props.supplier.Type
};
</script>

<template>
  <NavBar />
  <h2>Welcome to Ayman foods, <br> where dream meals become a reality!</h2>

  <button @click="supplierStore.fetchSuppliers()">fetch suppliers</button>

  <div>
    <div v-if="loading">Loading suppliers...</div>
    <div v-else-if="error">{{ error }}</div>
    <div v-else class="suppliers-view" @click="handleClick">
      <SupplierCard
        v-for="supplier in suppliers"
        :key="supplier.ID"
        :supplier="supplier"
      />
    </div>

    <div class="pagination-controls">
      <button @click="previousPage" :disabled="page === 1">Previous</button>
      <span>Page {{ page }}</span>
      <button @click="nextPage" :disabled="suppliers.length < pageSize">Next</button>
    </div>
  </div>

  <main class="main">


    <p class="d-inline-flex gap-1">
      <h3 class="mt-5 text-center">
        <a class="btn btn-primary" data-bs-toggle="collapse" href="#suppliersList" role="button" aria-expanded="false"
           aria-controls="collapseExample">
          Suppliers list (toggle)
        </a>
      </h3>
      <p class="text-center">Clicking the above button can toggle between showing or hiding the list of available food suppliers</p>
    </p>

    <div class="collapse" id="suppliersList">
      <div class="card card-body">

        <div class="list-group list-group-horizontal mx-3 mt-2">
          <a href="./products.html" class="list-group-item list-group-item-action m-2" aria-current="true">
            <div class="d-flex w-100 justify-content-between">
              <h5 class="mb-1">Restaurants</h5>
              <small>3 days ago</small>
            </div>
            <p class="mb-1">Some placeholder content in a paragraph.</p>
            <small>And some small print.</small>
          </a>
          <a href="./products.html" class="list-group-item list-group-item-action m-2">
            <div class="d-flex w-100 justify-content-between">
              <h5 class="mb-1">Supermarkets</h5>
              <small class="text-body-secondary">3 days ago</small>
            </div>
            <p class="mb-1">Some placeholder content in a paragraph.</p>
            <small class="text-body-secondary">And some muted small print.</small>
          </a>
          <a href="./products.html" class="list-group-item list-group-item-action m-2">
            <div class="d-flex w-100 justify-content-between">
              <h5 class="mb-1">Cafes</h5>
              <small class="text-body-secondary">3 days ago</small>
            </div>
            <p class="mb-1">Some placeholder content in a paragraph.</p>
            <small class="text-body-secondary">And some muted small print.</small>
          </a>
        </div>
      </div>
    </div>

    <p class="d-inline-flex gap-1">
      <h3 class="mt-5 text-center">
        <a class="btn btn-primary" data-bs-toggle="collapse" href="#cuisineCategories" role="button" aria-expanded="false"
           aria-controls="collapseExample">
          Cuisine Categories: (toggle)
        </a>
      </h3>
      <p class="text-center">Clicking the above button can toggle between showing or hiding the list of available cuisine categories </p>
    </p>


    <div class="collapse" id="cuisineCategories">
      <div class="card card-body">

        <div class="list-group list-group-horizontal mx-3 mt-2">
          <a href="./products.html" class="list-group-item list-group-item-action m-2" aria-current="true">
            <div class="d-flex w-100 justify-content-between">
              <h5 class="mb-1">Italian</h5>
              <small>3 days ago</small>
            </div>
            <p class="mb-1">Some placeholder content in a paragraph.</p>
            <small>And some small print.</small>
          </a>

          <a href="./products.html" class="list-group-item list-group-item-action m-2">
            <div class="d-flex w-100 justify-content-between">
              <h5 class="mb-1">Japanese</h5>
              <small class="text-body-secondary">3 days ago</small>
            </div>
            <p class="mb-1">Some placeholder content in a paragraph.</p>
            <small class="text-body-secondary">And some muted small print.</small>
          </a>
          <a href="./products.html" class="list-group-item list-group-item-action m-2">
            <div class="d-flex w-100 justify-content-between">
              <h5 class="mb-1">Chinese</h5>
              <small class="text-body-secondary">3 days ago</small>
            </div>
            <p class="mb-1">Some placeholder content in a paragraph.</p>
            <small class="text-body-secondary">And some muted small print.</small>
          </a>
        </div>

      </div>
    </div>



  </main>

</template>

<style scoped>
.pagination-controls {
  margin-top: 20px;
}
button {
  margin: 0 5px;
}

.suppliers-view {
  display: grid;
  grid-template-columns: repeat(3, 1fr); /* 3 equal-width columns */
  gap: 20px; /* Space between the cards */
}

@media (max-width: 768px) {
  /* For smaller screens, adjust the number of columns */
  .suppliers-view {
    grid-template-columns: repeat(2, 1fr); /* 2 columns on medium screens */
  }
}

@media (max-width: 480px) {
  .suppliers-view {
    grid-template-columns: 1fr; /* 1 column on small screens */
  }
}
</style>
