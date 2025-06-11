<template>
  <fieldset class="software-info">
    <legend>Windows 라이센스 상세 정보</legend>

    <div class="grid"><label>OS 이름:</label><span>{{ windowsInfo?.caption || '불러오는 중...' }}</span></div>
    <div class="grid"><label>버전:</label><span>{{ windowsInfo?.version || '불러오는 중...' }}</span></div>
    <div class="grid"><label>Serial Number:</label><span>{{ windowsInfo?.serial_number || '불러오는 중...' }}</span></div>
    <div class="grid"><label>Product Key:</label><span>{{ windowsInfo?.product_key || '불러오는 중...' }}</span></div>
    <div class="grid"><label>Partial Key:</label><span>{{ windowsInfo?.partial_product_key || '불러오는 중...' }}</span></div>

    <div class="grid"><label>정품 인증 상태:</label>
      <span v-if="windowsInfo">
        <span v-if="windowsInfo.licensed === 1">✅ 정품</span>
        <span v-else>❌ 미인증</span>
      </span>
    </div>

    <div class="grid"><label>라이센스 채널:</label><span>{{ windowsInfo?.product_channel || '불러오는 중...' }}</span></div>
    <div class="grid"><label>KMS 서버:</label><span>{{ windowsInfo?.kms_machine || '없음' }}</span></div>
    <div class="grid"><label>유예기간 남은일:</label><span>{{ windowsInfo?.grace_period_remaining ?? '불러오는 중...' }}</span></div>
    <div class="grid"><label>예상 만료일:</label><span>{{ windowsInfo?.estimated_expire_date || '불러오는 중...' }}</span></div>

    <div class="grid"><label>크랙 의심 여부:</label>
      <span v-if="windowsInfo">
        <span v-if="windowsInfo.cracked === 1">🚨 크랙 의심 🚨</span>
        <span v-else>✅ 정상</span>
      </span>
    </div>

    <button @click="fetchWindowsInfo">정보 다시 불러오기</button>

    <!-- 디버깅용 전체 데이터 -->
    <!-- <pre>{{ windowsInfo }}</pre> -->
  </fieldset>
</template>

<script lang="ts" setup>
import { ref, onMounted, defineEmits, defineExpose } from 'vue';
import { GetWindowsInfo } from '@wails/go/main/App';
import { waitWailsReady } from '@/utils/wailsReady'


interface WindowsInfo {
  caption: string;
  version: string;
  serial_number: string;
  product_key: string;
  partial_product_key: string;
  licensed: number;
  product_channel: string;
  kms_machine: string;
  grace_period_remaining: number;
  estimated_expire_date: string;
  cracked: number;
}

const windowsInfo = ref<WindowsInfo | null>(null);

const emit = defineEmits<{
  (e: 'loaded', info: WindowsInfo): void
}>()

defineExpose({
    windowsInfo
})

const fetchWindowsInfo = async () => {
  try {
    windowsInfo.value = await GetWindowsInfo();
    if(windowsInfo.value) emit('loaded', windowsInfo.value)
  } catch (error) {
    console.error('Windows 정보 불러오기 실패:', error);
  }
}

onMounted(async () => {
    await waitWailsReady();
    fetchWindowsInfo();
});
</script>

<style scoped>
.software-info {
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
}
.grid {
  display: flex;
  margin-bottom: 0.5rem;
}
.grid label {
  width: 200px;
  font-weight: bold;
}
pre {
  background: #eee;
  padding: 0.5rem;
  border-radius: 4px;
}
</style>
