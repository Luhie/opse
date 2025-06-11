<template>
  <fieldset class="software-info">
    <legend>한컴오피스 (HWP) 라이센스 정보</legend>

    <div v-if="hwpInfo === null">
      한컴오피스가 설치되어 있지 않습니다.
    </div>

    <div v-else class="hwp-card">
      <div class="grid"><label>제품명:</label><span>{{ hwpInfo.product_name }}</span></div>
      <div class="grid"><label>버전:</label><span>{{ hwpInfo.version }}</span></div>
      <div class="grid"><label>설치 경로:</label><span>{{ hwpInfo.install_path }}</span></div>
    </div>

    <button @click="fetchHwpInfo">정보 다시 불러오기</button>

    <pre>{{ hwpInfo }}</pre>
  </fieldset>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { GetHwpInfo } from '@wails/go/main/App';

interface HwpInfo {
  product_name: string;
  version: string;
  install_path: string;
}

const hwpInfo = ref<HwpInfo | null>(null);

const fetchHwpInfo = async () => {
  try {
    hwpInfo.value = await GetHwpInfo();
  } catch (error) {
    console.error('HWP 정보 불러오기 실패:', error);
  }
}

onMounted(() => {
  fetchHwpInfo();
});
</script>

<style scoped>
.software-info {
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
}
.hwp-card {
  margin-bottom: 1rem;
  padding: 1rem;
  border: 1px dashed #888;
  border-radius: 6px;
}
.grid {
  display: flex;
  margin-bottom: 0.5rem;
}
.grid label {
  width: 150px;
  font-weight: bold;
}
pre {
  padding: 0.5rem;
  border-radius: 4px;
}
</style>
