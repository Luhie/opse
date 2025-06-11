<template>
  <div class="cpu-info-wrap">
    <fieldset v-if="CPUInfo">
      <legend>CPU 정보</legend>
      <div class="grid">
        <label>모델명</label>
        <input type="text" :value="CPUInfo.model" readonly />
        <label>쿨럭</label>
        <input type="text" :value="(CPUInfo.clock / 1000).toFixed(1) + 'GHz'" readonly />
      </div>
      <div class="grid">
        <label>코어</label>
        <input type="text" :value="CPUInfo.cores" readonly />
        <label>쓰레드</label>
        <input type="text" :value="CPUInfo.threads" readonly />
      </div>
    </fieldset>

    <fieldset v-if="MotherboardInfo">
      <legend>Mainboard 정보</legend>
      <div class="grid">
        <label>모델명</label>
        <input type="text" :value="MotherboardInfo.model" readonly />
        <label>시리얼</label>
        <input type="text" :value="MotherboardInfo.serial" readonly />
      </div>
    </fieldset>

    <fieldset v-if="RAMInfo">
      <legend>RAM 정보</legend>
      <div class="grid">
        <label>총용량</label>
        <input type="text" :value="(RAMInfo.totalMB / 1024).toFixed(1) + ' GB'" readonly />
      </div>
      <hr />
      <div v-for="(ram, idx) in RAMInfo.modules" :key="idx">
        <div class="grid">
          <label>제조사</label>
          <input type="text" :value="ram.manufacturer" readonly />
        </div>
        <div class="grid">
          <label>모델명</label>
          <input type="text" :value="ram.model" readonly />
          <label>DDR</label>
          <input type="text" :value="ram.ddr" readonly />
        </div>
        <div class="grid">
          <label>용량</label>
          <input type="text" :value="(ram.capacityMB / 1024).toFixed(1) + ' GB'" readonly />
          <label>속도</label>
          <input type="text" :value="ram.speedMHz" readonly />
        </div>
        <hr />
      </div>
    </fieldset>

    <fieldset v-if="GPUInfo">
      <legend>GPU 정보</legend>
      <div v-for="(gpu, idx) in GPUInfo.cards" :key="idx">
        <div class="grid">
          <label>모델명</label>
          <input type="text" :value="gpu.model" readonly />
          <label>제조사</label>
          <input type="text" :value="gpu.manufacturer" readonly />
        </div>
        <div class="grid">
          <label>RAM 용량</label>
          <input type="text" :value="gpu.vramMB" readonly />
          <label>드라이버 버전</label>
          <input type="text" :value="gpu.driverVersion" readonly />
        </div>
        <hr />
      </div>
    </fieldset>

    <fieldset v-if="DiskInfo">
      <legend>DISK 정보</legend>
      <div v-for="(disk, idx) in DiskInfo.drives" :key="idx">
        <div class="grid">
          <label>모델명</label>
          <input type="text" :value="disk.model" readonly />
          <label>타입</label>
          <input type="text" :value="disk.type" readonly />
        </div>
        <div class="grid">
          <label>용량</label>
          <input type="text" :value="disk.sizeGB + ' GB'" readonly />
          <label>시리얼</label>
          <input type="text" :value="disk.serial" readonly />
        </div>
        <hr />
      </div>
    </fieldset>

    <fieldset v-if="networkInfo">
      <legend>Network 정보</legend>
      <div class="grid">
        <label>Hostname</label>
        <input type="text" :value="networkInfo.hostname" readonly />
      </div>
      <hr />
      <div v-for="(adapter, idx) in networkInfo.adapters" :key="adapter.mac">
        <div class="grid">
          <label>아뎁터 이름</label>
          <input type="text" :value="adapter.name" readonly />
          <label>MAC Address</label>
          <input type="text" :value="adapter.mac" readonly />
        </div>
        <div class="grid">
          <label>IP4</label>
          <input type="text" :value="adapter.ips?.[0]" readonly />
          <label>IP6</label>
          <input type="text" :value="adapter.ips?.[1]" readonly />
        </div>
        <hr />
      </div>
    </fieldset>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { waitWailsReady } from '@/utils/wailsReady'
import {
  GetCPUInfo,
  GetMotherboardInfo,
  GetRAMInfo,
  GetGPUInfo,
  GetDiskInfo,
  GetNetworkInfo,
} from '@wails/go/main/App'

const CPUInfo = ref<any>(null)
const MotherboardInfo = ref<any>(null)
const RAMInfo = ref<any>(null)
const GPUInfo = ref<any>(null)
const DiskInfo = ref<any>(null)
const networkInfo = ref<any>(null)

async function fetchAll() {
  try {
    await waitWailsReady();
    
    CPUInfo.value = await GetCPUInfo()
    MotherboardInfo.value = await GetMotherboardInfo()
    RAMInfo.value = await GetRAMInfo()
    GPUInfo.value = await GetGPUInfo()
    DiskInfo.value = await GetDiskInfo()
    networkInfo.value = await GetNetworkInfo()
  } catch (err) {
    console.error('정보 로드 실패:', err)
  }
}

defineExpose({
  CPUInfo,
  MotherboardInfo,
  RAMInfo,
  GPUInfo,
  DiskInfo,
  networkInfo,
})

onMounted(fetchAll)
</script>

<style scoped>
.cpu-info-wrap {
  display: flex;
  flex-direction: column;
  margin: 0 auto;

  fieldset {
  }
  input {
    background-color: rgb(205, 205, 205);
    font-size: 1.2rem;
  }
}
.grid {
  display: flex;
  align-items: center;
  width: 100%;
  padding: 0.1rem;

  label {
    width: 7rem;
  }
  input {
    flex: 1;
  }
}
</style>
