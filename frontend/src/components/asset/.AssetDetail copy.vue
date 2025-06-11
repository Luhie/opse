<template>
  <div class='cpu-info-wrap'>
    <fieldset v-if='CPUInfo'>
        <legend>CPU 정보</legend>
        <div class='grid'>
            <label>모델명</label>
            <input type='text' :value='CPUInfo.model' readonly />
            <label>쿨럭</label>
            <input type='text' :value="(CPUInfo.clock/1000).toFixed(1)+'GHz'" readonly />
        </div>
        <div class='grid'>
            <label>코어</label>
            <input type='text' v-model='CPUInfo.cores' readonly />
            <label>쓰레드</label>
            <input type='text' v-model='CPUInfo.threads' readonly />
        </div>
    </fieldset>
    <fieldset v-if='MotherboardInfo'>
        <legend>Mainboard 정보</legend>
        <div class='grid'>
            <label>모델명</label>
            <input type='text' v-model='MotherboardInfo.model' readonly/>
            <label>시리얼</label>
            <input type='text' v-model='MotherboardInfo.serial' readonly/>
        </div>
    </fieldset>
    <fieldset v-if='RAMInfo'>
        <legend>RAM 정보</legend>
        <div class='grid'>
          <label>총용량</label>
          <input type='text' :value='RAMInfo.totalMB/1024+" GB"' readonly />
        </div>
        <hr />
        <div v-for='ram in RAMInfo.modules'>
          <div class='grid'>
              <label>제조사</label>
              <input type='text' v-model='ram.manufacturer' readonly />
          </div>
          <div class='grid'>
              <label>모델명</label>
              <input type='text' v-model='ram.model' readonly />
              <label>DDR</label>
              <input type='text' v-model='ram.ddr' readonly />
          </div>
          <div class='grid'>
              <label>용량</label>
              <input type='text' :value='ram.capacityMB/1024+" GB"' readonly />
              <label>속도</label>
              <input type='text' :value='ram.speedMHz' readonly />
          </div>

        </div>
        <hr>
    </fieldset>
    <fieldset v-if='GPUInfo'>
        <legend>GPU 정보</legend>
        <div v-for='gpu in GPUInfo.cards'>
          <div class='grid'>
              <label>모델명</label>
              <input type='text' :value='gpu.model' readonly />
              <label>제조사</label>
              <input type='text' :value='gpu.manufacturer' readonly />
          </div>
          <div class='grid'>
              <label>RAM 용량</label>
              <input type='text' :value='gpu.vramMB' readonly />
              <label>드라이버 버전</label>
              <input type='text' :value='gpu.driverVersion' readonly />
          </div>
        <hr>
        </div>
    </fieldset>
    <fieldset v-if='DiskInfo'>
        <legend>DISK 정보</legend>
        <div v-for='disk in DiskInfo.drives'>
          <div class='grid'>
              <label>모델명</label>
              <input type='text' :value='disk.model' readonly />
              <label>타입</label>
              <input type='text' :value='disk.type' readonly />
          </div>
          <div class='grid'>
              <label>용량</label>
              <input type='text' :value='disk.sizeGB+" GB"' readonly />
              <label>시리얼</label>
              <input type='text' :value='disk.serial' readonly />
          </div>
        <hr>
        </div>
    </fieldset>
    <fieldset v-if='networkInfo'>
        <legend>Network 정보</legend>
        <div class='grid'>
            <label>Hostname</label>
            <input type='text' :value='networkInfo.hostname' readonly />
        </div>
        <hr />
        <div v-for='adapter in networkInfo.adapters' :key='adapter.mac'>
          <div class='grid'>
              <label>아뎁터 이름</label>
              <input type='text' :value='adapter.name' readonly />
              <label>MAC Address</label>
              <input type='text' :value='adapter.mac' readonly />
          </div>
          <div class='grid'>
              <label>IP4</label>
              <input type='text' :value='adapter.ips?.[0]' readonly />
              <label>IP6</label>
              <input type='text' :value='adapter.ips?.[1]' readonly />
          </div>
        <hr>
        </div>
    </fieldset>
  </div>
</template>
<script setup lang='ts'>
import { ref, onMounted } from 'vue'
import { GetCPUInfo,GetMotherboardInfo,GetRAMInfo,GetGPUInfo,GetDiskInfo,GetNetworkInfo } from '@wails/go/main/App'

const CPUInfo = ref<any>(null)
const MotherboardInfo = ref<any>(null)
const RAMInfo = ref<any>(null)
const GPUInfo = ref<any>(null)
const DiskInfo = ref<any>(null)

GetCPUInfo()
.then((result: any) => {
    console.log('>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>')
    console.log(result)
    CPUInfo.value = result
})
.catch((err: any) => {
    console.error('CPU 정보 가져오기 실패: ',err)
})

GetMotherboardInfo()
.then((result: any) => {
  MotherboardInfo.value = result
})
.catch((err: any) => {
    console.error('Motherboard 정보 가져오기 실패: ',err)
})

GetRAMInfo()
.then((result: any) => {
  RAMInfo.value = result
})
.catch((err: any) => {
    console.error('RAM 정보 가져오기 실패: ',err)
})

GetGPUInfo()
.then((result: any) => {
  GPUInfo.value = result
})
.catch((err: any) => {
    console.error('GPU 정보 가져오기 실패: ',err)
})

GetDiskInfo()
.then((result: any) => {
  DiskInfo.value = result
})
.catch((err: any) => {
    console.error('Disk 정보 가져오기 실패: ',err)
})

const networkInfo = ref<any>(null)
const loading = ref(false)
const error = ref<string | null>(null)

async function fetchNetwork() {
    loading.value = true
    try {
        networkInfo.value = await GetNetworkInfo()

    } catch (e) {
        error.value = String(e)
    } finally {
        loading.value = false
    }
}

// 부모(AssetCreateForm) 접근할 수 있도록 노출
defineExpose({
  CPUInfo,
  MotherboardInfo,
  RAMInfo,
  GPUInfo,
  DiskInfo,
  networkInfo
})

onMounted(fetchNetwork)
</script>


<style scoped>
.cpu-info-wrap {
  display:flex;
  flex-direction: column;
  /* width: 40vw; */
  margin:0 auto;
  /* 가로 가운데 정렬 */
  /* justify-content: center; */
  /* 세로 가운데 정렬 */
  /* align-items: center; */
  /* width:50%; */
  /* flex-wrap:wrap; */

  fieldset {
    /* flex: 1 1 40%; */
  }
  input {
    background-color:rgb(205, 205, 205);
    font-size:1.2rem;
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
    input
    {
        flex:1;
    }
}
</style>