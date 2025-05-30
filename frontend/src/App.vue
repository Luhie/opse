<script setup lang="ts">
import AssetCreateForm from './components/asset/AssetCreateForm.vue'

</script>
<template>
  <AssetCreateForm />
  <!-- <div class="wrapper">
    <div class="input-field">
      <label>이름</label>
      <input v-model="name" @keyup.enter="submit" placeholder="Write name" />
    </div>
    <button @click="submit">확인</button>
    <p v-if="message">{{  message  }}</P>
  </div> -->
</template>

<!-- <script setup lang="ts">
import { ref } from 'vue'

const name = ref('')
const message = ref('')

function submit() {
  message.value = `Hello, ${name.value}!`

}

</script> -->
<!-- 
<style scoped>
.input-field {
  border: 1px solid red;
  label {
    padding:10px;
    color:red;
  }
}
.wrapper {
  height: 100vh;

  display: flex;
  flex-direction: column;
  justify-content: center; /* 가운데 */
  align-items: center; /* 확장 안되고 item 가운데 */
  gap: 1rem;


  background-color: #1B2636;
  color: #FFFFFF;

}
input {
  padding: 0.5rem 1rem;
  font-size: 1rem;
  border-radius: 8px;
  border: none;
  outline: none;
}
button {
  padding: 0.5rem 1.5rem;
  font-size: 1rem;
  border-radius: 8px;
  background-color: #4CAF50;
  color: white;
  border: none;
  cursor: pointer;
}
button:hover {
  background-color: #45A049;
}
</style> -->