// stores/supplierStore.js
import { ref, computed } from 'vue';
import { defineStore } from 'pinia';

export const useSupplierStore = defineStore('supplierStore', () => {
  const suppliers = ref([]);
  const loading = ref(false);
  const error = ref(null);
  const page = ref(1);
  const pageSize = ref(10); // Default page size
  const selectedSupplier = ref(null);

  // Computed properties
  const hasNextPage = computed(() => suppliers.value.length === pageSize.value);
  const hasPreviousPage = computed(() => page.value > 1);

  // Actions
  const fetchSuppliers = async (newPage = 1, newPageSize = 9) => {
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
      suppliers.value = data; // Make sure the API response is compatible
    } catch (err) {
      console.error('Error fetching suppliers on the frontend:', err);
      error.value = 'Failed to fetch suppliers on the frontend';
    } finally {
      loading.value = false;
    }
  };

  const nextPage = () => {
    fetchSuppliers(page.value + 1, pageSize.value);
  };

  const previousPage = () => {
    if (page.value > 1) {
      fetchSuppliers(page.value - 1, pageSize.value);
    }
  };

  const setSelectedSupplier = (supplier) => {
    selectedSupplier.value = supplier; // Set the selected supplier
  };

  const setSuppliers = (supplierList) => {
    suppliers.value = supplierList; // Set the suppliers in the store
  };

  return {
    suppliers,
    loading,
    error,
    page,
    pageSize,
    fetchSuppliers,
    nextPage,
    previousPage,
    hasNextPage,
    hasPreviousPage,
    selectedSupplier,
    setSelectedSupplier,
    setSuppliers
  };
});
