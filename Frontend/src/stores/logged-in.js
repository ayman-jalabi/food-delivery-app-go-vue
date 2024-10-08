import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useLoggedInStore = defineStore('loggedIn', () => {
  //variable used to determine whether to display or provide logged-in functionality/accesses across the app:
  const loggedIn = ref(false)

  //I can use the loggedIn variable with access and refresh tokens. And with V-if "conditional rendering"
  //to determine what to render in the navbar, cart items, order info, etc...
  //https://vuejs.org/guide/essentials/conditional.html

  //single-line computed function syntax example:
  const doubleCount = computed(() => count.value * 2)

  // multi-line computed function syntax example:
  const publishedBooksMessage = computed(() => {
    return author.books.length > 0 ? 'Yes' : 'No'
})

  function increment() {
    count.value++
  }

  function decrement() {
    count.value--
  }

  function $reset(){
    count.value = 0
  }

  return { count, doubleCount, increment, decrement, $reset }
})
