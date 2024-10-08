<script setup>


// Props passed from the parent component
const props = defineProps({
  supplier: {
    type: Object,
    required: true
  }
});



// Image fallback logic
const handleImageError = (event) => {
  switch (props.supplier.Type) {
    case 'restaurant':
      event.target.src = 'src/assets/restaurant-placeholder.png';
      break;
    case 'coffee_shop':
      event.target.src = 'src/assets/coffee-shop-placeholder.png';
      break;
    case 'supermarket':
      event.target.src = 'src/assets/supermarket-placeholder.png';
      break;
    case 'bar':
      event.target.src = 'src/assets/bar-placeholder.png';
      break;
    default:
      event.target.src = 'src/assets/shop-placeholder.png';
  }
};
</script>

<template>
  <div class="supplier-card">
    <img
      :src="supplier.Image"
      @error="handleImageError"
      class="supplier-image"
      alt="Supplier Image"
    />
    <div class="supplier-info">
      <h3>{{ supplier.Name }}</h3>
      <p>Type: {{ supplier.Type }}</p>
      <p>Working Hours: {{ supplier.WorkingHours.Opening }} - {{ supplier.WorkingHours.Closing }}</p>
    </div>
  </div>
</template>

<style scoped>
.supplier-card {
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 16px;
  text-align: center;
  margin: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.supplier-image {
  width: 100px;
  height: 100px;
  object-fit: cover;
  border-radius: 4%;
  margin-bottom: 12px;
}

.supplier-info h3 {
  margin: 0;
  font-size: 18px;
  color: hotpink;
}

.supplier-info p {
  font-size: 14px;
  color: wheat;
  margin: 4px 0;
}

</style>
