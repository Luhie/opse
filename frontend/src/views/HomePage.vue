<script setup lang="ts">
import { ref, onMounted } from 'vue'
import SoftwareInfo from '@/components/asset/SoftwareInfo.vue';

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

const handleLoaded = (info: any) => {
    console.log(info.caption)
      try {
    // await InsertWindowsInfo(info);
    console.log('DB 저장 성공');
  } catch (e) {
    console.error('DB 저장 실패', e);
  }
}

</script>
<template>
    <h1>This is Home</h1>

    <SoftwareInfo @loaded='handleLoaded' />
</template>