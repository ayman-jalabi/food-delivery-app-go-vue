import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', {
  state: () => ({ count: 0, name: 'Eduardo' }),
  getters: {
    doubleCount: (state) => state.count * 2,
  },
  actions: {
    increment() {
      this.count++
    },
    decrement() {
      this.count--
    }
  },
})

//below is composition syntax/style of Pinia
// export const useCounterStore = defineStore('counter', () => {
//   const count = ref(0)
//   const doubleCount = computed(() => count.value * 2)
//   // another computed ref:
//   const publishedBooksMessage = computed(() => {
//     return author.books.length > 0 ? 'Yes' : 'No'
// })
//
//   function increment() {
//     count.value++
//   }
//
//   function decrement() {
//     count.value--
//   }
//
//   function $reset(){
//     count.value = 0
//   }
//
//   return { count, doubleCount, increment, decrement, $reset }
// })
