// stores/supplierStore.js
import { ref, computed } from 'vue';
import { defineStore } from 'pinia';

export const useCategoryStore = defineStore('categoryStore', () => {
  const categories = ref([]);
  const loading = ref(false);
  const error = ref(null);
  const page = ref(1);
  const pageSize = ref(10); // Default page size

  // Computed properties
  const hasNextPage = computed(() => categories.value.length === pageSize.value);
  const hasPreviousPage = computed(() => page.value > 1);

  // Actions
  const fetchCategories = async (newPage = 1, newPageSize = 9) => {
    loading.value = true;
    error.value = null;
    page.value = newPage;
    pageSize.value = newPageSize;

    try {
      // Corrected URL without the /api prefix
      const response = await fetch(`http://localhost:8082/suppliers?page=${newPage}&pageSize=${newPageSize}`);
      if (!response.ok) {
        throw new Error('Failed to fetch suppliers on the frontend');
      }
      const data = await response.json();
      categories.value = data; // Make sure the API response is compatible
    } catch (err) {
      console.error('Error fetching suppliers on the frontend:', err);
      error.value = 'Failed to fetch suppliers on the frontend';
    } finally {
      loading.value = false;
    }
  };

  const nextPage = () => {
    fetchCategories(page.value + 1, pageSize.value);
  };

  const previousPage = () => {
    if (page.value > 1) {
      fetchCategories(page.value - 1, pageSize.value);
    }
  };

  return {
    categories,
    loading,
    error,
    page,
    pageSize,
    fetchCategories,
    nextPage,
    previousPage,
    hasNextPage,
    hasPreviousPage
  };
});
