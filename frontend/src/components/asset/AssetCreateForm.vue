<script setup lang="ts">
import { reactive, ref, watch } from 'vue' 
import AssetDetail from '@/components/asset/AssetDetail.vue'

import OfficeInfo from '@/components/asset/OfficeInfo.vue';
import SoftwareInfo from '@/components/asset/SoftwareInfo.vue';
import { CreateAsset, TestAsset} from '@wails/go/main/App'

const assetDetailRef = ref<InstanceType<typeof AssetDetail> | null>(null)


const visiblePC = ref(false)
const visibleSoftware = ref(false)

const today = new Date()
const year = today.getFullYear()
const month = (today.getMonth()+1).toString().padStart(2,'0')
const date = today.getDate().toString().padStart(2,'0')
const day = `${year}-${month}-${date}`
const defaultForm = {
  corporation: '솜인터내셔널',
  department: '',
  team: '',
  name: '',
  position: '대리',
  location: '',
  userNote: '',

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
};

// 
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
interface OfficeInfo {
  name: string;
  version: string;
  partial_product_key: string;
  licensed: number;
  grace_period: number;
  estimated_expire_date: string;
  cracked: number;
}
// 소프트웨어 정보 저장용 state 변수
const windowsInfo = ref<WindowsInfo | null>(null);
const officeList = ref<OfficeInfo[]>([]);

const GetWindowsInfo = (info: WindowsInfo) => {
  console.log('windowInfo: ', info);
}
const GetOfficeInfoMation = (info: OfficeInfo) => {
  console.log('windowInfo: ', info);
}
const handleWindowsLoaded = (info: WindowsInfo) => {
  // GetWindowsInfo(info)
    windowsInfo.value = info;
}

const handleOfficeLoaded = (infoList: OfficeInfo[]) => {
  // infoList.forEach((info) => {
  //   GetOfficeInfoMation(info);
  // });
    officeList.value = infoList;
};

const form = reactive({ ...defaultForm })
function submitForm() {
  // 1) 자식 컴포넌트 인스턴스 가져오기
  const hw = assetDetailRef.value

  // 1-1) 인스턴스가 붙어 있지 않으면 경고하고 종료
  if (!hw) {
    alert("PC 정보를 먼저 불러와 주세요");
    return;
  }

    // 2) 디버깅용: 자식 인스턴스와 내부 ref 값 콘솔에 찍어 보기
  // console.log('>>> assetDetailRef.value (자식 인스턴스):', hw)
  // console.log('>>> CPUInfo ref 자체:', hw.CPUInfo)
    const systemInfo = {
    CPU: hw?.CPUInfo,
    Motherboard: hw?.MotherboardInfo,
    RAM: hw?.RAMInfo,
    GPU: hw?.GPUInfo,
    Disk: hw?.DiskInfo,
    Network: hw?.networkInfo
  }

  const finalPayload = {
    ...form,
    systemInfo,
    windowsInfo: windowsInfo.value,
    officeList: officeList.value

  }
  console.log(JSON.stringify(finalPayload))
    TestAsset(JSON.stringify(finalPayload))
  CreateAsset(JSON.stringify(finalPayload))
  .then(() => {
    alert('저장 완료!');
    Object.assign(form, defaultForm) // 폼 초기화
  })
  .catch((err: any) => {
    console.error('저장 오류:', err);
    alert('저장 중 오류 발생!');
  }); 
  console.log('제출된 폼:', JSON.stringify(form, null, 2))
  alert('폼이 제출되었습니다! (콘솔 확인)')
}

const sendSoftware = (event: MouseEvent) => {
  if(visibleSoftware.value){
    // event.preventDefault()
  }

}
</script>

<template>
  <div style="display:flex; justify-content:center; align-items:center; border:2px dotted red;">
  <form class="form-wrapper" @submit.prevent="submitForm">
    <!-- 사용자 정보 -->
    <fieldset>
      <legend>사용자 정보</legend>
      <div class="grid-user">
        <div class="grid-item">
          <label for="corporation">법인</label>
          <select id="corporation" v-model="form.corporation">
            <option>솜인터내셔널</option>
            <option>솜밸리</option>
            <option>딜라잇가든</option>
            <option>솜티앤엘</option>
          </select>
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
          <select id="position" v-model="form.position">
            <option>대표</option>
            <option>이사</option>
            <option>부장</option>
            <option>차장</option>
            <option>과장</option>
            <option>대리</option>
            <option>주임</option>
            <option>사원</option>
            <option>인턴</option>
          </select>
        </div>
        <div class="grid-item">
          <label for="location">사용위치</label>
          <input id="location" v-model="form.location" />
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

    <button :class="visibleSoftware ? 'btn-send-software' : ''" type="submit" id='submitBtn' @click="sendSoftware">제출</button>
    <button type="button" @click="visiblePC = !visiblePC">{{ visiblePC ? '닫기' : 'PC정보 가져오기' }}</button>
    <button type="button" @click="visibleSoftware = !visibleSoftware">{{ visibleSoftware ? '닫기' : 'Software 정보 가져오기' }}</button>
  </form>
  </div>
  <AssetDetail ref='assetDetailRef' v-show='visiblePC'/>

  <SoftwareInfo @loaded="handleWindowsLoaded" v-show='visibleSoftware' style='margin-top:20px;'/>
  <OfficeInfo @loaded="handleOfficeLoaded" v-show='visibleSoftware' style='margin-top:20px;' />

</template>

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

  input,
  select{
    flex:1;
  }
  &:nth-child(5){
    /* grid-row: 2/2; */
    /* grid-column: 2/4; */
    /* background-color:red; */
  }
  &:nth-child(7){
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


button.btn-send-software {
  background-color: #fdb0ce;
}
button.btn-send-software:hover {
  background-color: #ff9ac0;
}
</style>
