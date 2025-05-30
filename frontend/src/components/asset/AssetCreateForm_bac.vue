<template>
  <form class="form-wrapper" @submit.prevent="submitForm">
    <!-- 사용자 정보 -->
    <fieldset>
      <legend>사용자 정보</legend>
      <label>법인: <input v-model="form.corporation" /></label>
      <label>부서: <input v-model="form.department" /></label>
      <label>팀명: <input v-model="form.team" /></label>
      <label>성명: <input v-model="form.name" /></label>
      <label>직급: <input v-model="form.position" /></label>
      <label>비고: <input v-model="form.userNote" /></label>
    </fieldset>

    <!-- 자산 정보 -->
    <fieldset>
      <legend>자산 정보</legend>
      <label>구분: 
        <select v-model="form.category">
          <option>비품</option>
          <option>소모품</option>
        </select>
      </label>
      <label>종류: 
        <select v-model="form.itemType">
          <option>PC본체</option>
          <option>모니터</option>
          <option>무선브릿지</option>
          <option>AP</option>
        </select>
      </label>
      <label>수량: <input type="number" v-model.number="form.quantity" /></label>
      <label>모델명: <input v-model="form.model" /></label>
      <label>제조사: <input v-model="form.manufacturer" /></label>
      <label>구매일: <input type="date" v-model="form.purchaseDate" /></label>
      <label>사용유무: 
        <select v-model="form.usage">
          <option value="O">O</option>
          <option value="X">X</option>
          <option value="폐기">폐기</option>
        </select>
      </label>
      <label>비고: <input v-model="form.assetNote" /></label>
      <label>자산명칭: <input v-model="form.assetName" /></label>
      <label>관리번호: <input v-model="form.assetId" /></label>
    </fieldset>

    <button type="submit">제출</button>
  </form>
</template>

<script setup lang="ts">
import { reactive } from 'vue'

const form = reactive({
  // 사용자 정보
  corporation: '',
  department: '',
  team: '',
  name: '',
  position: '',
  userNote: '',

  // 자산 정보
  category: '비품',
  itemType: 'PC본체',
  quantity: 1,
  model: '',
  manufacturer: '',
  purchaseDate: '',
  usage: 'O',
  assetNote: '',
  assetName: '',
  assetId: '',
})

function submitForm() {
  console.log('제출된 폼:', JSON.stringify(form, null, 2))
  alert('폼이 제출되었습니다! (콘솔 확인)')
}
</script>

<style scoped>
.form-wrapper {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  padding: 2rem;
  background-color: #1b1b1b;
  color: white;
  max-width: 600px;
  margin: auto;
}
fieldset {
  border: 1px solid #444;
  border-radius: 8px;
  padding: 1rem;
}
legend {
  padding: 0 0.5rem;
  font-weight: bold;
  color: #ccc;
}
label {
  display: flex;
  flex-direction: column;
  margin-bottom: 1rem;
}
input, select {
  background-color: #333;
  color: white;
  border: 1px solid #555;
  border-radius: 4px;
  padding: 0.5rem;
}
button {
  background-color: #4CAF50;
  color: white;
  border: none;
  padding: 0.75rem;
  font-size: 1rem;
  border-radius: 8px;
  cursor: pointer;
}
button:hover {
  background-color: #45a049;
}
</style>
