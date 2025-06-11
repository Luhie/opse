<template>
  <fieldset class="software-info">
    <legend>Microsoft Office 라이센스 정보</legend>

    <div v-if="officeList.length === 0">Office 정보를 불러오는 중...</div>

    <div v-for="(office, index) in officeList" :key="index" class="office-card">
      <h4>Office {{ index + 1 }}</h4>

      <div class="grid"><label>제품명:</label><span>{{ office.name }}</span></div>
      <div class="grid"><label>버전:</label><span>{{ office.version }}</span></div>
      <div class="grid"><label>Partial Key:</label><span>{{ office.partial_product_key }}</span></div>

      <div class="grid"><label>정품 인증:</label>
        <span v-if="office.licensed === 1">✅ 정품</span>
        <span v-else-if="office.licensed === 0">❌ 미인증</span>
        <span v-else>⚠️ 판별불가</span>
      </div>

      <div class="grid"><label>유예기간 남은 일수:</label><span>{{ office.grace_period }}</span></div>
      <div class="grid"><label>예상 만료일:</label><span>{{ office.estimated_expire_date || '-' }}</span></div>

      <div class="grid"><label>크랙 의심:</label>
        <span v-if="office.cracked === 1">🚨 크랙 의심 🚨</span>
        <span v-else-if="office.cracked === 0">✅ 정상</span>
        <span v-else>⚠️ 판별불가</span>
      </div>
    </div>

    <button @click="fetchOfficeInfo">정보 다시 불러오기</button>

    <!-- <pre>{{ officeList }}</pre> -->
  </fieldset>
</template>

<script lang="ts" setup>
import { ref, onMounted, defineExpose, defineEmits } from 'vue';
import { GetOfficeInfo } from '@wails/go/main/App';
import { waitWailsReady } from '@/utils/wailsReady'

interface OfficeInfo {
  name: string;
  version: string;
  partial_product_key: string;
  licensed: number;
  grace_period: number;
  estimated_expire_date: string;
  cracked: number;
}

const officeList = ref<OfficeInfo[]>([]);
const emit = defineEmits<{ (e: 'loaded', info: OfficeInfo[]): void }>()
defineExpose({
    officeList
})

const fetchOfficeInfo = async () => {
  try {
    officeList.value = await GetOfficeInfo();
    if (officeList.value) {
      emit('loaded', officeList.value)
    }
  } catch (error) {
    console.error('Office 정보 불러오기 실패:', error);
  }
}

onMounted(async () => {
  await waitWailsReady();
  fetchOfficeInfo();
});
</script>

<style scoped>
.software-info {
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
}
.office-card {
  margin-bottom: 1.5rem;
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
