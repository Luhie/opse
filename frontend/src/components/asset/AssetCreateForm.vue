<template>
  <div style="min-height: 100vh; width: 100vw; display:flex; justify-content:center; align-items:center;">
  <form class="form-wrapper" @submit.prevent="submitForm">
    <!-- 사용자 정보 -->
    <fieldset>
      <legend>사용자 정보</legend>
      <div class="grid-user">
        <div class="grid-item">
          <label for="corporation">법인</label>
          <input id="corporation" v-model="form.corporation" />
        </div>
        <div class="grid-item">
          <label for="department">부서</label>
          <input id="department" v-model="form.department" />
        </div>
        <div class="grid-item">
          <label for="team">팀명</label>
          <input id="team" v-model="form.team" />
        </div>
        <div class="grid-item">
          <label for="name">성명</label>
          <input id="name" v-model="form.name" />
        </div>
        <div class="grid-item">
          <label for="position">직급</label>
          <input id="position" v-model="form.position" />
        </div>
        <div class="grid-item">
          <label for="userNote">비고</label>
          <input id="userNote" v-model="form.userNote" />
        </div>
      </div>
    </fieldset>

    <!-- 자산 정보 -->
    <fieldset>
      <legend>자산 정보</legend>
      <div class="grid-asset">
        <div class="grid-asset-item">
          <label>구분:</label>
          <select v-model="form.category">
            <option>비품</option>
            <option>소모품</option>
          </select>
        </div>
        <div class="grid-asset-item"><label>관리번호</label><input v-model="form.assetId" /></div>
        <div class="grid-asset-item"><label>자산명칭</label><input v-model="form.assetName" /></div>
        <div class="grid-asset-item">
          <label>종류:</label>
          <select v-model="form.itemType">
            <option>PC본체</option>
            <option>모니터</option>
            <option>무선브릿지</option>
            <option>AP</option>
          </select>
        </div>
        <div class="grid-asset-item"><label>수량</label><input type="number" v-model.number="form.quantity" /></div>
        <div class="grid-asset-item"><label>제조사</label><input v-model="form.manufacturer" /></div>
        <div class="grid-asset-item"><label>모델명</label><input v-model="form.model" /></div>
        <div class="grid-asset-item">
          <label>사용유무</label>
          <select v-model="form.usage">
            <option value="O">O</option>
            <option value="X">X</option>
            <option value="폐기">폐기</option>
          </select>
        </div>
        <div class="grid-asset-item"><label>구매일</label><input type="date" v-model="form.purchaseDate" /></div>
        <div class="grid-asset-item"><label>비고</label><input v-model="form.assetNote" /></div>
      </div>
    </fieldset>

    <button type="submit">제출</button>
  </form>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'

const today = new Date()
const year = today.getFullYear()
const month = (today.getMonth()+1).toString().padStart(2,'0')
const date = today.getDate().toString().padStart(2,'0')
const day = `${year}-${month}-${date}`

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
  purchaseDate: day,
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
  /* max-width: 1200px; */
  
  /* margin: auto; */
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

.grid-user {
  /* grid container */
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(30%,auto));
  gap: 0.5rem;
  /* grid-template-columns: repeat(2, 1fr);
  gap: 1.5rem; */
}

label {
  /* flex: 0 0 20%; */
  font-weight: 500;
  width: 7rem;
}

.grid-item {
  /* border: 1px solid pink; */
  display:flex;
  align-items: center;
  width: 100%;
  padding: 0.1rem;

  input{
    flex:1;
  }
  &:nth-child(5){
    /* grid-row: 2/2; */
    /* grid-column: 2/4; */
    /* background-color:red; */
  }
  &:nth-child(6){
    grid-column: 1/4;
    /* background-color:red; */
  }

}
.grid-asset {
  /* grid container */
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(30%,auto));
  gap: 0.5rem;
  /* grid-template-columns: repeat(2, 1fr);
  gap: 1.5rem; */
}
.grid-asset-item {
  display: flex;
  align-items:center;
  width: 100%;
  padding: 0.1rem;
  input,
  select{
    flex:1;
  }

  &:nth-child(10){
    grid-column: 1/4;
  }
}

input, select {
  background-color: #333;
  color: white;
  border: 1px solid #555;
  border-radius: 4px;
  padding: 0.5rem;
  width: 100%;
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
