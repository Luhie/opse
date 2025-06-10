<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

import { GetAssetList } from '@wails/go/main/App'

// interface Asset {
//   assetId: string
//   assetName: string
//   itemType: string
//   manufacturer: string
//   purchaseDate: string
// }

// const assets = ref<Asset[]>([])
const assets = ref()

const router = useRouter()

onMounted(async () => {
    try {
        assets.value = await GetAssetList()
    } catch (err) {
        console.log('자산 목록 불러오기 실패:', err)
    }
})

function goToDetail(id: string) {
  router.push({ name: 'AssetDetail', params: { id } })
}
</script>

<template>
  <div class="asset-list-page">
    <h1>자산 목록</h1>
    <table>
      <thead>
        <tr>
          <th>관리번호</th>
          <th>자산명</th>
          <th>종류</th>
          <th>제조사</th>
          <th>등록일</th>
          <th>상세</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in assets" :key="item.assetId">
          <td>{{ item.assetId }}</td>
          <td>{{ item.assetName }}</td>
          <td>{{ item.itemType }}</td>
          <td>{{ item.manufacturer }}</td>
          <td>{{ item.purchaseDate }}</td>
          <td>
            <button @click="goToDetail(item.assetId)">상세</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
  color: white;
  background-color: #222;
}
th, td {
  border: 1px solid #444;
  padding: 0.5rem;
  text-align: center;
}
th {
  background-color: #333;
}
button {
  background-color: #4CAF50;
  border: none;
  padding: 0.3rem 0.7rem;
  color: white;
  border-radius: 4px;
  cursor: pointer;
}
button:hover {
  background-color: #45a049;
}
</style>
